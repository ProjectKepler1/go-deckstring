package deckstring

import (
	"encoding/base64"
	"errors"
	"sort"
	"strings"
)

// A Deck of prototype card ids
type Deck struct {
	Version uint64   `json:"version"`
	God     string   `json:"god"`
	Protos  []uint64 `json:"protos"`
}

type cardCollection struct {
	frequency uint64
	protos    []uint64
}

// Encode a deck into a deckstring
func Encode(deck Deck) (string, error) {

	b := new(buffer)

	if deck.Version != 1 {
		return "", errors.New("invalid version")
	}

	// only support version 1
	b.appendVarint(deck.Version)

	god, ok := nameToGod[strings.ToLower(deck.God)]
	if !ok {
		return "", errors.New("invalid god")
	}

	b.appendVarint(god)

	ccs := collectCards(deck.Protos)

	for _, cc := range ccs {
		b.appendVarint(cc.frequency)
		b.appendVarint(uint64(len(cc.protos)))
		for _, proto := range cc.protos {
			b.appendVarint(proto)
		}
	}

	return base64.URLEncoding.EncodeToString(*b), nil
}

// gather cards into an ordered list of frequency, protos
// the same list of protos (in any order) should always produce the same deck string
func collectCards(protos []uint64) []cardCollection {

	// count the number of times each card appears in the list
	counts := make(map[uint64]uint64)

	for _, proto := range protos {
		counts[proto]++
	}

	// create arrays of cards by frequency
	arrays := make(map[uint64][]uint64)
	for k, v := range counts {
		arrays[v] = append(arrays[v], k)
	}

	// sort the protos in ascending order
	for _, v := range arrays {
		sort.Slice(v, func(i, j int) bool {
			return v[j] > v[i]
		})
	}

	// turn those maps into arrays
	var ccs []cardCollection
	for k := range arrays {
		ccs = append(ccs, cardCollection{
			frequency: k,
			protos:    arrays[k],
		})
	}

	sort.Slice(ccs, func(i, j int) bool {
		return ccs[j].frequency > ccs[i].frequency
	})

	return ccs
}

var (
	nameToGod = map[string]uint64{
		"nature":    1,
		"war":       2,
		"death":     3,
		"light":     4,
		"magic":     5,
		"deception": 6,
	}
	godToName = map[uint64]string{
		1: "nature",
		2: "war",
		3: "death",
		4: "light",
		5: "magic",
		6: "deception",
	}
)

// Decode a deckstring into a deck
func Decode(data string) (*Deck, error) {

	buff, err := base64.URLEncoding.DecodeString(data)
	if err != nil {
		return nil, err
	}

	b := newBuffer(buff)

	version, err := b.getVarint()
	if err != nil {
		return nil, err
	}

	if version != 1 {
		return nil, errors.New("unsupported deck string version")
	}

	god, err := b.getVarint()
	if err != nil {
		return nil, err
	}

	godName, ok := godToName[god]
	if !ok {
		return nil, errors.New("unsupported god")
	}

	protos := make([]uint64, 0)
	for b.len() > 0 {

		num, err := b.getVarint()
		if err != nil {
			return nil, err
		}

		len, err := b.getVarint()
		if err != nil {
			return nil, err
		}

		var proto uint64

		for i := uint64(0); i < len; i++ {

			proto, err = b.getVarint()
			if err != nil {
				return nil, err
			}

			for j := uint64(0); j < num; j++ {
				protos = append(protos, proto)
			}

		}
	}

	pd := Deck{
		God:     godName,
		Version: version,
		Protos:  protos,
	}

	return &pd, nil
}
