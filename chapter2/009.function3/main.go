package main

import "fmt"

func main() {
	// special variable: function
	// declaration
	var add func(float64, float64) float64
	add = addition
	fmt.Printf("%.2f\n", add(23.5, 45.24))
	subtract := func(num1 float64, num2 float64) float64 {
		return num1 - num2
	}
	fmt.Printf("%.2f\n", subtract(123.34, 43.23))
}
func addition(num1 float64, num2 float64) float64 {
	return num1 + num2
}
