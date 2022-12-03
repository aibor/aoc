package main

import (
	"testing"

	"github.com/aibor/aoc/goutils"
)

var input = goutils.SplitInput(`
vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw
`)

func TestPart1(t *testing.T) {
	if r := part01(input); r != "157" {
		t.Fatalf("got %s\n", r)
	}
}

func TestPart2(t *testing.T) {
	if r := part02(input); r != "70" {
		t.Fatalf("got %s\n", r)
	}
}
