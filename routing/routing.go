package routing

import (
	"github.com/chenminjian/go-bittorrent/common/addr"
	"github.com/chenminjian/go-bittorrent/p2p/peer"
)

type Routing interface {
	Bootstrap(peers []addr.Addr) error

	FindPeer(id peer.ID) error
}
