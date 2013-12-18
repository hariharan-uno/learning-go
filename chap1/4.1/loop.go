package main

import "fmt"

func main() {
	for i := 1; i < 101; i++ {
		for j := 0; j < i; j++ {
			fmt.Printf("A")
		}
		fmt.Println()
	}

}
