package peerstore

import "github.com/chenminjian/go-bittorrent/p2p/peer"

type PeerInfo struct {
	IP   string
	Port int
	ID   peer.ID
}
