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

func Encode(deck Deck) (string, error) {

	b := new(buffer)

	buff, err := base64.URLEncoding.DecodeString("GU==")
	if err != nil {
		return "", err
	}

	b.append(buff)

	b.appendVarint(deck.Version)

	god, ok := nameToGod[strings.ToLower(deck.God)]

	if !ok {
		return "", errors.New("unrecognised god")
	}

	b.appendVarint(god)

	encodeCards(b, deck.Protos)

	return base64.URLEncoding.EncodeToString(*b), nil
}

func encodeCards(b *buffer, protos []int64) {

	arrays := collectCards(protos)

	for k, v := range arrays {
		b.appendVarint(k)
		b.appendVarint(int64(len(v)))
		for _, proto := range v {
			b.appendVarint(proto)
		}
	}
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
