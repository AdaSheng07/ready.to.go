package gobmi

import (
	"github.com/shopspring/decimal"
	"testing"
)

func TestGiveOutSuggestions(t *testing.T) {
	inputHeight, inputWeight := 1.70, 65.0
	inputGender := "male"
	inputAge := 35
	expectedBMI := 22.49
	expectedBFR := 0.1883
	expectedSuggestion := "目前的体脂率水平处于【偏胖】，建议：少坐多动，注意健康饮食"

	t.Logf("start giving out suggestions: \ngender: %s \nage: %d \nheight: %.2f \nweight: %.2f \n", inputGender, inputAge, inputHeight, inputWeight)
	actualBMI, err1 := CalcBMI(inputWeight, inputHeight)
	if err1 != nil {
		t.Fatalf("major error emerges while calculating bmi: %v\n", err1)
	}
	actualFormatBMI, _ := decimal.NewFromFloat(actualBMI).RoundFloor(2).Float64()
	if expectedBMI != actualFormatBMI {
		t.Logf("expecting %f, but got %f.\n", expectedBMI, actualFormatBMI)
		t.Failed()
	}

	actualBFR, err2 := CalcBFR(actualBMI, inputAge, inputGender)
	if err2 != nil {
		t.Fatalf("major error emerges while calculating body fat rate: %v\n", err2)
	}
	actualFormatBFR, _ := decimal.NewFromFloat(actualBFR).RoundFloor(4).Float64()
	if expectedBFR != actualFormatBFR {
		t.Logf("expecting %f, but got %f.\n", expectedBFR, actualFormatBFR)
		t.Failed()

	}

	actualSuggestion := GiveOutSuggestions(actualBFR, inputGender, inputAge)
	if expectedSuggestion != actualSuggestion {
		t.Logf("expecting %s, but got %s.\n", expectedSuggestion, actualSuggestion)
		t.Failed()
	}
}
