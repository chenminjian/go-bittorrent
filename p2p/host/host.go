package host

import (
	"github.com/chenminjian/go-bittorrent/common/addr"
	p2pnet "github.com/chenminjian/go-bittorrent/p2p/network"
	"github.com/chenminjian/go-bittorrent/p2p/peer"
)

type Host interface {
	ID() peer.ID

	Network() p2pnet.Network

	SetPacketHandler(p2pnet.PacketHandler)

	SendMessage(message string, addr addr.Addr) error
}
