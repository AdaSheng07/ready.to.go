package calc

import "math"

func CalcBMI(height, weight float64) (bmi float64) {
	return weight / math.Pow(height, 2)
}
