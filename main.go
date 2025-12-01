package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type IO struct {
	fin     *os.File
	fout    *os.File
	scanner *bufio.Scanner
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

func (io *IO) Close() {
	io.scanner = nil
	io.fin.Close()
	io.fout.Close()
}

func (io *IO) Read(x interface{}) *IO {
	if !io.scanner.Scan() {
		x = nil
		return nil
	}
	_, err := fmt.Sscanf(io.scanner.Text(), "%v", x)
	if err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}
	return io
}

func (io *IO) Write(format string, a ...interface{}) {
	fmt.Printf(format, a...)
	fmt.Fprintf(io.fout, format, a...)
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
		"day1a": {Call{f: Day1A, input: "day1"}},
		"day1b": {Call{f: Day1B, input: "day1"}},
	}
	fs, ok := dayMap[target]
	if ok {
		for _, c := range fs {
			io := NewIO(c.input)
			c.f(io)
			io.Close()
		}
	} else {
		fmt.Printf("Target %s not found\n", target)
	}
}
