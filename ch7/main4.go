package main

import "fmt"

var Cache = make(map[int]int)

func main() {
	Cache[0] = 0
	Cache[1] = 1
	fmt.Println(fastFib(90))
	fmt.Println(rubbishFib(40))
}

func fastFib(n int) int {
	if value, ok := Cache[n]; ok {
		return value
	} else {
		Cache[n] = fastFib(n-1) + fastFib(n-2)
		return Cache[n]
	}
}

func rubbishFib(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	return rubbishFib(n-1) + rubbishFib(n-2)
}
