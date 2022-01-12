package main

import (
	"fmt"
)

// do calculation on integers: if we do not use structure...
func main() {
	var left, right int
	var operator string
	fmt.Scanln(&left)
	fmt.Scanln(&operator)
	fmt.Scanln(&right)
	switch operator {
	case "+":
		fmt.Println(left + right)
	case "-":
		fmt.Println(left - right)
	case "*":
		fmt.Println(left * right)
	case "/":
		fmt.Println(left / right)
	case "%":
		fmt.Println(left % right)
	// if we want to add another type of calculation, we need to add another case
	// switch-case makes a whole, which is not convenient to test
	default:
		fmt.Println("this calculation is not supported.")
	}
}

type Calculator struct {
	left, right int
	operator    string
	Add         func(a, b int) (sum int)
}
