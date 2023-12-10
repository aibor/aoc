package main

import (
	"strconv"
	"strings"

	"github.com/aibor/aoc/goutils"
)

func part1(input string) string {
	var result int

	a := parseArea(input)
	result = a.walkPipe() / 2

	return strconv.Itoa(result)
}

func part2(input string) string {
	var result int

	for _, line := range goutils.SplitInput(input) {
		_ = line
	}

	return strconv.Itoa(result)
}

type tile struct {
	line, col int
}

func (t *tile) up() tile {
	return tile{t.line - 1, t.col}
}

func (t *tile) down() tile {
	return tile{t.line + 1, t.col}
}

func (t *tile) left() tile {
	return tile{t.line, t.col - 1}
}

func (t *tile) right() tile {
	return tile{t.line, t.col + 1}
}

type area struct {
	tiles [][]rune
	start tile
}

func parseArea(input string) *area {
	var a area
	a.tiles = make([][]rune, 0, 256)
	for lineNum, line := range goutils.SplitInput(input) {
		a.tiles = append(a.tiles, []rune(line))
		if s := strings.IndexRune(line, 'S'); s != -1 {
			a.start = tile{lineNum, s}
		}
	}
	return &a
}

func (a *area) lookup(t tile) rune {
	if t.line < 0 || t.line >= len(a.tiles) {
		return 0
	}
	if t.col < 0 || t.col >= len(a.tiles[0]) {
		return 0
	}
	return a.tiles[t.line][t.col]
}

type moveDir int

const (
	moveNone moveDir = iota
	moveUp
	moveRight
	moveDown
	moveLeft
)

func (a *area) findNext(t tile) moveDir {
	if n := a.lookup(t.up()); n == '|' || n == 'F' || n == '7' {
		return moveUp
	}
	if n := a.lookup(t.right()); n == '-' || n == '7' || n == 'J' {
		return moveRight
	}
	if n := a.lookup(t.down()); n == '|' || n == 'J' || n == 'L' {
		return moveDown
	}
	if n := a.lookup(t.left()); n == '-' || n == 'L' || n == 'F' {
		return moveDown
	}
	return moveNone
}

func (a *area) walkPipe() (steps int) {
	move := a.findNext(a.start)
	cur := a.start

	for move != moveNone {
		steps++
		switch move {
		case moveUp:
			cur = cur.up()
		case moveRight:
			cur = cur.right()
		case moveDown:
			cur = cur.down()
		case moveLeft:
			cur = cur.left()
		}
		switch a.lookup(cur) {
		case 'F':
			switch move {
			case moveUp:
				move = moveRight
			case moveLeft:
				move = moveDown
			}
		case '7':
			switch move {
			case moveUp:
				move = moveLeft
			case moveRight:
				move = moveDown
			}
		case 'L':
			switch move {
			case moveDown:
				move = moveRight
			case moveLeft:
				move = moveUp
			}
		case 'J':
			switch move {
			case moveDown:
				move = moveLeft
			case moveRight:
				move = moveUp
			}
		case 'S':
			move = moveNone
		}
	}
	return
}
