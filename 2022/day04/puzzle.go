package main

import (
	"strconv"

	"github.com/aibor/aoc/goutils"
)

var (
	exampleResult1 = "2"
	exampleResult2 = "4"

	result1 = "582"
	result2 = "893"
)

func part1(input string) string {
	var result int

	var s [4]int
	for _, line := range goutils.SplitInput(input) {
		parseSections(line, &s)
		if s[0] <= s[2] && s[1] >= s[3] || s[0] >= s[2] && s[1] <= s[3] {
			result++
		}
	}

	return strconv.Itoa(result)
}

func part2(input string) string {
	var result int

	var s [4]int
	for _, line := range goutils.SplitInput(input) {
		parseSections(line, &s)
		if s[1] < s[2] || s[3] < s[0] {
			continue
		}
		result++
	}

	return strconv.Itoa(result)
}

func parseSections(line string, sections *[4]int) {
	// Sscanf works but quite slow, lots of allocs
	//var s1s, s1e, s2s, s2e int
	//fmt.Sscanf(line, "%d-%d,%d-%d", &s1s, &s1e, &s2s, &s2e)
	var prev, i int
	for idx, r := range line {
		if r == '-' || r == ',' {
			sections[i], _ = strconv.Atoi(line[prev:idx])
			prev = idx + 1
			i++
		}
	}
	sections[i], _ = strconv.Atoi(line[prev:])
}
