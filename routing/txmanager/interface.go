package txmanager

type TxManager interface {
	UniqueID() string

	Get(key string) (*TxInfo, error)

	Set(key string, val *TxInfo)

	Del(key string)
}
