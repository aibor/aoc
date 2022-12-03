package main

import (
	"fmt"

	"github.com/aibor/aoc/goutils"
)

func main() {
	input, _ := goutils.ReadInput()

	fmt.Println("Part 1:", part01(input))
	fmt.Println("Part 2:", part02(input))
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

func part01(input []string) string {
	var score int
	for _, rucksack := range input {
		score += priority(findCommonItemInRucksack(rucksack))
	}
	return fmt.Sprintf("%d", score)
}

func part02(input []string) string {
	var score int
	for i := 0; i < len(input); i += 3 {
		score += priority(findCommonItemInGroup(input[i : i+3]))
	}
	return fmt.Sprintf("%d", score)
}
