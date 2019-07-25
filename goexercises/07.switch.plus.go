package main

import (
	"fmt"
)

func main() {
	var number = 0
	fmt.Print("[Switch] Type a number between 1 to 10: ")
	fmt.Scanf("%d", &number)

	// switch number {
	// 	case 1:
	// 		fmt.Println("The number is 1")
	// 	default:
	// 		fmt.Println("The number isnt 1")
	// }

	// switch {
	// 	case number%2 == 0:
	// 		fmt.Println("The number is even")
	// 	default:
	// 		fmt.Println("The number is odd")
	// }

	switch {
	case number == 2:
		fmt.Println("The number is 2")
		fallthrough
	case number%2 == 0:
		fmt.Println("The number is even")
	default:
		fmt.Println("The number is odd")
	}

}
