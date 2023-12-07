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
	var result int

	for _, line := range goutils.SplitInput(input) {
		_ = line
	}

	return strconv.Itoa(result)
}

type scratchCard struct {
	winning []int
	numbers []int
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
