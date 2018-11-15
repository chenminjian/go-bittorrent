package config

type Config struct {
	port int
}
type Option func(cfg *Config) error
