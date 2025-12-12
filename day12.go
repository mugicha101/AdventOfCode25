package main

import "strings"

func Day12(io *IO) {
	// size of grid >= num of total needed tiles
	// if num of shapes <= number of 3x3 regions in the area, is trivial to fit
	var tilesPerShape [6]int
	var line string
	for i := 0; i < len(tilesPerShape); i++ {
		io.Readln(&line)
		for r := 0; r < 3; r++ {
			io.Readln(&line)
			for c := 0; c < 3; c++ {
				if line[c] == '#' {
					tilesPerShape[i]++
				}
			}
		}
		io.Readln(&line)
	}
	res := 0
	for io.Readln(&line) == nil {
		x := strings.Index(line, "x")
		c := strings.Index(line, ":")
		rows := stoi(line[:x])
		cols := stoi(line[x+1 : c])
		numShapes := 0
		numTiles := 0
		for i, val := range strings.Split(line[c+2:], " ") {
			cnt := stoi(val)
			numShapes += cnt
			numTiles += cnt * tilesPerShape[i]
		}

		if numShapes <= (rows/3)*(cols/3) {
			// possible
			res++
		} else if numTiles >= rows*cols {
			// impossible
		} else {
			// np hard
			panic("Nope!") // doesn't work for samples lol
		}
	}
	io.Write("%d\n", res)
}
