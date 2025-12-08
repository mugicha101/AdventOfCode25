package main

import (
	"sort"
	"strings"
)

func day8input(io *IO) ([][3]int64, [][3]int64) {
	var line string
	pts := make([][3]int64, 0)
	for io.Readln(&line) == nil {
		segs := strings.Split(line, ",")
		pts = append(pts, [3]int64{stoll(segs[0]), stoll(segs[1]), stoll(segs[2])})
	}
	pairs := make([][3]int64, 0, len(pts)*(len(pts)-1)/2)
	for i := 0; i < len(pts); i++ {
		for j := i + 1; j < len(pts); j++ {
			dx := abs(pts[i][0] - pts[j][0])
			dy := abs(pts[i][1] - pts[j][1])
			dz := abs(pts[i][2] - pts[j][2])
			pairs = append(pairs, [3]int64{int64(i), int64(j), dx*dx + dy*dy + dz*dz})
		}
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][2] < pairs[j][2]
	})
	return pts, pairs
}

func Day8A(io *IO) {
	pts, pairs := day8input(io)
	uf := NewUnionFind(len(pts))
	for i := 0; i < 1000; i++ {
		a, b := int(pairs[i][0]), int(pairs[i][1])
		uf.Merge(a, b)
	}
	top := make(MinPriorityQueue[int], 0)
	for i := 0; i < uf.Len(); i++ {
		if uf.Find(i) != i {
			continue
		}
		sz := uf.CompSize(i)
		top.Push(sz)
		if len(top) > 3 {
			top.Pop()
		}
	}
	res := 1
	for _, x := range top {
		res *= x
	}
	io.Write("%d\n", res)
}

func Day8B(io *IO) {
	pts, pairs := day8input(io)
	uf := NewUnionFind(len(pts))
	i := -1
	for uf.NumComps() > 1 {
		i++
		uf.Merge(int(pairs[i][0]), int(pairs[i][1]))
	}
	io.Write("%d\n", pts[pairs[i][0]][0]*pts[pairs[i][1]][0])
}
