package main

import (
	"math"
	"strconv"
	"strings"
)

func day2rep(x int64, reps int) int64 {
	var mag int64 = 10
	for mag <= x {
		mag *= 10
	}
	var mult int64 = mag
	var res int64 = x
	for i := 0; i < reps-1; i++ {
		res += x * mult
		mult *= mag
	}
	return res
}

func day2nextBase(x int64, reps int) int64 {
	var lo int64 = 1
	var hi int64 = 1
	for day2rep(hi, reps) < x {
		lo = hi
		hi *= 2
	}
	for lo < hi {
		m := (lo + hi) >> 1
		if day2rep(m, reps) >= x {
			hi = m
		} else {
			lo = m + 1
		}
	}
	return lo
}

func Day2A(io *IO) {
	var line string
	io.Readln(&line)
	input := strings.Split(line, ",")
	itvs := make([]Pair[int64, int64], 0, len(input))
	for _, itv := range input {
		segs := strings.Split(itv, "-")
		lo, _ := strconv.ParseInt(segs[0], 10, 64)
		hi, _ := strconv.ParseInt(segs[1], 10, 64)
		itvs = append(itvs, Pair[int64, int64]{A: lo, B: hi})
	}
	var res int64 = 0
	for _, itv := range itvs {
		start := day2nextBase(itv.A, 2)
		end := day2nextBase(itv.B+1, 2)
		for half := start; half < end; half++ {
			res += day2rep(half, 2)
		}
	}
	io.Write("%d\n", res)
}

func Day2B(io *IO) {
	var line string
	io.Readln(&line)
	input := strings.Split(line, ",")
	itvs := make([]Pair[int64, int64], 0, len(input))
	for _, itv := range input {
		segs := strings.Split(itv, "-")
		lo, _ := strconv.ParseInt(segs[0], 10, 64)
		hi, _ := strconv.ParseInt(segs[1], 10, 64)
		itvs = append(itvs, Pair[int64, int64]{A: lo, B: hi})
	}
	var res int64 = 0
	for _, itv := range itvs {
		repCap := int(math.Ceil(math.Log10(float64(itv.B))))
		seen := make(Set[int64])
		for reps := 2; reps <= repCap; reps++ {
			start := day2nextBase(itv.A, reps)
			end := day2nextBase(itv.B+1, reps)
			for half := start; half < end; half++ {
				val := day2rep(half, reps)
				if seen.Insert(val) {
					res += val
				}
			}
		}
	}
	io.Write("%d\n", res)
}
