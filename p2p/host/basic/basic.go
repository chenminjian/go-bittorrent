package basic

import (
	"github.com/chenminjian/go-bittorrent/common/addr"
	p2pnet "github.com/chenminjian/go-bittorrent/p2p/network"
	"github.com/chenminjian/go-bittorrent/p2p/peer"
)

type BasicHost struct {
	id      peer.ID
	port    int
	network p2pnet.Network
}

func New(network p2pnet.Network) *BasicHost {
	h := &BasicHost{
		network: network,
	}
	return h
}

func (h *BasicHost) ID() peer.ID {
	return ""
}

func (h *BasicHost) Network() p2pnet.Network {
	return h.network
}

func (h *BasicHost) SetPacketHandler(handler p2pnet.PacketHandler) {
	h.network.SetPacketHandler(handler)
}

func (h *BasicHost) SendMessage(message string, addr addr.Addr) error {
	return h.network.SendData([]byte(message), addr)
}
