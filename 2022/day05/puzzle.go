package main

import (
	"strconv"
	"strings"

	"github.com/aibor/aoc/goutils"
)

var (
	exampleResult1 = "CMZ"
	exampleResult2 = "MCD"

	result1 = "TGWSMRBPN"
	result2 = "TZLTLWRNF"
)

func part1(input string) string {
	yard, movements := newYard(input)
	for _, line := range movements {
		if line == "" {
			continue
		}
		amount, from, to := parseInst(line)
		for i := 0; i < amount; i++ {
			yard.move(from-1, to-1, 1)
		}
	}
	return yard.top()
}

func part2(input string) string {
	yard, movements := newYard(input)
	for _, line := range movements {
		if line == "" {
			continue
		}
		amount, from, to := parseInst(line)
		yard.move(from-1, to-1, amount)
	}
	return yard.top()
}

func parseInst(inst string) (amount, from, to int) {
	var prev, s int
	for idx, r := range inst {
		if r == ' ' {
			s++
			switch s {
			case 2:
				amount, _ = strconv.Atoi(inst[prev:idx])
			case 4:
				from, _ = strconv.Atoi(inst[prev:idx])
			}
			prev = idx + 1
		}
	}
	to, _ = strconv.Atoi(inst[prev:])
	return
}

type crateStack = goutils.Stack[rune]

type crateYard [10]crateStack

func newYard(input string) (crateYard, []string) {
	parts := strings.SplitN(input, "\n\n", 2)
	start := strings.SplitN(parts[0], "\n", 16)
	yard := crateYard{}
	s := 0
	for idx, r := range start[len(start)-1] {
		if r < '0' || r > '9' {
			continue
		}
		yard.reset(s)
		for i := len(start) - 2; i >= 0; i-- {
			if len(start[i]) < idx {
				continue
			}
			crate := rune(start[i][idx])
			if crate >= 'A' {
				yard[s] = append(yard[s], crate)
			}
		}
		s++
	}

	return yard, strings.SplitN(parts[1], "\n", 512)
}

func (y *crateYard) top() string {
	var b strings.Builder
	b.Grow(len(y))
	for _, stack := range y {
		crate, _ := stack.Top()
		if crate != 0 {
			b.WriteRune(crate)
		}
	}
	return b.String()
}

func (y *crateYard) move(from, to, n int) {
	if n == 1 {
		y[to].Push(y[from].Pop())
		return
	}
	l := len(y[from])
	if l < n {
		n = l
	}
	y[to] = append(y[to], y[from][l-n:]...)
	y[from] = y[from][:l-n]
}

func (y *crateYard) reset(s int) {
	y[s] = make(crateStack, 0, 64)
}
