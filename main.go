package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type Call struct {
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
	dayMap := map[string][]Call{
		"all":   {Call{f: Day1A, input: "day1"}, Call{f: Day1B, input: "day1"}, Call{f: Day2A, input: "day2"}, Call{f: Day2B, input: "day2"}, Call{f: Day3A, input: "day3"}, Call{f: Day3B, input: "day3"}, Call{f: Day4A, input: "day4"}, Call{f: Day4B, input: "day4"}, Call{f: Day5A, input: "day5"}, Call{f: Day5B, input: "day5"}, Call{f: Day6A, input: "day6"}, Call{f: Day6B, input: "day6"}, Call{f: Day7A, input: "day7"}, Call{f: Day7B, input: "day7"}},
		"day1":  {Call{f: Day1A, input: "day1"}, Call{f: Day1B, input: "day1"}},
		"day1a": {Call{f: Day1A, input: "day1"}},
		"day1b": {Call{f: Day1B, input: "day1"}},
		"day2":  {Call{f: Day2A, input: "day2"}, Call{f: Day2B, input: "day2"}},
		"day2a": {Call{f: Day2A, input: "day2"}},
		"day2b": {Call{f: Day2B, input: "day2"}},
		"day3":  {Call{f: Day3A, input: "day3"}, Call{f: Day3B, input: "day3"}},
		"day3a": {Call{f: Day3A, input: "day3"}},
		"day3b": {Call{f: Day3B, input: "day3"}},
		"day4":  {Call{f: Day4A, input: "day4"}, Call{f: Day4B, input: "day4"}},
		"day4a": {Call{f: Day4A, input: "day4"}},
		"day4b": {Call{f: Day4B, input: "day4"}},
		"day5":  {Call{f: Day5A, input: "day5"}, Call{f: Day5B, input: "day5"}},
		"day5a": {Call{f: Day5A, input: "day5"}},
		"day5b": {Call{f: Day5B, input: "day5"}},
		"day6":  {Call{f: Day6A, input: "day6"}, Call{f: Day6B, input: "day6"}},
		"day6a": {Call{f: Day6A, input: "day6"}},
		"day6b": {Call{f: Day6B, input: "day6"}},
		"day7":  {Call{f: Day7A, input: "day7"}, Call{f: Day7B, input: "day7"}},
		"day7a": {Call{f: Day7A, input: "day7"}},
		"day7b": {Call{f: Day7B, input: "day7"}},
	}
	fs, ok := dayMap[target]
	if ok {
		var totalTimeExec time.Duration
		var totalTimeIo time.Duration
		for _, c := range fs {
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
	} else {
		fmt.Printf("Target %s not found\n", target)
	}
}
