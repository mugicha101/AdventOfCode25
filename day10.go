package main

import (
	"math"
	"strings"
)

type machine struct {
	nlights int
	lights  uint64
	buttons []uint64
	joltage []int
}

func day10input(io *IO) []machine {
	var line string
	machines := make([]machine, 0)
	for io.Readln(&line) == nil {
		m := machine{}
		a := strings.Index(line, "]")
		b := strings.Index(line, "{")
		for i := 1; i < a; i++ {
			if line[i] == '#' {
				m.lights |= 1 << (i - 1)
			}
		}
		m.nlights = a - 1
		buttonSegs := strings.Split(line[a+3:b-2], ") (")
		m.buttons = make([]uint64, len(buttonSegs))
		for i, seg := range buttonSegs {
			bls := strings.Split(seg, ",")
			for _, bl := range bls {
				m.buttons[i] |= 1 << stoi(bl)
			}
		}
		joltageSegs := strings.Split(line[b+1:len(line)-1], ",")
		m.joltage = make([]int, len(joltageSegs))
		for i, jreq := range joltageSegs {
			m.joltage[i] = stoi(jreq)
		}
		machines = append(machines, m)
	}
	return machines
}

func Day10A(io *IO) {
	machines := day10input(io)
	res := 0
	for _, m := range machines {
		seen := make([]bool, 1<<m.nlights)
		q := make(Queue[uint64], 0)
		q.Push(0)
		seen[0] = true
		presses := 0
		for len(q) > 0 && !seen[m.lights] {
			presses++
			for qi := len(q); qi > 0; qi-- {
				curr := q.Pop()
				for _, bls := range m.buttons {
					next := curr ^ bls
					if seen[next] {
						continue
					}
					seen[next] = true
					q.Push(next)
				}
			}
		}
		res += presses
	}
	io.Write("%d\n", res)
}

func day10bdfs(patMap map[uint64][]uint64, m *machine, io *IO) int {
	// find all ways to make odd joltages even
	// once thats applied, all future button presses need to be even
	// thus divide joltage by 2 and recurse to find the number of presses / 2

	// all joltage 0
	totalJoltage := 0
	for _, j := range m.joltage {
		totalJoltage += j
	}
	if totalJoltage == 0 {
		return 0
	}

	// figure out target light pattern from joltage parity
	target := 0
	for i, j := range m.joltage {
		target |= (j & 1) << i
	}

	// try all subsets that achieve target pattern
	minPresses := math.MaxInt32
	for _, subset := range patMap[uint64(target)] {
		// figure out subsets effect on joltage
		presses := 0
		djoltage := make([]int, len(m.joltage))
		for i := 0; i < len(m.buttons); i++ {
			if (subset>>i)&1 == 0 {
				continue
			}
			presses++
			for j := 0; j < len(m.joltage); j++ {
				djoltage[j] += int((m.buttons[i] >> j) & 1)
			}
		}

		// check that joltages all non-negative
		valid := true
		for j := 0; valid && j < len(m.joltage); j++ {
			if djoltage[j] > m.joltage[j] {
				valid = false
			}
		}
		if !valid {
			continue
		}

		// apply joltage change and recurse
		for j := 0; j < len(m.joltage); j++ {
			m.joltage[j] = (m.joltage[j] - djoltage[j]) >> 1
		}
		recPresses := day10bdfs(patMap, m, io)
		if recPresses != math.MaxInt32 {
			minPresses = min(minPresses, presses+recPresses*2)
		}
		for j := 0; j < len(m.joltage); j++ {
			m.joltage[j] = (m.joltage[j] << 1) + djoltage[j]
		}
	}
	return minPresses
}

func Day10B(io *IO) {
	machines := day10input(io)
	res := 0
	// elegant solution from https://www.reddit.com/r/adventofcode/comments/1pk87hl/2025_day_10_part_2_bifurcate_your_way_to_victory/
	// avoids integer linear programming and gauss elimination
	for _, m := range machines {
		// map every lighting pattern to subsets of buttons satisfying it
		patMap := make(map[uint64][]uint64)
		for subset := uint64(0); subset < 1<<len(m.buttons); subset++ {
			presses := 0
			pattern := uint64(0)
			for i := 0; i < len(m.buttons); i++ {
				if (subset>>i)&1 == 1 {
					pattern ^= m.buttons[i]
					presses++
				}
			}
			if patMap[pattern] == nil {
				patMap[pattern] = make([]uint64, 0)
			}
			patMap[pattern] = append(patMap[pattern], subset)
		}

		res += day10bdfs(patMap, &m, io)
	}
	io.Write("%d\n", res)
}
