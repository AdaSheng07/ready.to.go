package bfr

import "math"

func CalcBMI(height, weight float64) (bmi float64, err error) {
	bmi = weight / math.Pow(height, 2)
	return
}
