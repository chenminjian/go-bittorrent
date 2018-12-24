package metric

import "sync"

type impl struct {
	pingNum         int
	findNodeNum     int
	getPeersNum     int
	announcePeerNum int

	blockNodeNum int

	mutex sync.RWMutex
}

func New() Reporter {
	return &impl{}
}

func (im *impl) PingNum() int {
	im.mutex.RLock()
	defer im.mutex.RUnlock()

	return im.pingNum
}

func (im *impl) FindNodeNum() int {
	im.mutex.RLock()
	defer im.mutex.RUnlock()

	return im.findNodeNum
}

func (im *impl) GetPeersNum() int {
	im.mutex.RLock()
	defer im.mutex.RUnlock()

	return im.getPeersNum
}

func (im *impl) AnnouncePeerNum() int {
	im.mutex.RLock()
	defer im.mutex.RUnlock()

	return im.announcePeerNum
}

func (im *impl) PingInc() {
	im.mutex.Lock()
	defer im.mutex.Unlock()

	im.pingNum++
}

func (im *impl) FindNodeInc() {
	im.mutex.Lock()
	defer im.mutex.Unlock()

	im.findNodeNum++
}

func (im *impl) GetPeersInc() {
	im.mutex.Lock()
	defer im.mutex.Unlock()

	im.getPeersNum++
}

func (im *impl) AnnouncePeerInc() {
	im.mutex.Lock()
	defer im.mutex.Unlock()

	im.announcePeerNum++
}

func (im *impl) BlockNodeNum() int {
	im.mutex.RLock()
	defer im.mutex.RUnlock()

	return im.blockNodeNum
}

func (im *impl) SetBlockNodeNum(num int) {
	im.mutex.Lock()
	defer im.mutex.Unlock()

	im.blockNodeNum = num
}
