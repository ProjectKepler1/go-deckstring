package deckstring

import (
	"encoding/binary"
	"errors"
)

type buffer []byte

func newBuffer(data []byte) *buffer {
	b := new(buffer)
	b.append(data)
	return b
}

func (b *buffer) append(data []byte) {
	*b = append(*b, data...)
}

func (b *buffer) appendVarint(id int64) {
	added := make([]byte, 10)
	count := binary.PutVarint(added, id)
	b.append(added[:count])
}

func (b *buffer) getVarint() (int64, error) {
	value, count := binary.Varint(*b)
	if 0 >= count {
		return 0, errors.New("invalid varint")
	}
	*b = (*b)[count:]
	return value, nil
}

func (b *buffer) len() int {
	return len(*b)
}
