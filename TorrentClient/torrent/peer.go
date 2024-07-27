package torrent

import (
    "fmt"
    "net"
)

type Peer struct {
    IP   net.IP
    Port uint16
}

func (p Peer) String() string {
    return fmt.Sprintf("%s:%d", p.IP, p.Port)
}

func Connect(peer Peer) (net.Conn, error) {
    conn, err := net.Dial("tcp", peer.String())
    if err != nil {
        return nil, err
    }
    return conn, nil
}
