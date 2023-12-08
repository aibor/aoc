package main

import (
	"fmt"
	"strconv"

	"github.com/aibor/aoc/goutils"
)

func part1(input string) string {
	var result int

	lines := goutils.SplitInput(input)

	nw := parseNetwork(lines[2:])
	result = nw.walk(lines[0])

	return strconv.Itoa(result)
}

func part2(input string) string {
	var result int

	for _, line := range goutils.SplitInput(input) {
		_ = line
	}

	return strconv.Itoa(result)
}

type node [2]string
type network map[string]node

func parseNetwork(input []string) *network {
	nw := make(network, 1024)
	for _, line := range input {
		var self, nodeL, nodeR string
		fmt.Sscanf(line, "%3s = (%3s, %3s)", &self, &nodeL, &nodeR)
		nw[self] = node{nodeL, nodeR}
	}

	return &nw
}

func (nw *network) walk(moves string) int {
	var steps int
	moveIter := goutils.NewIterator([]rune(moves))
	cur := "AAA"

	for moveIter.Next() {
		steps++
		switch moveIter.Value() {
		case 'L':
			cur = (*nw)[cur][0]
		case 'R':
			cur = (*nw)[cur][1]
		}
		if cur == "ZZZ" {
			break
		}
		if moveIter.LenLeft() == 0 {
			moveIter.Reset()
		}
	}

	return steps
}
