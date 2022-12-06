package main

import (
	"strconv"
	"strings"

	"github.com/aibor/aoc/goutils"
)

var (
	exampleResult1 = "7,5,6,10,11"
	exampleResult2 = "19,23,23,29,26"

	result1 = "1080"
	result2 = "3645"
)

func part1(input string) string {
	result := make([]string, 0, 8)

	for _, line := range goutils.SplitInput(input) {
		m := findMarker(line, 4)
		result = append(result, strconv.Itoa(m))
	}

	return strings.Join(result, ",")
}

func part2(input string) string {
	result := make([]string, 0, 8)

	for _, line := range goutils.SplitInput(input) {
		m := findMarker(line, 14)
		result = append(result, strconv.Itoa(m))
	}

	return strings.Join(result, ",")
}

func findMarker(s string, l int) int {
	for i := l; i < len(s); i++ {
		if isMarker(s[i-l:i]) {
			return i
		}
	}
	return 0
}

func isMarker(chunk string) bool {
	for idx, r := range chunk {
		if strings.LastIndexByte(chunk, byte(r)) != idx {
			return false
		}
	}
	return true
}
