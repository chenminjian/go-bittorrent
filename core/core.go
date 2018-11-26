package core

import (
	"github.com/chenminjian/go-bittorrent/p2p/host"
	"github.com/chenminjian/go-bittorrent/routing"
)

type BTNode struct {
	PeerHost host.Host
	Routing  routing.Routing
}
