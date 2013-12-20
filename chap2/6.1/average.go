package main

import "fmt"

func main() {
	var arr = [...]float64{1.2, 3.4, 4.5, 55.6, 6.7, 7.8, 8.9, 2.4, 5.3, 9.7}
	sl := arr[2:6]
	a := average(sl)
	fmt.Printf("%v\n", a)
}

func average(s []float64) float64 {
	sum := 0.0
	switch len(s) {
	case 0:
		return 0
	default:
		for _, k := range s {
			sum += k
		}
		return sum / float64(len(s))
	}
}
