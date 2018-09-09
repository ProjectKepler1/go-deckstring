package deckstring

import (
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

const maxVarintBytes = 10 // maximum length of a varint

func (b *buffer) appendVarint(x uint64) {
	var buf [maxVarintBytes]byte
	var n int
	for n = 0; x > 127; n++ {
		buf[n] = 0x80 | uint8(x&0x7F)
		x >>= 7
	}
	buf[n] = uint8(x)
	n++
	b.append(buf[0:n])
}

func (b *buffer) getVarint() (x uint64, err error) {
	n := 0
	for shift := uint(0); shift < 64; shift += 7 {
		if n >= b.len() {
			return 0, errors.New("invalid varint")
		}
		a := uint64((*b)[n])
		n++
		x |= (a & 0x7F) << shift
		if (a & 0x80) == 0 {
			*b = (*b)[n:]
			return x, nil
		}
	}

	// The number is too large to represent in a 64-bit value.
	return 0, errors.New("invalid varint")
}

func (b *buffer) len() int {
	return len(*b)
}
