package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	test := func(g int) int {
		return g + 1
	}
	fmt.Printf("%v", Mapfun(test, arr))
	fmt.Println()
}

func Mapfun(f func(int) int, l []int) []int {
	x := make([]int, len(l))
	for i, j := range l {
		x[i] = f(j)
	}
	return x
}
