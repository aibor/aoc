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

	result = ms["root"].yell()

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
		return func(a, b int) int { return a + b }
	case "-":
		return func(a, b int) int { return a - b }
	case "*":
		return func(a, b int) int { return a * b }
	case "/":
		return func(a, b int) int { return a / b }
	}
	panic("oh no")
}

type monkey struct {
	name       string
	num        int
	op         string
	dep1, dep2 *monkey
}

func (m *monkey) yell() int {
	if m.num == 0 {
		m.num = monkeyOps(m.op)(m.dep1.yell(), m.dep2.yell())
	}
	return m.num
}

type monkeys map[string]*monkey

func (ms *monkeys) get(name string) *monkey {
	m, e := (*ms)[name]
	if !e {
		m = &monkey{name: name}
		(*ms)[name] = m
	}
	return m
}

func parseMonkeys(lines []string) monkeys {
	ms := make(monkeys, 8192)
	for _, l := range lines {
		f := strings.Fields(l)
		name := f[0][:4]
		m := ms.get(name)
		if len(f) == 2 {
			m.num = goutils.MustBeInt(f[1])
			continue
		}
		m.op = f[2]
		m.dep1 = ms.get(f[1])
		m.dep2 = ms.get(f[3])
	}
	return ms
}
