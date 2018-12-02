package peer

import (
	"fmt"
	"testing"
)

func TestRandomID(t *testing.T) {
	id := RandomID()
	fmt.Println(id.Pretty())
}
