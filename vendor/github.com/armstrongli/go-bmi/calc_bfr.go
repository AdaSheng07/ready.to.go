package gobmi

import "fmt"

func CalcBFR(bmi float64, age int, gender string) (bfr float64, err error) {
	if bmi <= 0 {
		err = fmt.Errorf("bmi must be positive")
		return
	}
	if age <= 0 {
		err = fmt.Errorf("age must be a positive integer")
		return
	}
	if age > 150 {
		err = fmt.Errorf("age must be reasonable")
		return
	}
	if gender != "male" && gender != "female" {
		err = fmt.Errorf("gender must be male or female")
		return
	}

	bfr = (1.2*bmi + 0.23*float64(age) - 5.4 - 10.8*getGenderWeight(gender)) / 100
	fmt.Printf("BMI: %.2f%%\nBFR: %.2f%%\n", bmi, bfr*100)
	return
}
