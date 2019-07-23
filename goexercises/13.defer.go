package main

import (
	"fmt"
)

func deferTest () {
	fmt.Println("La función main ha culminado")
}

func main() {
	defer deferTest()

	fmt.Println("Hello world")
}