package torrent

import (
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
	//attempt to open the file from given path
	file, err := os.Open(path)	
	if err != nil {
		//If there is error in opening the file, return nil and error
		return nil, err
	}
	//Ensuring the file is closed when the function exits
	defer file.Close()
    //using the bencode library from github to decode the file contents to the torrentFile
	torrent := TorrentFile{}
	err = bencode.Unmarshal(file, &torrent)
	if err != nil {
		return nil, err
	}
	//return the populated torrentFile struct and nil error
	return &torrent, nil
}
