package main

import (
	"strconv"
	"strings"

	"github.com/aibor/aoc/goutils"
)

func part1(input string) string {
	var result int

	lines := goutils.SplitInput(input)
	alm := parseAlmanac(lines[2:])

	// Run each seed through the algo as described.
	for _, seed := range strings.Split(lines[0], " ")[1:] {
		value := alm.run(goutils.MustBeInt(seed))
		if result == 0 || value < result {
			result = value
		}
	}

	return strconv.Itoa(result)
}

func part2(input string) string {
	var seeds []valueRange

	lines := goutils.SplitInput(input)
	alm := parseAlmanac(lines[2:])

	// Parse seed ranges.
	i := goutils.NewStringFieldsIterator(lines[0])
	i.Next()
	for i.Next() {
		rg := valueRange{start: i.MustBeInt()}
		i.Next()
		rg.length = i.MustBeInt()
		seeds = append(seeds, rg)
	}

	// Run the algo reverse, recursively probing if the location maps to a
	// value in a seed range.
	for i := 0; i < 1<<32; i++ {
		value := alm.runReverse(i)
		for _, seed := range seeds {
			if seed.includes(value) {
				return strconv.Itoa(i)
			}
		}
	}

	return ""
}

type valueRange struct {
	start  int
	length int
}

func (r *valueRange) includes(value int) bool {
	return value >= r.start && value < r.start+r.length
}

type mapRange struct {
	length           int
	sourceStart      int
	destinationStart int
}

type valueMap struct {
	name   string
	ranges []mapRange
}

func (m *valueMap) run(value int, reverse bool) int {
	for _, r := range m.ranges {
		var inStart, outStart int
		if reverse {
			inStart = r.destinationStart
			outStart = r.sourceStart
		} else {
			inStart = r.sourceStart
			outStart = r.destinationStart
		}

		rg := valueRange{inStart, r.length}
		if rg.includes(value) {
			value = outStart + (value - inStart)
			break
		}
	}
	return value
}

type almanac []*valueMap

func (a *almanac) run(value int) int {
	for _, m := range *a {
		value = m.run(value, false)
	}
	return value
}

func (a *almanac) runReverse(value int) int {
	for i := len(*a) - 1; i >= 0; i-- {
		value = (*a)[i].run(value, true)
	}
	return value
}

func parseAlmanac(input []string) almanac {
	var curMap *valueMap
	var alm almanac

	for _, line := range input {
		switch {
		case line == "":
		case strings.HasSuffix(line, ":"):
			curMap = &valueMap{
				name: line[:len(line)-1],
			}
			alm = append(alm, curMap)
		default:
			fields := strings.Split(line, " ")
			curMap.ranges = append(curMap.ranges, mapRange{
				goutils.MustBeInt(fields[2]),
				goutils.MustBeInt(fields[1]),
				goutils.MustBeInt(fields[0]),
			})
		}
	}

	return alm
}
