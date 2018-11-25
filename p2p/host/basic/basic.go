package basic

import (
	"github.com/chenminjian/go-bittorrent/p2p/network"
	"github.com/chenminjian/go-bittorrent/p2p/peer"
)

type BasicHost struct {
	id      peer.ID
	port    int
	network net.Network
}

func New(network net.Network) *BasicHost {
	h := &BasicHost{
		network: network,
	}
	return h
}

func (h *BasicHost) ID() peer.ID {
	return ""
}

func (h *BasicHost) Network() net.Network {
	return h.network
}
