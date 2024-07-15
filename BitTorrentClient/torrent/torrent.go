package torrent

import (
	"bytes"
	"crypto/sha1"
	"fmt"
	"os"

	"github.com/jackpal/bencode-go"
)

type TorrentFile struct {
	Announce string   `bencode:"announce"`
	Info     InfoDict `bencode:"info"`
}

type InfoDict struct {
	PieceLength int    `bencode:"piece length"`
	Pieces      string `bencode:"pieces"`
	Name        string `bencode:"name"`
	Length      int    `bencode:"length"`
}
func Open(path string) (*TorrentFile, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	torrent := TorrentFile{}
	err = bencode.Unmarshal(file, &torrent)
	if err != nil {
		return nil, err
	}
	return &torrent, nil
}

func (t *TorrentFile) InfoHash() [20]byte {
	var buf bytes.Buffer
	bencode.Marshal(&buf, t.Info)
	return sha1.Sum(buf.Bytes())


	//Comment