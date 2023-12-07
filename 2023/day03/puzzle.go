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
			num := number{point{lineNum, match[0]}, match[1]}
			if em.check(num, em.isSymbol) != nil {
				result += goutils.MustBeInt(line[match[0]:match[1]])
			}
		}
	}

	return strconv.Itoa(result)
}

func part2(input string) string {
	var result int
	gears := make(map[point][]int)

	re := regexp.MustCompile(`\d+`)

	em := engineMap(goutils.SplitInput(input))
	for lineNum, line := range em {
		matches := re.FindAllStringIndex(line, -1)
		for _, match := range matches {
			num := number{point{lineNum, match[0]}, match[1]}
			if p := em.check(num, em.isGearSymbol); p != nil {
				gears[*p] = append(gears[*p], goutils.MustBeInt(line[match[0]:match[1]]))
			}
		}
	}

	for _, nums := range gears {
		if len(nums) == 2 {
			result += nums[0] * nums[1]
		}
	}

	return strconv.Itoa(result)
}

type engineMap []string

type point struct {
	line int
	col  int
}

type number struct {
	point
	colNext int
}

func (em *engineMap) check(num number, f func(int, int) bool) *point {
	if num.col > 0 && f(num.line, num.col-1) {
		return &point{num.line, num.col - 1}
	}
	if num.colNext < len((*em)[0]) && f(num.line, num.colNext) {
		return &point{num.line, num.colNext}
	}

	scanLine := func(l int) *point {
		for c := num.col - 1; c <= num.colNext; c++ {
			if c < 0 || c >= len((*em)[0]) {
				continue
			}
			if f(l, c) {
				return &point{l, c}
			}
		}
		return nil
	}

	if num.line > 0 {
		if p := scanLine(num.line - 1); p != nil {
			return p
		}
	}
	if num.line < len(*em)-1 {
		if p := scanLine(num.line + 1); p != nil {
			return p
		}
	}
	return nil
}

func (em *engineMap) isSymbol(lineNum, colNum int) bool {
	r := (*em)[lineNum][colNum]
	if r == '.' {
		return false
	}
	return !unicode.IsDigit(rune(r))
}

func (em *engineMap) isGearSymbol(lineNum, colNum int) bool {
	r := (*em)[lineNum][colNum]
	return r == '*'
}
