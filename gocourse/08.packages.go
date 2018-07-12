package main

import (
	"fmt"
	"goBasics/gocourse/mypackage"
)

func main() {
	firstName := mypackage.GetName()
	fmt.Println("Your name is: ", firstName)
}