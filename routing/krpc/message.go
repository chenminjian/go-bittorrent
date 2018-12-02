package krpc

import "github.com/chenminjian/go-bittorrent/p2p/peer"

const (
	Message_FIND_NODE = "find_node"
)

func NewFindNodeMessage(id peer.ID, target peer.ID, txID string) map[string]interface{} {
	a := map[string]interface{}{
		"id":     id.String(),
		"target": target.String(),
	}

	return map[string]interface{}{
		"t": txID,
		"y": "q",
		"q": Message_FIND_NODE,
		"a": a,
	}
}
