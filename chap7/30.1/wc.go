package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	f := string(bytes)
	fmt.Println()
	fmt.Printf("Number of characters: %d\n", len(f))
	line := strings.Split(f, "\n")
	fmt.Printf("Number of lines: %d\n", len(line))
	w := 0
	for _, x := range line {
		w += len(strings.Fields(x))
	}
	fmt.Printf("Number of words: %d\n", w)
}
