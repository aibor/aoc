package main

import (
  "fmt"
  "strconv"

  "github.com/aibor/aoc/goutils"
)

func main() {
  input, _ := goutils.ReadInput()

  fmt.Println("Part 1:", part01(input))
  fmt.Println("Part 2:", part02(input))
}

func part01(input []string) string {
	var cal, max int

	for _, line := range input {
		if cal > max {
			max = cal
		}
		if line == "" {
			cal = 0
		} else {
			c, _ := strconv.Atoi(line)
			cal += c
		}
	}

	return fmt.Sprintf("%d", max)
}

func part02(input []string) string {
	var cal, maxsum int
	cals := make([]int, 0)

	for _, line := range input {
		if line == "" {
			cals = append(cals, cal)
			cal = 0
		} else {
			c, _ := strconv.Atoi(line)
			cal += c
		}
	}
	cals = append(cals, cal)

	for j := 0; j < 3; j++ {
		var maxk int
		for k, c := range cals {
			if c > cals[maxk] {
				maxk = k
			}
		}
		maxsum += cals[maxk]
		cals[maxk] = 0
	}

	return fmt.Sprintf("%d", maxsum)
}
