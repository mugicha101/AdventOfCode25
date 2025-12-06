package main

import (
	"strconv"
	"strings"
)

func fetchIntervals(io *IO) OrderedPairList[int64, int64] {
	var line string
	itvs := make(OrderedPairList[int64, int64], 0)
	for io.Readln(&line) == nil && line != "" {
		s := strings.Index(line, "-")
		i, _ := strconv.ParseInt(line[:s], 10, 64)
		j, _ := strconv.ParseInt(line[s+1:], 10, 64)
		itvs = append(itvs, OrderedPair[int64, int64]{i, j})
	}
	itvs.Sort()
	j := 0
	for i := 1; i < len(itvs); i++ {
		if itvs[j].B+1 >= itvs[i].A {
			itvs[j].B = max(itvs[j].B, itvs[i].B)
		} else {
			j++
			itvs[j] = itvs[i]
		}
	}
	itvs = itvs[:j+1]
	return itvs
}

func Day5A(io *IO) {
	itvs := fetchIntervals(io)
	var x int64
	res := 0
	for io.Read(&x) == nil {
		l := 0
		r := len(itvs) - 1
		for l != r {
			m := (l + r + 1) / 2
			if itvs[m].A <= x {
				l = m
			} else {
				r = m - 1
			}
		}
		if itvs[l].A <= x && itvs[l].B >= x {
			res++
		}
	}

	io.Write("%d\n", res)
}

func Day5B(io *IO) {
	itvs := fetchIntervals(io)
	res := int64(0)
	for _, itv := range itvs {
		res += itv.B - itv.A + 1
	}
	io.Write("%d\n", res)
}
