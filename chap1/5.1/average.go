package main

import "fmt"

func main() {
	arr := [...]float64{1.2, 3.2, 4.5, 5.5, 2.5, 6.7, 4.56, 8.32, 4}
	s := arr[1:6]
	sum := 0.0
	switch len(s) {
	case 0:
		fmt.Printf("Average:0\n")
	default:
		for _, v := range s {
			sum += v
		}
		fmt.Printf("Average:%v\n", sum/float64(len(s)))
	}
}
