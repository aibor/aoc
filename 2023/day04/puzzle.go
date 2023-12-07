package main

import (
	"slices"
	"strconv"

	"github.com/aibor/aoc/goutils"
)

func part1(input string) string {
	var result int

	for _, line := range goutils.SplitInput(input) {
		sc := parseCard(line)
		if numbers := sc.winningNumbers(); numbers > 0 {
			result += 1 << (numbers - 1)
		}
	}

	return strconv.Itoa(result)
}

func part2(input string) string {
	var (
		result int
		cards  []scratchCard
	)

	for _, line := range goutils.SplitInput(input) {
		cards = append(cards, parseCard(line))
	}

	for idx, sc := range cards {
		for i := 1; i <= sc.winningNumbers(); i++ {
			cards[idx+i].copies += 1 + sc.copies
		}
		result += 1 + sc.copies
	}

	return strconv.Itoa(result)
}

type scratchCard struct {
	winning []int
	numbers []int
	copies  int
}

func parseCard(card string) scratchCard {
	var sc scratchCard
	i := goutils.NewStringFieldsIterator(card)
	i.Skip(1)

	for i.Next() && i.Value() != "|" {
		sc.winning = append(sc.winning, i.MustBeInt())
	}
	for i.Next() {
		sc.numbers = append(sc.numbers, i.MustBeInt())
	}

	return sc
}

func (sc *scratchCard) winningNumbers() int {
	var found int
	for _, number := range sc.winning {
		if slices.Contains(sc.numbers, number) {
			found++
		}
	}
	return found
}
