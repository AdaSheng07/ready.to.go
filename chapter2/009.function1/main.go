package main

import (
	"fmt"
)

func main() {
	// fixed-length & unfixed-length parameters
	helloToOnePerson := constructHello1("Tom")
	helloToPersons := constructHello2("Tom", "David")
	fmt.Println(helloToOnePerson)
	fmt.Println(helloToPersons)

	bmi := []float64{0.203, 0.242, 0.261}
	aveBmi1 := calculateAverage1(bmi)
	fmt.Println("Average BMI:", aveBmi1)

	// we can send no parameters to functions w unfixed-length parameter
	aveBmi2 := calculateAverage2()
	fmt.Println(aveBmi2) // NaN
	aveBmi2 = calculateAverage2(bmi...)
	fmt.Println(aveBmi2)

	// naming return values is a good habit
	// 1. named return values can be used directly in main function
	// 2. named return values also enables people who use this piece of code understand it quickly
	sexOfStudent, ageOfStudent, heightOfStudent := getInfoOfStudents("Tom")
	fmt.Println("Tom:\nsex:", sexOfStudent, "\nage:", ageOfStudent, "\nheight:", heightOfStudent, "m")

}

func getInfoOfStudents(stdName string) (sex string, age int, height float64) {
	return "Male", 18, 1.70
}

func constructHello1(name string) string {
	return fmt.Sprintf("hello, %s", name)
}
func constructHello2(name ...string) string {
	return fmt.Sprintf("hello, %s", name)
}

func calculateAverage1(nums []float64) float64 {
	var total float64 = 0
	for _, num := range nums {
		total += num
	}
	return total / float64(len(nums))
}

func calculateAverage2(nums ...float64) float64 {
	var total float64 = 0
	for _, num := range nums {
		total += num
	}
	return total / float64(len(nums))
}
