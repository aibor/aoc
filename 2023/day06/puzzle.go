package main

import (
	"strconv"

	"github.com/aibor/aoc/goutils"
)

func part1(input string) string {
	result := 1

	races := parseRaces(input)
	for _, r := range races {
		// Detect the minimum hold time required to beat the record. Since the
		// distance peaks at the median of race time, the number of options is
		// 2 times the difference between the minimum holdTime and median race
		// time.
		holdTime := 1
		for r.distance(holdTime) <= r.record {
			holdTime++
		}
		result *= r.time + 1 - 2*holdTime
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

func parseRaces(input string) []race {
	var races []race
	lines := goutils.SplitInput(input)
	timeIter := goutils.NewStringFieldsIterator(lines[0])
	distIter := goutils.NewStringFieldsIterator(lines[1])
	timeIter.Next()
	distIter.Next()
	for timeIter.Next() {
		distIter.Next()
		races = append(races, race{timeIter.MustBeInt(), distIter.MustBeInt()})
	}
	return races
}

type race struct {
	time   int
	record int
}

func (r *race) distance(holdTime int) int {
	if holdTime >= r.time {
		return 0
	}
	return holdTime * (r.time - holdTime)
}
