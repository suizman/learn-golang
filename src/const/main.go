package main

import (
	"fmt"
)

const (
	message = "My first constant with iota %d %d\n"
	iota1 = iota
	iota2
)

func main() {
	fmt.Printf(message, iota1, iota2)
}
