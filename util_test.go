package deckstring

import "testing"

func TestAppendVarintEmpty(t *testing.T) {

	buffer := make([]byte, 0)

	buffer = appendVarint(buffer, 1)

	if len(buffer) != 1 {
		t.Error(buffer)
	}
}

func TestAppendVarintFull(t *testing.T) {

	buffer := make([]byte, 0)

	buffer = appendVarint(buffer, 1)

	if len(buffer) != 1 {
		t.Error(buffer)
	}

	buffer = appendVarint(buffer, 2)

	if len(buffer) != 2 {
		t.Error(buffer)
	}
}

func TestGetVarintSingle(t *testing.T) {

	buffer := make([]byte, 0)

	buffer = appendVarint(buffer, 1)

	value, buffer, err := getVarint(buffer)

	if err != nil {
		t.Error(err)
	}

	if len(buffer) != 0 {
		t.Errorf("wrong buffer length: %d", len(buffer))
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

	if len(*b) != 0 {
		t.Errorf("wrong buffer length: %d", len(*b))
	}

	if value != 2 {
		t.Errorf("wrong value: %d", value)
	}
}
