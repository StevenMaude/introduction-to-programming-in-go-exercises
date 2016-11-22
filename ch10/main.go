package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("FOO")
	Sleep(5)
	fmt.Println("BAR")
}

func Sleep(n int) {
	select {
	case <-time.After(time.Duration(n) * time.Second):
	}
}
