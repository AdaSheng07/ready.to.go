package main

import "fmt"

func main() {
	// callback function
	// use a function as a formal parameter in another function
	// example: addition, subtraction, multiplication, division
	var num1, num2 float64
	fmt.Println("Input number 1 and number 2 values: ")
	_, _ = fmt.Scan(&num1, &num2)
	operation(num1, num2, add)      // use function add as a formal parameter
	operation(num1, num2, subtract) // use function subtract as a formal parameter
	operation(num1, num2, multiply) // use function multiply as a formal parameter
	operation(num1, num2, divide)   // use function divide as a formal parameter

	result1 := compoundOperation1()
	fmt.Println(result1(num1, num2))
	result2 := compoundOperation2(2)
	fmt.Println(result2(num1, num2))
}

func operation(a float64, b float64, f func(a float64, b float64)) {
	f(a, b)
}
func add(a float64, b float64) {
	fmt.Println("a + b = ", a+b)
}
func subtract(a float64, b float64) {
	fmt.Println("a - b = ", a-b)
}
func multiply(a float64, b float64) {
	fmt.Println("a * b = ", a*b)
}
func divide(a float64, b float64) {
	fmt.Println("a / b = ", a/b)
}

// return a function 1
func compoundOperation1() func(a float64, b float64) float64 {
	// custom-made
	return func(a float64, b float64) float64 {
		return (a + b) * (a - b)
	}
}

// return a function 2: delta is a formal parameter for compoundOperation2() to generate a new function
func compoundOperation2(delta float64) func(a float64, b float64) float64 {
	// custom-made
	return func(a float64, b float64) float64 {
		return (a+b)*(a+b) - 2*a*b + delta
	}
}
