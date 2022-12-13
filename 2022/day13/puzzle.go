package main

import (
	"encoding/json"
	_ "fmt"
	"sort"
	"strconv"

	"github.com/aibor/aoc/goutils"
)

var (
	exampleResult1 = "13"
	exampleResult2 = "140"

	result1 = "4809"
	result2 = "22600"
)

func part1(input string) string {
	var result int

	pair := 1
	lines := goutils.SplitInput(input)
	lines = append(lines, "")
	for len(lines) > 0 {
		var left, right []any
		json.Unmarshal([]byte(lines[0]), &left)
		json.Unmarshal([]byte(lines[1]), &right)
		e := compare(left, right)
		if e == -1 {
			result += pair
		}

		lines = lines[3:]
		pair++
	}

	return strconv.Itoa(result)
}

func part2(input string) string {
	var result int = 1

	lines := goutils.SplitInput(input)
	packets := append(
		make(Packets, 0, len(lines)),
		Packet{[]any{[]any{2.0}}, true},
		Packet{[]any{[]any{6.0}}, true},
	)
	for _, line := range lines {
		if line == "" {
			continue
		}
		var p Packet
		json.Unmarshal([]byte(line), &p.data)
		packets = append(packets, p)
	}

	sort.Sort(packets)

	for i, p := range packets {
		if p.divider {
			result *= i + 1
		}
	}

	return strconv.Itoa(result)
}

func compare(left, right []any) int {
	for len(left) > 0 {
		if len(right) < 1 {
			return 1
		}
		l, r := left[0], right[0]
		left, right = left[1:], right[1:]
		lInt, lIntOk := l.(float64)
		rInt, rIntOk := r.(float64)
		ls, _ := l.([]any)
		rs, _ := r.([]any)
		switch {
		case lIntOk && rIntOk:
			switch {
			case lInt < rInt:
				return -1
			case lInt > rInt:
				return 1
			}
		case lIntOk:
			ls = []any{lInt}
		case rIntOk:
			rs = []any{rInt}
		}
		if o := compare(ls, rs); o != 0 {
			return o
		}

	}
	if len(right) == len(left) {
		return 0
	}
	return -1
}

type Packet struct {
	data    []any
	divider bool
}

type Packets []Packet

func (p Packets) Len() int           { return len(p) }
func (p Packets) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p Packets) Less(i, j int) bool { return compare(p[i].data, p[j].data) == -1 }
