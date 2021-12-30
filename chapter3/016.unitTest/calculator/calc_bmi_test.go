package calculator

import (
	"github.com/shopspring/decimal"
	"testing"
)

func TestCalcBMI(t *testing.T) {
	inputHeight, inputWeight := 1.70, 65.0
	expectedOutput := 22.49
	t.Logf("start calculating bmi: input height = %.2f, weight = %.2f, expecting bmi = %.2f.\n", inputHeight, inputWeight, expectedOutput)
	actualOutput, err1 := CalcBMI(inputHeight, inputWeight)
	if err1 != nil {
		t.Fatalf("major error: bmi is wrongly calculated, expecting no error but got %v", err1)
	}
	actualFormatOutput, _ := decimal.NewFromFloat(actualOutput).RoundFloor(2).Float64()
	if expectedOutput != actualFormatOutput {
		t.Logf("expecting %f, but got %f.\n", expectedOutput, actualFormatOutput)
		t.Failed()
		// 或者
		// t.Errorf("expecting %f, but got %f.\n", expectedOutput, actualFormatOutput)
	}
}
