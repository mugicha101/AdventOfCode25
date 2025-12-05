package main

import (
	"strconv"
)

func Day1A(io *IO) {
	var line string
	pos := 50
	res := 0
	for io.Readln(&line) == nil {
		amt, _ := strconv.Atoi(line[1:])
		amt %= 100
		if line[0] == 'L' {
			amt = 100 - amt
		}
		pos = (pos + amt) % 100
		if pos == 0 {
			res++
		}
	}
	io.Write("%d\n", res)
}

func Day1B(io *IO) {
	var line string
	pos := 50
	res := 0
	for io.Readln(&line) == nil {
		amt, _ := strconv.Atoi(line[1:])

		// calculate number of zero passes
		rem := amt
		if line[0] == 'L' {
			if pos == 0 {
				rem -= 100
			} else {
				rem -= pos
			}
		} else {
			rem -= 100 - pos
		}
		if rem >= 0 {
			res += rem/100 + 1
		}

		// move position
		amt %= 100
		if line[0] == 'L' {
			amt = 100 - amt
		}
		pos = (pos + amt) % 100
	}
	io.Write("%d\n", res)
}
