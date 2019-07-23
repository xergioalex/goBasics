package main

import "fmt"
import "time"

func helloGo(index int, text string) {
	fmt.Printf("%s: Hello, I am a go routine print #%d\n", text, index)
}

func forGo(n int, text string) {
	for i := 0; i < n; i++ {
		go helloGo(i, text)
	}
}

func main () {
	forGo(500, "For 1")
	forGo(400, "For 2")
	time.Sleep(1000 * time.Millisecond);
}