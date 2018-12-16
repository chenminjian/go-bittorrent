package kbucket

import (
	"sort"
	"sync"

	"github.com/chenminjian/go-bittorrent/p2p/peer"
)

type RoutingTable struct {
	// ID of the local peer
	local peer.ID

	// Buckets define all the fingers to other nodes.
	Buckets []*Bucket

	bucketsize int

	mutex sync.RWMutex
}

func NewRoutingTable(bucketsize int, localID peer.ID) *RoutingTable {
	return &RoutingTable{
		local:      localID,
		Buckets:    []*Bucket{newBucket()},
		bucketsize: bucketsize,
	}
}

// Update adds or
// moves the given peer to the front of its respective bucket
func (rt *RoutingTable) Update(p peer.ID) {
	cpl := commonPrefixLen(p, rt.local)

	rt.mutex.Lock()
	defer rt.mutex.Unlock()

	bucketID := cpl
	if bucketID >= len(rt.Buckets) {
		bucketID = len(rt.Buckets) - 1
	}

	bucket := rt.Buckets[bucketID]
	if bucket.Has(p) {
		// If the peer is already in the table, move it to the front.
		// This signifies that it it "more active" and the less active nodes
		// Will as a result tend towards the back of the list
		bucket.MoveToFront(p)
		return
	}

	// New peer, add to bucket
	bucket.PushFront(p)

	// Are we past the max bucket size?
	if bucket.Len() > rt.bucketsize {
		// If this bucket is the rightmost bucket, and its full
		// we need to split it and create a new bucket
		if bucketID == len(rt.Buckets)-1 {
			rt.nextBucket()
		} else {
			// If the bucket cant split kick out least active node
			bucket.PopBack()
		}
	}
}

// NearestPeers returns a list of the 'count' closest peers to the given ID
func (rt *RoutingTable) NearestPeers(id peer.ID, count int) []peer.ID {
	cpl := commonPrefixLen(id, rt.local)

	rt.mutex.RLock()

	// Get bucket at cpl index or last bucket
	var bucket *Bucket
	if cpl >= len(rt.Buckets) {
		cpl = len(rt.Buckets) - 1
	}
	bucket = rt.Buckets[cpl]

	peerArr := make(peerSorterArr, 0, count)
	peerArr = copyPeersFromList(id, peerArr, bucket.list)
	if len(peerArr) < count {
		// In the case of an unusual split, one bucket may be short or empty.
		// if this happens, search both surrounding buckets for nearby peers
		if cpl > 0 {
			plist := rt.Buckets[cpl-1].list
			peerArr = copyPeersFromList(id, peerArr, plist)
		}

		if cpl < len(rt.Buckets)-1 {
			plist := rt.Buckets[cpl+1].list
			peerArr = copyPeersFromList(id, peerArr, plist)
		}
	}
	rt.mutex.RUnlock()

	// Sort by distance to local peer
	sort.Sort(peerArr)

	if count < len(peerArr) {
		peerArr = peerArr[:count]
	}

	out := make([]peer.ID, 0, len(peerArr))
	for _, p := range peerArr {
		out = append(out, p.p)
	}

	return out
}

// Size returns the total number of peers in the routing table
func (rt *RoutingTable) Size() int {
	var tot int
	rt.mutex.RLock()
	for _, buck := range rt.Buckets {
		tot += buck.Len()
	}
	rt.mutex.RUnlock()
	return tot
}

// generate new bucket
func (rt *RoutingTable) nextBucket() {
	bucket := rt.Buckets[len(rt.Buckets)-1]
	newBucket := bucket.Split(len(rt.Buckets)-1, rt.local)
	rt.Buckets = append(rt.Buckets, newBucket)
	if newBucket.Len() > rt.bucketsize {
		rt.nextBucket()
	}

	// If all elements were on left side of split...
	if bucket.Len() > rt.bucketsize {
		bucket.PopBack()
	}
}
