package gobmi

func getGenderWeight(gender string) (genderWeight float64) {
	if gender == "男" {
		genderWeight = 1.0
	} else {
		genderWeight = 0
	}
	return
}
