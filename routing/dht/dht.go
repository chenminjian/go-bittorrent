package dht

import (
	"fmt"

	"github.com/chenminjian/go-bittorrent/p2p/host"
	"github.com/chenminjian/go-bittorrent/p2p/network"
)

type DHT struct {
	host host.Host
}

func New(h host.Host) *DHT {
	dht := &DHT{
		host: h,
	}

	h.SetPacketHandler(dht.handlePacket)

	return dht
}

func (dht *DHT) Bootstrap() error {
	return nil
}

func (dht *DHT) handlePacket(packet net.Packet) {
	fmt.Printf("receive:%s", string(packet.Data()))
}
