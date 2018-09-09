package deckstring

import (
	"encoding/base64"
	"errors"
	"log"
	"sort"
	"strings"
)

// A Deck of prototype card ids
type Deck struct {
	Version int64   `json:"version"`
	God     string  `json:"god"`
	Protos  []int64 `json:"protos"`
}

type cardCollection struct {
	frequency int64
	protos    []int64
}

var (
	nameToGod = map[string]int64{
		"nature":    1,
		"war":       2,
		"death":     3,
		"light":     4,
		"magic":     5,
		"deception": 6,
	}

	godToName = map[int64]string{
		1: "nature",
		2: "war",
		3: "death",
		4: "light",
		5: "magic",
		6: "deception",
	}
)

// Encode a deck into a deckstring
func Encode(deck Deck) (string, error) {

	b := new(buffer)

	b.appendVarint(deck.Version)

	god, ok := nameToGod[strings.ToLower(deck.God)]

	if !ok {
		return "", errors.New("unrecognised god")
	}

	b.appendVarint(god)

	ccs := collectCards(deck.Protos)

	for _, cc := range ccs {
		b.appendVarint(cc.frequency)
		b.appendVarint(int64(len(cc.protos)))
		for _, proto := range cc.protos {
			b.appendVarint(proto)
		}
	}

	return base64.URLEncoding.EncodeToString(*b), nil
}

func appendPrefix(b *buffer) {

	buff, err := base64.URLEncoding.DecodeString("GU==")

	if err != nil {
		log.Fatal(err)
	}

	b.append(buff)
}

// gather cards into an ordered list of frequency, protos
// the same list of protos (in any order) should always produce the same deck string
func collectCards(protos []int64) []cardCollection {

	// count the number of times each card appears in the list
	counts := make(map[int64]int64)

	for _, proto := range protos {
		counts[proto]++
	}

	// create arrays of cards by frequency
	arrays := make(map[int64][]int64)
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

	god, err := b.getVarint()
	if err != nil {
		return nil, err
	}

	godName, ok := godToName[god]
	if !ok {
		return nil, errors.New("invalid god")
	}

	protos := make([]int64, 0)
	for b.len() > 0 {

		num, err := b.getVarint()
		if err != nil {
			return nil, err
		}

		len, err := b.getVarint()
		if err != nil {
			return nil, err
		}

		var proto int64

		for i := int64(0); i < len; i++ {

			proto, err = b.getVarint()
			if err != nil {
				return nil, err
			}

			for j := int64(0); j < num; j++ {
				protos = append(protos, proto)
			}

		}
	}

	pd := Deck{
		Version: version,
		God:     godName,
		Protos:  protos,
	}

	return &pd, nil
}
