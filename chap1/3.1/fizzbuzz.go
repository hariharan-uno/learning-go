package main

import "fmt"

func main() {
	var p bool
	for i := 1; i < 101; i++ {
		p = false
		if i%3 == 0 {
			fmt.Printf("Fizz")
			p = true
		}
		if i%5 == 0 {
			fmt.Printf("Buzz")
			p = true
		}
		if !p {
			fmt.Printf("%d", i)
		}
		fmt.Println()
	}

}
