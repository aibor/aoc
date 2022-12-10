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

func findMarker(message string, length int) int {
	var idx int
	for i := length; i < len(message); i += idx + 1 {
		idx = lastFirstNonUniqueIndex(message[i-length : i])
		if idx == length {
			return i
		}
	}
	return 0
}

func lastFirstNonUniqueIndex(chunk string) int {
	length := len(chunk)
	for i := length - 1; i > 0; i-- {
		if j := strings.IndexByte(chunk, chunk[i]); j != i {
			return j
		}
	}
	return length
}
