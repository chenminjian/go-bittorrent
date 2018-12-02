package net

import (
	"github.com/chenminjian/go-bittorrent/common/addr"
	"github.com/chenminjian/go-bittorrent/p2p/peer"
)

type Network interface {
	LocalPeer() peer.ID

	Listen() error

	SetPacketHandler(PacketHandler)

	SendData(data []byte, addr addr.Addr) error
}

type Packet interface {
	IP() string

	Port() int

	Data() []byte
}

type PacketHandler func(Packet)
