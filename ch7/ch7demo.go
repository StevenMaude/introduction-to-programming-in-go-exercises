package main

import "fmt"

var cache = make(map[int64]int64)

func fib(x int64) int64 {
	if cached, ok := cache[x]; ok {
		return cached
	}
	cache[x] = fib(x-1) + fib(x-2)
	return cache[x]
}

func main() {
	cache[0] = 0
	cache[1] = 1
	fmt.Println(fib(600))
}
