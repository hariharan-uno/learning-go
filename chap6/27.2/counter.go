package main

import "fmt"

var c chan int

func count(c chan int, quit chan bool) {
	for {
		select {
		case i := <-c:
			fmt.Println(i)
		case <-quit:
			break
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan bool)
	go count(c, quit)
	for i := 0; i < 10; i++ {
		c <- i

	}
	quit <- false //Signal the end
}
