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

func part1(input string) string {
	g := newGrid(512, 2)
	g.process(input)
	return strconv.Itoa(g.tailVisitedCount)
}

func part2(input string) string {
	g := newGrid(512, 10)
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

type grid struct {
	size             int
	start            pos
	parts            []pos
	tailVisited      [][]bool
	tailVisitedCount int
}

func newGrid(start int, length int) grid {
	g := grid{
		size:        start * 2,
		start:       pos{start, start},
		parts:       make([]pos, length, 16),
		tailVisited: make([][]bool, start*2),
	}
	for i := range g.parts {
		g.parts[i] = g.start
	}
	g.markVisited(g.parts[0])
	return g
}

func (g *grid) process(input string) {
	insts := strings.Fields(input)
	for len(insts) > 0 {
		g.move(insts[0], insts[1])
		insts = insts[2:]
	}
}

func (g *grid) move(dir string, num string) {
	for n, _ := strconv.Atoi(num); n > 0; n-- {
		head := &g.parts[0]
		switch dir[0] {
		case 'L':
			head.x--
		case 'R':
			head.x++
		case 'U':
			head.y--
		case 'D':
			head.y++
		}
		for i, c := range g.parts[1:] {
			p := &g.parts[i+1]
			d := g.parts[i].distance(c)
			if d.x <= 1 && d.x >= -1 && d.y <= 1 && d.y >= -1 {
				continue
			}
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
			if i == len(g.parts)-2 {
				g.markVisited(*p)
			}
		}
	}
}

func (g *grid) markVisited(p pos) {
	if g.tailVisited[p.y] == nil {
		g.tailVisited[p.y] = make([]bool, g.size, g.size)
	}
	if g.tailVisited[p.y][p.x] == false {
		g.tailVisitedCount++
		g.tailVisited[p.y][p.x] = true
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
