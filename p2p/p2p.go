package p2p

import (
	"github.com/chenminjian/go-bittorrent/p2p/config"
	"github.com/chenminjian/go-bittorrent/p2p/host"
	"github.com/chenminjian/go-bittorrent/p2p/host/basic"
)

type Option = config.Option

func New() host.Host {
	return basic.New()
}
