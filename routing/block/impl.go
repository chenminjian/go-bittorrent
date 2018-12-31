package block

import (
	"sync"
	"time"

	"github.com/chenminjian/go-bittorrent/p2p/peer"
)

type list struct {
	peers map[peer.ID]time.Time
	mutex sync.RWMutex
}

func New() List {
	l := &list{
		peers: make(map[peer.ID]time.Time),
	}

	l.gc(time.Hour)

	return l
}

func (l *list) Add(id peer.ID) {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	l.peers[id] = time.Now()
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

func (l *list) gc(period time.Duration) {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	now := time.Now()
	for k, v := range l.peers {
		if now.Sub(v) > period {
			delete(l.peers, k)
		}
	}
}
