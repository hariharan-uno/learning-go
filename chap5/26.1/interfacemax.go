package main

import "fmt"

/*
Big returns a bool value to specify if the first of the two given inputs is greater or not.
*/
func Big(i, j interface{}) bool { //Empty interface to keep the function generic
	switch i.(type) {
	case int:
		if _, ok := j.(int); ok {
			return i.(int) > j.(int) //Convert into type int
		}
	case float32:
		if _, ok := j.(float32); ok {
			return i.(float32) > j.(float32) //Convert into type float
		}
	}
	return false
}

func main() {
	a := []int{5, 10, 13}
	b := []float32{4.34, 5.65, 7.23}
	maxa := a[0]
	maxb := b[0]

	for _, x := range a {
		if Big(x, maxa) {
			maxa = x
		}
	}

	for _, x := range b {
		if Big(x, maxb) {
			maxb = x
		}
	}
	fmt.Println(maxa, maxb)
}
