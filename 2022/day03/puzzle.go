package main

import (
	"fmt"
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
	lines := goutils.SplitInput(input)
	for len(lines) > 0 {
		score += priority(findCommonItemInGroup(lines[0:3]))
		lines = lines[3:]
	}
	return fmt.Sprintf("%d", score)
}

func priority(c rune) int {
	// Uppercase letters
	if c < 0x60 {
		return int(c - 38)
	}
	// Lowercase letters
	return int(c - 96)
}

func findCommonItemInRucksack(content string) rune {
	// Check every item in first compartment for presence in second compartment
	// and return the first item found in both compartments.
	l := len(content) / 2
	for i := 0; i < l; i++ {
		for j := l; j < 2*l; j++ {
			if content[i] == content[j] {
				return rune(content[i])
			}
		}
	}

	return 0
}

func findCommonItemInGroup(rucksacks []string) rune {
	// Walk items in first rucksack and search second rucksack for them. Search
	// third rucksack for items found in the first two and return the first
	// item found in all three.
	for _, r0 := range rucksacks[0] {
		for _, r1 := range rucksacks[1] {
			if r0 != r1 {
				continue
			}
			for _, r2 := range rucksacks[2] {
				if r0 == r2 {
					return r0
				}
			}
		}
	}

	return 0
}
