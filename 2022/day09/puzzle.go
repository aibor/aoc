package main

import (
	"fmt"
	"strconv"
	"strings"
)

var (
	exampleResult1 = "88"
	exampleResult2 = "36"

	result1 = "6181"
	result2 = "2386"
)

const size = 1024

func part1(input string) string {
	g := newGrid(size, 2)
	g.process(input)
	return strconv.Itoa(g.tailVisitedCount)
}

func part2(input string) string {
	g := newGrid(size, 10)
	g.process(input)
	return strconv.Itoa(g.tailVisitedCount)
}

type pos struct {
	x int
	y int
}

func (p *pos) is(x, y int) bool {
	return p.x == x && p.y == y
}

func (p *pos) distance(q pos) pos {
	return pos{p.x - q.x, p.y - q.y}
}

func (p *pos) move(d pos) {
	switch d.x {
	case 2, -2:
		p.x += d.x / 2
	case 1, -1:
		p.x += d.x
	}
	switch d.y {
	case 2, -2:
		p.y += d.y / 2
	case 1, -1:
		p.y += d.y
	}
}

type grid struct {
	size             int
	start            pos
	parts            []pos
	tailVisited      [][]bool
	tailVisitedCount int
	head             *pos
	tail             *pos
}

func newGrid(size int, length int) grid {
	g := grid{
		size:        size,
		start:       pos{size / 2, size / 2},
		parts:       make([]pos, length),
		tailVisited: make([][]bool, size),
	}
	for i := range g.parts {
		g.parts[i] = g.start
	}
	g.head = &g.parts[0]
	g.tail = &g.parts[length-1]
	g.markVisited()
	return g
}

func (g *grid) process(input string) {
	insts := strings.Fields(input)
	for len(insts) > 0 {
		n, _ := strconv.Atoi(insts[1])
		g.move(insts[0][0], n)
		insts = insts[2:]
	}
}

func (g *grid) move(dir byte, num int) {
	for ; num > 0; num-- {
		switch dir {
		case 'L':
			g.head.x--
		case 'R':
			g.head.x++
		case 'U':
			g.head.y--
		case 'D':
			g.head.y++
		}
		g.moveParts()
	}
}

func (g *grid) moveParts() {
	for i, c := range g.parts[1:] {
		d := g.parts[i].distance(c)
		if d.x <= 1 && d.x >= -1 && d.y <= 1 && d.y >= -1 {
			return
		}
		g.parts[i+1].move(d)
	}
	g.markVisited()
}

func (g *grid) markVisited() {
	if g.tailVisited[g.tail.y] == nil {
		g.tailVisited[g.tail.y] = make([]bool, g.size, g.size)
	}
	if g.tailVisited[g.tail.y][g.tail.x] == false {
		g.tailVisitedCount++
		g.tailVisited[g.tail.y][g.tail.x] = true
	}
}

func (g *grid) PrintVisited() {
	for y, line := range g.tailVisited {
		if line == nil {
			continue
		}
		for x, b := range line {
			s := "."
			switch {
			case g.start.is(x, y):
				s = "s"
			case b:
				s = "#"
			}
			fmt.Printf(s)
		}
		fmt.Println()
	}
}

func (g *grid) PrintState() {
	for y := 0; y < g.size; y++ {
		for x := 0; x < g.size; x++ {
			s := "."
			if g.start.is(x, y) {
				s = "s"
			}
			for n, p := range g.parts {
				if p.is(x, y) {
					s = strconv.Itoa(n)
					if n == 0 {
						s = "H"
					}
					break
				}
			}
			fmt.Print(s)
		}
		fmt.Println()
	}
	fmt.Println()
}
