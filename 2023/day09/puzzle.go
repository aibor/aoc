package main

import (
	"slices"
	"strconv"

	"github.com/aibor/aoc/goutils"
)

func part1(input string) string {
	var result int

	for _, line := range goutils.SplitInput(input) {
		var sequence []int
		i := goutils.NewStringFieldsIterator(line)
		for i.Next() {
			sequence = append(sequence, i.MustBeInt())
		}
		result += extrapolate(sequence)
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

func extrapolate(sequence []int) (next int) {
	isNonZero := func(e int) bool { return e != 0 }
	for slices.ContainsFunc(sequence, isNonZero) {
		next += sequence[len(sequence)-1]
		sequence = derive(sequence)
	}
	return
}

func derive(in []int) (out []int) {
	i := goutils.NewIterator(in)
	i.Next()
	a := i.Value()
	for i.Next() {
		b := i.Value()
		out = append(out, b-a)
		a = b
	}
	return
}
