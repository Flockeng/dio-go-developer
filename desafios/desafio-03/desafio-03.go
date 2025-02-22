package main

import (
	"fmt"
	"time"
)

func ping(c chan string) {
	for {
		c <- "ping"
	}
}

func pong(c chan string) {
	for {
		c <- "pong"
	}
}

func print(c chan string) {
	for {
		fmt.Println(<-c)
		time.Sleep(time.Second)
	}
}

func main() {
	c := make(chan string)

	go ping(c)
	go print(c)
	go pong(c)

	var entrada string
	fmt.Scanln(&entrada)
}
