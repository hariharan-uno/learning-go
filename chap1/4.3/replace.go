package main

import "fmt"

func main() {
	s := "dsksdlndsn  µ kjbds"
	r := []rune(s)
	fmt.Printf("Before:%s\n", s)
	copy(r[4:4+3], []rune("abc"))
	fmt.Printf("After:%s\n", string(r))
}
