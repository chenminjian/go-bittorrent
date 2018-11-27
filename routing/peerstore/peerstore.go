package peerstore

import (
	"errors"

	"github.com/chenminjian/go-bittorrent/p2p/peer"
)

type peerstore struct {
	addrs map[peer.ID]PeerInfo
}

func NewPeerStore() PeerStore {
	return &peerstore{
		addrs: make(map[peer.ID]PeerInfo),
	}
}

func (ps *peerstore) AddAddr(info PeerInfo) {
	ps.addrs[info.ID] = info
}

func (ps *peerstore) PeerInfo(id peer.ID) (PeerInfo, error) {
	peerInfo, ok := ps.addrs[id]
	if !ok {
		return PeerInfo{}, errors.New("peer not found")
	}
	return peerInfo, nil
}

func (ps *peerstore) Peers() []peer.ID {
	ids := make([]peer.ID, 0, len(ps.addrs))
	for k, _ := range ps.addrs {
		ids = append(ids, k)
	}

	return ids
}
