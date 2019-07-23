package main

import "fmt"

func main() {
	var name string
	name = "MyName"
	lastname := "MyLastname"
	var number = 100
	var (
		a = 1
		b = 2
		c = 3
	)
	const helloWorld = "Hello World"
	fmt.Printf("Name: %s %s\n", name, lastname)
	fmt.Printf("Number: %d\n", number)
	fmt.Print(a, b, c)
}