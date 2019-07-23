package main

import (
	"fmt"
)

func deferTest () {
	fmt.Println("The main function finish")
}

func orderDefer() { // Defers works as stack
	fmt.Println("Start")
	defer defered("1")
	defer defered("2")
	defer defered("3")
	defer defered("4")
	defer defered("5")
	fmt.Println("End")
}

func defered(text string) {
	fmt.Printf("%s\n", text)
}

func main() {
	defer deferTest() // defer allow do something after the current function context finish

	fmt.Println("Hello world")
	orderDefer()
}

