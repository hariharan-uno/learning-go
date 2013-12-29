package main

import "fmt"

var c chan int

func count(c chan int) {
	for {
		i := <-c
		fmt.Println(i)
	}
}

func main() {
	c = make(chan int)
	go count(c)
	for i := 0; i < 10; i++ {
		c <- i
	}
}
