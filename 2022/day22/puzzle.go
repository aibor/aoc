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

	b, insts := parseInput(input)
	l := pos{150, 200}
	if len(insts) < 100 {
		l.x, l.y = 16, 12
	}
	b.wrapFunc = func(n *pos) {
		n.x = (n.x + l.x) % l.x
		n.y = (n.y + l.y) % l.y
	}
	result = b.findPassword(insts)

	return strconv.Itoa(result)
}

func part2(input string) string {
	var result int

	return strconv.Itoa(result)
}

type pos struct {
	x, y int
}

func (p pos) add(q pos) pos {
	return pos{p.x + q.x, p.y + q.y}
}

type board struct {
	fields   [][]byte
	dir      int
	p        pos
	wrapFunc func(*pos)
}

func (b *board) turnLeft() {
	b.dir = (b.dir + 3) % 4
}

func (b *board) turnRight() {
	b.dir = (b.dir + 1) % 4
}

func (b *board) move(dist int) {
	n := b.p
	for dist > 0 {
		b.fields[b.p.y][b.p.x] = dirChar[b.dir]
		n = n.add(dirPos[b.dir])
		b.wrapFunc(&n)
		if b.fields[n.y][n.x] == '#' {
			break
		}
		if b.fields[n.y][n.x] > '#' {
			dist--
			b.p = n
		}
	}
}

func (b *board) findPassword(insts string) int {
	for len(insts) > 0 {
		instEnd := 1
		switch insts[0] {
		case 'L':
			b.turnLeft()
		case 'R':
			b.turnRight()
		default:
			instEnd = strings.IndexAny(insts, "LR")
			if instEnd == -1 {
				instEnd = len(insts)
			}
			b.move(goutils.MustBeInt(insts[:instEnd]))
		}
		insts = insts[instEnd:]
	}

	return 1000*(b.p.y+1) + 4*(b.p.x+1) + b.dir
}

var dirChar = map[int]byte{
	0: '>',
	1: 'v',
	2: '<',
	3: '^',
}

var dirPos = map[int]pos{
	0: {1, 0},
	1: {0, 1},
	2: {-1, 0},
	3: {0, -1},
}

func parseInput(input string) (board, string) {
	parts := strings.Split(strings.TrimRight(input, "\n"), "\n\n")
	ll := len(parts[0])
	var b board
	for _, l := range strings.Split(parts[0], "\n") {
		k := make([]byte, ll)
		copy(k, l)
		b.fields = append(b.fields, k)
	}
	for ; b.fields[b.p.y][b.p.x] != '.'; b.p.x++ {
	}
	return b, parts[1]
}
