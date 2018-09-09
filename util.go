package deckstring

import (
	"encoding/binary"
	"errors"
)

type buffer []byte

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
	if count < 0 {
		return 0, errors.New("invalid varint")
	}
	*b = (*b)[count:]
	return value, nil
}
