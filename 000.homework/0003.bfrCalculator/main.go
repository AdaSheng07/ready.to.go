package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"learn.go.tools"
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
		Use:   "calcBFR",
		Short: "A body fat rate calculator based on BMI value.",
		Long: "Given a person's name, gender, height, weight and age, calculate this person's BMI and BFR. \n" +
			"Based on this person's gender and age, set up the BFR benchmark for this person. \n" +
			"Check for this peron's healthy status and give out specialized suggestions.",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("name:", name)
			fmt.Println("gender:", gender)
			fmt.Println("height:", height)
			fmt.Println("weight:", weight)
			fmt.Println("age:", age)

			bmi, bfr := learn_go_tools.CalcBfr(age, gender, weight, height)
			fmt.Println("bmi:", bmi, "\nbfr:", bfr)
		},
	}
	cmd.Flags().StringVar(&name, "name", "", "input name")
	cmd.Flags().StringVar(&gender, "gender", "", "input gender")
	cmd.Flags().Float64Var(&height, "height", 0, "input height")
	cmd.Flags().Float64Var(&weight, "weight", 0, "input weight")
	cmd.Flags().IntVar(&age, "age", 0, "input age")

	_ = cmd.Execute()
}
