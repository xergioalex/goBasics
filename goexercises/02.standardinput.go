package main

import "fmt"

func main () {
	var name string
	fmt.Print("Type your name: ")
	fmt.Scanf("%s", &name)
	fmt.Printf("Hey %s, welcome", name)
}
