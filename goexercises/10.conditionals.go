package main

import "fmt"

func main() {
	ifTest()
}

func ifTest() {
	var number = 0
	fmt.Print("Type a number between 1 to 10: ")
	fmt.Scanf("%d", &number)
	if number%2 == 0 {
		fmt.Println("The number is even")
	} else {
		fmt.Println("The number is odd")
	}

	if number2 :=3; number2 ==3 {
		fmt.Println("Here!")
	}
}