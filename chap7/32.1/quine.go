package main

import "fmt"

func main() {
	fmt.Printf("%s%c%s%c\n", q, 0x60, q, 0x60)
}

var q = `
package main

import "fmt"

func main() {
	fmt.Printf("%s%c%s%c\n", q, 0x60, q, 0x60)
}
var q = `

//This solution is from Russ Cox. It was posted to the Go Nuts mailing list.
