package main

import (
	"slices"
	"strconv"

	"github.com/aibor/aoc/goutils"
)

func part1(input string) string {
	var result int

	for _, line := range goutils.SplitInput(input) {
		sequence := parseSequence(line)
		result += extrapolate(sequence)
	}

	return strconv.Itoa(result)
}

func part2(input string) string {
	var result int

	for _, line := range goutils.SplitInput(input) {
		sequence := parseSequence(line)
		slices.Reverse(sequence)
		result += extrapolate(sequence)
	}

	return strconv.Itoa(result)
}

func parseSequence(input string) (sequence []int) {
	i := goutils.NewStringFieldsIterator(input)
	for i.Next() {
		sequence = append(sequence, i.MustBeInt())
	}
	return
}

func extrapolate(sequence []int) (next int) {
	for slices.ContainsFunc(sequence, isNonZero) {
		next += sequence[len(sequence)-1]
		sequence = derive(sequence)
	}
	return
}

func derive(in []int) (out []int) {
	a := in[0]
	for _, b := range in[1:] {
		out = append(out, b-a)
		a = b
	}
	return
}

func isNonZero(e int) bool {
	return e != 0
}
