package main

import (
	"strconv"
	"strings"

	"github.com/aibor/aoc/goutils"
)

var (
	exampleResult1 = "152"
	exampleResult2 = "301"

	result1 = "38914458159166"
	result2 = "3665520865940"
)

func part1(input string) string {
	var result int

	root := parseMonkeys(goutils.SplitInput(input))
	result = root.yell()

	return strconv.Itoa(result)
}

func part2(input string) string {
	var result int

	root := parseMonkeys(goutils.SplitInput(input))
	root.op = "="
	result = root.track("humn", 0)

	return strconv.Itoa(result)
}

func monkeyOp(op string, a, b int) int {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	case "=":
		return b
	}
	panic("oh no")
}

var reverseOp = map[string]string{
	"+": "-",
	"-": "+",
	"*": "/",
	"/": "*",
	"=": "=",
}

type monkey struct {
	name       string
	num        int
	op         string
	dep1, dep2 *monkey
}

func (m *monkey) yell() int {
	if m.num == 0 {
		m.num = monkeyOp(m.op, m.dep1.yell(), m.dep2.yell())
	}
	return m.num
}

func (m *monkey) track(name string, val int) int {
	if m.name == name {
		return val
	}

	var newVal int
	if m.dep1.hasDep(name) {
		newVal = monkeyOp(reverseOp[m.op], val, m.dep2.yell())
		return m.dep1.track(name, newVal)
	}
	if m.dep2.hasDep(name) {
		if m.op == "/" || m.op == "-" {
			newVal = monkeyOp(m.op, m.dep1.yell(), val)
		} else {
			newVal = monkeyOp(reverseOp[m.op], val, m.dep1.yell())
		}
		return m.dep2.track(name, newVal)
	}
	panic("failtrack")
}

func (m *monkey) hasDep(name string) bool {
	if m == nil {
		return false
	}
	return m.name == name || m.dep1.hasDep(name) || m.dep2.hasDep(name)
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

func parseMonkeys(lines []string) *monkey {
	ms := make(monkeys, len(lines))
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
	return ms["root"]
}
