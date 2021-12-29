package main

import (
	"fmt"
	"learn.go/chapter3/015.errorExposure2/bfr"
	"runtime/debug"
)

func main() {
	defer recoverMainBody()

	name, gender, height, weight, age, err := bfr.InputInfo()
	fmt.Printf("Your input:\nname:%s\n", name)
	fmt.Printf("gender(male/female): %s\n", gender)
	fmt.Printf("height: %.2f\n", height)
	fmt.Printf("weight: %.2f\n", weight)
	fmt.Printf("age: %d\n", age)
	fmt.Println("error:", err)

	fatRate, err := bfr.CalcBFR(height, weight, age, gender)
	if err != nil {
		fmt.Println("warning: bfr is wrongly calculated, stop this procedure", err)
	} else if fatRate <= 0 {
		panic("body fat rate must be positive")
	} else {
		fmt.Printf("%s's body fat rate is %v%%.", name, fatRate*100)
	}
}

func recoverMainBody() {
	if re := recover(); re != nil {
		fmt.Printf("warning: catch major error: %v\n", re)
		debug.PrintStack()
	}
}
