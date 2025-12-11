package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Call struct {
	name  string
	f     func(*IO)
	input string
}

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("No target specified")
		return
	}
	target := strings.ToLower(args[1])
	calls := [][]Call{
		{Call{name: "day1a", f: Day1A, input: "day1"}, Call{name: "day1b", f: Day1B, input: "day1"}},
		{Call{name: "day2a", f: Day2A, input: "day2"}, Call{name: "day2b", f: Day2B, input: "day2"}},
		{Call{name: "day3a", f: Day3A, input: "day3"}, Call{name: "day3b", f: Day3B, input: "day3"}},
		{Call{name: "day4a", f: Day4A, input: "day4"}, Call{name: "day4b", f: Day4B, input: "day4"}},
		{Call{name: "day5a", f: Day5A, input: "day5"}, Call{name: "day5b", f: Day5B, input: "day5"}},
		{Call{name: "day6a", f: Day6A, input: "day6"}, Call{name: "day6b", f: Day6B, input: "day6"}},
		{Call{name: "day7a", f: Day7A, input: "day7"}, Call{name: "day7b", f: Day7B, input: "day7"}},
		{Call{name: "day8a", f: Day8A, input: "day8"}, Call{name: "day8b", f: Day8B, input: "day8"}},
		{Call{name: "day9a", f: Day9A, input: "day9"}, Call{name: "day9b", f: Day9B, input: "day9"}},
		{Call{name: "day10a", f: Day10A, input: "day10"}, Call{name: "day10b", f: Day10B, input: "day10"}},
		{Call{name: "day11a", f: Day11A, input: "day11"}, Call{name: "day11b", f: Day11B, input: "day11"}},
	}
	targetCalls := make([]Call, 0)
	if target == "all" {
		for _, dayCalls := range calls {
			targetCalls = append(targetCalls, dayCalls...)
		}
	} else if strings.HasPrefix(target, "day") && len(target) >= 4 {
		last := len(target)
		includeA := true
		includeB := true
		if target[len(target)-1] == 'a' {
			last--
			includeB = false
		} else if target[len(target)-1] == 'b' {
			last--
			includeA = false
		}
		if last > 3 {
			dayNum, ok := strconv.Atoi(target[3:last])
			if ok == nil && dayNum >= 1 && dayNum <= len(calls) {
				if includeA {
					targetCalls = append(targetCalls, calls[dayNum-1][0])
				}
				if includeB {
					targetCalls = append(targetCalls, calls[dayNum-1][1])
				}
			}
		}
	}
	if len(targetCalls) == 0 {
		fmt.Printf("Target %s not found\n", target)
		return
	}

	var totalTimeExec time.Duration
	var totalTimeIo time.Duration
	for _, c := range targetCalls {
		fmt.Printf("%s:\n", c.name)
		io := NewIO(c.input)
		start := time.Now()
		c.f(io)
		elapsed := time.Since(start)
		elapsed -= io.ioTime
		io.Close()
		fmt.Printf("      Time: %v exec\n            %v io\n", elapsed, io.ioTime)
		totalTimeExec += elapsed
		totalTimeIo += io.ioTime
	}
	fmt.Printf("Total Time: %v exec\n            %v io\n", totalTimeExec, totalTimeIo)
}
