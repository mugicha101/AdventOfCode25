package main

import (
	"strconv"
	"strings"
)

func Day6A(io *IO) {
	var line string
	var vals = make([][]int64, 0)
	res := int64(0)
	for io.Readln(&line) == nil {
		segs := strings.Fields(line)
		if segs[0][0] >= '0' && segs[0][0] <= '9' {
			vals = append(vals, make([]int64, len(segs)))
			for i, seg := range segs {
				v, _ := strconv.ParseInt(seg, 10, 64)
				vals[len(vals)-1][i] = v
			}
		} else {
			for i, seg := range segs {
				var op func(int64, int64) int64
				var a int64
				switch seg[0] {
				case '+':
					a = 0
					op = func(a, b int64) int64 { return a + b }
				case '-':
					a = 0
					op = func(a, b int64) int64 { return a - b }
				case '*':
					a = 1
					op = func(a, b int64) int64 { return a * b }
				default:
					panic("unknown operator " + seg)
				}
				for _, row := range vals {
					a = op(a, row[i])
				}
				res += a
			}
		}
	}
	io.Write("%d\n", res)
}

func Day6B(io *IO) {
	var line string
	var lines = make([]string, 0)
	res := int64(0)
	for io.Readln(&line) == nil && !(line[0] == '*' || line[0] == '+' || line[0] == '-') {
		lines = append(lines, line)
	}
	ops := make(PairList[rune, int], 0)
	for i, c := range line {
		if c != ' ' {
			ops = append(ops, Pair[rune, int]{First: c, Second: i})
		}
	}
	ops = append(ops, Pair[rune, int]{First: ' ', Second: len(line) + 1})
	for i := 0; i < len(ops)-1; i++ {
		var op func(int64, int64) int64
		var a int64
		switch ops[i].First {
		case '+':
			a = 0
			op = func(a, b int64) int64 { return a + b }
		case '-':
			a = 0
			op = func(a, b int64) int64 { return a - b }
		case '*':
			a = 1
			op = func(a, b int64) int64 { return a * b }
		default:
			panic("unknown operator " + string(ops[i].First))
		}
		for c := ops[i+1].Second - 2; c >= ops[i].Second; c-- {
			v := int64(0)
			for _, row := range lines {
				if row[c] != ' ' {
					v = v*10 + int64(row[c]-'0')
				}
			}
			a = op(a, v)
		}
		res += a
	}
	io.Write("%d\n", res)
}
