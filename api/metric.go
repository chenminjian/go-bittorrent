package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Metric struct {
	Receive Receive `json:"receive"`
	Block   Block   `json:"block"`
}

type Receive struct {
	Ping         int `json:"ping"`
	FindNode     int `json:"find_node"`
	GetPEERS     int `json:"get_peers"`
	AnnouncePeer int `json:"announce_peer"`
}

type Block struct {
	Size int `json:"size"`
}

func (api *Api) Metric(c *gin.Context) {
	m := Metric{
		Receive: Receive{
			Ping:         api.node.Reporter.PingNum(),
			FindNode:     api.node.Reporter.FindNodeNum(),
			GetPEERS:     api.node.Reporter.GetPeersNum(),
			AnnouncePeer: api.node.Reporter.AnnouncePeerNum(),
		},
		Block: Block{
			Size: api.node.Reporter.BlockNodeNum(),
		},
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "",
		"data":    m,
	})
}
