package main

import (
	"math"
	"strings"
)

func day8input(io *IO) ([][3]int64, int64) {
	var line string
	pts := make([][3]int64, 0)
	for io.Readln(&line) == nil {
		segs := strings.Split(line, ",")
		pts = append(pts, [3]int64{stoll(segs[0]), stoll(segs[1]), stoll(segs[2])})
	}
	k := int64(0)
	for _, p := range pts {
		k = max(k, p[0], p[1], p[2])
	}
	return pts, k
}

// spatial hashing:
// fix the distance D and find all pairs within distance D
// group points into chunks with side lengths D
// this means the min distance from a chunk to a point outside an adjacent chunk is D+epsilon
// thus all points within D distance of a point in the chunk is contained within the adjacent chunks
// increase size until enough connections are made
// assume the input is a uniform distribution of n points within a bounding cube of size k
// all points are connected when D is roughly k/(n^(1/3)) * some constant (probably)
// so the number of iterations should be roughly constant with a good choice of D
// although each check has a worst case of O(n^2), D should be pretty small
// thus we can expect the runtime to scale roughly with the runtime of a spatial hashing iteration, which is O(D^3 + n) ~= O(k^3/n + n)
// this is very hand-wavy but at least shows roughly why its better than naive O(N^2) approach
// part A is bounded by part B so we don't need to worry about it

// finds and sorts all pairs within d distance of each other
func pairsWithinDist(pts [][3]int64, k, d int64) [][2]int {
	n := int((k + d - 1) / d) // diameter in chunks
	grid := Mat3D[[]int](n, n, n)
	maxChunk := 0
	for i, p := range pts {
		x := p[0] / d
		y := p[1] / d
		z := p[2] / d
		grid[x][y][z] = append(grid[x][y][z], i)
		if len(grid[x][y][z]) > maxChunk {
			maxChunk++
		}
	}
	pairs := make([][2]int, 0, maxChunk*maxChunk*9)
	for x := 0; x < n; x++ {
		for y := 0; y < n; y++ {
			for z := 0; z < n; z++ {
				cpts := grid[x][y][z]
				for i, a := range cpts {
					// compare with points in same chunk
					for j := i + 1; j < len(cpts); j++ {
						dx := pts[a][0] - pts[cpts[j]][0]
						dy := pts[a][1] - pts[cpts[j]][1]
						dz := pts[a][2] - pts[cpts[j]][2]
						if dx*dx+dy*dy+dz*dz <= d*d {
							pairs = append(pairs, [2]int{cpts[i], cpts[j]})
						}
					}

					// compare with points in adjacent chunks
					for ox := -1; ox <= 1; ox++ {
						if x+ox < 0 || x+ox >= n {
							continue
						}
						for oy := -1; oy <= 1; oy++ {
							if y+oy < 0 || y+oy >= n {
								continue
							}
							for oz := -1; oz <= 1; oz++ {
								if z+oz < 0 || z+oz >= n || ox|oy|oz == 0 {
									continue
								}

								// to prevent duplicates, assert that min index in the pair is in this chunk
								apts := grid[x+ox][y+oy][z+oz]
								for _, j := range apts {
									if cpts[i] > j {
										continue
									}
									dx := pts[cpts[i]][0] - pts[j][0]
									dy := pts[cpts[i]][1] - pts[j][1]
									dz := pts[cpts[i]][2] - pts[j][2]
									if dx*dx+dy*dy+dz*dz <= d*d {
										pairs = append(pairs, [2]int{cpts[i], j})
									}
								}
							}
						}
					}
				}
			}
		}
	}

	// sort pairs by distance
	order := make([]Pair[int, int64], len(pairs))
	for i := 0; i < len(pairs); i++ {
		dx := pts[pairs[i][0]][0] - pts[pairs[i][1]][0]
		dy := pts[pairs[i][0]][1] - pts[pairs[i][1]][1]
		dz := pts[pairs[i][0]][2] - pts[pairs[i][1]][2]
		order[i] = Pair[int, int64]{i, dx*dx + dy*dy + dz*dz}
	}
	QSort(order)
	sortedPairs := make([][2]int, len(pairs))
	for i := 0; i < len(order); i++ {
		sortedPairs[i] = pairs[order[i].A]
	}
	return sortedPairs
}

func Day8A(io *IO) {
	pts, k := day8input(io)
	var pairs [][2]int
	dd := int64(math.Ceil(float64(k) / math.Cbrt(float64(len(pts)))))
	d := int64(0)
	for len(pairs) < 1000 {
		d += dd
		pairs = pairsWithinDist(pts, k, d)
	}
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
	pts, k := day8input(io)
	dd := int64(math.Ceil(float64(k) / math.Cbrt(float64(len(pts)))))
	d := int64(0)
	for {
		d += dd
		pairs := pairsWithinDist(pts, k, d)
		uf := NewUnionFind(len(pts))
		i := -1
		for i+1 < len(pairs) && uf.NumComps() > 1 {
			i++
			uf.Merge(int(pairs[i][0]), int(pairs[i][1]))
		}
		if uf.NumComps() == 1 {
			io.Write("%d\n", pts[pairs[i][0]][0]*pts[pairs[i][1]][0])
			break
		}
	}
}
