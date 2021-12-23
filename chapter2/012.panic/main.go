package main

import "fmt"

func main() {
	var name string
	var weight float64
	var height float64
	var age int
	fmt.Println("input your name: ")
	_, _ = fmt.Scanln(&name)
	fmt.Println("input your weight(kg): ")
	_, _ = fmt.Scanln(&weight)
	// use panic to stop the procedure if conditions are not satisfied
	if weight <= 0 {
		panic("weight cannot be negative, please input the right value. ")
	}
	fmt.Println("input your height(m): ")
	_, _ = fmt.Scanln(&height)
	// use panic to stop the procedure if conditions are not satisfied
	if height <= 0 {
		panic("height cannot be negative, please input the right value. ")
	}
	fmt.Println("input your age: ")
	_, _ = fmt.Scanln(&age)
	// use panic to stop the procedure if conditions are not satisfied
	if age <= 0 {
		panic("age cannot be negative, please input the right value. ")
	}
	fmt.Printf("name: %s, age: %d, height: %.2fm, weight: %2.fkg.", name, age, height, weight)
}
