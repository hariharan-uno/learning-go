package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	str := "asSASA ¿¡ ddd dsjkdsjs dk"
	fmt.Printf("Characters:%d\nBytes:%d\n", utf8.RuneCountInString(str), len(str))
}
