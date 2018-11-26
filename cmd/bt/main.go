package main

import (
	"time"

	"github.com/chenminjian/go-bittorrent/config"
	"github.com/chenminjian/go-bittorrent/core"
)

func main() {
	conf := config.Default()
	_, err := core.NewNode(conf)
	if err != nil {
		panic(err)
	}

	for {
		time.Sleep(time.Minute)
	}
}
