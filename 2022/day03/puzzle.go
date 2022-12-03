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
	var score int
	for _, rucksack := range goutils.SplitInput(input) {
		score += priority(findCommonItemInRucksack(rucksack))
	}
	return fmt.Sprintf("%d", score)
}

func part2(input string) string {
	var score int
	for lines := goutils.SplitInput(input); len(lines) > 0; lines = lines[3:] {
		score += priority(findCommonItemInGroup(lines[0:3]))
	}
	return fmt.Sprintf("%d", score)
}

func priority(c rune) int {
	// Uppercase letters
	if c < 'a' {
		return int(c - 'A' + 27)
	}
	// Lowercase letters
	return int(c - 'a' + 1)
}

func findCommonItemInRucksack(content string) rune {
	l := len(content) / 2
	for _, item := range content[:l] {
		if strings.ContainsRune(content[l:], item) {
			return item
		}
	}

	return 0
}

func findCommonItemInGroup(rucksacks []string) rune {
	for _, item := range rucksacks[0] {
		if strings.ContainsRune(rucksacks[1], item) {
			if strings.ContainsRune(rucksacks[2], item) {
				return item
			}
		}
	}

	return 0
}
