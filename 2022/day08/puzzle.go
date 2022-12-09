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

	f := forest(goutils.SplitInput(input))
	size := f.size()

	// keep top values seen for early exit when processing lower trees
	var topy byte
	topx := make([]byte, size)

	// edge visible, no need to process
	result += size*2 + (size-2)*2

	for y, l := range f[1 : size-1] {
		y := y + 1
		topy = 0
		for x := range l[1 : size-1] {
			x := x + 1
			t := f[y][x]
			if t > topy && f.visible(x, y, -1, 0) ||
				f.visible(x, y, 1, 0) ||
				t > topx[x] && f.visible(x, y, 0, -1) ||
				f.visible(x, y, 0, 1) {
				result += 1
			}
			if t > topy {
				topy = t
			}
			if t > topx[x] {
				topx[x] = t
			}
		}
	}

	return strconv.Itoa(result)
}

func part2(input string) string {
	var result, dist int

	f := forest(goutils.SplitInput(input))
	size := f.size()
	// skip edge that is to close to the border to have a good view
	edge := size / 10

	for y, l := range f[edge : size-edge] {
		y := y + edge
		for x := range l[edge : size-edge] {
			x := x + edge
			dist = f.distance(x, y, -1, 0) *
				f.distance(x, y, 1, 0) *
				f.distance(x, y, 0, -1) *
				f.distance(x, y, 0, 1)
			if dist > result {
				result = dist
			}
		}
	}

	return strconv.Itoa(result)
}

type forest []string

func (f *forest) size() int {
	return len(*f)
}

func (f *forest) visible(x, y, dirx, diry int) bool {
	t := (*f)[y][x]
	for x, y := x+dirx, y+diry; x >= 0 && x < f.size() && y >= 0 && y < f.size(); x, y = x+dirx, y+diry {
		if (*f)[y][x] >= t {
			return false
		}
	}
	return true
}

func (f *forest) distance(x, y, dirx, diry int) int {
	var dist int
	t := (*f)[y][x]
	for x, y := x+dirx, y+diry; x >= 0 && x < f.size() && y >= 0 && y < f.size(); x, y = x+dirx, y+diry {
		dist++
		if (*f)[y][x] >= t {
			break
		}
	}
	return dist
}
