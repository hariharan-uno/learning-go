package main

import "fmt"

func main() {
	a, b := ordering(7, 4)
	c, d := ordering(3, 4)
	fmt.Printf("%v,%v\n", a, b)
	fmt.Printf("%v,%v\n", c, d)

}

func ordering(x, y int) (small, large int) {
	if x < y {
		small = x
		large = y
	} else {
		small = y
		large = x
	}

	return
}
