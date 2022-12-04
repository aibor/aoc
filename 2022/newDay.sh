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

mkdir -v "$dir"
ln -s ../_puzzle_test.go "$dir/puzzle_test.go"

for f in input example
do
  touch "$dir/$f"
done

cat > "$dir/puzzle.go" <<EOF
package main

import (
	"strconv"
	"strings"
)

var (
	exampleResult1 = "0"
	exampleResult2 = "0"

	result1 = "0"
	result2 = "0"
)

func part1(input string) string {
	var result int

	for _, line := range strings.Split(input, "\n") {
		_ = line
	}

	return strconv.Itoa(result)
}

func part2(input string) string {
	var result int

	for _, line := range strings.Split(input, "\n") {
		_ = line
	}

	return strconv.Itoa(result)
}
EOF
