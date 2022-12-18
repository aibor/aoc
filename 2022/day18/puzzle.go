package main

import (
	"strconv"
	"strings"

	"github.com/aibor/aoc/goutils"
)

var (
	exampleResult1 = "64"
	exampleResult2 = "58"

	result1 = "4608"
	result2 = "2652"
)

func part1(input string) string {
	var result int

	var conn int
	s := make(space, 1024)
	for _, line := range goutils.SplitInput(input) {
		c := s.addCube(line)
		for _, n := range c.adjacent() {
			if s[n] {
				conn++
			}
		}
	}
	result = len(s)*6 - conn*2

	return strconv.Itoa(result)
}

func part2(input string) string {
	var result int

	s := make(space, 1024)
	for _, line := range goutils.SplitInput(input) {
		s.addCube(line)
	}
	for c := range s {
		for _, a := range c.adjacent() {
			if s[a] {
				continue
			}
			if s.checkOpen(a, make(map[cube]bool, 4096)) {
				result++
			}
		}
	}

	return strconv.Itoa(result)
}

type space map[cube]bool

type cube struct {
	x, y, z int
}

var dirs = [6]cube{
	{-1, 0, 0},
	{1, 0, 0},
	{0, -1, 0},
	{0, 1, 0},
	{0, 0, -1},
	{0, 0, 1},
}

func (c *cube) adjacent() [6]cube {
	var adj [6]cube
	for i, n := range dirs {
		adj[i] = cube{c.x + n.x, c.y + n.y, c.z + n.z}
	}
	return adj
}

func (s *space) addCube(line string) cube {
	l := strings.Split(line, ",")
	c := cube{
		goutils.MustBeInt(l[0]),
		goutils.MustBeInt(l[1]),
		goutils.MustBeInt(l[2]),
	}
	(*s)[c] = true
	return c
}

func (s *space) checkOpen(c cube, path map[cube]bool) bool {
	if (*s)[c] {
		return false
	}
	if c.x == 0 || c.x == 32 || c.y == 0 || c.y == 32 || c.z == 0 || c.z == 32 {
		return true
	}
	for _, d := range dirs {
		n := cube{c.x + d.x, c.y + d.y, c.z + d.z}
		if path[n] {
			continue
		}
		path[n] = true
		if s.checkOpen(n, path) {
			return true
		}
	}
	return false
}
