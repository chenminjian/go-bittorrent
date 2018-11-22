package p2p

import (
	"github.com/chenminjian/go-bittorrent/p2p/config"
	"github.com/chenminjian/go-bittorrent/p2p/host"
	"github.com/chenminjian/go-bittorrent/p2p/host/basic"
	"github.com/chenminjian/go-bittorrent/p2p/network/swarm"
)

type Option = config.Option

func New(config *config.Config) host.Host {
	network := swarm.New(config)
	return basic.New(network)
}
