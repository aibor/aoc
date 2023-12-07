package main

import (
	"sort"
	"strconv"

	"github.com/aibor/aoc/goutils"
)

func part1(input string) string {
	var (
		result int
		hands  []*camelHand
	)

	i := goutils.NewStringFieldsIterator(input)
	for i.Next() {
		h := i.Value()
		i.Next()
		hands = append(hands, newCamelHand(h, i.MustBeInt()))
	}

	sort.Slice(hands, func(i, j int) bool {
		return hands[j].strongerThan(hands[i])
	})

	for idx, hand := range hands {
		result += hand.bid * (idx + 1)
	}

	return strconv.Itoa(result)
}

func part2(input string) string {
	var result int

	for _, line := range goutils.SplitInput(input) {
		_ = line
	}

	return strconv.Itoa(result)
}

var cardRank = map[rune]int{
	'2': 0,
	'3': 1,
	'4': 2,
	'5': 3,
	'6': 4,
	'7': 5,
	'8': 6,
	'9': 7,
	'T': 8,
	'J': 9,
	'Q': 10,
	'K': 11,
	'A': 12,
}

type handType int

const (
	handTypeHighCard handType = iota
	handTypeOnePair
	handTypeTwoPair
	handTypeThreeOfAKind
	handTypeFullHouse
	handTypeFourOfAKind
	handTypeFiveOfAKind
)

type camelHand struct {
	hand string
	typ  handType
	bid  int
}

func newCamelHand(hand string, bid int) *camelHand {
	return &camelHand{
		hand: hand,
		typ:  classify(hand),
		bid:  bid,
	}
}

func classify(hand string) handType {
	// Get count of distinct cards labels of the given hand.
	cardLabels := make(map[rune]int, 5)
	for _, r := range hand {
		cardLabels[r]++
	}

	// Classify hand based on the combinations of label counts.
	switch len(cardLabels) {
	case 1:
		return handTypeFiveOfAKind
	case 2:
		for _, count := range cardLabels {
			if count > 3 || count < 2 {
				return handTypeFourOfAKind
			}
			break
		}
		return handTypeFullHouse
	case 3:
		for _, count := range cardLabels {
			if count > 2 {
				return handTypeThreeOfAKind
			}
		}
		return handTypeTwoPair
	case 4:
		return handTypeOnePair
	default:
		return handTypeHighCard
	}
}

func (h *camelHand) strongerThan(other *camelHand) bool {
	switch {
	case h.typ > other.typ:
		return true
	case h.typ < other.typ:
		return false
	}
	for idx, r := range h.hand {
		ca := cardRank[r]
		cb := cardRank[[]rune(other.hand)[idx]]
		switch {
		case ca > cb:
			return true
		case ca < cb:
			return false
		}
	}
	return true
}
