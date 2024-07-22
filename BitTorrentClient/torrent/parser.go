package torrent

import (
	"bytes"
	"crypto/sha1"
	"github.com/jackpal/bencode-go"
)

//Infohash calculates and return SHA-1 hash of the 'info' dictionary from the .torrent file
func (t *TorrentFile) InfoHash() [20]byte {
	var buf bytes.Buffer
	//Marshall the 'info' dictionary into buffer using bencode
	bencode.Marshal(&buf, t.Info)
    //Compute the SHA-1 hash of the buffer's content
	return sha1.Sum(buf.Bytes())
}
