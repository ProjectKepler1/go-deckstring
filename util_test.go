package deckstring

import "testing"

func TestAppendVarintEmpty(t *testing.T) {

	buffer := new(buffer)

	buffer.appendVarint(1)

	if buffer.len() != 1 {
		t.Error(buffer)
	}
}

func TestAppendVarintFull(t *testing.T) {

	buffer := new(buffer)

	buffer.appendVarint(1)

	if buffer.len() != 1 {
		t.Error(buffer)
	}

	buffer.appendVarint(2)

	if buffer.len() != 2 {
		t.Error(buffer)
	}
}

func TestGetVarintSingle(t *testing.T) {

	buffer := new(buffer)

	buffer.appendVarint(1)

	value, err := buffer.getVarint()

	if err != nil {
		t.Error(err)
	}

	if buffer.len() != 0 {
		t.Errorf("wrong buffer length: %d", buffer.len())
	}

	if value != 1 {
		t.Errorf("wrong value: %d", value)
	}

}

func TestGetVarintMultiple(t *testing.T) {

	b := new(buffer)

	b.appendVarint(1)
	b.appendVarint(2)

	value, err := b.getVarint()

	if err != nil {
		t.Error(err)
	}

	if value != 1 {
		t.Errorf("wrong value: %d", value)
	}

	value, err = b.getVarint()

	if err != nil {
		t.Error(err)
	}

	if b.len() != 0 {
		t.Errorf("wrong buffer length: %d", b.len())
	}

	if value != 2 {
		t.Errorf("wrong value: %d", value)
	}
}
