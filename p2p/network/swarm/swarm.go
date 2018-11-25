package swarm

import (
	"fmt"
	"net"

	"github.com/chenminjian/go-bittorrent/p2p/config"
	"github.com/chenminjian/go-bittorrent/p2p/peer"
)

type Swarm struct {
	id   peer.ID
	port int
	conn *net.UDPConn
}

func New(config *config.Config) *Swarm {
	swm := &Swarm{
		id:   config.Pid,
		port: config.Port,
	}
	return swm
}

func (s *Swarm) Listen() error {
	addr := fmt.Sprintf(":%d", s.port)
	listener, err := net.ListenPacket("udp", addr)
	if err != nil {
		return err
	}

	s.conn = listener.(*net.UDPConn)
	return nil
}
