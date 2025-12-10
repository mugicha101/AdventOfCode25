package main

import (
	"math"
	"strconv"
	"strings"

	"github.com/chriso345/gspl/lp"
	"github.com/chriso345/gspl/solver"
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

func Day10B(io *IO) {
	machines := day10input(io)
	res := 0
	for _, m := range machines {
		// reduces to a linear equation
		// x_i = presses of button i
		// b_i = counter i
		// a_ij = 1 if button j increments counter i
		// minimize sum x_i
		// s.t. for all i: sum_j a_ij * x_j = b_i
		// TODO: replace with z3, this library may have bugs
		x := make([]lp.LpVariable, len(m.buttons))
		objTerms := make([]lp.LpTerm, len(m.buttons))
		for i := range m.buttons {
			x[i] = lp.NewVariable("x"+strconv.Itoa(i), lp.LpCategoryContinuous)
			objTerms[i] = lp.NewTerm(1, x[i])
		}
		obj := lp.NewExpression(objTerms)
		prog := lp.NewLinearProgram("ILP", x)
		prog.AddObjective(lp.LpMinimise, obj)
		io.Write("%d %d\n", len(m.buttons), len(m.joltage))
		for i, jreq := range m.joltage {
			conTerms := make([]lp.LpTerm, 0)
			for j, bl := range m.buttons {
				if (bl>>i)&1 == 0 {
					continue
				}
				// button j increases counter i
				conTerms = append(conTerms, lp.NewTerm(1, x[j]))
			}
			prog.AddConstraint(lp.NewExpression(conTerms), lp.LpConstraintEQ, float64(jreq))
		}
		solver.Solve(&prog)
		// prog.PrintSolution()
		io.Write("%v\n", prog.Solution)
		if prog.Status == lp.LpStatusOptimal {
			res -= int(math.Floor(prog.Solution))
		}
	}
	io.Write("%d\n", res)
}
