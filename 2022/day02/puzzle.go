package main

import (
	"fmt"

	"github.com/aibor/aoc/goutils"
)

var (
	exampleResult1 = "15"
	exampleResult2 = "12"

	result1 = "13446"
	result2 = "13509"
)

func part1(input string) string {
	var score int
	s := scores()

	for _, line := range goutils.SplitInput(input) {
		op := s[line[0]]
		me := s[line[2]]
		score += me
		switch me - op {
		case 0:
			score += 3
		case 1, -2:
			score += 6
		}
	}

	return fmt.Sprintf("%d", score)
}

func part2(input string) string {
	var score int

	s := scores()

	for _, line := range goutils.SplitInput(input) {
		var me int
		op := s[line[0]]
		res := rune(line[2])
		switch res {
		case 'X':
			me = op - 1
		case 'Y':
			score += 3
			me = op
		case 'Z':
			score += 6
			me = op + 1
		}
		switch me {
		case 0:
			score += 3
		case 4:
			score += 1
		default:
			score += me
		}
	}

	return fmt.Sprintf("%d", score)
}

func scores() map[byte]int {
	return map[byte]int{
		byte('A'): 1,
		byte('B'): 2,
		byte('C'): 3,
		byte('X'): 1,
		byte('Y'): 2,
		byte('Z'): 3,
	}
}
