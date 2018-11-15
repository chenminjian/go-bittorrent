package basic

import "github.com/chenminjian/go-bittorrent/p2p/peer"

type BasicHost struct {
	id   peer.ID
	port int
}

func New() *BasicHost {
	h := &BasicHost{}
	return h
}

func (h *BasicHost) ID() peer.ID {
	return ""
}
