package main

import "fmt"

func main() {
	vararg(1, 2, 3, 4, 5, 6)
}

func vararg(arg ...int) {
	for _, n := range arg {
		fmt.Printf("%d\n", n)
	}
}
