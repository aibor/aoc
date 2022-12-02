package main

import (
	"fmt"

	"github.com/aibor/aoc/goutils"
)

func main() {
	input, _ := goutils.ReadInput()

	fmt.Println("Part 1:", part01(input))
	fmt.Println("Part 2:", part02(input))
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

func part01(input []string) string {
	var score int
	s := scores()

	for _, line := range input {
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

func part02(input []string) string {
	var score int

	s := scores()

	for _, line := range input {
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
