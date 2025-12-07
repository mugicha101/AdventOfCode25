package main

import "strings"

func Day7A(io *IO) {
	var line string
	io.Readln(&line)
	n := len(line)
	dp := make([]bool, n)
	dp[strings.Index(line, "S")] = true
	res := 0
	for io.Readln(&line) == nil {
		for i := 1; i < n-1; i++ {
			if dp[i] && line[i] == '^' {
				res++
				dp[i] = false
				dp[i-1] = true
				dp[i+1] = true
			}
		}
	}
	io.Write("%d\n", res)
}

func Day7B(io *IO) {
	var line string
	io.Readln(&line)
	n := len(line)
	dp := make([]int64, n)
	dp[strings.Index(line, "S")] = 1
	for io.Readln(&line) == nil {
		for i := 1; i < n-1; i++ {
			if line[i] == '^' {
				dp[i-1] += dp[i]
				dp[i+1] += dp[i]
				dp[i] = 0
			}
		}
	}
	res := 0
	for _, v := range dp {
		res += int(v)
	}
	io.Write("%d\n", res)
}
