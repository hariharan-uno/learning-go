/*This package implements a stack datastructure
and simple push and pop functions for stack manipulations*/
package stack

type Stack struct {
	i    int
	data [10]int
}

/*We use a pointer here so that the function works
on the actual variable and not a copy of it*/

//Push pushes an int on the stack 
func (s *Stack) Push(k int) {
	if s.i+1 > 9 {
		return
	}
	s.data[s.i] = k
	s.i++
}

//Pop pops out an item from the stack
func (s *Stack) Pop() int {
	s.i--
	return s.data[s.i]
}
