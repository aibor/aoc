package main

import (
	"testing"

	"github.com/aibor/aoc/goutils"
)

var input = goutils.SplitInput(`
1000
2000
3000

4000

5000
6000

7000
8000
9000

10000
`)

func TestPart1(t *testing.T) {
	if r := part01(input); r != "24000" {
		t.Fatalf("got %s\n", r)
	}
}

func TestPart2(t *testing.T) {
	if r := part02(input); r != "45000" {
		t.Fatalf("got %s\n", r)
	}
}
