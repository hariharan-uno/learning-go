package main

import "fmt"

func main() {
	a := []string{"h", "h", "h", "a", "a", "r", "i", "i", "i"}
	b := []string{}
	fmt.Println(a)
	var prev string
	for _, x := range a {
		if x != prev {
			b = append(b, x)
			prev = x
		}
	}
	fmt.Println(b)
}
