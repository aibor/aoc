package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

var (
	exampleResult1 = "10605"
	exampleResult2 = "2713310158"

	result1 = "316888"
	result2 = "35270398814"
)

const itemsSize = 64

func part1(input string) string {
	monkeys := parseMonkeys(input)
	worryDecrease := func(i int) int {
		return i / 3
	}
	for i := 1; i <= 20; i++ {
		monkeys.playRound(worryDecrease)
	}
	return strconv.Itoa(monkeys.businessLevel())
}

func part2(input string) string {
	monkeys := parseMonkeys(input)
	n := monkeys.divisor()
	worryDecrease := func(i int) int {
		return i % n
	}
	for i := 1; i <= 10000; i++ {
		monkeys.playRound(worryDecrease)
	}
	return strconv.Itoa(monkeys.businessLevel())
}

func parseMonkeys(input string) monkeys {
	var curMonkey *monkey
	all := make(monkeys, 0, 16)

	iter := stringIterator(strings.Fields(input))
	for {
		switch iter.value() {
		case "Monkey":
			curMonkey = &monkey{}
			curMonkey.resetItems()
			all = append(all, curMonkey)
		case "Starting":
			for iter.next() {
				// Peek only to not consume the breaking token so oit will be processed in the main loop porperly as well.
				n, err := strconv.Atoi(strings.Replace(iter.peek(), ",", "", 1))
				if err != nil {
					break
				}
				curMonkey.addItem(n)
			}
		case "Operation:":
			iter.skip(4)
			operator := iter.value()
			iter.next()
			if n, err := iter.num(); err == nil {
				switch operator {
				case "*":
					curMonkey.inspect = func(i int) int { return i * n }
				case "+":
					curMonkey.inspect = func(i int) int { return i + n }
				}
			} else {
				switch operator {
				case "*":
					curMonkey.inspect = func(i int) int { return i * i }
				case "+":
					curMonkey.inspect = func(i int) int { return i * 2 }
				}
			}
		case "Test:":
			iter.skip(3)
			curMonkey.testDivisor = iter.mustNum()
			iter.skip(6)
			curMonkey.targetTrue = iter.mustNum()
			iter.skip(6)
			curMonkey.targetFalse = iter.mustNum()
		}

		if !iter.next() {
			break
		}
	}

	return all
}

type stringIterator []string

func (i *stringIterator) next() bool {
	if len(*i) > 0 {
		*i = (*i)[1:]
		return true
	}
	return false
}

func (i *stringIterator) skip(n int) {
	l := len(*i)
	if l < n {
		n = l
	}
	*i = (*i)[n:]
}

func (i *stringIterator) value() string {
	if len(*i) > 0 {
		return (*i)[0]
	}
	return ""
}

func (i *stringIterator) num() (int, error) {
	return strconv.Atoi(i.value())
}

func (i *stringIterator) mustNum() int {
	n, err := i.num()
	if err != nil {
		panic("must parse Atoi")
	}
	return n
}

func (i *stringIterator) peek() string {
	if len(*i) > 1 {
		return (*i)[1]
	}
	return ""
}

type worryOp func(int) int

type monkey struct {
	items       []int
	inspected   int
	inspect     worryOp
	testDivisor int
	targetTrue  int
	targetFalse int
}

func (m *monkey) addItem(i int) {
	m.items = append(m.items, i)
}

func (m *monkey) resetItems() {
	if m.items == nil {
		m.items = make([]int, 0, itemsSize)
	} else {
		m.items = m.items[:0]
	}
}

func (m *monkey) String() string {
	f := "items: % d; inspected: %d; test: (x/%d)*%d == x ? %d : %d"
	return fmt.Sprintf(f, m.items, m.inspected, m.testDivisor, m.testDivisor, m.targetTrue, m.targetFalse)
}

type monkeys []*monkey

func (all *monkeys) playRound(w worryOp) {
	for _, m := range *all {
		for _, i := range m.items {
			i = w(m.inspect(i))
			m.inspected++
			if i%m.testDivisor == 0 {
				(*all)[m.targetTrue].addItem(i)
			} else {
				(*all)[m.targetFalse].addItem(i)
			}
		}
		m.resetItems()
	}
}

func (all *monkeys) businessLevel() int {
	var inspected []int
	for _, m := range *all {
		inspected = append(inspected, m.inspected)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(inspected)))
	return inspected[0] * inspected[1]
}
func (all *monkeys) String() string {
	var out string
	for i, m := range *all {
		out += fmt.Sprintf("monkey %d: %s\n\n", i, m)
	}
	return out
}

func (all *monkeys) divisor() int {
	out := 1
	for _, m := range *all {
		if m.testDivisor != 0 {
			out *= m.testDivisor
		}
	}
	return out
}
