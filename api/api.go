package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/chenminjian/go-bittorrent/core"
	"github.com/gin-gonic/gin"
)

type Api struct {
	node   *core.BTNode
	router *gin.Engine
}

func New(node *core.BTNode) *Api {
	router := gin.New()
	router.Use(gin.Recovery())

	api := &Api{
		node:   node,
		router: router,
	}
	api.init()

	return api
}

func (api *Api) init() {
	api.router.GET("metric", api.Metric)
}

func (api *Api) Start() error {
	hs := &http.Server{
		Addr:           fmt.Sprintf(":%d", 9527),
		Handler:        api.router,
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	if err := hs.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
