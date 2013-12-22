package main

import (
	"fmt"
	"github.com/hariharan-uno/learning-go/chap3/stack"
)

func main() {
	s := new(stack.Stack) //Remember this! Usage of exported types
	s.Push(25)
	fmt.Printf("stack %v\n", s)
	s.Push(14)
	fmt.Printf("stack %v\n", s)
	fmt.Printf("%d\n", s.Pop())
}
