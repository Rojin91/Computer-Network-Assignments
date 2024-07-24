package torrent

import(
	"bytes"
	"crypto/sha1"
	"fmt"
	"io"
	"net"
	"os"

	"TorrentClient/bitfield"
    "TorrentClient/handshake"
    "TorrentClient/message"
)

//Download manages the overall downloading process for a torrent file
func (t *TorrentFile) Download(outputPath string) error {
	//Retrieve peers list from peer.go
	peers, err := getPeers(t)
	if err != nil {
		return err
	}
	fmt.Printf("Found %d peers\n", len(peers))
	
	//Try downloading from each peer until successfull
	for _, peer := range peers {
		err := downloadFromPeer(peer, t, outputPath)
		if err != nil {
			fmt.Printf("Failed to download from peer %s : %v\n",peer,err)
		} else {
			fmt.Printf("Successfully downloaded from peer %s\n",peer)
		}
	}
	return fmt.Errorf("Failed to download from any peers!\n")
}

//Attempts to download from the single peer
func downloadFromPeer(peer Peer,t *TorrentFile, outputPath string) error {
	//Connect to the peer
	conn, err := Connect(peer)
	if err != nil {
		return err
	}
	defer conn.Close()

	//Perform the BitTorrent Handshake
	infoHash := t.InfoHash()
	peerID := generatePeerId()
	err = performHandshake(conn,infoHash,peerID)
	if err != nil {
		return err
	}

	//Receive and parse the peer's bitfield
	bf, err := receiveBitfield(conn)
	if err != nil {
		return err
	}

	//Download pieces from peer
	return downloadPieces(conn, bf, t, outputPath)
}

//performs the BitTorrent handshake with the peer
func performHandshake(conn net.Conn, infoHash [20]byte, peerID [20]byte) error {
	//create and send the handshake
	h := handshake.New(infoHash,peerID)
	//serialized handshake is sent to the peer over the established connection
	_, err := conn.Write(h.Serialize())
	if err != nil {
		return err
	}

	resp, err := handshake.Read(conn)]
	if err != nil {
		return err
	}
	if !bytes.Equal(resp.InfoHash[:],infoHash[:]) {
		return fmt.Errorf("infohash mismatch")
	}
	return nil
}

//receives and returns the bitfield to the peer
func receiveBitfield(conn net.Conn) (bitfield.Bitfield, error) {
	//read the message
	msg, err := message.Read(conn)
	if err != nil {
		return nil, err
	}

	//ensure the message is a bitfield message.
	if msg.ID != message.MsgBitfield {
        return nil, fmt.Errorf("expected bitfield, got ID %d", msg.ID)
    }

	return msg.Payload, nil
}

func downloadPieces(conn net.Conn, bf bitfield.Bitfield,t *TorrentFile, outputPath string) error {
	//open the output file
	file, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer file.Close()

	//Create a buffer to hold the piece data
	pieceCount := len(t.Info.Pieces)/sha1.Size
	pieceLength := t.Info.PieceLength
	for i := 0 ; i < pieceCount ; i++ {
		if !bf.HasPiece(i) {
			continue
		}
		//request the piece from the peer
		piece, err := requestPiece(conn, i, pieceLength)
		if err != nil {
			return err
		}

		//write the piece to the output file
		_, err := file.Write(piece)
		if err != nil {
			return err
		}
	}

	return nil
}

//request a piece from the user 
func requestPiece(conn net.Conn, index int, pieceLength int)([]byte, error){
	//create a message request
	req := message.Request{
		Index: index,
		Begin: 0,
		Length: pieceLength
	}
	_,err := conn.Write(req.Serialize())
	if err != nil {
		return err
	}

	//read the response message
	msg, err := message.Read(conn)
	if err != nil {
		return nil, err
	}

	//ensure the message is a piece message
	if msg.ID != message.MsgPiece {
		return nil, fmt.Errorf("expected piece, got ID %d", msg.ID)
	}
	return msg.Payload, nil
}

// generatePeerID generates a unique peer ID.
func generatePeerId() [20]byte {
	var peerID [20]byte
	copy(peerID[:],"-BT0001-123456789012")
	return peerID
}




