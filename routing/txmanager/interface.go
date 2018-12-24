package txmanager

type GCHandler func(info *TxInfo)

type TxManager interface {
	UniqueID() string

	Get(key string) (*TxInfo, error)

	Set(key string, val *TxInfo)

	Del(key string)

	SetGCHandler(handler GCHandler)
}
