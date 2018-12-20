package krpc

import "github.com/chenminjian/go-bittorrent/p2p/peer"

const (
	Message_FIND_NODE     = "find_node"
	Message_PING          = "ping"
	Message_GET_PEERS     = "get_peers"
	Message_ANNOUNCE_PEER = "announce_peer"
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

func NewPingRespMessage(localID peer.ID, txID string) map[string]interface{} {
	r := map[string]interface{}{
		"id": localID.String(),
	}

	return map[string]interface{}{
		"t": txID,
		"y": "r",
		"r": r,
	}
}
