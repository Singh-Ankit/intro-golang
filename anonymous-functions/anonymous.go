package main

import (
	"fmt"
)

func main() {
	// EXAMPLE 1: anonymous function doesn't have any name
	func(input string) {
		fmt.Println(input)
	}("This is input to the function")

	//EXAMPLE 2: anonymous function without any input
	func() {
		fmt.Println("There is no input to this function")
	}()

	//EXAMPLE 3:  Assigning an anonymous to a variable, end of function doesn't have parenthesis()
	var value func() = func() {
		fmt.Println("This is anonymous to a variable")
	}
	value()
}
