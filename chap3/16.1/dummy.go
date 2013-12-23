package main

import (
	"fmt"
	"github.com/hariharan-uno/learning-go/chap3/16.1/stack"
)

func main() {
	s := new(stack.Stack) //Remember this! Usage of exported types
	s.Push(25)
	fmt.Printf("%d\n", s.Data[0])
	s.Push(14)
	fmt.Printf("stack %v\n", s)
	fmt.Printf("%d\n", s.Pop())
}
