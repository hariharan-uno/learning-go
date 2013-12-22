package main

import "fmt"

func main() {
	m := []int{1, 3, 4, 2, 6, 4}
	fmt.Printf("%v\n", minval(m))

}

func minval(i []int) (min int) {
	min = i[0]
	for _, x := range i {
		if x < min {
			min = x
		}
	}
	return
}
