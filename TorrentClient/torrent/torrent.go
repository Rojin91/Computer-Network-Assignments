package torrent

import (
    "bytes"
    "crypto/sha1"
    "fmt"
    "net"
    "os"
    "path/filepath"

    "TorrentClient/bitfield"
    "TorrentClient/handshake"
    "TorrentClient/message"
)

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
    if err != nil {
        return err
    }
    fmt.Printf("Found %d peers\n", len(peers))

    // Ensure the output directory exists
    err = os.MkdirAll(filepath.Dir(outputPath), os.ModePerm)
    if err != nil {
        return err
    }

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

        for {
            piece, err := requestPiece(conn, i, pieceLength)
            if err != nil {
                fmt.Printf("Failed to download piece %d: %v\n", i, err)
                continue
            }

            _, err = file.Write(piece)
            if err != nil {
                return err
            }

            break
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
