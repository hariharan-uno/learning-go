package main

import "fmt"

func main() {
	p := plusTwo()
	fmt.Printf("%v\n", p(2))
}

func plusTwo() func(int) int {
	return func(x int) int { return x + 2 }
}
