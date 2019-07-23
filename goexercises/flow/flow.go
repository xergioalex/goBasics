package flow

import (
	"fmt"
	"runtime"
)

func IfTest() {
	var number = 0

	fmt.Print("Ingresa un número del 1 al 10:")
	fmt.Scanf("%d", &number)

	if number%2 == 0 {
		fmt.Println("El número es par")
	} else {
		fmt.Println("El número es impar")
	}
}

func SwitchTest() {
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


	var number = 0
	fmt.Print("[Switch] Ingresa un nubero del 1 al 10: ")
	fmt.Scanf("%d", &number)

	switch number {
		case 1:
			fmt.Println("El número es 1")
		default:
			fmt.Println("El número no es 1")
	}

	switch {
		case number%2 == 0:
			fmt.Println("El número es par")
		default:
			fmt.Println("El número es impar")
	}
}