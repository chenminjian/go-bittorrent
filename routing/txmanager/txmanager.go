package txmanager

import (
	"encoding/binary"
	"sync"
)

type txManager struct {
	idCount uint32
	txMap   map[string]*TxInfo
	mutex   sync.RWMutex
}

func New() TxManager {
	return &txManager{
		txMap: make(map[string]*TxInfo),
	}
}

func (tm *txManager) UniqueID() string {
	tm.mutex.Lock()
	defer tm.mutex.RLock()

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
