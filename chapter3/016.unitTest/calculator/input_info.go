package calculator

import "fmt"

func InputInfo() (name string, gender string, height float64, weight float64, age int, err error) {
	name = ""
	gender = ""
	weight = 0
	height = 0
	age = 0
	fmt.Println("input your name:")
	_, _ = fmt.Scanln(&name)
	if name == "" {
		err = fmt.Errorf("there must be a name")
		return
	}

	fmt.Print("input your gender: ")
	_, _ = fmt.Scanln(&gender)
	if gender != "male" && gender != "female" {
		err = fmt.Errorf("gender: %s is neither male nor female", gender)
		return
	}

	fmt.Print("input your weight: ")
	_, _ = fmt.Scanln(&weight)
	if weight <= 10 || weight > 200 {
		err = fmt.Errorf("weight: %.2f is out of range", weight)
		return
	}

	fmt.Print("input your height: ")
	_, _ = fmt.Scanln(&height)
	if height <= 0.8 || height > 2.5 {
		err = fmt.Errorf("height: %.2f is out of range", height)
		return
	}

	fmt.Print("input your age: ")
	_, _ = fmt.Scanln(&age)
	if age <= 18 || age > 150 {
		err = fmt.Errorf("age: %d is out of range", age)
		return
	}
	return
}
