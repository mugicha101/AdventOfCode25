package main

func nbs(grid [][]bool, r, c int) int {
	rows := len(grid)
	cols := len(grid[0])
	nbs := 0
	if grid[r][c] {
		nbs -= 1
	}
	for nr := max(0, r-1); nr <= min(rows-1, r+1); nr++ {
		for nc := max(0, c-1); nc <= min(cols-1, c+1); nc++ {
			if grid[nr][nc] {
				nbs++
			}
		}
	}
	return nbs
}

func Day4A(io *IO) {
	var line string
	grid := make([][]bool, 0)
	for io.Read(&line) != nil {
		grid = append(grid, make([]bool, len(line)))
		for i, c := range line {
			grid[len(grid)-1][i] = c == '@'
		}
	}
	rows := len(grid)
	cols := len(grid[0])
	res := 0
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] && nbs(grid, r, c) < 4 {
				res++
			}
		}
	}
	io.Write("%d\n", res)
}

func Day4B(io *IO) {
	var line string
	grid := make([][]bool, 0)
	for io.Read(&line) != nil {
		grid = append(grid, make([]bool, len(line)))
		for i, c := range line {
			grid[len(grid)-1][i] = c == '@'
		}
	}
	rows := len(grid)
	cols := len(grid[0])
	q := make([][2]int, 0)
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] && nbs(grid, r, c) < 4 {
				grid[r][c] = false
				q = append(q, [2]int{r, c})
			}
		}
	}
	res := 0
	for len(q) > 0 {
		res++
		r := q[0][0]
		c := q[0][1]
		q = q[1:]
		for nr := max(0, r-1); nr <= min(rows-1, r+1); nr++ {
			for nc := max(0, c-1); nc <= min(cols-1, c+1); nc++ {
				if grid[nr][nc] && nbs(grid, nr, nc) < 4 {
					grid[nr][nc] = false
					q = append(q, [2]int{nr, nc})
				}
			}
		}
	}
	io.Write("%d\n", res)
}
