package torrent

import(
	"fmt"
	"net"
	"time"
)

type Peer struct{
	IP net.IP
	Port uint16
}

func (p Peer) String() string {
	return fmt.Sprintf("%s:%d",p.IP.String(),p.Port)
}

func Connect(peer Peer) (net.Conn, error) {
	//Set a 3-second timer to connect to the peer
 	conn, err := net.DialTimeout("tcp",peer.String(),3*time.Second)
	if err != nil {
		return nil, err
	}
	return conn, err
}
