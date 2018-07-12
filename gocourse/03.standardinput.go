package main

import "fmt"

const saludo string = "Hola %s %s, esta es tu bienvenida al mundo de GO.\n"

func main() {
	nombre := getName()
	apellido := "GutiGu"
	fmt.Print("Lo primero q vi. \n")
	fmt.Printf("Hola %s, tu bienvenida al mundo de GO.\n", nombre)
	fmt.Printf(saludo, nombre, apellido)
}

func getName() string {
	var nombre string
	nombre = "Dchuss"
	fmt.Print("Escribe tu nombre: ")
	fmt.Scanf("%s", &nombre)
	return nombre
}