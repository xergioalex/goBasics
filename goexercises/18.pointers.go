package main

import "fmt"

func pointerTest() {
	a := 100
	var b * int
	b = &a
	fmt.Println("Without modifications")
	fmt.Println(a, *b)
	fmt.Println(&a, b)
	pointerModify(b)
	fmt.Println("With modifications")
	fmt.Println(a, *b)
	fmt.Println(&a, b)
}

func pointerModify(c  *int) {
	*c = 10
}

func main() {
	pointerTest()
}