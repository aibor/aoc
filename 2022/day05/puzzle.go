package main

import (
	"fmt"
	"strings"
)

var (
	exampleResult1 = "CMZ"
	exampleResult2 = "MCD"

	result1 = "TGWSMRBPN"
	result2 = "TZLTLWRNF"
)

func part1(input string) string {
	parts := strings.Split(input, "\n\n")
	start := strings.Split(parts[0], "\n")
	yard := newYard(start)

	var amount, from, to int
	for _, line := range strings.Split(parts[1], "\n") {
		if n, _ := fmt.Sscanf(line, "move %d from %d to %d", &amount, &from, &to); n != 3 {
			continue
		}
		for i := 0; i < amount; i++ {
			yard.move(from - 1, to - 1)
		}
	}

	return yard.top()
}

func part2(input string) string {
	parts := strings.Split(input, "\n\n")
	start := strings.Split(parts[0], "\n")
	yard := newYard(start)

	var amount, from, to int
	for _, line := range strings.Split(parts[1], "\n") {
		if n, _ := fmt.Sscanf(line, "move %d from %d to %d", &amount, &from, &to); n != 3 {
			continue
		}
		crates := make([]rune,0 )
		for i := 0; i < amount; i++ {
			crates = append(crates, yard[from - 1].pop())
		}
		for i := len(crates) - 1; i >= 0; i-- {
			yard[to - 1].push(crates[i])
		}
	}

	return yard.top()
}

type stacks []*stack

func newYard(start []string) stacks {
	yard := stacks{}

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

	return yard
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
	e := &stack{}
	*s = append(*s, e)
	return e
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
