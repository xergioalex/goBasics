package main

import "fmt"
import "time"

func helloGo(index int) {
	fmt.Println("Hola soy un print en la Go routine #", index)
}

func forGo(n int) {
	for i := 0; i < n; i++ {
		go helloGo(i)
	}
}

func main () {
	forGo(500)
	forGo(400)
	time.Sleep(1000 * time.Millisecond);
}