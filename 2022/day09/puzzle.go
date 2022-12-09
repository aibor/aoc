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
	insts := strings.Fields(input)
	for len(insts) > 0 {
		g.move(insts[0], insts[1])
		insts = insts[2:]
	}
	return strconv.Itoa(g.tailVisitedCount)
}

func part2(input string) string {
	g := newGrid(512, 10)
	insts := strings.Fields(input)
	for len(insts) > 0 {
		g.move(insts[0], insts[1])
		insts = insts[2:]
	}
	return strconv.Itoa(g.tailVisitedCount)
}

type grid struct {
	size             int
	start            pos
	parts            []pos
	tailVisitedCount int
	tailVisited      [][]bool
}

func (g *grid) PrintVisited() {
	for y, line := range g.tailVisited {
		if line == nil {
			continue
		}
		for x, b := range line {
			if g.start.is(x, y) {
				fmt.Print("s")
			} else if b {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Println()
	}
}

func (g *grid) PrintState() {
	for y := 0; y < g.size; y++ {
		for x := 0; x < g.size; x++ {
			s := -1
			for n, p := range g.parts {
				if p.is(x, y) {
					s = n
					break
				}
			}
			if s == -1 {
				if g.start.is(x, y) {
					fmt.Print("s")
				} else {
					fmt.Print(".")
				}
			} else if s == 0 {
				fmt.Print("H")
			} else {
				fmt.Print(s)
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func newGrid(start int, length int) grid {
	g := grid{
		start:       pos{start, start},
		size:        start * 2,
		parts:       make([]pos, length),
		tailVisited: make([][]bool, start*2),
	}
	for i := range g.parts {
		g.parts[i] = g.start
	}
	g.markVisited()

	return g
}

func (g *grid) move(dir string, num string) {
	for n, _ := strconv.Atoi(num); n > 0; n-- {
		switch dir {
		case "R":
			g.parts[0].x++
		case "L":
			g.parts[0].x--
		case "U":
			g.parts[0].y--
		case "D":
			g.parts[0].y++
		}
		for i, c := range g.parts[1:] {
			b := g.parts[i]
			d := b.distance(c)
			if d.x > 1 || d.x < -1 {
				g.parts[i+1].x += d.x / 2
				if d.y > 1 || d.y < -1 {
					g.parts[i+1].y += d.y / 2
				} else {
					g.parts[i+1].y += d.y
				}
			} else if d.y > 1 || d.y < -1 {
				g.parts[i+1].y += d.y / 2
				if d.x > 1 || d.x < -1 {
					g.parts[i+1].x += d.x / 2
				} else {
					g.parts[i+1].x += d.x
				}
			}
		}
		g.markVisited()
	}
}

func (g *grid) markVisited() {
	p := g.parts[len(g.parts)-1]
	if g.tailVisited[p.y] == nil {
		g.tailVisited[p.y] = make([]bool, len(g.tailVisited))
	}
	if g.tailVisited[p.y][p.x] == false {
		g.tailVisitedCount++
		g.tailVisited[p.y][p.x] = true
	}
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
