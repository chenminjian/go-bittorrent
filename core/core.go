package core

import (
	"fmt"
	"github.com/chenminjian/go-bittorrent/config"
	"github.com/chenminjian/go-bittorrent/p2p/host"
	"github.com/chenminjian/go-bittorrent/routing"
	pstore "github.com/chenminjian/go-bittorrent/routing/peerstore"
	"strconv"
	"strings"
)

type BTNode struct {
	PeerHost host.Host
	Routing  routing.Routing
}

func (node *BTNode) Bootstrap(conf *config.Config) error {

	trans := func(conf *config.Config) []pstore.PeerInfo {

		bootstrapPeers := make([]pstore.PeerInfo, 0, len(conf.Bootstrap))
		for i := 0; i < len(conf.Bootstrap); i++ {
			bs := conf.Bootstrap[i]
			parts := strings.Split(bs, ":")
			port, err := strconv.Atoi(parts[1])
			if err != nil {
				fmt.Println(err)
				continue
			}

			peerInfo := pstore.PeerInfo{IP: parts[0], Port: port}
			bootstrapPeers = append(bootstrapPeers, peerInfo)
		}

		return bootstrapPeers
	}

	bootstrapPeers := trans(conf)
	if err := node.Routing.Bootstrap(bootstrapPeers); err != nil {
		return err
	}

	return nil
}
