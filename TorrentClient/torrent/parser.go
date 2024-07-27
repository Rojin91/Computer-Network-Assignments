package torrent

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"io"
	"os"
)

// ParseInfoHash parses the info hash from the .torrent file
func ParseInfoHash(r io.Reader) ([20]byte, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return [20]byte{}, err
	}

	var infoDict map[string]interface{}
	if err := Unmarshal(bytes.NewReader(data), &infoDict); err != nil {
		return [20]byte{}, err
	}
	info, ok := infoDict["info"]
	if !ok {
		return [20]byte{}, errors.New("no info dictionary found")
	}
	bencodedInfo, err := Marshal(info)
	if err != nil {
		return [20]byte{}, err
	}
	return sha1.Sum(bencodedInfo), nil
}

// InfoHashFromFile computes the info hash from a .torrent file
func InfoHashFromFile(path string) ([20]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return [20]byte{}, err
	}
	defer file.Close()
	return ParseInfoHash(file)
}

// InfoHashToString converts the info hash to a string
func InfoHashToString(infoHash [20]byte) string {
	return hex.EncodeToString(infoHash[:])
}
