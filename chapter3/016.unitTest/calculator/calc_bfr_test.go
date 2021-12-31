package calculator

import (
	"github.com/shopspring/decimal"
	"testing"
)

func TestCalcBFR(t *testing.T) {
	inputHeight := 1.70
	inputWeight := 65.00
	inputAge := 35
	inputGender := "male"
	expectedOutput := 0.1883
	t.Logf("start calculating bfr(body fat rate): \ngender: %s \nage: %d \nheight = %.2f \nweight = %.2f \nexpectedOutput: %.4f ", inputGender, inputAge, inputHeight, inputWeight, expectedOutput)
	actualOutput, err := CalcBFR(inputHeight, inputWeight, inputAge, inputGender)
	if err != nil {
		t.Fatalf("major error: bfr is wrongly calculated, expecting no error but got %v", err)
	}
	actualFormatOutput, _ := decimal.NewFromFloat(actualOutput).RoundFloor(4).Float64()
	if actualFormatOutput != expectedOutput {
		t.Logf("expecting %f, but got %f.\n", expectedOutput, actualFormatOutput)
		t.Failed()
	}
}
