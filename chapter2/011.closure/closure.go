package main

import "fmt"

var counter int
var count int

// function calcSum is bound with its context/environment, including sum and counter
// that is why sum and counter are accessible in main() and function countUsage()
func calcSum(nums ...int) (sum int) {
	for _, num := range nums {
		sum += num
	}
	counter++
	return
}

// when we call calcSum function, its context has been associated with itself, e.g. counter
// so the value of counter is accessible
func returnCallTimes() {
	fmt.Printf("function calcSum is used %v times.\n", counter)
}

// what if a function is returned?
func generateImprovementOfCpu() func(percentage float64) {
	base := 1000.00
	// when declared
	return func(percentage float64) {
		base = base * (1 + percentage)
		count++
		fmt.Println(base)
	}

}

func closureMain() {
	// call function calSum defined
	fmt.Println(calcSum(1, 2, 3, 5, 4, 6, 45, 23, 5))
	fmt.Println(calcSum(1, 2, 3, 5, 4, 6, 45, 23, 5))
	fmt.Println(calcSum(1, 2, 3, 5, 4, 6, 45, 23, 5))
	fmt.Println(calcSum(1, 2, 3, 5, 4, 6, 45, 23, 5))
	fmt.Println(calcSum(1, 2, 3, 5, 4, 6, 45, 23, 5))
	fmt.Println(counter)
	// when we call calcSum function, its context has been associated with itself, e.g. counter
	// so the value of counter is accessible
	returnCallTimes()

	calcImprovement := generateImprovementOfCpu()
	calcImprovement(0.1)
	calcImprovement(0.1)
	calcImprovement(0.1)
	calcImprovement(0.1)
	fmt.Println(count)
}
