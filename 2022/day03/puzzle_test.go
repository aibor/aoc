package main

import (
	_ "embed"
	"testing"
)

//go:embed example
var example string

//go:embed input
var input string

func assert(t *testing.T, result, expectation string) {
	t.Helper()

	if result != expectation {
		t.Fatalf("got %s\n", result)
	}
}

func TestPart1Example(t *testing.T) {
	assert(t, part1(example), exampleResult1)
}

func TestPart2Example(t *testing.T) {
	assert(t, part2(example), exampleResult2)
}

func TestPart1Input(t *testing.T) {
	assert(t, part1(input), result1)
}

func TestPart2Input(t *testing.T) {
	assert(t, part2(input), result2)
}

func BenchmarkPart1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		part1(input)
	}
}

func BenchmarkPart2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		part2(input)
	}
}
