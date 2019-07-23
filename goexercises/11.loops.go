package main

import "fmt"

func main() {
	forLoop()
}

func forLoop() {
	for i := 0; i < 5; i++ {
		fmt.Println("[FOR] ", i)
	}

	c := 100
	for c > 0 {
		c -= 10
		fmt.Println("[While] c > 0: ", c)
	}

	s := 1000
	for {
		s -= 1
		if s == 0 {
			fmt.Println("[While] infinite until break condition")
			break
		}
	}
}