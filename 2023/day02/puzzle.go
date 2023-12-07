package main

import (
	"strconv"
	"strings"

	"github.com/aibor/aoc/goutils"
)

func part1(input string) string {
	var result int

	valid := func(g game) bool {
		for _, s := range g.sets {
			if s.red > 12 || s.green > 13 || s.blue > 14 {
				return false
			}
		}
		return true
	}

	for _, line := range goutils.SplitInput(input) {
		g := parseGame(line)
		if valid(g) {
			result += g.id
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

type set struct {
	red   int
	green int
	blue  int
}

type game struct {
	id   int
	sets []set
}

func parseGame(input string) game {
	i := goutils.NewStringFieldsIterator(input)
	i.Skip(1)
	g := game{
		id:   goutils.MustBeInt(i.Value()[:len(i.Value())-1]),
		sets: make([]set, 1),
	}
	s := &g.sets[0]

	for i.Next() {
		num := i.MustBeInt()
		i.Next()
		switch {
		case strings.HasPrefix(i.Value(), "red"):
			s.red = num
		case strings.HasPrefix(i.Value(), "green"):
			s.green = num
		case strings.HasPrefix(i.Value(), "blue"):
			s.blue = num
		}
		if strings.HasSuffix(i.Value(), ";") {
			g.sets = append(g.sets, set{})
			s = &g.sets[len(g.sets)-1]
		}
	}

	return g
}
