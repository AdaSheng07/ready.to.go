package main

import (
	"fmt"
	gobmi "github.com/armstrongli/go-bmi"
	"github.com/spf13/cobra"
)

func main() {
	var (
		name   string
		gender string
		height float64
		weight float64
		age    int
	)
	cmd := cobra.Command{
		Use:   "CheckHealthyStatus",
		Short: "A body fat rate calculator based on BMI value.",
		Long: "Given a person's name, gender, height, weight and age, calculate this person's BMI and BFR. \n" +
			"Based on this person's gender and age, set up the BFR benchmark for this person. \n" +
			"Check for this peron's healthy status and give out specialized suggestions.",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Name:", name)
			fmt.Println("Gender(male/female):", gender)
			fmt.Println("Height(m):", height)
			fmt.Println("Weight(kg):", weight)
			fmt.Println("Age:", age)

			// calculate bmi & calculator and print them out
			bmi, _ := gobmi.CalcBMI(weight, height)
			bfr, _ := gobmi.CalcBFR(bmi, age, gender)

			// healthiness assessment & suggestions and print them out
			gobmi.GiveOutSuggestions(bfr, gender, age)
		},
	}
	cmd.Flags().StringVar(&name, "name", "", "input name")
	cmd.Flags().StringVar(&gender, "gender", "", "input gender")
	cmd.Flags().Float64Var(&height, "height", 0, "input height")
	cmd.Flags().Float64Var(&weight, "weight", 0, "input weight")
	cmd.Flags().IntVar(&age, "age", 0, "input age")

	_ = cmd.Execute()
}
