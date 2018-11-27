package dht

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/chenminjian/go-bittorrent/common/addr"
	"github.com/chenminjian/go-bittorrent/p2p/host"
	p2pnet "github.com/chenminjian/go-bittorrent/p2p/network"
	"github.com/chenminjian/go-bittorrent/p2p/peer"
	"github.com/chenminjian/go-bittorrent/routing/bencoded"
	"github.com/chenminjian/go-bittorrent/routing/kbucket"
	"github.com/chenminjian/go-bittorrent/routing/krpc"
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

func (dht *DHT) FindPeer(target peer.ID) error {
	ids := dht.routingTable.NearestPeers(target, 3)

	for _, id := range ids {
		if err := dht.findPeerSingle(id, target); err != nil {
			fmt.Printf("findPeerSingle error: %s\n", err)
		}
	}

	return nil
}

func (dht *DHT) handlePacket(packet p2pnet.Packet) {
	fmt.Printf("receive:%s", string(packet.Data()))
}

func (dht *DHT) doBoostrap() {
	for {
		select {
		case <-time.After(time.Second * 5):
			dht.boostrapWorker()
		}
	}
}

func (dht *DHT) boostrapWorker() {
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

func (dht *DHT) findPeerSingle(id peer.ID, target peer.ID) error {
	msg := krpc.NewFindNodeMessage(dht.host.ID(), target)
	encodeMsg := bencoded.Encode(msg)

	peer, err := dht.peerstore.PeerInfo(id)
	if err != nil {
		return err
	}

	// send msg
	err = dht.host.SendMessage(encodeMsg, addr.Addr{IP: peer.IP, Port: peer.Port})
	if err != nil {
		return err
	}

	return nil
}
