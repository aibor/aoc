package main

import (
	"fmt"
	"strconv"

	"github.com/aibor/aoc/goutils"
)

func part1(input string) string {
	var result int

	lines := goutils.SplitInput(input)

	nw := parseNetwork(
		lines[2:],
		func(s string) bool { return s == "AAA" },
		func(s string) bool { return s == "ZZZ" },
	)
	result = nw.walk(lines[0])

	return strconv.Itoa(result)
}

func part2(input string) string {
	var result int

	if input == example1 {
		input = example2
	}

	lines := goutils.SplitInput(input)

	nw := parseNetwork(
		lines[2:],
		func(s string) bool { return s[2] == 'A' },
		func(s string) bool { return s[2] == 'Z' },
	)
	result = nw.walk(lines[0])

	return strconv.Itoa(result)
}

type node struct {
	name  string
	end   bool
	links [2]*node
	done  bool
}

type network struct {
	start []*node
	end   []*node
}

func (n *node) String() string {
	return fmt.Sprintf("%s", n.name)
}

func parseNetwork(input []string, start, end func(string) bool) *network {
	var nw network
	links := make(map[string][2]string, 1024)
	nodes := make(map[string]*node, 1024)

	for _, line := range input {
		var self, nodeL, nodeR string
		fmt.Sscanf(line, "%3s = (%3s, %3s)", &self, &nodeL, &nodeR)

		links[self] = [2]string{nodeL, nodeR}
		nodes[self] = &node{
			name: self,
			end:  end(self),
		}
		if start(self) {
			nw.start = append(nw.start, nodes[self])
		}
		if end(self) {
			nw.end = append(nw.end, nodes[self])
		}
	}
	for n, lr := range links {
		nodes[n].links = [2]*node{nodes[lr[0]], nodes[lr[1]]}
	}

	return &nw
}

func (nw *network) walk(moves string) int {
	var steps int
	var ends []int

	cur := nw.start

	// Run until for each track an end node is reached at the end of move
	// instructions. With those values, calculate the result where they all
	// line up by calculating least common multiple.
	moveIter := goutils.NewIterator([]rune(moves))
	for moveIter.Next() {
		var idx int

		steps++

		switch moveIter.Value() {
		case 'L':
			idx = 0
		case 'R':
			idx = 1
		}

		for i := 0; i < len(cur); i++ {
			if cur[i].done {
				continue
			}
			cur[i] = cur[i].links[idx]
			if cur[i].end {
				if moveIter.LenLeft() == 0 {
					cur[i].done = true
					ends = append(ends, steps)
				}
			}
		}

		if len(ends) == len(nw.start) {
			break
		}

		if moveIter.LenLeft() == 0 {
			moveIter.Reset()
		}
	}

	return goutils.LCM(ends...)
}
