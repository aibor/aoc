package main

import (
	_ "embed"
	"testing"
)

//go:embed example
var example string

//go:embed input
var input string

func TestPart1(t *testing.T) {
	if r := part1(example); r != exampleResult1 {
		t.Fatalf("example: got %s\n", r)
	}
	if r := part1(input); r != result1 {
		t.Fatalf("input: got %s\n", r)
	}
}

func TestPart2(t *testing.T) {
	if r := part2(example); r != exampleResult2 {
		t.Fatalf("example: got %s\n", r)
	}
	if r := part2(input); r != result2 {
		t.Fatalf("input: got %s\n", r)
	}
}
