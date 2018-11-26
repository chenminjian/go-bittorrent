package core

import (
	"github.com/chenminjian/go-bittorrent/config"
	"github.com/chenminjian/go-bittorrent/p2p"
	p2pconfig "github.com/chenminjian/go-bittorrent/p2p/config"
	"github.com/chenminjian/go-bittorrent/p2p/host"
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

	p2pConfig := &p2pconfig.Config{Port: config.Port, Pid: config.Pid}
	n.PeerHost = p2p.New(p2pConfig)

	startListening(n.PeerHost)

	n.Routing = dht.New(n.PeerHost)

	return nil
}

func startListening(peerHost host.Host) error {
	if err := peerHost.Network().Listen(); err != nil {
		return err
	}

	return nil
}
