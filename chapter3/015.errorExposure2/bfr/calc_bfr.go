package bfr

import "fmt"

func CalcBFR(height, weight float64, age int, gender string) (bfr float64, err error) {
	bmi, errorOut := CalcBMI(height, weight)
	if errorOut != nil {
		fmt.Println("warning: bmi is not calculated correctly, stop this procedure", errorOut)
	} else {
		bfr = (1.2*bmi + 0.23*float64(age) - 5.4 - 10.8*getGenderWeight(gender)) / 100
	}
	return
}

func getGenderWeight(gender string) (genderWeight float64) {
	if gender == "male" {
		genderWeight = 1.0
	} else {
		genderWeight = 0
	}
	return
}
