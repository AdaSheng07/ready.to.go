package main

import "fmt"

func main() {

	// round 1
	fmt.Println("Round 1")
	var fruit1 = "six apples"
	watermelon := true
	if watermelon {
		fruit1 = "an apple"
	}
	fmt.Printf("I bought %s.\n", fruit1)

	// round 2
	fmt.Println("Round 2")
	var fruit2 = "six apples"
	banana := false
	if !banana {
		fruit2 = "a watermelon and two bananas"
	} else {
		fruit2 = "four apples"
	}
	fmt.Printf("I bought %s.\n", fruit2)

	// round 3
	fmt.Println("Round 3")
	var fruit3 = "strawberries"
	var applePrice = 4
	if applePrice >= 7 {
		fmt.Printf("I only bought %s.\n", fruit3)
	} else if applePrice > 5 {
		fmt.Printf("I bought %s and two apples.\n", fruit3)
	} else {
		fmt.Printf("I bought %s and four apples.\n", fruit3)
	}
}
