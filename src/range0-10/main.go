package main

import(
	"fmt"
)

func main() {
	myString := "This is a tring to test range function!"
	fmt.Printf("Explicitly set index to zero: %s\n", myString[0:10])
	fmt.Printf("Default go index to zero:     %s\n", myString[:10])
}
