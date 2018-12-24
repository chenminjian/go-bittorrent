package dht

import (
	"errors"
	"fmt"
	"time"

	"github.com/chenminjian/go-bittorrent/common/addr"
	"github.com/chenminjian/go-bittorrent/metric"
	"github.com/chenminjian/go-bittorrent/p2p/host"
	p2pnet "github.com/chenminjian/go-bittorrent/p2p/network"
	"github.com/chenminjian/go-bittorrent/p2p/peer"
	"github.com/chenminjian/go-bittorrent/routing/bencoded"
	"github.com/chenminjian/go-bittorrent/routing/block"
	"github.com/chenminjian/go-bittorrent/routing/kbucket"
	"github.com/chenminjian/go-bittorrent/routing/krpc"
	pstore "github.com/chenminjian/go-bittorrent/routing/peerstore"
	"github.com/chenminjian/go-bittorrent/routing/txmanager"
)

var unknownMsgTypeErr = errors.New("unknown message type")

type DHT struct {
	host         host.Host
	routingTable *kbucket.RoutingTable
	peerstore    pstore.PeerStore
	txMgr        txmanager.TxManager
	reporter     metric.Reporter
	blockList    block.List
}

func New(h host.Host, reporter metric.Reporter) *DHT {
	dht := &DHT{
		host:         h,
		routingTable: kbucket.NewRoutingTable(20, h.ID()),
		peerstore:    pstore.NewPeerStore(),
		txMgr:        txmanager.New(),
		reporter:     reporter,
		blockList:    block.New(),
	}

	h.SetPacketHandler(dht.handlePacket)

	dht.txMgr.SetGCHandler(dht.handleTxGC)

	return dht
}

func (dht *DHT) Bootstrap(peers []addr.Addr) error {

	go dht.doBootstrap(peers)

	return nil
}

func (dht *DHT) FindPeer(target peer.ID) error {
	ids := dht.routingTable.NearestPeers(target, 3)

	fmt.Printf("routing table size:%d\n", dht.routingTable.Size())

	for _, id := range ids {
		if err := dht.findPeerSingle(id, target); err != nil {
			fmt.Printf("findPeerSingle error: %s\n", err)
		}
	}

	return nil
}

// handlePacket handles network packet.
func (dht *DHT) handlePacket(packet p2pnet.Packet) {
	handler := func() error {
		dict, err := bencoded.Decode(packet.Data())
		if err != nil {
			return err
		}

		switch dict["y"] {
		case "q":
			fmt.Println("receive query")
			switch dict["q"] {
			case krpc.Message_PING:
				dht.reporter.PingInc()
				tx, ok := dict["t"].(string)
				if !ok {
					return errors.New("tx is not string")
				}

				content, ok := dict["a"].(map[string]interface{})
				if !ok {
					return errors.New("content is not map[string]string")
				}

				addr := addr.Addr{IP: packet.IP(), Port: packet.Port()}

				return dht.handlePing(tx, content, addr)
			case krpc.Message_FIND_NODE:
				fmt.Println("receive find_node")
				dht.reporter.FindNodeInc()
			case krpc.Message_GET_PEERS:
				fmt.Println("receive get_peers")
				dht.reporter.GetPeersInc()
			case krpc.Message_ANNOUNCE_PEER:
				fmt.Println("receive announce_peer")
				dht.reporter.AnnouncePeerInc()
			default:
				return errors.New("unsupport query")
			}
		case "r":
			txID, ok := dict["t"].(string)
			if !ok {
				return errors.New("tx id format error")
			}

			txInfo, err := dht.txMgr.Get(txID)
			if err != nil {
				return err
			}

			dht.txMgr.Del(txID)

			switch txInfo.Message["q"].(string) {
			case krpc.Message_FIND_NODE:
				content, ok := dict["r"].(map[string]interface{})
				if !ok {
					return errors.New("content is not map[string]string")
				}

				return dht.handleFindNodeResp(content)
			default:
				return errors.New("unreachable")
			}
		default:
			return err
		}
		return nil
	}

	if err := handler(); err != nil {
		fmt.Printf("handle msg error:%v\n", err)
	}
}

func (dht *DHT) handlePing(txID string, content map[string]interface{}, addr addr.Addr) error {
	id, ok := content["id"].(string)
	if !ok {
		return errors.New("id is not string")
	}

	// TODO: record ID
	fmt.Printf("receive ping request, id:%s\n", peer.ID(id).Pretty())

	msg := krpc.NewPingRespMessage(dht.host.ID(), txID)
	encodeMsg := bencoded.Encode(msg)

	if err := dht.host.SendMessage(encodeMsg, addr); err != nil {
		return err
	}

	return nil
}

func (dht *DHT) handleFindNodeResp(content map[string]interface{}) error {
	id, ok := content["id"].(string)
	if !ok {
		return errors.New("id is not string")
	}
	fmt.Printf("receive find_node resp,id:%s\n", peer.ID(id).Pretty())

	nodesStr, ok := content["nodes"].(string)
	if len(nodesStr)%26 != 0 {
		return errors.New("nodes' length is 26's multiple")
	}

	num := len(nodesStr) / 26
	for i := 0; i < num; i++ {
		nodeStr := string(nodesStr[i*26 : (i+1)*26])
		pi, err := pstore.Decode(nodeStr)
		if err != nil {
			return err
		}

		// add peer
		dht.addPeer(pi)
	}

	return nil
}

func (dht *DHT) addPeer(info *pstore.PeerInfo) {
	if dht.blockList.Contain(info.ID) {
		return
	}

	dht.peerstore.AddAddr(*info)
	dht.routingTable.Update(info.ID)
}

func (dht *DHT) doBootstrap(peers []addr.Addr) {
	var count int
	for {
		select {
		case <-time.After(time.Second):
			if count%30 == 0 {
				go dht.connectBootstrapPeers(peers)
			}

			go dht.bootstrapWorker()

			count++
		}
	}
}

func (dht *DHT) connectBootstrapPeers(peers []addr.Addr) {
	fmt.Println("connect bootstrap peers")

	id := peer.RandomID()
	for i := 0; i < len(peers); i++ {
		dht.bootstrapFindPeer(pstore.PeerInfo{Addr: peers[i], ID: "useless"}, id)
	}
}

// bootstrapWorker sends find_peer to three random peer.
func (dht *DHT) bootstrapWorker() {
	id := peer.RandomID()
	if err := dht.FindPeer(id); err != nil {
		fmt.Println(err)
	}
}

func (dht *DHT) findPeerSingle(id peer.ID, target peer.ID) error {
	txID := dht.txMgr.UniqueID()
	msg := krpc.NewFindNodeMessage(dht.host.ID(), target, txID)
	encodeMsg := bencoded.Encode(msg)

	peer, err := dht.peerstore.PeerInfo(id)
	if err != nil {
		return err
	}

	// record tx
	dht.txMgr.Set(msg["t"].(string), &txmanager.TxInfo{
		PeerInfo: peer,
		Message:  msg,
		Curr:     time.Now(),
	})

	// send msg
	err = dht.host.SendMessage(encodeMsg, addr.Addr{IP: peer.IP, Port: peer.Port})
	if err != nil {
		fmt.Printf("ip:%s, port:%d\n", peer.IP, peer.Port)
		return err
	}

	return nil
}

func (dht *DHT) bootstrapFindPeer(receiver pstore.PeerInfo, target peer.ID) error {
	txID := dht.txMgr.UniqueID()
	msg := krpc.NewFindNodeMessage(dht.host.ID(), target, txID)
	encodeMsg := bencoded.Encode(msg)

	// record tx
	dht.txMgr.Set(msg["t"].(string), &txmanager.TxInfo{
		PeerInfo: receiver,
		Message:  msg,
	})

	// send msg
	if err := dht.host.SendMessage(encodeMsg, receiver.Addr); err != nil {
		return err
	}

	return nil
}

func (dht *DHT) handleTxGC(info *txmanager.TxInfo) {
	dht.blockList.Add(info.ID)
}
