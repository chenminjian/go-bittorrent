package swarm

import (
	"github.com/chenminjian/go-bittorrent/p2p/config"
	"github.com/chenminjian/go-bittorrent/p2p/peer"
)

type Swarm struct {
	id   peer.ID
	port int
}

func New(config *config.Config) *Swarm {
	swm := &Swarm{
		id:   config.Pid,
		port: config.Port,
	}
	return swm
}

func (s *Swarm) Listen(port int) error {
	return nil
}
