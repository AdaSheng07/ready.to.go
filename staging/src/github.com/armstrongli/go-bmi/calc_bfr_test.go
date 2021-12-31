package gobmi

import (
	"github.com/shopspring/decimal"
	"testing"
)

func TestCalcBFR(t *testing.T) {
	inputBmi := 22.49
	inputAge := 35
	inputGender := "male"
	expectedOutput := 0.1883
	t.Logf("start calculating bfr(body fat rate): \ngender: %s \nage: %d \nbmi: %.2f \nexpectedOutput: %.4f ", inputGender, inputAge, inputBmi, expectedOutput)
	actualOutput, err := CalcBFR(inputBmi, inputAge, inputGender)
	if err != nil {
		t.Fatalf("major error while calculating bfr: %v\n", err)
	}
	actualFormatOutput, _ := decimal.NewFromFloat(actualOutput).RoundFloor(4).Float64()
	if actualFormatOutput != expectedOutput {
		t.Logf("expecting %f, but got %f.\n", expectedOutput, actualFormatOutput)
		t.Failed()

	}

}
