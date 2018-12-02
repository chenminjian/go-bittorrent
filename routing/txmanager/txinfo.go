package txmanager

import (
	"errors"

	"github.com/chenminjian/go-bittorrent/routing/peerstore"
)

type TxInfo struct {
	peerstore.PeerInfo
	Message map[string]interface{}
}

var txNotFound = errors.New("tx not found")
