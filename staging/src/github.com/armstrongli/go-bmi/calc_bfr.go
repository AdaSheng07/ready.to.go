package gobmi

import "fmt"

func CalcBFR(bmi float64, age int, gender string) (bfr float64) {
	bfr = (1.2*bmi + 0.23*float64(age) - 5.4 - 10.8*getGenderWeight(gender)) / 100
	fmt.Printf("BMI: %.2f%%\nBFR: %.2f%%\n", bmi, bfr*100)
	return
}
