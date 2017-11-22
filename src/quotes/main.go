package main

import (
	"fmt"
)

func main() {

	myString1 := `Using back "quotes". This is a literal.\n`
	fmt.Printf(myString1)

	myString2 := "\nUsing double quotes.\n"
	fmt.Printf(myString2)
}
