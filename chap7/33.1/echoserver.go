package main

import (
	"io"
	"log"
	"net"
)

func main() {
	l, err := net.Listen("tcp", ":8053")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
	for {
		if c, err := l.Accept(); err == nil {
			Echo(c)
		}
	}
}

func Echo(c net.Conn) {
	io.Copy(c, c)
	c.Close()
}
