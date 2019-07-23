package main

import "fmt"

const saludo string = "Hey %s %s, welcome to Go world.\n"

func main() {
	a, b, c := getVars()
	fmt.Println(a, b, c)
}

func getVars() (int, int, int) {
	return 1, 2, 3;
}