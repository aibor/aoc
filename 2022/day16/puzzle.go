package main

import (
	"sort"
	"strconv"

	"github.com/aibor/aoc/goutils"
)

var (
	exampleResult1 = "1651"
	//exampleResult2 = "1707"
	exampleResult2 = "0"

	result1 = "2029"
	result2 = "0"
)

func part1(input string) string {
	var result int

	t := parseNodes(input)

	c := t["AA"]
	result = t.bestPath(c, []string{}, 30, 0)

	return strconv.Itoa(result)
}

func part2(input string) string {
	var result int

	//	t := parseNodes(input)

	return strconv.Itoa(result)
}

type node struct {
	name    string
	tunnels []*node
	rate    int
	dists   map[string]int
}

func (n *node) dist(d *node, visited []string, di int) int {
	if d == n {
		return di
	}
	if di == 0 {
		if c, e := n.dists[d.name]; e {
			return c
		}
	}
	if visited == nil {
		visited = []string{n.name}
	}
	min := 10000
main:
	for _, m := range n.tunnels {
		for _, e := range visited {
			if e == m.name {
				continue main
			}
		}
		r := m.dist(d, append(visited, m.name), di+1)
		if r < min {
			min = r
		}
	}
	if di == 0 {
		n.dists[d.name] = min
	}
	return min
}

func newNode(name string, n *node) *node {
	nn := node{
		name:    name,
		tunnels: make([]*node, 0, 64),
		dists:   make(map[string]int, 256),
	}
	if n != nil {
		nn.tunnels = append(nn.tunnels, n)
	}

	return &nn
}

type nodes map[string]*node

func (ns *nodes) get(name string) *node {
	n, exist := (*ns)[name]
	if !exist {
		n = newNode(name, n)
		(*ns)[name] = n
	}
	return n
}

func (ns *nodes) bestPath(start *node, opened []string, t int, v int) int {
	nr := make([]*node, 0, 64)
	for m, e := range *ns {
		var open bool
		for _, o := range opened {
			if m == o {
				open = true
			}
		}
		if !open && e.rate > 0 {
			nr = append(nr, e)
		}
	}
	if len(nr) < 1 || t < 1 {
		return v
	}
	sort.SliceStable(nr, func(i, j int) bool {
		return nr[i].rate > nr[j].rate
	})

	var max int
	for _, e := range nr {
		u := t - start.dist(e, nil, 0)
		u--
		nv := ns.bestPath(e, append(opened, e.name), u, v+u*e.rate)
		if nv > max {
			max = nv
		}
	}
	return max
}

func parseNodes(input string) nodes {
	var n *node
	ns := make(nodes, 128)
	iter := goutils.NewStringFieldsIterator(input)
	for iter.Next() {
		iter.Next()
		n = ns.get(iter.Value())
		iter.Skip(3)
		n.rate = goutils.MustBeInt(iter.Value()[5 : len(iter.Value())-1])
		iter.Skip(4)
		for iter.Next() {
			n.tunnels = append(n.tunnels, ns.get(iter.Value()[:2]))
			if len(iter.Value()) == 2 {
				break
			}
		}
	}
	return ns
}
