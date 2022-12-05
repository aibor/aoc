package main

import (
	"fmt"
	"strconv"
	"strings"
)

var (
	exampleResult1 = "CMZ"
	exampleResult2 = "MCD"

	result1 = "TGWSMRBPN"
	result2 = "TZLTLWRNF"
)

func part1(input string) string {
	yard, movements := newYard(input)

	var amount, from, to int
	for _, line := range movements {
		if line == "" {
			continue
		}
		amount, from, to = parseInst(line)
		for i := 0; i < amount; i++ {
			yard.move(from - 1, to - 1)
		}
	}

	return yard.top()
}

func part2(input string) string {
	yard, movements := newYard(input)

	var amount, from, to int
	crates := make([]rune, 0, 64)
	for _, line := range movements {
		if line == "" {
			continue
		}
		amount, from, to = parseInst(line)
		crates = crates[:0]
		for i := 0; i < amount; i++ {
			crates = append(crates, yard[from - 1].pop())
		}
		for i := len(crates) - 1; i >= 0; i-- {
			yard[to - 1].push(crates[i])
		}
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

type stacks []*stack

func newYard(input string) (stacks, []string) {
	parts := strings.SplitN(input, "\n\n", 2)
	start := strings.SplitN(parts[0], "\n", 16)
	yard := make(stacks, 0, 10)

	for idx, r := range start[len(start) - 1] {
		if r < '0' || r > '9' {
			continue
		}
		s := yard.add()
		for i := len(start) - 2; i >= 0; i-- {
			if len(start[i]) < idx {
				continue
			}
			crate := rune(start[i][idx])
			if crate >= 'A' {
				s.push(crate)
			}
		}
	}

	return yard, strings.SplitN(parts[1], "\n", 512)
}

func (s *stacks) top() string {
	var t string
	for _, stack := range *s {
		crate := stack.pop()
		if crate != 0 {
			t += string(crate)
		}
	}
	return t
}

func (s *stacks) add() *stack {
	e := make(stack, 0, 64)
	*s = append(*s, &e)
	return &e
}

func (s *stacks) move(from, to int) {
	crate := (*s)[from].pop()
	(*s)[to].push(crate)
}

type stack []rune

func(s *stack) String() string {
	return fmt.Sprintf("% q", *s)
}

func (s *stack) push(r rune) {
	*s = append(*s, r)
}

func (s *stack) pop() rune {
	if len(*s) < 1 {
		return 0
	}
	r := (*s)[len(*s) - 1]
	*s = append((*s)[:len(*s) -1])
	return r
}
