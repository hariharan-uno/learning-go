package main

import (
	"errors"
	"fmt"
)

type value int

type node struct {
	value
	next *node
	prev *node
}

type linkedlist struct {
	head, tail *node
}

func (l *linkedlist) Front() *node {
	return l.head
}

func (n *node) Next() *node {
	return n.next
}

func (l *linkedlist) Push(v value) *linkedlist {
	p := &node{value: v}
	if l.head == nil {
		l.head = p
	} else {
		l.tail.next = p
		p.prev = l.tail
	}
	l.tail = p
	return l
}

var errEmpty = errors.New("List is empty")

func (l *linkedlist) Pop() (v value, err error) {
	if l.tail == nil {
		err = errEmpty
	} else {
		v = l.tail.value
		l.tail = l.tail.prev
		if l.tail == nil {
			l.head = nil
		}
	}
	return
}

func main() {
	l := new(linkedlist)
	l.Push(1)
	l.Push(2)
	l.Push(4)
	for n := l.Front(); n != nil; n = n.Next() {
		fmt.Printf("%v\n", n.value)
	}
}
