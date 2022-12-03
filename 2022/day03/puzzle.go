package main

import (
	"fmt"
	"strings"

	"github.com/aibor/aoc/goutils"
)

var (
	exampleResult1 = "157"
	exampleResult2 = "70"

	result1 = "8240"
	result2 = "2587"
)

func part1(input string) string {
	var result int

	for _, rucksack := range goutils.SplitInput(input) {
		half := len(rucksack) / 2
		for _, item := range rucksack[:half] {
			if strings.ContainsRune(rucksack[half:], item) {
				result += priority(item)
				break
			}
		}
	}

	return fmt.Sprintf("%d", result)
}

func part2(input string) string {
	var result int

	rucksacks := goutils.SplitInput(input)
	for ; len(rucksacks) > 0; rucksacks = rucksacks[3:] {
		for _, item := range rucksacks[0] {
			if strings.ContainsRune(rucksacks[1], item) &&
				strings.ContainsRune(rucksacks[2], item) {
				result += priority(item)
				break
			}
		}
	}

	return fmt.Sprintf("%d", result)
}

func priority(c rune) int {
	// Uppercase letters
	if c < 'a' {
		return int(c - 'A' + 27)
	}
	// Lowercase letters
	return int(c - 'a' + 1)
}
