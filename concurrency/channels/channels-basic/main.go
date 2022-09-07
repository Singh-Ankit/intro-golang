//package channels_basic --> while running this result in go run: cannot run non-main package
package main

import (
	"fmt"
)

func Example1() {
	//creating a channel variable using make
	ch := make(chan int, 1)
	// insert a value into a channel using '<-'
	ch <- 6
	fmt.Println(<-ch)
}

func main() {
	fmt.Println("Channels basic demo")
	messages := make(chan string, 1)

	func() {
		messages <- "ping"
	}()
	fmt.Println(<-messages)
	Example1()
}
