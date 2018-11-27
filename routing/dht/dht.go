package dht

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/chenminjian/go-bittorrent/p2p/host"
	"github.com/chenminjian/go-bittorrent/p2p/network"
	"github.com/chenminjian/go-bittorrent/p2p/peer"
	"github.com/chenminjian/go-bittorrent/routing/kbucket"
	pstore "github.com/chenminjian/go-bittorrent/routing/peerstore"
)

type DHT struct {
	host         host.Host
	routingTable *kbucket.RoutingTable
	peerstore    pstore.PeerStore
}

func New(h host.Host) *DHT {
	dht := &DHT{
		host:         h,
		routingTable: kbucket.NewRoutingTable(),
		peerstore:    pstore.NewPeerStore(),
	}

	h.SetPacketHandler(dht.handlePacket)

	return dht
}

func (dht *DHT) Bootstrap(peers []pstore.PeerInfo) error {
	for i := 0; i < len(peers); i++ {
		info := peers[i]
		dht.routingTable.Add(info.ID)
		dht.peerstore.AddAddr(info)
	}

	go dht.doBoostrap()

	return nil
}

func (dht *DHT) FindPeer(id peer.ID) error {
	return nil
}

func (dht *DHT) handlePacket(packet net.Packet) {
	fmt.Printf("receive:%s", string(packet.Data()))
}

func (dht *DHT) doBoostrap() {
	for {
		select {
		case <-time.After(time.Second * 10):
			dht.boostrapWorker()
		}
	}
}

func (dht *DHT) boostrapWorker() {
	fmt.Println("test")
	randId := func() peer.ID {
		data := make([]byte, 160)
		rand.Read(data)
		return peer.ID(data)
	}

	id := randId()
	if err := dht.FindPeer(id); err != nil {
		fmt.Println(err)
	}
}
