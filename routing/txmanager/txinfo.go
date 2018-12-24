package txmanager

import (
	"errors"
	"time"

	"github.com/chenminjian/go-bittorrent/routing/peerstore"
)

type TxInfo struct {
	peerstore.PeerInfo
	Message map[string]interface{}
	Curr    time.Time
}

var txNotFound = errors.New("tx not found")
