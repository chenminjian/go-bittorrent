package peerstore

import (
	"github.com/chenminjian/go-bittorrent/common/addr"
	"github.com/chenminjian/go-bittorrent/p2p/peer"
)

type PeerInfo struct {
	addr.Addr
	ID peer.ID
}
