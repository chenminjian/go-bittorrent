package peerstore

import "github.com/chenminjian/go-bittorrent/p2p/peer"

type PeerStore interface {
	AddAddr(PeerInfo)

	PeerInfo(peer.ID) (PeerInfo, error)

	Peers() []peer.ID
}
