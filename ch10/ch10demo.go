package main

import (
	"fmt"
	"time"
)

func pinger(c chan<- string) {
	for {
		c <- "ping"
	}
}

func printer(c <-chan string) {
	for {
		fmt.Println(<-c)
		time.Sleep(time.Second * 1)
	}
}

func ponger(c chan<- string) {
	for {
		c <- "pong"
	}
}

func main() {
	c := make(chan string)
	go pinger(c)
	go ponger(c)
	go printer(c)
	var input string
	fmt.Scanln(&input)
}
