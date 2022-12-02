package main

import (
	"testing"

	"github.com/aibor/aoc/goutils"
)

var input = goutils.SplitInput(`
A Y
B X
C Z
`)

func TestPart1(t *testing.T) {
	if r := part01(input); r != "15" {
		t.Fatalf("got %s\n", r)
	}
}

func TestPart2(t *testing.T) {
	if r := part02(input); r != "12" {
		t.Fatalf("got %s\n", r)
	}
}
