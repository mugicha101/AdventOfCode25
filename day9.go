package main

import "strings"

func day9input(io *IO) [][2]int64 {
	var line string
	pos := make([][2]int64, 0)
	for io.Readln(&line) == nil {
		mid := strings.Index(line, ",")
		c := stoll(line[:mid])
		r := stoll(line[mid+1:])
		pos = append(pos, [2]int64{r, c})
	}
	return pos
}

func Day9A(io *IO) {
	pts := day9input(io)
	res := int64(0)
	for i := 0; i < len(pts); i++ {
		for j := i + 1; j < len(pts); j++ {
			h := abs(pts[i][0]-pts[j][0]) + 1
			w := abs(pts[i][1]-pts[j][1]) + 1
			res = max(res, w*h)
		}
	}
	io.Write("%d\n", res)
}

func Day9B(io *IO) {
	pts := day9input(io)
	rs := make([]int64, len(pts))
	cs := make([]int64, len(pts))
	for i, p := range pts {
		rs[i] = p[0]
		cs[i] = p[1]
	}

	// coordinate compression
	rs, rrank := RankMap(rs)
	cs, crank := RankMap(cs)
	for i := 0; i < len(pts); i++ {
		pts[i][0] = int64(rrank[pts[i][0]]) + 1
		pts[i][1] = int64(crank[pts[i][1]]) + 1
	}

	// difference array for each row
	rows := len(rs) + 2
	cols := len(cs) + 2
	grid := Mat[uint32](rows, cols)
	prev := pts[len(pts)-1]

	for i := 0; i < len(pts); i++ {
		curr := pts[i]
		if curr[0] == prev[0] {
			// horizontal line
			r := curr[0]
			ca := min(curr[1], prev[1])
			cb := max(curr[1], prev[1])
			grid[r][ca] = max(grid[r][ca], 1)
			for c := ca + 1; c < cb; c++ {
				grid[r][c] = 1
			}
			grid[r][cb] = max(grid[r][cb], 1)
		} else {
			// vertical line
			c := curr[1]
			ra := min(curr[0], prev[0])
			rb := max(curr[0], prev[0])
			for r := ra + 1; r <= rb; r++ {
				grid[r][c] = 2
			}
		}
		prev = curr
	}

	// rasterize
	for r := 1; r < rows; r++ {
		in := uint32(0)
		for c := 1; c < cols; c++ {
			switch grid[r][c] {
			case 1: // on a horizontal line
			case 2: // on a vertical line
				in = 1 - in
				grid[r][c] = 1
			default: // not on a line
				grid[r][c] = in
			}

			// convert to psum
			grid[r][c] = (grid[r][c] + grid[r][c-1] + grid[r-1][c]) - grid[r-1][c-1]
		}
	}

	// check all pairs, query psum to validate that pair within polygon
	res := int64(0)
	for i := 0; i < len(pts); i++ {
		for j := i + 1; j < len(pts); j++ {
			ar := uint32(min(pts[i][0], pts[j][0]))
			ac := uint32(min(pts[i][1], pts[j][1]))
			br := uint32(max(pts[i][0], pts[j][0]))
			bc := uint32(max(pts[i][1], pts[j][1]))
			h := abs(rs[br-1]-rs[ar-1]) + 1
			w := abs(cs[bc-1]-cs[ac-1]) + 1
			area := w * h
			if area > res && grid[br][bc]-grid[ar-1][bc]-grid[br][ac-1]+grid[ar-1][ac-1] == (br+1-ar)*(bc+1-ac) {
				res = area
			}
		}
	}
	io.Write("%d\n", res)
}
