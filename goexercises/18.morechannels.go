package main

import "fmt"
import "time"

func ping(ball chan<- int, action chan<- string) {
	ball <-1
	action <- "Player ping"
}

func pong(ball chan<- int, action chan<- string) {
	ball <- 2
	action <- "Player pong"
}

func referee(action <-chan string) {
	for {
		fmt.Println("Action: ", <- action)
	}
}

func pingpong() {
	ball := make(chan int)
	action := make(chan string)
	go referee(action)
	go ping(ball, action)

	for {
		value := <- ball
		switch value {
			case 1:
				go pong(ball, action)
			case 2:
				go ping(ball, action)
		}
	}
}

func main () {
	go pingpong()
	time.Sleep(10 * time.Second);
}