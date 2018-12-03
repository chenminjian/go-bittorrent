package peerstore

import (
	"errors"
	"sync"

	"github.com/chenminjian/go-bittorrent/p2p/peer"
)

type peerstore struct {
	addrs map[peer.ID]PeerInfo
	mutex sync.RWMutex
}

func NewPeerStore() PeerStore {
	return &peerstore{
		addrs: make(map[peer.ID]PeerInfo),
	}
}

func (ps *peerstore) AddAddr(info PeerInfo) {
	ps.mutex.Lock()
	defer ps.mutex.Unlock()

	ps.addrs[info.ID] = info
}

func (ps *peerstore) PeerInfo(id peer.ID) (PeerInfo, error) {
	ps.mutex.RLock()
	defer ps.mutex.RUnlock()

	peerInfo, ok := ps.addrs[id]
	if !ok {
		return PeerInfo{}, errors.New("peer not found")
	}
	return peerInfo, nil
}

func (ps *peerstore) Peers() []peer.ID {
	ps.mutex.RLock()
	defer ps.mutex.RUnlock()

	ids := make([]peer.ID, 0, len(ps.addrs))
	for k, _ := range ps.addrs {
		ids = append(ids, k)
	}

	return ids
}
