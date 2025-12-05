package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

type IO struct {
	fin    *os.File
	fout   *os.File
	reader *bufio.Reader
	ioTime time.Duration
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

	return &IO{fin: fin, fout: fout, reader: bufio.NewReader(fin)}
}

func (o *IO) Close() {
	o.reader = nil
	o.fin.Close()
	o.fout.Close()
}

func (o *IO) Read(x ...interface{}) error {
	start := time.Now()
	defer func() { o.ioTime += time.Since(start) }()

	_, err := fmt.Fscan(o.reader, x...)
	return err
}

func (o *IO) Readln(s *string) error {
	start := time.Now()
	defer func() { o.ioTime += time.Since(start) }()

	*s = ""
	isPref := true
	for isPref {
		var line []byte
		var err error
		line, isPref, err = o.reader.ReadLine()
		if err != nil {
			return err
		}
		*s += string(line)
	}
	return nil
}

func (o *IO) Write(format string, a ...interface{}) {
	start := time.Now()
	fmt.Printf(format, a...)
	fmt.Fprintf(o.fout, format, a...)
	o.ioTime += time.Since(start)
}
