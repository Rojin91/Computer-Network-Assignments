package message

import (
	"encoding/binary"
	"io"
	"net"
)

const (
	MsgChoke         = 0
	MsgUnchoke       = 1
	MsgInterested    = 2
	MsgNotInterested = 3
	MsgHave          = 4
	MsgBitfield      = 5
	MsgRequest       = 6
	MsgPiece         = 7
	MsgCancel        = 8
)

type Message struct {
	Length  int
	ID      byte
	Payload []byte
}

func Read(conn net.Conn) (*Message, error) {
	lengthBuf := make([]byte, 4)
	_, err := io.ReadFull(conn, lengthBuf)
	if err != nil {
		return nil, err
	}
	length := binary.BigEndian.Uint32(lengthBuf)
	if length == 0 {
		return &Message{Length: 0}, nil
	}

	messageBuf := make([]byte, length)
	_, err = io.ReadFull(conn, messageBuf)
	if err != nil {
		return nil, err
	}

	return &Message{
		Length:  int(length),
		ID:      messageBuf[0],
		Payload: messageBuf[1:],
	}, nil
}

func (m *Message) Serialize() []byte {
	length := uint32(len(m.Payload) + 1)
	buf := make([]byte, 4+length)
	binary.BigEndian.PutUint32(buf[:4], length)
	buf[4] = m.ID
	copy(buf[5:], m.Payload)
	return buf
}

type Request struct {
	Index  int
	Begin  int
	Length int
}

func (r *Request) Serialize() []byte {
	buf := make([]byte, 13)
	binary.BigEndian.PutUint32(buf[0:4], uint32(r.Index))
	binary.BigEndian.PutUint32(buf[4:8], uint32(r.Begin))
	binary.BigEndian.PutUint32(buf[8:12], uint32(r.Length))
	return buf
}
