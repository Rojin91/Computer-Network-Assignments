package torrent

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"strconv"
)

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
