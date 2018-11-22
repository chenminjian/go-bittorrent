package main

import (
	"fmt"

	"github.com/chenminjian/go-bittorrent/config"
	"github.com/chenminjian/go-bittorrent/core"
)

func main() {
	conf := &config.Config{Pid: "", Port: 9527}
	_, err := core.NewNode(conf)
	if err != nil {
		panic(err)
	}

	fmt.Println("before select")
	select {}
}
