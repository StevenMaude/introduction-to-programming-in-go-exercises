package main

import "fmt"

func main() {
	fmt.Println(half(1))
	fmt.Println(half(2))
}

func half(n int) (halfN int, even bool) {
	halfN = n / 2
	if halfN == 0 {
		even = true
	} else {
		even = false
	}
	return halfN, even
}
