package main

import "fmt"

func main() {
	p := plusX(5)
	fmt.Printf("%v\n", p(2))
}

func plusX(a int) func(int) int {
	return func(x int) int { return x + a }
}
