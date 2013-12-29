package main

import "fmt"

type e interface{}

func plus2(f e) e {
	switch f.(type) {
	case int:
		return f.(int) + 2
	case string:
		return f.(string) + string(2)
	}
	return f
}

func Map(n []e, f func(e) e) []e {
	m := make([]e, len(n))
	for k, x := range n {
		m[k] = f(x)
	}
	return m
}

func main() {
	m := []e{1, 2, 3, 4}
	s := []e{"a", "b", "c", "d"}
	mm := Map(m, plus2)
	ms := Map(s, plus2)
	fmt.Printf("%v\n", mm)
	fmt.Printf("%v\n", ms)
}
