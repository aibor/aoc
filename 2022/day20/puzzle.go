package main

import (
	"container/ring"
	"strconv"

	"github.com/aibor/aoc/goutils"
)

var (
	exampleResult1 = "3"
	exampleResult2 = "1623178306"

	result1 = "27726"
	result2 = "4275451658004"
)

func part1(input string) string {
	var result int

	nl := parseNumbers(input, 1)
	nl.mix(1)
	result = nl.coordinates()

	return strconv.Itoa(result)
}

func part2(input string) string {
	var result int

	nl := parseNumbers(input, 811589153)
	nl.mix(10)
	result = nl.coordinates()

	return strconv.Itoa(result)
}

type numberList struct {
	zero  *ring.Ring
	order []*ring.Ring
}

func parseNumbers(input string, decryptionKey int) numberList {
	iter := goutils.NewStringFieldsIterator(input)
	nl := numberList{}
	r := ring.New(iter.Length())
	for iter.Next() {
		r.Value = iter.MustBeInt() * decryptionKey
		nl.order = append(nl.order, r)
		if r.Value == 0 {
			nl.zero = r
		}
		r = r.Next()
	}
	return nl
}

func (nl *numberList) mix(count int) {
	l := len(nl.order)
	for i := 0; i < count; i++ {
		for _, r := range nl.order {
			v := r.Value.(int)
			p := r.Prev()
			t := p.Unlink(1)
			p.Move(v % (l - 1)).Link(t)
		}
	}
}

func (nl *numberList) coordinates() int {
	var c int
	r := nl.zero
	for i := 1; i <= 3000; i++ {
		r = r.Next()
		if i%1000 == 0 {
			v := r.Value.(int)
			c += v
		}
	}
	return c
}
