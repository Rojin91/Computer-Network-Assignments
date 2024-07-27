package torrent

import (
    "bytes"
    "crypto/sha1"
    "fmt"
    "net"
    "net/http"
    "net/url"
    "os"
    "strconv"
    "io"
    "encoding/binary"

    "TorrentClient/bitfield"
    "TorrentClient/handshake"
    "TorrentClient/message"
)

// TorrentFile represents the structure of a torrent file.
type TorrentFile struct {
    Announce string
    Info     InfoDict
}

// InfoDict represents the 'info' dictionary within a torrent file.
type InfoDict struct {
    PieceLength int
    Pieces      string
    Name        string
    Length      int
}

// LoadTorrentFile loads a torrent file from the given path and parses it.
func LoadTorrentFile(path string) (*TorrentFile, error) {
    file, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var torrentFile TorrentFile
    err = Unmarshal(file, &torrentFile)
    if err != nil {
        return nil, err
    }
    return &torrentFile, nil
}

// InfoHash calculates and returns the SHA-1 hash of the bencoded info dictionary.
func (t *TorrentFile) InfoHash() [20]byte {
    var buf bytes.Buffer
    bencodedInfo, err := Marshal(t.Info)
    if err != nil {
        panic(err)
    }
    buf.Write(bencodedInfo)
    return sha1.Sum(buf.Bytes())
}

// Download manages the overall downloading process for a torrent file.
func (t *TorrentFile) Download(outputPath string) error {
    peers, err := getPeers(t)
    if (err != nil) {
        return err
    }
    fmt.Printf("Found %d peers\n", len(peers))

    for _, peer := range peers {
        err := downloadFromPeer(peer, t, outputPath)
        if err != nil {
            fmt.Printf("Failed to download from peer %s: %v\n", peer, err)
        } else {
            fmt.Printf("Successfully downloaded from peer %s\n", peer)
            return nil
        }
    }
    return fmt.Errorf("failed to download from any peers")
}

func downloadFromPeer(peer Peer, t *TorrentFile, outputPath string) error {
    conn, err := Connect(peer)
    if err != nil {
        return err
    }
    defer conn.Close()

    infoHash := t.InfoHash()
    peerID := generatePeerID()
    err = performHandshake(conn, infoHash, peerID)
    if err != nil {
        return err
    }

    bf, err := receiveBitfield(conn)
    if err != nil {
        return err
    }

    return downloadPieces(conn, bf, t, outputPath)
}

func performHandshake(conn net.Conn, infoHash [20]byte, peerID [20]byte) error {
    h := handshake.New(infoHash, peerID)
    _, err := conn.Write(h.Serialize())
    if err != nil {
        return err
    }

    resp, err := handshake.Read(conn)
    if err != nil {
        return err
    }

    if !bytes.Equal(resp.InfoHash[:], infoHash[:]) {
        return fmt.Errorf("infohash mismatch")
    }

    return nil
}

func receiveBitfield(conn net.Conn) (bitfield.Bitfield, error) {
    msg, err := message.Read(conn)
    if err != nil {
        return nil, err
    }

    if msg.ID != message.MsgBitfield {
        return nil, fmt.Errorf("expected bitfield, got ID %d", msg.ID)
    }

    return bitfield.Bitfield(msg.Payload), nil
}

func downloadPieces(conn net.Conn, bf bitfield.Bitfield, t *TorrentFile, outputPath string) error {
    file, err := os.Create(outputPath)
    if err != nil {
        return err
    }
    defer file.Close()

    pieceCount := len(t.Info.Pieces) / sha1.Size
    pieceLength := t.Info.PieceLength
    for i := 0; i < pieceCount; i++ {
        if !bf.HasPiece(i) {
            continue
        }

        piece, err := requestPiece(conn, i, pieceLength)
        if err != nil {
            return err
        }

        _, err = file.Write(piece)
        if err != nil {
            return err
        }
    }

    return nil
}

func requestPiece(conn net.Conn, index int, pieceLength int) ([]byte, error) {
    req := message.Request{
        Index:  index,
        Begin:  0,
        Length: pieceLength,
    }
    _, err := conn.Write(req.Serialize())
    if err != nil {
        return nil, err
    }

    msg, err := message.Read(conn)
    if err != nil {
        return nil, err
    }

    if msg.ID != message.MsgPiece {
        return nil, fmt.Errorf("expected piece, got ID %d", msg.ID)
    }

    return msg.Payload, nil
}

func generatePeerID() [20]byte {
    var peerID [20]byte
    copy(peerID[:], "-BT0001-123456789012")
    return peerID
}

// getPeers communicates with the tracker to retrieve a list of peers for the torrent.
func getPeers(t *TorrentFile) ([]Peer, error) {
    base, err := url.Parse(t.Announce)
    if err != nil {
        return nil, fmt.Errorf("error parsing announce URL: %w", err)
    }

    infoHash := t.InfoHash()
    params := url.Values{
        "info_hash":  {string(infoHash[:])},
        "peer_id":    {"-BT0001-123456789012"},
        "port":       {"6881"},
        "uploaded":   {"0"},
        "downloaded": {"0"},
        "left":       {strconv.Itoa(t.Info.Length)},
        "compact":    {"1"},
    }

    base.RawQuery = params.Encode()

    resp, err := http.Get(base.String())
    if err != nil {
        return nil, fmt.Errorf("error making GET request to tracker: %w", err)
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, fmt.Errorf("error reading response body: %w", err)
    }

    trackerResponse := make(map[string]interface{})
    err = Unmarshal(bytes.NewReader(body), &trackerResponse)
    if err != nil {
        return nil, fmt.Errorf("error unmarshalling tracker response: %w", err)
    }

    fmt.Printf("Tracker Response: %+v\n", trackerResponse)

    peersValue, ok := trackerResponse["peers"]
    if !ok {
        return nil, fmt.Errorf("tracker response does not contain 'peers' field")
    }

    if peersValue == nil {
        return nil, fmt.Errorf("tracker response contains nil 'peers' field")
    }

    fmt.Printf("Peers Value: %v (Type: %T)\n", peersValue, peersValue)

    peersBinary, ok := peersValue.([]byte)
    if !ok {
        return nil, fmt.Errorf("invalid peers format: expected []byte, got %T", peersValue)
    }

    return UnmarshalPeers(peersBinary)
}

// UnmarshalPeers parses a list of peers from a binary representation.
func UnmarshalPeers(peersBinary []byte) ([]Peer, error) {
    const peerSize = 6
    numPeers := len(peersBinary) / peerSize
    peers := make([]Peer, numPeers)

    for i := 0; i < numPeers; i++ {
        peers[i].IP = net.IP(peersBinary[i*peerSize : i*peerSize+4])
        peers[i].Port = binary.BigEndian.Uint16(peersBinary[i*peerSize+4 : i*peerSize+6])
    }
    return peers, nil
}
