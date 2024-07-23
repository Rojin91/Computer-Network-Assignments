package torrent

import (
    "net"
    "net/http"
    "net/url"
    "strconv"
    "io"
    "bytes"
	"encoding/binary"
    "github.com/jackpal/bencode-go"
)

//getPeers communicates with the tracker to retrieve a list of peers for the torrent.
func getPeers(t *TorrentFile) ([]Peer, error) {
	 // Parse the tracker announce URL from the torrent file.
    base, err := url.Parse(t.Announce)
    if err != nil {
        return nil, err
    }

	// Store the result of t.InfoHash() in a variable for later use.
	infoHash := t.InfoHash()

	 // Construct the query parameters for the tracker request.
    params := url.Values{
        "info_hash":  {string(infoHash[:])},
        "peer_id":    {"-BT0001-123456789012"},
        "port":       {"6881"},
        "uploaded":   {"0"},
        "downloaded": {"0"},
        "left":       {strconv.Itoa(t.Info.Length)},
        "compact":    {"1"},
    }

	// Encode the query parameters and attach them to the base URL.
    base.RawQuery = params.Encode()

	// Make an HTTP GET request to the tracker's announce URL.
    resp, err := http.Get(base.String())
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

	// Read the response body.
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

	// Unmarshal the bencoded response into a map.
    trackerResponse := make(map[string]interface{})
    err = bencode.Unmarshal(bytes.NewReader(body), &trackerResponse)
    if err != nil {
        return nil, err
    }

	// Parse the peer list from the tracker response.
    peers, err := parsePeers(trackerResponse["peers"].([]byte))
    if err != nil {
        return nil, err
    }
    return peers, nil
}

// parsePeers converts the compact peer list into a slice of Peer structs.
func parsePeers(peersBinary []byte) ([]Peer, error) {
    const peerSize = 6
    numPeers := len(peersBinary) / peerSize
    peers := make([]Peer, numPeers)

	// Iterate over the peer list and extract IP and port for each peer.
    for i := 0; i < numPeers; i++ {
        peers[i].IP = net.IP(peersBinary[i*peerSize : i*peerSize+4])
        peers[i].Port = binary.BigEndian.Uint16(peersBinary[i*peerSize+4 : i*peerSize+6])
    }
    return peers, nil
}
