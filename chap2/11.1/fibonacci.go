package main

import "fmt"

func main() {
	for _, j := range fib(4) {
		fmt.Printf("%d ", j)
	}
	fmt.Println()
}

func fib(n int) []int {
	x := make([]int, n)
	x[0], x[1] = 1, 1
	for i := 2; i < n; i++ {
		x[i] = x[i-1] + x[i-2]
	}
	return x
}
