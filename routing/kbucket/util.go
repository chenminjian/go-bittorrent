package kbucket

import (
	"math/bits"

	"github.com/chenminjian/go-bittorrent/p2p/peer"
)

func commonPrefixLen(a, b peer.ID) int {
	return zeroPrefixLen(xor([]byte(a), []byte(b)))
}

// zeroPrefixLen returns the number of consecutive zeroes in a byte slice.
func zeroPrefixLen(id []byte) int {
	for i, b := range id {
		if b != 0 {
			return i*8 + bits.LeadingZeros8(uint8(b))
		}
	}
	return len(id) * 8
}

// xor takes two byte slices, XORs them together, returns the resulting slice.
func xor(a, b []byte) []byte {
	c := make([]byte, len(a))
	for i := 0; i < len(a); i++ {
		c[i] = a[i] ^ b[i]
	}
	return c
}
