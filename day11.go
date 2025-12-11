package main

import (
	"slices"
	"strings"
)

func day11input(io *IO) ([][]int, []int, map[string]int) {
	var line string
	adj := make([][]int, 0)
	nameMap := make(map[string]int)
	indeg := make([]int, 0)
	getId := func(name string) int {
		id, ok := nameMap[name]
		if !ok {
			id = len(nameMap)
			nameMap[name] = id
			adj = append(adj, make([]int, 0))
			indeg = append(indeg, 0)
		}
		return id
	}
	for io.Readln(&line) == nil {
		s := strings.Index(line, ":")
		src := getId(line[:s])
		for _, dstName := range strings.Split(line[s+2:], " ") {
			dst := getId(dstName)
			adj[src] = append(adj[src], dst)
			indeg[dst]++
		}
	}
	order := make([]int, 0, len(adj))
	for i := 0; i < len(indeg); i++ {
		if indeg[i] == 0 {
			order = append(order, i)
		}
	}
	for i := 0; i < len(order); i++ {
		curr := order[i]
		for _, next := range adj[curr] {
			indeg[next]--
			if indeg[next] == 0 {
				order = append(order, next)
			}
		}
	}
	return adj, order, nameMap
}

// counts paths from src to dst
func day11paths(adj [][]int, topo []int, src, dst int) int64 {
	paths := make([]int64, len(adj))
	paths[src] = 1
	for _, curr := range topo {
		if paths[curr] == 0 {
			continue
		}
		for _, next := range adj[curr] {
			paths[next] += paths[curr]
		}
	}
	return paths[dst]
}

func Day11A(io *IO) {
	adj, topo, nameMap := day11input(io)
	io.Write("%d\n", day11paths(adj, topo, nameMap["you"], nameMap["out"]))
}

func Day11B(io *IO) {
	adj, topo, nameMap := day11input(io)
	svr := nameMap["svr"]
	a := nameMap["dac"]
	b := nameMap["fft"]
	out := nameMap["out"]
	if slices.Index(topo, a) > slices.Index(topo, b) {
		a, b = b, a
	}
	io.Write("%d\n", day11paths(adj, topo, svr, a)*day11paths(adj, topo, a, b)*day11paths(adj, topo, b, out))
}
