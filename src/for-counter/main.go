package main

import (
	"fmt"
)

func main() {
	var counter int

	counter = 0

	fmt.Printf("Clasic counter:\n")
	for counter < 5 {
		fmt.Printf("Hello, World!\n")

		counter += 1
	}

	fmt.Print("\nMinimal counter:\n")
	for i := 0; i < 5; i++ {
		fmt.Printf("Hello, World!\n")
	}

	fmt.Print("\nSimultaneous vars counter:\n")
	for i2, j := 0, 1; i2 < 5; i2, j = i2+1, j*2 {
		fmt.Printf("%d Hello, World!\n", j)
	}

	fmt.Print("\nBoolean counter:\n")
	var stop bool

	i := 0

	for !stop {
		fmt.Printf("Hello, World!\n")

		i += 1
		if i == 5 {
			stop = true
		}
	}

}
