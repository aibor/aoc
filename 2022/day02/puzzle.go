package main

import (
	"strconv"

	"github.com/aibor/aoc/goutils"
)

var (
	exampleResult1 = "15"
	exampleResult2 = "12"

	result1 = "13446"
	result2 = "13509"
)

func part1(input string) string {
	var result int

	for _, line := range goutils.SplitInput(input) {
		op := int(line[0] - 'A' + 1)
		me := int(line[2] - 'X' + 1)
		result += me
		switch me - op {
		case 0:
			result += 3
		case 1, -2:
			result += 6
		}
	}

	return strconv.Itoa(result)
}

func part2(input string) string {
	var result int

	for _, line := range goutils.SplitInput(input) {
		op := int(line[0] - 'A' + 1)
		// result value -- 0: lose, 1: draw, 2: win
		res := int(line[2] - 'X')
		// add round outcome
		result += res * 3
		// add shape value
		result += (op+1+res)%3 + 1
	}

	return strconv.Itoa(result)
}
