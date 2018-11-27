package kbucket

import (
	"container/list"
	"sync"

	"github.com/chenminjian/go-bittorrent/p2p/peer"
)

type RoutingTable struct {
	list  *list.List
	mutex sync.RWMutex
}

func NewRoutingTable() *RoutingTable {
	return &RoutingTable{
		list: list.New(),
	}
}

func (rt *RoutingTable) Add(id peer.ID) {
	rt.mutex.Lock()
	rt.list.PushBack(id)
	rt.mutex.Unlock()
}

// TODO: id is unused.
func (rt *RoutingTable) NearestPeers(id peer.ID, count int) []peer.ID {
	rt.mutex.RLock()
	defer rt.mutex.RUnlock()

	ps := make([]peer.ID, 0, count)
	i := 0
	for e := rt.list.Front(); e != nil; e = e.Next() {
		id := e.Value.(peer.ID)
		ps = append(ps, id)

		i++
		if i >= count {
			break
		}
	}
	return ps
}
