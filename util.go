package deckstring

import (
	"encoding/binary"
	"errors"
)

func appendVarint(buffer []byte, id int64) []byte {
	added := make([]byte, 10)
	count := binary.PutVarint(added, id)
	return append(buffer, added[:count]...)
}

func getVarint(buffer []byte) (int64, []byte, error) {
	value, count := binary.Varint(buffer)
	if count < 0 {
		return 0, buffer, errors.New("invalid varint")
	}
	return value, buffer[count:], nil
}
