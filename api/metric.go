package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Metric struct {
	Receive Receive `json:"receive"`
}

type Receive struct {
	Ping         int `json:"ping"`
	FindNode     int `json:"find_node"`
	GetPEERS     int `json:"get_peers"`
	AnnouncePeer int `json:"announce_peer"`
}

func (api *Api) Metric(c *gin.Context) {
	r := Receive{
		Ping:         api.node.Reporter.PingNum(),
		FindNode:     api.node.Reporter.FindNodeNum(),
		GetPEERS:     api.node.Reporter.GetPeersNum(),
		AnnouncePeer: api.node.Reporter.AnnouncePeerNum(),
	}

	m := Metric{
		Receive: r,
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "",
		"data":    m,
	})
}
