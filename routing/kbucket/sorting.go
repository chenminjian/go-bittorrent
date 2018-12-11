package kbucket

import (
	"bytes"
	"container/list"

	"github.com/chenminjian/go-bittorrent/p2p/peer"
)

// A helper struct to sort peers by their distance to the local node
type peerDistance struct {
	p        peer.ID
	distance []byte
}

// peerSorterArr implements sort.Interface to sort peers by xor distance
type peerSorterArr []*peerDistance

func (p peerSorterArr) Len() int      { return len(p) }
func (p peerSorterArr) Swap(a, b int) { p[a], p[b] = p[b], p[a] }
func (p peerSorterArr) Less(a, b int) bool {
	return bytes.Compare(p[a].distance, p[b].distance) < 0
}

// copy all peerList's peer to peerArr
func copyPeersFromList(target peer.ID, peerArr peerSorterArr, peerList *list.List) peerSorterArr {
	if cap(peerArr) < len(peerArr)+peerList.Len() {
		newArr := make(peerSorterArr, 0, len(peerArr)+peerList.Len())
		copy(newArr, peerArr)
		peerArr = newArr
	}
	for e := peerList.Front(); e != nil; e = e.Next() {
		p := e.Value.(peer.ID)
		pd := peerDistance{
			p:        p,
			distance: xor([]byte(target), []byte(p)),
		}
		peerArr = append(peerArr, &pd)
	}
	return peerArr
}
