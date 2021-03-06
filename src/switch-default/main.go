package main

import (
	"fmt"
	"os"
)

func main() {
	n, err := fmt.Printf("Hello, World!")

	switch {
	case err != nil:
		os.Exit(1)
	case n == 0:
		fmt.Printf("No bytes output")
	case n != 13:
		fmt.Printf("Wrong number of characters: %d", n)
	default:
		fmt.Printf("OK!")
	}

	fmt.Printf("\n")
}
