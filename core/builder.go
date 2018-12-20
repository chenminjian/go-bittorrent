package core

import (
	"github.com/chenminjian/go-bittorrent/config"
	"github.com/chenminjian/go-bittorrent/metric"
	"github.com/chenminjian/go-bittorrent/p2p"
	p2pconfig "github.com/chenminjian/go-bittorrent/p2p/config"
	"github.com/chenminjian/go-bittorrent/p2p/host"
	"github.com/chenminjian/go-bittorrent/p2p/peer"
	"github.com/chenminjian/go-bittorrent/routing/dht"
)

func NewNode(config *config.Config) (*BTNode, error) {
	node := &BTNode{}

	if err := setupNode(node, config); err != nil {
		return nil, err
	}

	return node, nil
}

func setupNode(n *BTNode, config *config.Config) error {
	pid, err := peer.IDB58Decode(config.Pid)
	if err != nil {
		return err
	}

	p2pConfig := &p2pconfig.Config{Port: config.Port, Pid: pid}
	n.PeerHost = p2p.New(p2pConfig)

	startListening(n.PeerHost)

	n.Reporter = metric.New()

	n.Routing = dht.New(n.PeerHost, n.Reporter)

	if err := n.Bootstrap(config); err != nil {
		return err
	}

	return nil
}

func startListening(peerHost host.Host) error {
	if err := peerHost.Network().Listen(); err != nil {
		return err
	}

	return nil
}
