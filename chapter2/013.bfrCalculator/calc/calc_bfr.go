package calc

func CalcBFR(bmi float64, age int, sexWeight int) (bfr float64) {
	bfr = (1.2*bmi + getAgeWeight(age)*float64(age) - 5.4 - 10.8*float64(sexWeight)) / 100
	return
}

// 计算公式中的数字是常量，需要把比重提取为函数
func getAgeWeight(age int) (ageWeight float64) {
	ageWeight = 0.23
	if age >= 30 && age <= 40 {
		ageWeight = 0.22
	}
	return
}
