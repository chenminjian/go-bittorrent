package block

import (
	"sync"

	"github.com/chenminjian/go-bittorrent/p2p/peer"
)

type list struct {
	peers map[peer.ID]interface{}
	mutex sync.RWMutex
}

func New() List {
	return &list{
		peers: make(map[peer.ID]interface{}),
	}
}

func (l *list) Add(id peer.ID) {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	l.peers[id] = struct{}{}
}

func (l *list) Size() int {
	l.mutex.RLock()
	defer l.mutex.RUnlock()

	return len(l.peers)
}

func (l *list) Contain(id peer.ID) bool {
	l.mutex.RLock()
	defer l.mutex.RUnlock()

	if _, ok := l.peers[id]; ok {
		return true
	}
	return false
}
