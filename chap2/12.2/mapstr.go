package main

import (
	"fmt"
	"strings"
)

func main() {
	arr := []string{"hari", "haran", "m"}
	test := func(s string) string {
		return strings.ToUpper(s)
	}
	fmt.Printf("%v", Mapfun(test, arr))
	fmt.Println()
}

func Mapfun(f func(string) string, l []string) []string {
	x := make([]string, len(l))
	for i, j := range l {
		x[i] = f(j)
	}
	return x
}
