package main

import (
	"fmt"
	"strings"
)

func main() {
	strings2()
}

func strings2() {
	var text = "Hello world, Hello Go"
	fmt.Println(strings.ToUpper(text))
	fmt.Println(strings.ToLower(text))
	fmt.Println(strings.Replace(text, "Hello", "Hola", -1))
	fmt.Println(strings.Split(text, ","))
	fmt.Println("Hello"[0]) // Get one character code
	fmt.Println(string("Hello"[0])) // Get one character
	fmt.Println(len("Hello")) // Get string length

	stringUTF8 := getUnicode()
	fmt.Println("String with UTF8", stringUTF8)
}

func getUnicode() string {
	// Go supports UTF8
	return "こんにちは世界"
}