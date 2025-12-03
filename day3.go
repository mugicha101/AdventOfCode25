package main

func Day3A(io *IO) {
	var line string
	res := 0
	for io.Read(&line) != nil {
		mv := 0
		for i := 0; i < len(line); i++ {
			for j := i + 1; j < len(line); j++ {
				mv = max(mv, int(line[i]-'0')*10+int(line[j]-'0'))
			}
		}
		res += mv
	}
	io.Write("%d\n", res)
}

func Day3B(io *IO) {
	var line string
	var res int64 = 0
	for io.Read(&line) != nil {
		// dp: (index, num vals left) -> max value
		var dp [13]int64
		for _, c := range line {
			v := int64(c - '0')
			for i := 0; i < 12; i++ {
				dp[i] = max(dp[i], dp[i+1]*10+v)
			}
		}
		mv := int64(0)
		for i := 0; i < 12; i++ {
			mv = max(mv, dp[i])
		}
		res += mv
	}
	io.Write("%d\n", res)
}
