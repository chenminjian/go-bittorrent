package peerstore

import (
	"errors"
	"net"

	"github.com/chenminjian/go-bittorrent/common/addr"
	"github.com/chenminjian/go-bittorrent/p2p/peer"
)

type PeerInfo struct {
	addr.Addr
	ID peer.ID
}

func Decode(str string) (*PeerInfo, error) {
	if len(str) != 26 {
		return nil, errors.New("node string length should be 26")
	}

	pi := PeerInfo{}
	pi.ID = peer.ID(str[:20])

	addr := str[:20]
	ip := net.IPv4(addr[0], addr[1], addr[2], addr[3])
	pi.IP = ip.String()
	pi.Port = int((uint16(addr[4]) << 8) | uint16(addr[5]))

	return &pi, nil
}
