package main

import (
	"strconv"
	"strings"

	"github.com/aibor/aoc/goutils"
)

var (
	exampleResult1 = "95437"
	exampleResult2 = "24933642"

	result1 = "1915606"
	result2 = "5025657"
)

func part1(input string) string {
	var result int
	sizes := getSizes(goutils.SplitInput(input))

	for _, size := range sizes {
		if size <= 100000 {
			result += size
		}
	}

	return strconv.Itoa(result)
}

func part2(input string) string {
	sizes := getSizes(goutils.SplitInput(input))
	l := "/"
	unused := 70000000 - sizes[l]

	for d, size := range sizes {
		if unused+size >= 30000000 && size < sizes[l] {
			l = d
		}
	}

	return strconv.Itoa(sizes[l])
}

type stack []string

func (s stack) String() string {
	return strings.Join(s, "/")
}

func (s *stack) push(e string) {
	*s = append(*s, e)
}

func (s *stack) pop() string {
	l := len(*s)
	if l < 1 {
		return ""
	}
	e := (*s)[l-1]
	*s = append((*s)[:l-1])
	return e
}

func (s *stack) last() string {
	l := len(*s)
	if l < 1 {
		return ""
	}
	return (*s)[l-1]
}

func getSizes(input []string) map[string]int {
	var line string
	var size int
	dirstack := stack{}
	sizes := make(map[string]int, 64)

	for _, line = range input {
		switch {
		case strings.Contains(line, " cd "):
			if line[5:] == ".." {
				dirstack.pop()
			} else {
				dirstack.push(line[5:])
			}
		case line[0] >= '0' && line[0] <= '9':
			size, _ = strconv.Atoi(line[:strings.Index(line, " ")])
			for l := len(dirstack); l > 0; l-- {
				sizes[dirstack[:l].String()] += size
			}
		}
	}
	return sizes
}
