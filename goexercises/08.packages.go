package main

import (
	"fmt"
	"goBasics/goexercises/mypackage"
)

func main() {
	firstName := mypackage.GetName()
	fmt.Println("Your name is: ", firstName)
}