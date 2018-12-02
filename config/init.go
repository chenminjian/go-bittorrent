package config

func Default() *Config {
	return &Config{
		Pid:  "2A4XMxQwz7oPW9wvEe8t5PhyQow2",
		Port: 9527,
		Bootstrap: []string{
			"router.bittorrent.com:6881",
			"router.utorrent.com:6881",
			"dht.transmissionbt.com:6881",
		},
	}
}
