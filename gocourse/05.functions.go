package main

import "fmt"

const saludo string = "Hola %s %s, esta es tu bienvenida al mundo de GO.\n"

func main() {
	a, b, c := getVariables()
	fmt.Println(a, b, c)
}

func getVariables() (int, int, int) {
	return 1, 2, 3;
}