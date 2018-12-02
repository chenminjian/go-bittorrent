package peer

import (
	"math/rand"

	"github.com/mr-tron/base58"
)

func RandomID() ID {
	buff := make([]byte, 20)
	rand.Read(buff)
	return ID(buff)
}

func IDB58Decode(s string) (ID, error) {
	buf, err := base58.Decode(s)
	if err != nil {
		return ID(""), err
	}
	return ID(buf), nil
}
