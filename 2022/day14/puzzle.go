package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/aibor/aoc/goutils"
)

var (
	exampleResult1 = "24"
	exampleResult2 = "93"

	result1 = "768"
	result2 = "26686"
)

const sizex = 250
const sizey = 170

func part1(input string) string {
	var result int

	m := parseMap(input)
	p := pos{sizex, 0}
	for m.dropSand(p) {
		result++
	}

	return strconv.Itoa(result)
}

func part2(input string) string {
	var result int

	m := parseMap(input)
	m.maxy += 2
	for i := range m.fields {
		m.fields[i][m.maxy] = 1
	}

	p := pos{sizex, 0}
	for m.fields[p.x][p.y] == 0 && m.dropSand(p) {
		result++
	}

	return strconv.Itoa(result)
}

type pos struct {
	x, y int
}

type Map struct {
	fields [sizex*2 + 1][sizey]uint8
	maxy   int
}

func (m *Map) Print() {
	for y := 0; y < sizey; y++ {
		for x := 0; x <= sizex*2; x++ {
			if y == 0 && x == sizex {
				fmt.Printf("+")
			} else if m.fields[x][y] == 1 {
				fmt.Printf("#")
			} else if m.fields[x][y] == 2 {
				fmt.Printf("o")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Println()
	}
}

func parseMap(input string) Map {
	var m Map
	iter := goutils.NewStringFieldsIterator(input)
	split := func(token string) pos {
		p := strings.Index(token, ",")
		return pos{
			goutils.MustBeInt(token[:p]) - (500 - sizex),
			goutils.MustBeInt(token[p+1:]),
		}
	}
	for iter.Next() {
		pos1 := split(iter.Value())
		if !iter.Next() {
			break
		}
		if iter.Value() != "->" {
			iter.Prev()
			continue
		}
		if !iter.Next() {
			break
		}
		pos2 := split(iter.Value())
		iter.Prev()
		if m.maxy == 0 || pos1.y > m.maxy {
			m.maxy = pos1.y
		}
		if pos2.y > m.maxy {
			m.maxy = pos2.y
		}
		var d pos
		switch {
		case pos1.x == pos2.x:
			d.y = pos2.y - pos1.y
		case pos1.y == pos2.y:
			d.x = pos2.x - pos1.x
		}
		m.fields[pos1.x][pos1.y] = 1
		for d.x != 0 || d.y != 0 {
			m.fields[pos1.x+d.x][pos1.y+d.y] = 1
			switch {
			case d.x > 0:
				d.x--
			case d.x < 0:
				d.x++
			case d.y > 0:
				d.y--
			case d.y < 0:
				d.y++
			}
		}
	}
	return m
}

func (m *Map) dropSand(p pos) bool {
	if p.y >= m.maxy+1 {
		return false
	} else if m.fields[p.x][p.y+1] == 0 {
		p.y++
	} else if m.fields[p.x-1][p.y+1] == 0 {
		p.x--
		p.y++
	} else if m.fields[p.x+1][p.y+1] == 0 {
		p.x++
		p.y++
	} else {
		m.fields[p.x][p.y] = 2
		return true
	}
	return m.dropSand(p)
}
