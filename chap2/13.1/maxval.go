package main

import "fmt"

func main() {
	m := []int{1, 3, 4, 2, 6, 4}
	fmt.Printf("%v\n", maxval(m))

}

func maxval(i []int) (max int) {
	max = i[0] //Important when the slice contains all negative values
	for _, x := range i {
		if x > max {
			max = x
		}
	}
	return
}
