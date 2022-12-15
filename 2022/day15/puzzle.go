package main

import (
	"math"
	"strconv"

	"github.com/aibor/aoc/goutils"
)

var (
	exampleResult1 = "26"
	exampleResult2 = "56000011"

	result1 = "4873353"
	result2 = "11600823139120"
)

func part1(input string) string {
	var result int

	line := 2000000
	if len(input) < 1024 {
		line = 10
	}
	m := parseMap(input)
	result = m.noBeaconCount(line)

	return strconv.Itoa(result)
}

func part2(input string) string {
	var result int

	m := parseMap(input)
	max := 4000000
	if len(input) < 1024 {
		max = 20
	}
	p := m.noBeaconPos(0, max)
	result = p.x*4000000 + p.y

	return strconv.Itoa(result)
}

type pos struct {
	x, y int
}

func (p *pos) move(x, y int) {
	p.x += x
	p.y += y
}

func (p pos) manhattanDist(q pos) int {
	return int(math.Abs(float64(p.x-q.x)) + math.Abs(float64(p.y-q.y)))
}

type sensor struct {
	pos
	beacon pos
	dist   int
}

func (s *sensor) init() {
	s.dist = s.manhattanDist(s.beacon)
}

type Map struct {
	sensors []sensor
	minx    int
	maxx    int
}

func (m *Map) noBeaconCount(y int) int {
	var i int
	p := pos{m.minx, y}
	for p.x <= m.maxx {
		s := m.checkPos(p)
		add := 1
		if s != nil && s.beacon != p {
			// skip to the other side
			if s.x > p.x {
				add = (s.x - p.x) * 2
			}
			i += add
		}
		p.move(add, 0)
	}
	return i
}

func (m *Map) noBeaconPos(min, max int) pos {
	for _, s := range m.sensors {
		p := pos{s.x, s.y - s.dist - 1}
		validPos := func(p pos) bool {
			return p.x >= min && p.x <= max && p.y >= min && p.y <= max
		}
		for validPos(p) && p.y < s.y {
			if m.checkPos(p) == nil {
				return p
			}
			p.move(1, 1)
		}
		for validPos(p) && p.x > s.x {
			if m.checkPos(p) == nil {
				return p
			}
			p.move(-1, 1)
		}
		for validPos(p) && p.y > s.y {
			if m.checkPos(p) == nil {
				return p
			}
			p.move(-1, -1)
		}
		for validPos(p) && p.x < s.x {
			if m.checkPos(p) == nil {
				return p
			}
			p.move(1, -1)
		}
	}
	panic("not found")
}

func (m *Map) checkPos(p pos) *sensor {
	for _, s := range m.sensors {
		if s.manhattanDist(p) <= s.dist {
			return &s
		}
	}
	return nil
}

func parseMap(input string) Map {
	m := Map{
		sensors: make([]sensor, 0, 64),
	}
	iter := goutils.NewStringFieldsIterator(input)
	for iter.Skip(2) {
		var s sensor
		s.x = goutils.MustBeInt(iter.Value()[2 : len(iter.Value())-1])
		iter.Next()
		s.y = goutils.MustBeInt(iter.Value()[2 : len(iter.Value())-1])
		iter.Skip(5)
		s.beacon.x = goutils.MustBeInt(iter.Value()[2 : len(iter.Value())-1])
		iter.Next()
		s.beacon.y = goutils.MustBeInt(iter.Value()[2:])
		s.init()
		if m.minx == 0 || s.x-s.dist < m.minx {
			m.minx = s.x - s.dist
		}
		if m.maxx == 0 || s.x+s.dist > m.maxx {
			m.maxx = s.x + s.dist
		}
		m.sensors = append(m.sensors, s)
		iter.Next()
	}
	return m
}
