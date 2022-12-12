package main

import (
	"strconv"
	"strings"

	"github.com/aibor/aoc/goutils"
)

var (
	exampleResult1 = "31"
	exampleResult2 = "29"

	result1 = "361"
	result2 = "354"
)

func part1(input string) string {
	var result int

	m := newMap(goutils.SplitInput(input))
	start := m.findNodes('S')
	result = m.search(start[0])

	return strconv.Itoa(result)
}

func part2(input string) string {
	var result int

	m := newMap(goutils.SplitInput(input))
	for _, start := range m.findNodes('a') {
		r := m.search(start)
		if result == 0 || r < result {
			result = r
		}
	}

	return strconv.Itoa(result)
}

func newMap(lines []string) *Map {
	return &Map{
		fields: lines,
		goalTest: func(b byte) bool {
			return b == 'E'
		},
		neighborTest: func(a, b byte) bool {
			if a == 'S' {
				a = 'a' - 1
			} else if b == 'E' {
				b = 'z' + 1
			}
			return a+1 >= b
		},
	}
}

type pos struct {
	x, y int
}

type node struct {
	pos
	cost int
}

func (p *node) neighbors() [4]node {
	return [4]node{
		{pos{p.x - 1, p.y}, p.cost + 1},
		{pos{p.x + 1, p.y}, p.cost + 1},
		{pos{p.x, p.y - 1}, p.cost + 1},
		{pos{p.x, p.y + 1}, p.cost + 1},
	}
}

type Map struct {
	fields       []string
	goalTest     func(byte) bool
	neighborTest func(byte, byte) bool
}

func (m *Map) findNodes(r rune) []node {
	var nodes []node
	for y, line := range m.fields {
		if x := strings.IndexRune(line, r); x != -1 {
			nodes = append(nodes, node{pos{x, y}, 0})
		}
	}
	return nodes
}

func (m *Map) validNeighbor(p, n pos) bool {
	if n.y < 0 || n.y >= len(m.fields) {
		return false
	}
	if n.x < 0 || n.x >= len(m.fields[n.y]) {
		return false
	}
	return m.neighborTest(m.fields[p.y][p.x], m.fields[n.y][n.x])
}

func (m *Map) search(p node) int {
	var cur node
	queue := make([]node, 1, 128)
	queue[0] = p
	seen := make(map[pos]bool, 8192)

	for len(queue) > 0 {
		cur = queue[0]
		queue = queue[1:]

		if m.goalTest(m.fields[cur.y][cur.x]) {
			return cur.cost
		}
		seen[cur.pos] = true

	neigh:
		for _, n := range cur.neighbors() {
			if seen[n.pos] || !m.validNeighbor(cur.pos, n.pos) {
				continue
			}
			for i, q := range queue {
				if q.pos == n.pos {
					if n.cost < q.cost {
						queue[i] = n
					}
					continue neigh
				}
			}
			queue = append(queue, n)
		}
	}
	return 0xffff
}
