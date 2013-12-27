package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

var showNum = flag.Bool("n", false, "flag to output the line numbers")

func readlines(r io.Reader) {
	rd := bufio.NewReader(r)
	lineNumber := 1
	for {
		line, err := rd.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if *showNum {
			fmt.Printf("\r%4d %s", lineNumber, line)
		} else {
			fmt.Printf("%s", line)
		}
		lineNumber++
	}
}

func cat(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	readlines(file)
}

func main() {
	flag.Parse()
	for i := 0; i < flag.NArg(); i++ {
		cat(flag.Arg(i))
	}
}
