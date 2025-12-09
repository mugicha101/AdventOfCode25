package main

import "strings"

func day9input(io *IO) [][2]int64 {
	var line string
	pos := make([][2]int64, 0)
	for io.Readln(&line) == nil {
		mid := strings.Index(line, ",")
		x := stoll(line[:mid])
		y := stoll(line[mid+1:])
		pos = append(pos, [2]int64{x, y})
	}
	return pos
}

func pip(x, y int64, xLines, yLines [][3]int64) bool {
	// cast ray downwards
	// io.Write("pip %d,%d\n", x, y)
	for _, line := range xLines {
		if line[0] == x && line[1] <= y && line[2] >= y {
			// io.Write(" on x line\n")
			return true // on line
		}
	}
	cnt := 0
	for _, line := range yLines {
		if line[0] == y && line[1] <= x && line[2] >= x {
			// io.Write(" on y line\n")
			return true // on line
		}
		if line[0] < y && line[1] <= x && line[2] > x {
			// io.Write("crosses y line %v\n", line)
			cnt++ // crosses line
		}
	}
	// io.Write(" %v\n", cnt)
	return cnt&1 == 1
}

// checks if x line in poly
func lip(x, ya, yb int64, xLines, yLines [][3]int64) bool {
	if ya > yb {
		ya, yb = yb, ya
	}

	// ensure both endpoints in poly
	if !pip(x, ya, xLines, yLines) || !pip(x, yb, xLines, yLines) {
		return false
	}

	// check that line crosses no y line between endpoints
	for _, line := range yLines {
		if line[0] <= ya || line[0] >= yb {
			continue
		}
		if line[1] < x && line[2] > x {
			return false
		}
		if (line[1] == x || line[2] == x) && !(pip(x, line[0]+1, xLines, yLines) && pip(x, line[0]-1, xLines, yLines)) {
			return false
		}
	}

	return true
}

func Day9A(io *IO) {
	pos := day9input(io)
	res := int64(0)
	for i := 0; i < len(pos); i++ {
		for j := i + 1; j < len(pos); j++ {
			dx := pos[i][0] - pos[j][0]
			dy := pos[i][1] - pos[j][1]
			res = max(res, (dx+1)*(dy+1))
		}
	}
	io.Write("%d\n", res)
}

func Day9B(io *IO) {
	// rect valid if all perimeter points in shape
	// can use line in poly to detect this
	pos := day9input(io)
	res := int64(0)
	xLines := make([][3]int64, 0) // (x, y1, y2)
	yLines := make([][3]int64, 0) // (y, x1, x2)
	prev := pos[len(pos)-1]
	for _, curr := range pos {
		if prev[0] == curr[0] {
			xLines = append(xLines, [3]int64{curr[0], min(prev[1], curr[1]), max(prev[1], curr[1])})
		} else {
			yLines = append(yLines, [3]int64{curr[1], min(prev[0], curr[0]), max(prev[0], curr[0])})
		}
		prev = curr
	}
	lipx := func(x, ya, yb int64) bool {
		return lip(x, ya, yb, xLines, yLines)
	}
	lipy := func(y, xa, xb int64) bool {
		return lip(y, xa, xb, yLines, xLines)
	}
	for i := 0; i < len(pos); i++ {
		ax := pos[i][0]
		ay := pos[i][1]
		for j := i + 1; j < len(pos); j++ {
			bx := pos[j][0]
			by := pos[j][1]
			area := (abs(bx-ax) + 1) * (abs(by-ay) + 1)
			if area > res && lipx(ax, ay, by) && lipx(bx, ay, by) && lipy(ay, ax, bx) && lipy(by, ax, bx) {
				res = area
			}
		}
	}
	io.Write("%d\n", res)
}
