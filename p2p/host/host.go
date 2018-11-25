package host

import (
	"github.com/chenminjian/go-bittorrent/p2p/network"
	"github.com/chenminjian/go-bittorrent/p2p/peer"
)

type Host interface {
	ID() peer.ID

	Network() net.Network
}
