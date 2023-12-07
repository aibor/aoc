package main

import (
	"regexp"
	"strconv"
	"unicode"

	"github.com/aibor/aoc/goutils"
)

func part1(input string) string {
	var result int

	re := regexp.MustCompile(`\d+`)

	em := engineMap(goutils.SplitInput(input))
	for lineNum, line := range em {
		matches := re.FindAllStringIndex(line, -1)
		for _, match := range matches {
			if em.isPartNumber(lineNum, match[0], match[1]) {
				result += goutils.MustBeInt(line[match[0]:match[1]])
			}
		}
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

type engineMap []string

func (em *engineMap) isPartNumber(lineNum, colStart, colNext int) bool {
	if em.isSymbol(lineNum, colStart-1) || em.isSymbol(lineNum, colNext) {
		return true
	}
	for _, l := range []int{lineNum - 1, lineNum + 1} {
		for c := colStart - 1; c <= colNext; c++ {
			if em.isSymbol(l, c) {
				return true
			}
		}
	}
	return false
}

func (em *engineMap) isSymbol(lineNum, colNum int) bool {
	if lineNum < 0 || lineNum >= len(*em) {
		return false
	}
	if colNum < 0 || colNum >= len((*em)[0]) {
		return false
	}
	r := (*em)[lineNum][colNum]
	if r == '.' {
		return false
	}
	return !unicode.IsDigit(rune(r))
}
