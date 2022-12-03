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
EOF
