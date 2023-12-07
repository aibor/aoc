package main

import (
	"strconv"
	"strings"
	"sync"

	"github.com/aibor/aoc/goutils"
)

func part1(input string) string {
	var result int

	lines := goutils.SplitInput(input)
	alm := parseAlmanac(lines[2:])
	for _, seed := range strings.Split(lines[0], " ")[1:] {
		value := alm.run(goutils.MustBeInt(seed))
		if result == 0 || value < result {
			result = value
		}
	}

	return strconv.Itoa(result)
}

func part2(input string) string {
	var (
		result int
		mu     sync.Mutex
		wg     sync.WaitGroup
	)

	lines := goutils.SplitInput(input)
	alm := parseAlmanac(lines[2:])
	i := goutils.NewStringFieldsIterator(lines[0])
	i.Next()
	for i.Next() {
		start := i.MustBeInt()
		i.Next()
		length := i.MustBeInt()
		wg.Add(1)
		go func(start, length int) {
			defer wg.Done()
			for i := start; i < start+length; i++ {
				value := alm.run(i)
				mu.Lock()

				if result == 0 || value < result {

					result = value
				}
				mu.Unlock()
			}
		}(start, length)
	}

	wg.Wait()

	return strconv.Itoa(result)
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

func (m *valueMap) run(value int) int {
	for _, r := range m.ranges {
		if value < r.sourceStart {
			continue
		}
		if value >= r.sourceStart+r.length {
			continue
		}
		value = r.destinationStart + (value - r.sourceStart)
		break
	}
	return value
}

type almanac []*valueMap

func (a *almanac) run(value int) int {
	for _, m := range *a {
		value = m.run(value)
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
