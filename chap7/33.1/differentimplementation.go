package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	l, err := net.Listen("tcp", ":8053")
	if err != nil {
		log.Fatal(err)
	}
	for {
		if c, err := l.Accept(); err == nil {
			Echo(c)
		}
	}
}

func Echo(c net.Conn) {
	defer c.Close()
	line, err := bufio.NewReader(c).ReadString('\n')
	if err != nil {
		fmt.Println("Failure to read: %s", err.Error())
		return
	}
	_, err = c.Write([]byte(line))
	if err != nil {
		fmt.Println("Failure to write: %s", err.Error())
		return
	}
}
