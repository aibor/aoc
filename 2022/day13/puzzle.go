package main

import (
	"encoding/json"
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
	var pair, result int
	var left, right []any

	lines := append(goutils.SplitInput(input), "")
	for len(lines) > 0 {
		pair++
		json.Unmarshal([]byte(lines[0]), &left)
		json.Unmarshal([]byte(lines[1]), &right)
		lines = lines[3:]
		if compare(left, right) == -1 {
			result += pair
		}
	}

	return strconv.Itoa(result)
}

func part2(input string) string {
	var result int

	lines := goutils.SplitInput(input)
	dividerPackets := []*Packet{
		{[]any{[]any{2.0}}, 0},
		{[]any{[]any{6.0}}, 0},
	}
	packets := append(make(Packets, 0, len(lines)), dividerPackets...)
	for _, line := range lines {
		if line != "" {
			var p Packet
			json.Unmarshal([]byte(line), &p.data)
			packets = append(packets, &p)
		}
	}
	sort.Sort(packets)
	result = (dividerPackets[0].index + 1) * (dividerPackets[1].index + 1)

	return strconv.Itoa(result)
}

func compare(left, right []any) int {
	for len(left) > 0 {
		if len(right) < 1 {
			return 1
		}
		l, r := left[0], right[0]
		left, right = left[1:], right[1:]
		lNum, lNumOk := l.(float64)
		rNum, rNumOk := r.(float64)
		switch {
		case lNumOk && rNumOk:
			switch {
			case lNum == rNum:
				continue
			case lNum < rNum:
				return -1
			case lNum > rNum:
				return 1
			}
		case lNumOk:
			l = []any{lNum}
		case rNumOk:
			r = []any{rNum}
		}
		if o := compare(l.([]any), r.([]any)); o != 0 {
			return o
		}

	}
	if len(right) == len(left) {
		return 0
	}
	return -1
}

type Packet struct {
	data  []any
	index int
}

type Packets []*Packet

func (p Packets) Len() int           { return len(p) }
func (p Packets) Less(i, j int) bool { return compare(p[i].data, p[j].data) == -1 }
func (p Packets) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
	p[i].index = i
	p[j].index = j
}
