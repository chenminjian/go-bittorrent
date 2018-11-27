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
