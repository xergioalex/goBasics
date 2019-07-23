package main

import "fmt"

const saludo string = "Hello %s %s, this is your welcome to Go world using text format.\n"

func main() {
	nombre := getName()
	apellido := "Florez"
	fmt.Print("Dummy text. \n")
	fmt.Printf("Hello %s, welcome to GO world getting name from standar input.\n", nombre)
	fmt.Printf(saludo, nombre, apellido)
}

func getName() string {
	var nombre string
	nombre = "Dchuss"
	fmt.Print("Type your name: ")
	fmt.Scanf("%s", &nombre)
	return nombre
}