package main

import (
	"fmt"
	gobmi "github.com/armstrongli/go-bmi"
)

func main() {

	persons := []Person{
		{
			"小强",
			"男",
			1.7,
			70,
			35,
		},
	}
	for _, item := range persons {
		bmi, err := gobmi.CalcBMI(item.weight, item.height)
		fmt.Println(bmi, err)
	}

}

type Person struct {
	name   string
	gender string
	height float64
	weight float64
	age    int
}
