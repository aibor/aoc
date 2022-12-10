package main

import (
	"strconv"
	"strings"
)

var (
	exampleResult1 = "13140"
	exampleResult2 = `
##..##..##..##..##..##..##..##..##..##..
###...###...###...###...###...###...###.
####....####....####....####....####....
#####.....#####.....#####.....#####.....
######......######......######......####
#######.......#######.......#######.....`

	result1 = "14720"
	result2 = `
####.####.###..###..###..####.####.####.
#.......#.#..#.#..#.#..#.#.......#.#....
###....#..###..#..#.###..###....#..###..
#.....#...#..#.###..#..#.#.....#...#....
#....#....#..#.#....#..#.#....#....#....
#....####.###..#....###..#....####.#....`
)

func part1(input string) string {
	var result int

	c := cpu{x: 1}
	c.eachTick = func() {
		switch c.cycle {
		case 20, 60, 100, 140, 180, 220:
			result += c.cycle * c.x
		}
	}
	c.process(input)

	return strconv.Itoa(result)
}

func part2(input string) string {
	s := screen{
		width: 40,
		rows:  6,
		cpu:   cpu{x: 1},
	}
	s.cpu.eachTick = func() {
		char := '.'
		if s.drawing() {
			char = '#'
		}
		s.draw(byte(char))
	}
	s.cpu.process(input)

	return "\n" + s.String()
}

type instIterator []string

func (i *instIterator) next() bool {
	if len(*i) > 0 {
		*i = (*i)[1:]
		return true
	}
	return false
}

func (i *instIterator) value() string {
	return (*i)[0]
}

func (c *cpu) process(input string) {
	i := instIterator(strings.Fields(input))
	for len(i) > 0 {
		switch i.value() {
		case "noop":
			c.tick()
		case "addx":
			i.next()
			x, _ := strconv.Atoi(i.value())
			c.addx(x)
		}
		if !i.next() {
			break
		}
	}
}

type cpu struct {
	x        int
	cycle    int
	eachTick func()
}

func (c *cpu) tick() {
	c.cycle++
	if c.eachTick != nil {
		c.eachTick()
	}
}

func (c *cpu) addx(x int) {
	c.tick()
	c.tick()
	c.x += x
}

type screen struct {
	width int
	rows  int
	state []byte
	cpu   cpu
}

func (s *screen) sprite() int {
	return s.cpu.x
}

func (s *screen) index() int {
	return s.cpu.cycle - 1
}

func (s *screen) pixel() int {
	return s.index() % s.width
}

func (s *screen) row() int {
	return s.index() / s.width
}

func (s *screen) drawing() bool {
	diff := s.sprite() - s.pixel()
	return diff >= -1 && diff <= 1
}

func (s *screen) draw(char byte) {
	if s.state == nil {
		s.state = make([]byte, s.width*s.rows)
	}
	s.state[s.index()] = char
}

func (s *screen) String() string {
	out := make([]string, s.rows)
	for i := 0; len(s.state) > i*s.width && i < s.rows; i++ {
		out[i] = string(s.state[i*s.width : (i+1)*s.width])
	}
	return strings.Join(out, "\n")
}
