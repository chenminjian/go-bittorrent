package swarm

import (
	"fmt"
	"net"

	"github.com/chenminjian/go-bittorrent/common/addr"
	"github.com/chenminjian/go-bittorrent/p2p/config"
	p2pnet "github.com/chenminjian/go-bittorrent/p2p/network"
	"github.com/chenminjian/go-bittorrent/p2p/peer"
)

type Swarm struct {
	id        peer.ID
	port      int
	conn      *net.UDPConn
	receivedC chan Packet
	packetH   p2pnet.PacketHandler
}

func New(config *config.Config) *Swarm {
	swm := &Swarm{
		id:        config.Pid,
		port:      config.Port,
		receivedC: make(chan Packet),
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

	go func() {
		buff := make([]byte, 8192)
		for {
			n, raddr, err := s.conn.ReadFromUDP(buff)
			if err != nil {
				fmt.Printf("ReadFromUDP error:%s\n", err)
				continue
			}

			s.receivedC <- Packet{data: buff[:n], addr: raddr}
		}
	}()

	go s.startHandlingPackets()

	return nil
}

func (s *Swarm) SetPacketHandler(handler p2pnet.PacketHandler) {
	s.packetH = handler
}

func (s *Swarm) SendData(data []byte, addr addr.Addr) error {
	raddr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%d", addr.IP, addr.Port))
	if err != nil {
		return err
	}

	// fmt.Printf("send find_node to: %s\n", raddr)

	if _, err := s.conn.WriteToUDP(data, raddr); err != nil {
		return err
	}
	return nil
}

func (s *Swarm) startHandlingPackets() {
	for p := range s.receivedC {
		if s.packetH == nil {
			fmt.Println("packet handler unset")

			continue
		}

		s.packetH(&p)
	}
}
