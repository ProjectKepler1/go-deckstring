package deckstring

import (
	"encoding/base64"
	"errors"
	"sort"
	"strings"
)

type Deck struct {
	Version int64   `json:"version"`
	God     string  `json:"god"`
	Protos  []int64 `json:"protos"`
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

	/*buff, err := base64.URLEncoding.DecodeString("GU==")
	if err != nil {
		return "", err
	}

	b.append(buff)

	log.Println(*b)*/

	b.appendVarint(deck.Version)

	god, ok := nameToGod[strings.ToLower(deck.God)]

	if !ok {
		return "", errors.New("unrecognised god")
	}

	b.appendVarint(god)

	arrays := collectCards(deck.Protos)

	for k, v := range arrays {
		b.appendVarint(k)
		b.appendVarint(int64(len(v)))
		for _, proto := range v {
			b.appendVarint(proto)
		}
	}

	return base64.URLEncoding.EncodeToString(*b), nil
}

func collectCards(protos []int64) map[int64][]int64 {

	counts := make(map[int64]int64)

	for _, proto := range protos {
		counts[proto]++
	}

	arrays := make(map[int64][]int64)

	for k, v := range counts {
		arrays[v] = append(arrays[v], k)
	}

	for _, v := range arrays {
		sort.Slice(v, func(i, j int) bool {
			return v[j] > v[i]
		})
	}
	return arrays
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
