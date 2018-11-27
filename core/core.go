package core

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/chenminjian/go-bittorrent/common/addr"
	"github.com/chenminjian/go-bittorrent/config"
	"github.com/chenminjian/go-bittorrent/p2p/host"
	"github.com/chenminjian/go-bittorrent/p2p/peer"
	"github.com/chenminjian/go-bittorrent/routing"
	pstore "github.com/chenminjian/go-bittorrent/routing/peerstore"
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

			id := peer.ID(fmt.Sprintf("%d", i))
			peerInfo := pstore.PeerInfo{Addr: addr.Addr{IP: parts[0], Port: port}, ID: id}
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
