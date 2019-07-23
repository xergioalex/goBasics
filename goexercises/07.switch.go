package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
		case "darwin":
			fmt.Println("OS X.")
		case "linux":
			fmt.Println("Linux.")
		default:
			// freebsd, openbsd,
			// plan9, windows...
			fmt.Printf("%s.", os)
	}

	switchTest()
}

func switchTest() {
	var number = 0
	fmt.Print("[Switch] Type a number between 1 to 10: ")
	fmt.Scanf("%d", &number)

	switch number {
		case 1:
			fmt.Println("The number is 1")
		default:
			fmt.Println("The number isnt 1")
	}

	switch {
		case number%2 == 0:
			fmt.Println("The number is even")
		default:
			fmt.Println("The number is odd")
	}
}