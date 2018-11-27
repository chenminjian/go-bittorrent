package peer

type ID string

func (id ID) String() string {
	b := []byte(id)
	return string(b)
}
