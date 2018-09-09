package deckstring

import (
	"sort"
	"testing"
)

func TestDeckStringNoProtos(t *testing.T) {

	deck := Deck{
		Version: 2,
		God:     "war",
	}

	ds, err := Encode(deck)

	if err != nil {
		t.Error(err)
	}

	after, err := Decode(ds)

	if err != nil {
		t.Error(err)
	}

	if deck.Version != after.Version {
		t.Errorf("wrong version, expected: %d, was %d", deck.Version, after.Version)
	}

	if deck.God != after.God {
		t.Errorf("wrong god, expected: %s, was %s", deck.God, after.God)
	}
}

func TestDeckStringSingleProtosSorted(t *testing.T) {

	deck := Deck{
		Version: 2,
		God:     "war",
		Protos:  []int64{1, 2, 3, 4, 5},
	}

	ds, err := Encode(deck)

	if err != nil {
		t.Error(err)
	}

	after, err := Decode(ds)

	if err != nil {
		t.Error(err)
	}

	if deck.Version != after.Version {
		t.Errorf("wrong version, expected: %d, was %d", deck.Version, after.Version)
	}

	if deck.God != after.God {
		t.Errorf("wrong god, expected: %s, was %s", deck.God, after.God)
	}

	if len(deck.Protos) != len(after.Protos) {
		t.Fatalf("wrong proto length, expected: %d, was %d", len(deck.Protos), len(after.Protos))
	}

	sort.Slice(deck.Protos, func(i, j int) bool {
		return deck.Protos[j] > deck.Protos[i]
	})

	for i, p := range deck.Protos {
		if p != after.Protos[i] {
			t.Errorf("wrong proto %d, expected %d, was %d", i, p, after.Protos[i])
		}
	}
}

func TestDeckStringSingleProtosUnsorted(t *testing.T) {

	deck := Deck{
		Version: 2,
		God:     "war",
		Protos:  []int64{3, 5, 4, 2, 1},
	}

	ds, err := Encode(deck)

	if err != nil {
		t.Error(err)
	}

	after, err := Decode(ds)

	if err != nil {
		t.Error(err)
	}

	if deck.Version != after.Version {
		t.Errorf("wrong version, expected: %d, was %d", deck.Version, after.Version)
	}

	if deck.God != after.God {
		t.Errorf("wrong god, expected: %s, was %s", deck.God, after.God)
	}

	if len(deck.Protos) != len(after.Protos) {
		t.Fatalf("wrong proto length, expected: %d, was %d", len(deck.Protos), len(after.Protos))
	}

	sort.Slice(deck.Protos, func(i, j int) bool {
		return deck.Protos[j] > deck.Protos[i]
	})

	for i, p := range deck.Protos {
		if p != after.Protos[i] {
			t.Errorf("wrong proto %d, expected %d, was %d", i, p, after.Protos[i])
		}
	}
}

func TestDeckStringSingleAndDoubleProtos(t *testing.T) {

	deck := Deck{
		Version: 2,
		God:     "war",
		Protos:  []int64{1, 2, 3, 4, 5, 2, 3},
	}

	ds, err := Encode(deck)

	if err != nil {
		t.Error(err)
	}

	after, err := Decode(ds)

	if err != nil {
		t.Error(err)
	}

	if deck.Version != after.Version {
		t.Errorf("wrong version, expected: %d, was %d", deck.Version, after.Version)
	}

	if deck.God != after.God {
		t.Errorf("wrong god, expected: %s, was %s", deck.God, after.God)
	}

	if len(deck.Protos) != len(after.Protos) {
		t.Fatalf("wrong proto length, expected: %d, was %d", len(deck.Protos), len(after.Protos))
	}

	expected := []int64{1, 4, 5, 2, 2, 3, 3}

	for i, p := range expected {
		if p != after.Protos[i] {
			t.Errorf("wrong proto %d, expected %d, was %d", i, p, after.Protos[i])
		}
	}
}
