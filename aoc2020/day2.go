package aoc2020

import (
	"strconv"
	"strings"
)

type Foo struct {
	min    uint32
	max    uint32
	pivot  byte
	target string
	plen   uint32
}

func _calculatePivotLength(pivot byte, buf string) uint32 {
	var numMatches uint32 = 0

	for _, bchar := range buf {
		if byte(bchar) == pivot {
			numMatches++
		}
	}

	return numMatches
}

func _buildDictionaryLookup(data []string) []Foo {
	dl := make([]Foo, 0)

	for _, d := range data {
		splitted := strings.Split(d, " ")
		minMaxRange := strings.Split(splitted[0], "-")
		pivot := splitted[1][0]
		target := splitted[2]

		min, _ := strconv.Atoi(minMaxRange[0])
		max, _ := strconv.Atoi(minMaxRange[1])

		dl = append(dl, Foo{
			min:    uint32(min),
			max:    uint32(max),
			pivot:  pivot,
			target: target,
			plen:   _calculatePivotLength(pivot, target),
		})
	}

	return dl
}

func PasswordPhilosophy(data []string) uint32 {
	dl := _buildDictionaryLookup(data)

	var numMatches uint32 = 0

	for _, _dl := range dl {
		if _dl.plen >= _dl.min && _dl.plen <= _dl.max {
			numMatches++
		}
	}

	return numMatches
}
