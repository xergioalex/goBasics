package main

import"fmt"

func main() {
	var (
		num1  int32
		num2  int32
		suma  int32
		resta int32
		mult  int32
		divis float32
	)

	fmt.Printf("\n-------------------\n")
	fmt.Printf("    CALCULADORA    \n")
	fmt.Printf("-------------------\n\n")

	fmt.Print("Ingresa el primer operando: ")
	fmt.Scanf("%d", &num1)
	fmt.Print("Ingresa el segundo operando: ")
	fmt.Scanf("%d", &num2)

	suma, resta, mult, divis = calcular(num1, num2)

	fmt.Printf("\n-------------------\n")
	fmt.Printf("    RESULTADOS    \n")
	fmt.Printf("-------------------\n\n")
	fmt.Printf("Suma: %d\nResta: %d\nMultiplicacion: %d \nDivision: %.3f \n", suma, resta, mult, divis)
}

func calcular(a int32, b int32) (int32, int32, int32, float32) {
	return (a + b), (a - b), (a * b), (float32(a) / float32(b))
}