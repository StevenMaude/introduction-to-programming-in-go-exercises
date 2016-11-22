package main

import "fmt"

func swap(x *int, y *int) {
	*x, *y = *y, *x
}

func main() {
	x, y := 5, 6
	fmt.Println(x, y)
	swap(&x, &y)
	fmt.Println(x, y)
}
