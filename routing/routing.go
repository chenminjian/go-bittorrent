package routing

import (
	"github.com/chenminjian/go-bittorrent/p2p/peer"
	pstore "github.com/chenminjian/go-bittorrent/routing/peerstore"
)

type Routing interface {
	Bootstrap(peers []pstore.PeerInfo) error

	FindPeer(id peer.ID) error
}
