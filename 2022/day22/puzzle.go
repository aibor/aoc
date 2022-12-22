package main

import (
	"strconv"
	"strings"

	"github.com/aibor/aoc/goutils"
)

var (
	exampleResult1 = "6032"
	exampleResult2 = "0"

	result1 = "196134"
	result2 = "0"
)

func part1(input string) string {
	var result int

	parts := strings.Split(strings.TrimRight(input, "\n"), "\n\n")
	ll := len(parts[0])
	var m [][]byte
	for _, l := range strings.Split(parts[0], "\n") {
		k := make([]byte, ll)
		copy(k, l)
		m = append(m, k)
	}
	p := pos{0, 0}
	for ; m[p.y][p.x] != '.'; p.x++ {
	}
	dir := 0
	insts := parts[1]
	c := map[int]byte{
		0: '>',
		1: 'v',
		2: '<',
		3: '^',
	}
	d := map[int]pos{
		0: {1, 0},
		1: {0, 1},
		2: {-1, 0},
		3: {0, -1},
	}
	for len(insts) > 0 {
		instEnd := 1
		switch insts[0] {
		case 'L':
			dir = (dir + 3) % 4
		case 'R':
			dir = (dir + 1) % 4
		default:
			instEnd = strings.IndexAny(insts, "LR")
			if instEnd == -1 {
				instEnd = len(insts)
			}
			dist := goutils.MustBeInt(insts[:instEnd])
			n := p
			for dist > 0 {
				m[p.y][p.x] = c[dir]
				n = n.add(d[dir])
				if n.y >= len(m) {
					n.y = 0
					for len(m[n.y]) < n.x {
						n.y++
					}
				} else if n.y < 0 {
					n.y = len(m) - 1
					for len(m[n.y]) < n.x {
						n.y--
					}
				} else if n.x >= len(m[n.y]) {
					n.x = 0
				} else if n.x < 0 {
					n.x = len(m[n.y]) - 1
				}
				if m[n.y][n.x] == '#' {
					break
				}
				if m[n.y][n.x] == '.' || m[n.y][n.x] == '>' || m[n.y][n.x] == '<' || m[n.y][n.x] == 'v' || m[n.y][n.x] == '^' {
					dist--
					p = n
				}
				n.add(d[dir])
			}
		}
		insts = insts[instEnd:]
	}

	result = 1000*(p.y+1) + 4*(p.x+1) + dir

	return strconv.Itoa(result)
}

func part2(input string) string {
	var result int

	for _, line := range goutils.SplitInput(input) {
		_ = line
	}

	return strconv.Itoa(result)
}

type pos struct {
	x, y int
}

func (p pos) add(q pos) pos {
	return pos{p.x + q.x, p.y + q.y}
}
