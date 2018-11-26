package config

func Default() *Config {
	return &Config{
		Pid:  "",
		Port: 9527,
		Bootstrap: []string{
			"",
		},
	}
}
