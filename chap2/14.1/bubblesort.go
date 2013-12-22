package main

import "fmt"

var x int

func main() {

	m := []int{1, 4, 2, 6, 8, 3, 4, 2, 9, 45, 2, 78, 2, 3, 8, 5}
	fmt.Printf("%v\n", bubblesort(m))

}

func bubblesort(l []int) []int {

Test:
	for i, _ := range l {
		if i != 0 {
			if l[i-1] > l[i] {
				l[i-1], l[i] = l[i], l[i-1]
				x += 1
				goto Test
			}
		}
	}
	defer fmt.Printf("%d\n", x) //Count number of times goto loops have run
	return l
}
