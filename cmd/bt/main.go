package main

import (
	"fmt"
	"github.com/chenminjian/go-bittorrent/core"
)

func main() {
	_, err := core.NewNode()
	if err != nil {
		panic(err)
	}

	fmt.Println("before select")
	select {}
}
