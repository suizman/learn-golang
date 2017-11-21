package main

import(
	"fmt"
)

func main() {
	myString := "This is my first test using for loop in golang!\n"

	fmt.Printf("Print index and value:\n")
	for i, r := range myString {
		fmt.Printf("%d %c\n", i, r)
	}

	fmt.Printf("Print index only:\n")
	for i := range myString {
		fmt.Printf("%d\n", i)
	}
}
