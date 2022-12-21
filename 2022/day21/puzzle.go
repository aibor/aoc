package main

import (
	"strconv"
	"strings"

	"github.com/aibor/aoc/goutils"
)

var (
	exampleResult1 = "152"
	exampleResult2 = "0"

	result1 = "38914458159166"
	result2 = "0"
)

func part1(input string) string {
	var result int

	ms := parseMonkeys(goutils.SplitInput(input))

	var e bool
	for !e {
		for n, m := range ms.waiting {
			d1, f1 := ms.solved[m.deps[0]]
			d2, f2 := ms.solved[m.deps[1]]
			if !f1 || !f2 {
				continue
			}
			ms.solved[n] = m.op(d1, d2)
			delete(ms.waiting, n)
		}
		result, e = ms.solved["root"]
	}

	return strconv.Itoa(result)
}

func part2(input string) string {
	var result int

	for _, line := range goutils.SplitInput(input) {
		_ = line
	}

	return strconv.Itoa(result)
}

type monkeyOp func(int, int) int

func monkeyOps(op string) monkeyOp {
	switch op {
	case "+":
		return func(a, b int) int {return a + b}
	case "-":
		return func(a, b int) int {return a - b}
	case "*":
		return func(a, b int) int {return a * b}
	case "/":
		return func(a, b int) int {return a / b}
	}
	panic("oh no")
}

type monkey struct {
	deps [2]string
	op monkeyOp
}

type monkeys struct {
	solved map[string]int
	waiting map[string]monkey
}

func parseMonkeys(lines []string) monkeys {
	m := monkeys{
		solved: make(map[string]int, 8192),
		waiting: make(map[string]monkey, 8192),
	}
	for _, l := range lines {
		f := strings.Fields(l)
		name := f[0][:4]
		if len(f) == 2 {
			m.solved[name] = goutils.MustBeInt(f[1])
			continue
		}
		m.waiting[name] = monkey{
			deps: [2]string{f[1], f[3]},
			op: monkeyOps(f[2]),
		}
	}
	return m
}
