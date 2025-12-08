package main

import (
	"strings"
)

func day8input(io *IO) ([][3]int64, [][2]int) {
	var line string
	pts := make([][3]int64, 0)
	for io.Readln(&line) == nil {
		segs := strings.Split(line, ",")
		pts = append(pts, [3]int64{stoll(segs[0]), stoll(segs[1]), stoll(segs[2])})
	}
	pairs := make([][2]int, len(pts)*(len(pts)-1)/2)
	k := 0
	for i := 0; i < len(pts); i++ {
		for j := i + 1; j < len(pts); j++ {
			pairs[k] = [2]int{i, j}
			k++
		}
	}
	order := make([]Pair[int, int64], len(pairs))
	for i := 0; i < len(pairs); i++ {
		dx := pts[pairs[i][0]][0] - pts[pairs[i][1]][0]
		dy := pts[pairs[i][0]][1] - pts[pairs[i][1]][1]
		dz := pts[pairs[i][0]][2] - pts[pairs[i][1]][2]
		order[i] = Pair[int, int64]{i, dx*dx + dy*dy + dz*dz}
	}
	order = QSort(order)
	sortedPairs := make([][2]int, len(pairs))
	for i := 0; i < len(order); i++ {
		sortedPairs[i] = pairs[order[i].A]
	}

	return pts, sortedPairs
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
