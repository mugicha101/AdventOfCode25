package main

import "strings"

type shape struct {
	grid [3][3]bool
	size int
}

type region struct {
	rows, cols int
	shapeCount [6]int
}

func day12input(io *IO) ([6]shape, []region) {
	var shapes [6]shape
	var line string
	for i := 0; i < len(shapes); i++ {
		io.Readln(&line)
		for r := 0; r < 3; r++ {
			io.Readln(&line)
			for c := 0; c < 3; c++ {
				if line[c] == '#' {
					shapes[i].grid[r][c] = true
					shapes[i].size++
				}
			}
		}
		io.Readln(&line)
	}
	regions := make([]region, 0)
	for io.Readln(&line) == nil {
		io.Write("%s\n", line)
		x := strings.Index(line, "x")
		c := strings.Index(line, ":")
		reg := region{}
		reg.rows = stoi(line[:x])
		reg.cols = stoi(line[x+1 : c])
		for i, val := range strings.Split(line[c+2:], " ") {
			reg.shapeCount[i] = stoi(val)
		}
		regions = append(regions, reg)
	}
	return shapes, regions
}

func Day12(io *IO) {
	// size of grid >= num of total needed tiles
	// if num of shapes <= number of 3x3 regions in the area, is trivial to fit
	shapes, regions := day12input(io)
	res := 0
	for _, reg := range regions {
		numShapes := 0
		numTiles := 0
		for i, cnt := range reg.shapeCount {
			numShapes += cnt
			numTiles += cnt * shapes[i].size
		}
		io.Write("numShapes: %d, numTiles: %d, area: %d\n", numShapes, numTiles, reg.rows*reg.cols)
		area := reg.rows * reg.cols
		if numShapes <= (reg.rows/3)*(reg.cols/3) {
			// possible
			res++
		} else if numTiles >= area {
			// impossible
		} else {
			// np hard
			panic("Nope!")
		}
	}
	io.Write("%d\n", res)
}
