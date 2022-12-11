package main

import (
	"strconv"
	"strings"

	"github.com/aibor/aoc/goutils"
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

	s := makeScreen()
	s.cpu.eachTick = func() {
		switch s.cpu.cycle {
		case 20, 60, 100, 140, 180, 220:
			result += s.signalStrength()
		}
	}
	s.process(input)

	return strconv.Itoa(result)
}

func part2(input string) string {
	s := makeScreen()
	s.cpu.eachTick = func() {
		char := byte('.')
		if s.drawing() {
			char = byte('#')
		}
		s.draw(char)
	}
	s.process(input)

	return "\n" + s.String()
}

func (s *screen) process(input string) {
	i := goutils.NewStringFieldsIterator(input)
	for i.Next() {
		switch i.Value() {
		case "noop":
			s.cpu.tick()
		case "addx":
			i.Next()
			s.cpu.addx(goutils.MustBeInt(i.Value()))
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

func makeScreen() screen {
	s := screen{
		width: 40,
		rows:  6,
		cpu:   cpu{x: 1},
	}
	s.state = make([]byte, s.width*s.rows)
	return s
}

func (s *screen) signalStrength() int {
	return s.cpu.cycle * s.cpu.x
}

func (s *screen) spritePos() int {
	return s.cpu.x
}

func (s *screen) stateIndex() int {
	if s.cpu.cycle == 0 {
		return 0
	}
	return s.cpu.cycle - 1
}

func (s *screen) pixelPos() int {
	return s.stateIndex() % s.width
}

func (s *screen) currentRow() int {
	return s.stateIndex() / s.width
}

func (s *screen) drawing() bool {
	diff := s.spritePos() - s.pixelPos()
	return diff >= -1 && diff <= 1
}

func (s *screen) draw(char byte) {
	s.state[s.stateIndex()] = char
}

func (s *screen) String() string {
	var b strings.Builder
	b.Grow((s.width + 1) * s.rows)
	for i := 0; i < s.rows; i++ {
		if i > 0 {
			b.WriteString("\n")
		}
		b.Write(s.state[i*s.width : (i+1)*s.width])
	}
	return b.String()
}
