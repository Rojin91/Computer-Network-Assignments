package handshake

import (
	"fmt"
	"io"
	"net"
)

const ProtocolIdentifier = "BitTorrent protocol"

// Handshake represents a BitTorrent handshake message.
type Handshake struct {
	Pstr     string
	InfoHash [20]byte
	PeerID   [20]byte
}

// New creates a new Handshake.
func New(infoHash, peerID [20]byte) *Handshake {
	return &Handshake{
		Pstr:     ProtocolIdentifier,
		InfoHash: infoHash,
		PeerID:   peerID,
	}
}

// Serialize serializes the handshake into a byte slice.
func (h *Handshake) Serialize() []byte {
	buf := make([]byte, len(h.Pstr)+49)
	buf[0] = byte(len(h.Pstr))
	copy(buf[1:], h.Pstr)
	copy(buf[1+len(h.Pstr)+8:], h.InfoHash[:])
	copy(buf[1+len(h.Pstr)+8+20:], h.PeerID[:])
	return buf
}

// Read reads a handshake from the given connection.
func Read(conn net.Conn) (*Handshake, error) {
	lengthBuf := make([]byte, 1)
	_, err := io.ReadFull(conn, lengthBuf)
	if err != nil {
		return nil, err
	}
	pstrlen := int(lengthBuf[0])

	if pstrlen == 0 {
		return nil, fmt.Errorf("invalid pstrlen")
	}

	handshakeBuf := make([]byte, pstrlen+48)
	_, err = io.ReadFull(conn, handshakeBuf)
	if err != nil {
		return nil, err
	}

	h := &Handshake{}
	h.Pstr = string(handshakeBuf[:pstrlen])
	copy(h.InfoHash[:], handshakeBuf[pstrlen+8:pstrlen+28])
	copy(h.PeerID[:], handshakeBuf[pstrlen+28:])

	return h, nil
}
