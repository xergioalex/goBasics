package main

import"fmt"

func main() {
  var (
    num1  int32
    num2  int32
    sumValue  int32
    substrationValue int32
    multValue  int32
    divisionValue float32
  )

  fmt.Printf("\n-------------------\n")
  fmt.Printf("    CALCULATOR    \n")
  fmt.Printf("-------------------\n\n")

  fmt.Print("Type the first value: ")
  fmt.Scanf("%d", &num1)
  fmt.Print("Typer the second value: ")
  fmt.Scanf("%d", &num2)

  sumValue, substrationValue, multValue, divisionValue = calc(num1, num2)

  fmt.Printf("\n-------------------\n")
  fmt.Printf("    RESULTADOS    \n")
  fmt.Printf("-------------------\n\n")
  fmt.Printf("Sum(+): %d\nSubscrtation(-): %d\nMult(*): %d \nDivision(/): %.3f \n", sumValue, substrationValue, multValue, divisionValue)
}

func calc(a int32, b int32) (int32, int32, int32, float32) {
  return sum(a, b), substration(a, b), mult(a, b), division(a, b)
}

func sum(a int32, b int32) int32 {
  return a + b
}

func substration(a int32, b int32) int32 {
  return a - b
}

func mult(a int32, b int32) int32 {
  return int32(a) * int32(b)
}

func division(a int32, b int32) float32 {
  return float32(a) / float32(b)
}