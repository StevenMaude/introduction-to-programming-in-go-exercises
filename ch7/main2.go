package main

import "fmt"

func main() {
	fmt.Println(greatest(5, 3, 36, 2))
	fmt.Println(greatest(1, 2, 3, 4))
}

func greatest(ints ...int) int {
	greatest := ints[0]
	for _, value := range ints[1:] {
		if value > greatest {
			greatest = value
		}

	}
	return greatest
}
