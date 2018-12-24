package block

import "github.com/chenminjian/go-bittorrent/p2p/peer"

type List interface {
	Add(id peer.ID)
	Size() int
	Contain(id peer.ID) bool
}
