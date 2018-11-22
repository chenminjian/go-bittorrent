package config

import "github.com/chenminjian/go-bittorrent/p2p/peer"

type Config struct {
	Pid  peer.ID
	Port int
}
type Option func(cfg *Config) error
