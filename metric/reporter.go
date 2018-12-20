package metric

type Reporter interface {
	PingNum() int

	FindNodeNum() int

	GetPeersNum() int

	AnnouncePeerNum() int

	PingInc()

	FindNodeInc()

	GetPeersInc()

	AnnouncePeerInc()
}
