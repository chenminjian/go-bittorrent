package core

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/chenminjian/go-bittorrent/common/addr"
	"github.com/chenminjian/go-bittorrent/config"
	"github.com/chenminjian/go-bittorrent/metric"
	"github.com/chenminjian/go-bittorrent/p2p/host"
	"github.com/chenminjian/go-bittorrent/routing"
)

type BTNode struct {
	PeerHost host.Host
	Routing  routing.Routing
	Reporter metric.Reporter
}

func (node *BTNode) Bootstrap(conf *config.Config) error {

	trans := func(conf *config.Config) []addr.Addr {

		bootstrapPeers := make([]addr.Addr, 0, len(conf.Bootstrap))
		for i := 0; i < len(conf.Bootstrap); i++ {
			bs := conf.Bootstrap[i]
			parts := strings.Split(bs, ":")
			port, err := strconv.Atoi(parts[1])
			if err != nil {
				fmt.Println(err)
				continue
			}

			address := addr.Addr{IP: parts[0], Port: port}
			bootstrapPeers = append(bootstrapPeers, address)
		}

		return bootstrapPeers
	}

	bootstrapPeers := trans(conf)
	if err := node.Routing.Bootstrap(bootstrapPeers); err != nil {
		return err
	}

	return nil
}
