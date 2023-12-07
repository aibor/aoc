package main

import (
	"sort"
	"strconv"

	"github.com/aibor/aoc/goutils"
)

func part1(input string) string {
	var result int

	hands := parseHands(input, false)

	sort.Slice(hands, func(i, j int) bool {
		return hands[j].strongerThan(hands[i], false)
	})

	for idx, hand := range hands {
		result += hand.bid * (idx + 1)
	}

	return strconv.Itoa(result)
}

func part2(input string) string {
	var result int

	hands := parseHands(input, true)

	sort.Slice(hands, func(i, j int) bool {
		return hands[j].strongerThan(hands[i], true)
	})

	for idx, hand := range hands {
		result += hand.bid * (idx + 1)
	}

	return strconv.Itoa(result)
}

func parseHands(input string, withJoker bool) []*camelHand {
	var hands []*camelHand
	i := goutils.NewStringFieldsIterator(input)
	for i.Next() {
		hand := i.Value()
		i.Next()
		hands = append(hands, &camelHand{
			hand: hand,
			typ:  classify(hand, withJoker),
			bid:  i.MustBeInt()},
		)
	}

	return hands
}

const joker = 'J'

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

func getCardRank(card rune, withJoker bool) int {
	if withJoker && card == joker {
		return -1
	}
	return cardRank[card]
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

func classify(hand string, withJoker bool) handType {
	// Get count of distinct cards labels of the given hand.
	cardLabels := make(map[rune]int, 5)
	for _, r := range hand {
		cardLabels[r]++
	}

	if withJoker {
		jokers, exist := cardLabels[joker]
		delete(cardLabels, joker)
		if exist {
			var (
				maxLabel rune
				maxCount int
			)
			for label, count := range cardLabels {
				if maxLabel == 0 || count > maxCount {
					maxLabel = label
					maxCount = count
				}
			}
			cardLabels[maxLabel] += jokers
		}
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

func (h *camelHand) strongerThan(other *camelHand, withJoker bool) bool {
	switch {
	case h.typ > other.typ:
		return true
	case h.typ < other.typ:
		return false
	}
	for idx, r := range h.hand {
		ca := getCardRank(r, withJoker)
		cb := getCardRank([]rune(other.hand)[idx], withJoker)
		switch {
		case ca > cb:
			return true
		case ca < cb:
			return false
		}
	}
	return true
}
