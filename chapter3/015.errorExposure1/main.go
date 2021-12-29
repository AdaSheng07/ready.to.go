package main

import "fmt"

func main() {
	gender, weight, err := inputInfo()
	fmt.Printf("Your input:\ngender: %s\n", gender)
	fmt.Printf("weight: %.2f\n", weight)
	fmt.Println("error:", err)
}

func inputInfo() (sex string, weight float64, err error) {
	sex = ""
	weight = 0
	fmt.Print("input your gender: ")
	_, _ = fmt.Scanln(&sex)
	if sex != "male" && sex != "female" {
		err = fmt.Errorf("gender: %s is neither male nor female", sex)
		return
	}
	fmt.Print("input your weight: ")
	_, _ = fmt.Scanln(&weight)
	if weight <= 10 || weight > 200 {
		err = fmt.Errorf("weight: %.2f is out of range", weight)
		return
	}
	return
}
