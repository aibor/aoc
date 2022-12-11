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

func part1(input string) string {
	monkeys := parseMonkeys(input)
	worryDecrease := func(i *int) {
		*i /= 3
	}
	for i := 1; i <= 20; i++ {
		monkeys.playRound(worryDecrease)
	}
	return strconv.Itoa(monkeys.businessLevel())
}

func part2(input string) string {
	monkeys := parseMonkeys(input)
	n := monkeys.divisor()
	worryDecrease := func(i *int) {
		*i %= n
	}
	for i := 1; i <= 10000; i++ {
		monkeys.playRound(worryDecrease)
	}
	return strconv.Itoa(monkeys.businessLevel())
}

func parseMonkeys(input string) monkeys {
	var curMonkey *monkey
	var all monkeys

	iter := stringIterator(strings.Fields(input))
mainLoop:
	for {
		switch iter.value() {
		case "Monkey":
			curMonkey = &monkey{items: make(queue, 0, 16)}
			all = append(all, curMonkey)
		case "Starting":
			iter.next() // "items:"
			for iter.next() {
				n, err := strconv.Atoi(strings.Replace(iter.value(), ",", "", 1))
				if err != nil {
					// Continue with main loop to process this token as well.
					continue mainLoop
				}
				curMonkey.items.push(n)
			}
		case "Operation:":
			iter.skip(4)
			operator := iter.value()
			iter.next()
			if n, err := strconv.Atoi(iter.value()); err == nil {
				switch operator {
				case "*":
					curMonkey.inspect = func(i *int) { *i *= n }
				case "+":
					curMonkey.inspect = func(i *int) { *i += n }
				}
			} else {
				switch operator {
				case "*":
					curMonkey.inspect = func(i *int) { *i *= *i }
				case "+":
					curMonkey.inspect = func(i *int) { *i += *i }
				}
			}
		case "Test:":
			iter.skip(3)
			n, _ := strconv.Atoi(iter.value())
			curMonkey.testDivisor = n
			iter.skip(6)
			n, _ = strconv.Atoi(iter.value())
			curMonkey.targetTrue = n
			iter.skip(6)
			n, _ = strconv.Atoi(iter.value())
			curMonkey.targetFalse = n
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

func (i *stringIterator) peek() string {
	if len(*i) > 1 {
		return (*i)[1]
	}
	return ""
}

type queue []int

func (q *queue) push(i int) {
	(*q) = append(*q, i)
}

func (q *queue) empty() bool {
	return len(*q) == 0
}

func (q *queue) pop() int {
	i := (*q)[0]
	*q = (*q)[1:]
	return i
}

type worryOp func(*int)

type monkey struct {
	items       queue
	inspected   int
	inspect     worryOp
	testDivisor int
	targetTrue  int
	targetFalse int
}

func (m *monkey) String() string {
	f := "items: % d; inspected: %d; test: (x/%d)*%d == x ? %d : %d"
	return fmt.Sprintf(f, m.items, m.inspected, m.testDivisor, m.testDivisor, m.targetTrue, m.targetFalse)
}

type monkeys []*monkey

func (all *monkeys) playRound(w worryOp) {
	for _, m := range *all {
		for !m.items.empty() {
			i := m.items.pop()
			m.inspect(&i)
			m.inspected++
			w(&i)
			if i%m.testDivisor == 0 {
				(*all)[m.targetTrue].items.push(i)
			} else {
				(*all)[m.targetFalse].items.push(i)
			}
		}
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
