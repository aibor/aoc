package main

import (
	_ "embed"
	"strings"
	"testing"
)

var (
	//go:embed example
	example_raw string

	//go:embed input
	input_raw string

	example = strings.TrimSpace(example_raw)
	input   = strings.TrimSpace(input_raw)
)

func assert(t *testing.T, result, expectation string) {
	t.Helper()

	if result != expectation {
		t.Fatalf("got %s\n", result)
	}
}

func TestPart1Example(t *testing.T) {
	assert(t, part1(example_raw), exampleResult1)
}

func TestPart2Example(t *testing.T) {
	assert(t, part2(example_raw), exampleResult2)
}

func TestPart1Input(t *testing.T) {
	assert(t, part1(input_raw), result1)
}

func TestPart2Input(t *testing.T) {
	assert(t, part2(input_raw), result2)
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
