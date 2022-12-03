#!/bin/bash

set -euo pipefail

declare -i day=${1:?give day number}

if [[ $day -lt 1 || $day -gt 24 ]]
then
  echo "invalid day number: $day"
  exit 1
fi

basedir=$(dirname "${BASH_SOURCE[0]}")
dir=$(printf "%s/day%02d" "$basedir" "$day")

mkdir "$dir"

for f in input example
do
  touch "$dir/$f"
done

cat > "$dir/puzzle.go" <<EOF
package main

import (
	"fmt"

	"github.com/aibor/aoc/goutils"
)

var (
	exampleResult1 = ""
	exampleResult2 = ""

	result1 = ""
	result2 = ""
)

func part1(input string) string {
	var result int

	for _, line := range goutils.SplitInput(input) {
		_ = line
	}

	return fmt.Sprintf("%d", result)
}

func part2(input string) string {
	var result int

	for _, line := range goutils.SplitInput(input) {
		_ = line
	}

	return fmt.Sprintf("%d", result)
}
EOF

cat > "$dir/puzzle_test.go" <<EOF
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
EOF
