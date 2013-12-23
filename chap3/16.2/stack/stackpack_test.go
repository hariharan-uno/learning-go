package stack

import "testing"

func TestPush(t *testing.T) {
	s := new(Stack)
	s.Push(25)
	if s.Data[0] != 25 {
		t.Log("Not Pushing")
		t.Fail()
	}
}

func TestPop(t *testing.T) {
	s := new(Stack)
	s.Push(25)
	s.Push(14)
	if s.Pop() != 14 {
		t.Log("Not Popping")
		t.Fail()
	}
}
