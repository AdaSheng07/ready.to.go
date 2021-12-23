package main

import (
	"fmt"
	"time"
)

var times int

// what if we return a function defined in another function body and using variables outside its scope?
func calcSumFunc() func(nums ...int) {
	var sum int
	var weightedSum float64
	weight := 0.5
	return func(nums ...int) {
		for _, value := range nums {
			sum += value
		}
		weightedSum = float64(sum) * weight
		times++
		fmt.Println(sum, weightedSum)
	}
}

// function calcSum is bound with its context/environment, including sum and counter
// that is why sum and counter are accessible in main() and function countUsage()
func main() {
	closureMain()
	fmt.Println("sleep some while")
	time.Sleep(5 * time.Second)

	// a function variable is declared to calcWeightedSum, calcWeightedSum is a closure
	// calcWeightedSum keeps its citation to weight, sum and times.
	// weight, sum and times seems to escape, but their life cycles have not ended yet
	calcWeightedSum := calcSumFunc()
	calcWeightedSum(1, 2, 3, 4, 5, 6, 7, 8, 9) // 45 22.5
	calcWeightedSum(1, 2, 3, 4, 5, 6, 7, 8, 9) // 90 45
	calcWeightedSum(1, 2, 3, 4, 5, 6, 7, 8, 9) // 135 67.5
	calcWeightedSum(1, 2, 3, 4, 5, 6, 7, 8, 9) // 180 90
	calcWeightedSum(1, 2, 3, 4, 5, 6, 7, 8, 9) // 225 112.5
	calcWeightedSum(1, 2, 3, 4, 5, 6, 7, 8, 9) // 270 135
	fmt.Println(times)                         // 6
	// here we call calcSumFunc for another five times
	// they return five closures, they cite different sums and weights (even though they have the same values)
	// times still accumulates because it is a global variable
	calcSumFunc()(1, 2, 3, 4, 5, 6, 7, 8, 9) // 45 22.5
	calcSumFunc()(1, 2, 3, 4, 5, 6, 7, 8, 9) // 45 22.5
	calcSumFunc()(1, 2, 3, 4, 5, 6, 7, 8, 9) // 45 22.5
	calcSumFunc()(1, 2, 3, 4, 5, 6, 7, 8, 9) // 45 22.5
	calcSumFunc()(1, 2, 3, 4, 5, 6, 7, 8, 9) // 45 22.5
	fmt.Println(times)                       // 11
}
