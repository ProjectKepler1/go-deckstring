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
		Protos:  []uint64{1, 2, 3, 4, 5},
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
		Protos:  []uint64{3, 5, 4, 2, 1},
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
		Protos:  []uint64{1, 2, 3, 4, 5, 2, 3},
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

	expected := []uint64{1, 4, 5, 2, 2, 3, 3}

	for i, p := range expected {
		if p != after.Protos[i] {
			t.Errorf("wrong proto %d, expected %d, was %d", i, p, after.Protos[i])
		}
	}
}

func TestCollectCards(t *testing.T) {

	protos := []uint64{1, 2, 3, 4, 5, 2, 3}

	ccs := collectCards(protos)

	if len(ccs) != 2 {
		t.Fatalf("wrong frequency: %d", len(ccs))
	}

	if ccs[0].frequency != 1 {
		t.Fatalf("wrong frequency: %d", ccs[0].frequency)
	}

	if len(ccs[0].protos) != 3 {
		t.Fatalf("wrong proto length: %d", len(ccs[0].protos))
	}

	if ccs[1].frequency != 2 {
		t.Fatalf("wrong frequency: %d", ccs[1].frequency)
	}

	if len(ccs[1].protos) != 2 {
		t.Fatalf("wrong proto length: %d", len(ccs[1].protos))
	}
}

func TestFullDeck(t *testing.T) {

	protos := []uint64{
		290, 17, 201, 201, 80, 80, 93, 93, 64, 64, 185, 185, 55, 55, 97, 331, 281, 281, 252, 252, 330,
		330, 280, 202, 202, 265, 265, 37, 94, 94,
	}

	ds, err := Encode(Deck{Version: 1, God: "deception", Protos: protos})

	if err != nil {
		t.Error(err)
	}

	_, err = Decode(ds)

	if err != nil {
		t.Error(err)
	}

}
