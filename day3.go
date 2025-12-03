package main

func solve(io *IO, digs int) {
	var line string
	var res int64 = 0
	for io.Read(&line) != nil {
		// dp: (index, num vals left) -> max value
		var dp = make([]int64, digs+1)
		mv := int64(0)
		for _, c := range line {
			v := int64(c - '0')
			for i := 0; i < digs; i++ {
				dp[i] = max(dp[i], dp[i+1]*10+v)
				mv = max(mv, dp[i])
			}
		}
		res += mv
	}
	io.Write("%d\n", res)
}

func Day3A(io *IO) {
	solve(io, 2)
}

func Day3B(io *IO) {
	solve(io, 12)
}
