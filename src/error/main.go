package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Printf("Using self defined variables with if statement:\n")
	if numberBytes1, theError1 := fmt.Printf("Hello World!\n"); theError1 != nil {
		os.Exit(1)
	} else {
		fmt.Printf("Printed %d bytes\n\n", numberBytes1)
	}

	fmt.Printf("Using manually defined variables with if statement:\n")

	var numberBytes2 int
	var theError2 error

	if numberBytes2, theError2 = fmt.Printf("Hello World 2!\n"); theError2 != nil {
		os.Exit(1)
	} else {
		fmt.Printf("Printed %d bytes\n\n", numberBytes2)
	}

	fmt.Printf("Using self defined variables clear syntax with if statement:\n")
	numberBytes3, theError3 := fmt.Printf("Hello World 3!\n")

	if theError3 != nil {
		os.Exit(1)
	} else {
		fmt.Printf("Printed %d bytes\n\n", numberBytes3)
	}

	fmt.Printf("Just an error with exit code 1:\n")
	if _, theError4 := fmt.Printf("Hello World 4!\n"); theError4 != nil {
		os.Exit(1)
	}

}
