package main

import (
	"fmt"
	"strconv"
)

type stack struct {
	i    int
	data [10]int
}

func main() {
	var s stack
	s.push(25)
	fmt.Printf("stack %v\n", s)
	s.push(14)
	s.push(174)
	fmt.Printf("stack %v\n", s.String())
}

/*We use a pointer here so that the function works
on the actual variable and not a copy of it*/
func (s *stack) push(k int) {
	if s.i+1 > 9 {
		return
	}
	s.data[s.i] = k
	s.i++
}

func (s *stack) pop() int {
	s.i--
	return s.data[s.i]
}

func (s *stack) String() string {
	var str string
	for i := 0; i < s.i; i++ {
		str = str + "[" + strconv.Itoa(i) + " : " + strconv.Itoa(s.data[i]) + "]  "
	}
	return str
}
