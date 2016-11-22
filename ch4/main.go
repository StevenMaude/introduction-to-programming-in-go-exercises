package main

import "fmt"

func main() {
	fmt.Printf("Enter a Fahrenheit temperature: ")
	var fahrenheit float64
	fmt.Scanf("%f", &fahrenheit)
	celsius := (fahrenheit - 32) * 5 / 9
	fmt.Printf("%v F is %v Celsius\n", fahrenheit, celsius)
}
