package gobmi

import (
	"github.com/shopspring/decimal"
	"testing"
)

func TestCalcBMI(t *testing.T) {
	inputHeight, inputWeight := 1.70, 65.0
	expectedOutput := 22.49
	t.Logf("start calculating bmi: \nheight: %.2f \nweight: %.2f \nexpecting bmi = %.2f.\n", inputHeight, inputWeight, expectedOutput)
	actualOutput, err := CalcBMI(inputWeight, inputHeight)
	if err != nil {
		t.Fatalf("major error emerges while calculating bmi: %v\n", err)
	}
	actualFormatOutput, _ := decimal.NewFromFloat(actualOutput).RoundFloor(2).Float64()
	if expectedOutput != actualFormatOutput {
		t.Logf("expecting %f, but got %f.\n", expectedOutput, actualFormatOutput)
		t.Failed()
	}
}
