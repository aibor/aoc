package main

import (
	"strconv"

	"github.com/aibor/aoc/goutils"
)

var (
	exampleResult1 = "21"
	exampleResult2 = "8"

	result1 = "1688"
	result2 = "410400"
)

func part1(input string) string {
	var result int

	lines := goutils.SplitInput(input)
	rows := len(lines)
	cols := len(lines[0])
	// edge visible
	result += rows*2 + (cols-2)*2

	for i, l := range lines[1 : rows-1] {
		for j, r := range l[1 : cols-1] {
			if visibleX(r, l, j+1, cols, -1) || visibleX(r, l, j+1, cols, 1) ||
				visibleY(r, lines, j+1, i+1, rows, -1) || visibleY(r, lines, j+1, i+1, rows, 1) {
				result += 1
			}
		}
	}

	return strconv.Itoa(result)
}

func part2(input string) string {
	var result, dist int

	lines := goutils.SplitInput(input)
	rows := len(lines)
	cols := len(lines[0])

	for i, l := range lines {
		for j, r := range l {
			dist = distanceX(r, l, j, cols, -1) * distanceX(r, l, j, cols, 1) *
				distanceY(r, lines, j, i, rows, -1) * distanceY(r, lines, j, i, rows, 1)
			if dist > result {
				result = dist
			}
		}
	}

	return strconv.Itoa(result)
}

func visibleX(r rune, line string, start, max, dir int) bool {
	for x := start + dir; x >= 0 && x < max; x += dir {
		if line[x] >= byte(r) {
			return false
		}
	}
	return true
}

func visibleY(r rune, lines []string, x int, start, max, dir int) bool {
	for y := start + dir; y >= 0 && y < max; y += dir {
		if lines[y][x] >= byte(r) {
			return false
		}
	}
	return true
}

func distanceX(r rune, line string, start, max, dir int) int {
	var dist int
	for x := start + dir; x >= 0 && x < max; x += dir {
		dist++
		if line[x] >= byte(r) {
			break
		}
	}
	return dist
}

func distanceY(r rune, lines []string, x int, start, max, dir int) int {
	var dist int
	for y := start + dir; y >= 0 && y < max; y += dir {
		dist++
		if lines[y][x] >= byte(r) {
			break
		}
	}
	return dist
}
