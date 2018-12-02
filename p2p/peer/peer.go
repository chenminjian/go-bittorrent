package peer

import (
	"github.com/mr-tron/base58"
)

type ID string

func (id ID) String() string {
	b := []byte(id)
	return string(b)
}

func (id ID) Pretty() string {
	chk := base58.Encode([]byte(id))
	return chk
}
