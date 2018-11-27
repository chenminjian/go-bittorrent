package net

import (
	"github.com/chenminjian/go-bittorrent/common/addr"
)

type Network interface {
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
