package mypackage

import "fmt"

// GetName: Get and return name
func GetName() string {
	var name string
	name = "Dchuss"
	fmt.Print("Write your name: ")
	fmt.Scanf("%s", &name)
	return name
}