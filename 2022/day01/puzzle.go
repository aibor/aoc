package main

import (
	"fmt"
	"strconv"
	"strings"
)

var (
	exampleResult1 = "24000"
	exampleResult2 = "45000"

	result1 = "68467"
	result2 = "203420"
)

func part1(input string) string {
	var cal, result int

	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			cal = 0
		} else {
			c, _ := strconv.Atoi(line)
			cal += c
			if cal > result {
				result = cal
			}
		}
	}

	return fmt.Sprintf("%d", result)
}

func part2(input string) string {
	var i, result int

	cals := make([]int, 1, 256)
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			cals = append(cals, 0)
			i++
		} else {
			c, _ := strconv.Atoi(line)
			cals[i] += c
		}
	}

	for j := 0; j < 3; j++ {
		var maxk int
		for k, c := range cals {
			if c > cals[maxk] {
				maxk = k
			}
		}
		result += cals[maxk]
		cals[maxk] = 0
	}

	return fmt.Sprintf("%d", result)
}
