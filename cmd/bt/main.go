package main

import (
	"github.com/chenminjian/go-bittorrent/api"
	"github.com/chenminjian/go-bittorrent/config"
	"github.com/chenminjian/go-bittorrent/core"
)

func main() {
	if err := execute(); err != nil {
		panic(err)
	}
}

func execute() error {
	conf := config.Default()
	node, err := core.NewNode(conf)
	if err != nil {
		return err
	}

	api := api.New(node)

	if err := api.Start(); err != nil {
		return err
	}

	return nil
}
