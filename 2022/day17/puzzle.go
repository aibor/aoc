package main

import (
	"fmt"
	"strconv"

	"github.com/aibor/aoc/goutils"
)

var (
	exampleResult1 = "3068"
	//exampleResult2 = "1514285714288"
	exampleResult2 = "0"

	result1 = "3083"
	result2 = "0"
)

func part1(input string) string {
	var result int

	c := chamber{
		movements: []byte(input),
	}
	c.drop(2022)
	c.print()
	result = c.top()

	return strconv.Itoa(result)
}

func part2(input string) string {
	var result int

	c := chamber{
		movements: []byte(input),
	}
	//c.drop(1000000000000)
	result = c.top()

	return strconv.Itoa(result)
}

type shape struct {
	ind    rune
	fields [4][4]bool
	x      int
	y      int
	width  int
	height int
}

func newShape(i int) shape {
	var s shape
	switch i {
	case 0:
		s = shape{
			ind:    '-',
			width:  4,
			height: 1,
			fields: [4][4]bool{
				{true, true, true, true},
			},
		}
	case 1:
		s = shape{
			ind:    '+',
			width:  3,
			height: 3,
			fields: [4][4]bool{
				{false, true, false},
				{true, true, true},
				{false, true, false},
			},
		}
	case 2:
		s = shape{
			ind:    'L',
			width:  3,
			height: 3,
			fields: [4][4]bool{
				{false, false, true},
				{false, false, true},
				{true, true, true},
			},
		}
	case 3:
		s = shape{
			ind:    '|',
			width:  1,
			height: 4,
			fields: [4][4]bool{
				{true},
				{true},
				{true},
				{true},
			},
		}
	case 4:
		s = shape{
			ind:    '¤',
			width:  2,
			height: 2,
			fields: [4][4]bool{
				{true, true},
				{true, true},
			},
		}
	}
	return s
}

type chamber struct {
	cols      [7][4096]rune
	tops      [7]int
	movements []byte
}

func (c *chamber) drop(rocks int) {
	iter := goutils.NewIterator(c.movements)
	si := 0
	for r := 0; r < rocks; r++ {
		s := newShape(si % 5)
		si++
		s.x, s.y = 2, c.top()+4+s.height
		for c.dist(s) > 1 {
			s.y--
			if !iter.Next() || iter.Value() == '\n' {
				iter.Reset()
				iter.Next()
			}
			c.push(&s, iter.Value())
		}
		c.put(s)
	}
}

func (c *chamber) print() {
	y := c.top()
	for y > 0 {
		for x := range c.cols {
			r := c.cols[x][y]
			if r == 0 {
				r = ' '
			}
			fmt.Printf("%c", r)
		}
		fmt.Println()
		y--
	}
}

func (c *chamber) dist(s shape) int {
	min := 1000000
	for x := 0; x < s.width; x++ {
		for y := s.height - 1; y >= 0; y-- {
			if !s.fields[y][x] {
				continue
			}
			d := 1
			for s.y-y-d > 0 && c.cols[s.x+x][s.y-y-d] == 0 {
				d++
			}
			if d < min {
				min = d
				if min <= 1 {
					return min
				}
			}
			break
		}
	}
	return min
}

func (c *chamber) put(s shape) {
	for y, line := range s.fields[:s.height] {
		for x, field := range line[:s.width] {
			if !field {
				continue
			}
			c.cols[s.x+x][s.y-y] = s.ind
			if c.tops[s.x+x] < s.y-y {
				c.tops[s.x+x] = s.y - y
			}
		}
	}
}

func (c *chamber) top() int {
	var max int
	for _, t := range c.tops {
		if t > max {
			max = t
		}
	}
	return max
}

func (c *chamber) push(s *shape, dir byte) {
	switch dir {
	case '<':
		if s.x > 0 {
			for y, line := range s.fields[:s.height] {
				for x, field := range line[:s.width] {
					if field && c.cols[s.x+x-1][s.y-y] != 0 {
						return
					}
				}
			}
			s.x--
		}
	case '>':
		if s.x+s.width < 7 {
			for y := range s.fields[:s.height] {
				for x := s.width - 1; x >= 0; x-- {
					if s.fields[y][x] && c.cols[s.x+x+1][s.y-y] != 0 {
						return
					}
				}
			}
			s.x++
		}
	}
}
