package net

type Network interface {
	Listen() error

	SetPacketHandler(PacketHandler)
}

type Packet interface {
	IP() string

	Port() int

	Data() []byte
}

type PacketHandler func(Packet)
