package txmanager

import (
	"encoding/binary"
	"sync"
	"time"
)

type txManager struct {
	idCount   uint32
	txMap     map[string]*TxInfo
	gcHandler GCHandler
	gcPeriod  time.Duration
	mutex     sync.RWMutex
}

func New() TxManager {
	mgr := &txManager{
		txMap:    make(map[string]*TxInfo),
		gcPeriod: 2 * time.Minute,
	}

	go mgr.gcWorker()

	return mgr
}

func (tm *txManager) UniqueID() string {
	tm.mutex.Lock()
	defer tm.mutex.Unlock()

	tm.idCount += 1
	buf := make([]byte, 4)
	binary.BigEndian.PutUint32(buf, tm.idCount)
	return string(buf)
}

func (tm *txManager) Get(key string) (*TxInfo, error) {
	tm.mutex.RLock()
	defer tm.mutex.RUnlock()

	val, ok := tm.txMap[key]
	if !ok {
		return nil, txNotFound
	}
	return val, nil
}

func (tm *txManager) Set(key string, val *TxInfo) {
	tm.mutex.Lock()
	defer tm.mutex.Unlock()

	tm.txMap[key] = val
}

func (tm *txManager) Del(key string) {
	tm.mutex.Lock()
	defer tm.mutex.Unlock()

	delete(tm.txMap, key)
}

func (tm *txManager) SetGCHandler(handler GCHandler) {
	tm.mutex.Lock()
	defer tm.mutex.Unlock()

	tm.gcHandler = handler
}

func (tm *txManager) gcWorker() {
	for {
		select {
		case <-time.After(time.Second * 30):
			tm.gc()
		}
	}
}

func (tm *txManager) gc() {
	tm.mutex.Lock()
	defer tm.mutex.Unlock()

	zero := time.Time{}

	curr := time.Now()

	for k, v := range tm.txMap {
		if v.Curr == zero {
			continue
		}

		if curr.Sub(v.Curr) > tm.gcPeriod {
			delete(tm.txMap, k)

			if tm.gcHandler != nil {
				tm.gcHandler(v)
			}
		}
	}
}
