package main

import (
	"fmt"
	"strconv"

	"github.com/aibor/aoc/goutils"
)

var (
	exampleResult1 = "95437"
	exampleResult2 = "24933642"

	result1 = "1915606"
	result2 = "5025657"
)

func part1(input string) string {
	var result int
	dir := parse(input)

	dir.walk(func(d *Dir) {
		if d.size <= 100000 {
			result += d.size
		}
	})

	return strconv.Itoa(result)
}

func part2(input string) string {
	dir := parse(input)
	unused := 70000000 - dir.size

	dir.walk(func(d *Dir) {
		if unused+d.size >= 30000000 && d.size < dir.size {
			dir = d
		}
	})

	return strconv.Itoa(dir.size)
}

func parse(input string) *Dir {
	var dir *Dir
	iter := goutils.NewStringFieldsIterator(input)

	for iter.Next() {
		if iter.Value() == "cd" {
			iter.Next()
			if iter.Value() == ".." {
				dir = dir.parent
			} else {
				dir = dir.addSub(iter.Value())
			}
		} else if iter.Value()[0] >= '0' && iter.Value()[0] <= '9' {
			dir.addSize(goutils.MustBeInt(iter.Value()))
			// skip file name
			iter.Next()
		}
	}
	return dir.root()
}

type Dir struct {
	size    int
	name    string
	parent  *Dir
	subDirs []*Dir
}

func (d *Dir) addSub(name string) *Dir {
	n := &Dir{
		name:   name,
		parent: d,
	}
	if d != nil {
		d.subDirs = append(d.subDirs, n)
	}
	return n
}

func (d *Dir) print() {
	var prefix string
	d.walk(func(d *Dir) {
		fmt.Println(d.name, d.size)
		prefix += "-"
	})
}

func (d *Dir) root() *Dir {
	if d.parent == nil {
		return d
	}
	return d.parent.root()
}

func (d *Dir) addSize(size int) {
	d.size += size
	if d.parent != nil {
		d.parent.addSize(size)
	}
}

func (d *Dir) walk(f func(*Dir)) {
	f(d)
	for _, c := range d.subDirs {
		c.walk(f)
	}
}
