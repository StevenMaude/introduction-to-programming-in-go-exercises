package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		line := ""
		if i%3 == 0 {
			line += "Fizz"
		}
		if i%5 == 0 {
			line += "Buzz"
		}
		if line != "" {
			fmt.Println(line)
		} else {
			fmt.Println(i)
		}
	}
}
