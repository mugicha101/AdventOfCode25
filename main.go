package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

type IO struct {
	fin     *os.File
	fout    *os.File
	scanner *bufio.Scanner
	ioTime  time.Duration
}

func NewIO(name string) *IO {
	fin, err := os.Open("input/" + name + ".txt")
	if err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}

	fout, err := os.Create("output/" + name + ".txt")
	if err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}

	return &IO{fin: fin, fout: fout, scanner: bufio.NewScanner(fin)}
}

func (o *IO) Close() {
	o.scanner = nil
	o.fin.Close()
	o.fout.Close()
}

func (o *IO) Read(x interface{}) *IO {
	start := time.Now()
	if !o.scanner.Scan() {
		x = nil
		return nil
	}
	_, err := fmt.Sscanf(o.scanner.Text(), "%v", x)
	if errors.Is(err, io.EOF) {
		x = nil
		return nil
	}
	if err != nil {
		log.Fatal("io.Read error: ", err)
		os.Exit(-1)
	}
	o.ioTime += time.Since(start)
	return o
}

func (o *IO) Write(format string, a ...interface{}) {
	start := time.Now()
	fmt.Printf(format, a...)
	fmt.Fprintf(o.fout, format, a...)
	o.ioTime += time.Since(start)
}

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
		"all":   {Call{f: Day1A, input: "day1"}, Call{f: Day1B, input: "day1"}, Call{f: Day2A, input: "day2"}, Call{f: Day2B, input: "day2"}, Call{f: Day3A, input: "day3"}, Call{f: Day3B, input: "day3"}, Call{f: Day4A, input: "day4"}, Call{f: Day4B, input: "day4"}, Call{f: Day5A, input: "day5"}, Call{f: Day5B, input: "day5"}},
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
