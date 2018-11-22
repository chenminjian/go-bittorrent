package net

type Network interface {
	Listen(port int) error
}
