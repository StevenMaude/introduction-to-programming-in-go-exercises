package main

import "fmt"
import "mymath"

func main() {
	xs := []float64{1, 2, 3, 4, 6, 3, 2, 8, 9, 6}
	avg := mymath.Average(xs)
	min := mymath.Min(xs)
	max := mymath.Max(xs)
	fmt.Println(avg, min, max)
}
