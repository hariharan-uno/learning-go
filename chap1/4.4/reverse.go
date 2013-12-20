package main

import "fmt"

func main() {
	str := "foobar"
	var reverse string
	x := []rune(str)
	for i, j := 0, len(x)-1; i < j; i, j = i+1, j-1 {
		x[i], x[j] = x[j], x[i]
	}
	reverse = string(x)
	fmt.Printf("%s\n", reverse)

}
