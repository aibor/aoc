package main

import (
	"fmt"
	"strconv"
	"strings"
	"text/scanner"
)

var (
	exampleResult1 = "95437"
	exampleResult2 = "24933642"

	result1 = "1915606"
	result2 = "5025657"
)

func part1(input string) string {
	var result int
	sizes := getSizes(input)

	for _, size := range sizes {
		if size <= 100000 {
			result += size
		}
	}

	return strconv.Itoa(result)
}

func part2(input string) string {
	sizes := getSizes(input)
	l := "[/]"
	unused := 70000000 - sizes[l]

	for d, size := range sizes {
		if unused + size >= 30000000 && size < sizes[l] {
			l = d
		}
	}

	return strconv.Itoa(sizes[l])
}

type stack []string

func(s stack) String() string {
	return fmt.Sprintf("% s", []string(s))
}

func (s *stack) push(e string) {
	*s = append(*s, e)
}

func (s *stack) pop() string {
	l := len(*s)
	if l < 1 {
		return ""
	}
	e := (*s)[l - 1]
	*s = append((*s)[:l -1])
	return e
}

func (s *stack) last() string {
	l := len(*s)
	if l < 1 {
		return ""
	}
	return (*s)[l-1]
}


func getSizes(input string) map[string]int {
	var inst scanner.Scanner
	var t, last string
	inst.Init(strings.NewReader(strings.TrimSpace(input)))
	dirstack := stack{}
	sizes := make(map[string]int,64)


	for tok := inst.Scan(); tok != scanner.EOF; tok = inst.Scan() {
		t = inst.TokenText()
		switch {
		case last == "cd":
			if t == "." {
				if inst.Peek() == '.' {
					dirstack.pop()
				}
			} else {
				dirstack.push(t)
			}
		case t[0]>= '0' && t[0]<='9':
			size, _ := strconv.Atoi(t)
			for l := len(dirstack); l > 0; l-- {
				sizes[dirstack[:l].String()] += size
			}
		}

		last = t
	}
	return sizes
}
