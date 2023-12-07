package main

import (
	"strconv"
	"strings"

	"github.com/aibor/aoc/goutils"
)

func part1(input string) string {
	var result int

	for _, line := range goutils.SplitInput(input) {
		first := strings.IndexAny(line, "0123456789")
		last := strings.LastIndexAny(line, "0123456789")
		result += goutils.MustBeInt(string([]byte{line[first], line[last]}))
	}

	return strconv.Itoa(result)
}

func part2(input string) string {
	var result int

	if input == example1 {
		input = example2
	}

	for _, line := range goutils.SplitInput(input) {
		var firstDigit, lastDigit word

		firstDigit.index = strings.IndexAny(line, "0123456789")
		if firstDigit.index >= 0 {
			firstDigit.value = string(line[firstDigit.index])
		} else {
			firstDigit.index = len(line)
		}

		lastDigit.index = strings.LastIndexAny(line, "0123456789")
		if lastDigit.index >= 0 {
			lastDigit.value = string(line[lastDigit.index])
		}

		for word := range words {
			if i := strings.Index(line, word); i >= 0 && i < firstDigit.index {
				firstDigit.index = i
				firstDigit.value = words[word]
			}
			if i := strings.LastIndex(line, word); i >= 0 && i > lastDigit.index {
				lastDigit.index = i
				lastDigit.value = words[word]
			}
		}
		result += goutils.MustBeInt(firstDigit.value + lastDigit.value)
	}

	return strconv.Itoa(result)
}

type word struct {
	index int
	value string
}

var words = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}
