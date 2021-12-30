package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"learn.go/chapter2/013.bfrCalculator/calc" // package name is also cobra
)

func main() {
	var (
		name   string
		sex    string
		height float64
		weight float64
		age    int
	)
	cmd := cobra.Command{
		Use:   "bfrCheck",    // use: set a name for this command
		Short: "基于BMI的体脂计算器", // short: a short description for this command
		// long: a detailed explanation for this command
		Long: "录入姓名、性别、身高、体重和年龄，计算他们的BMI值，基于他们的性别和年龄生成他们的体脂率标准，判断他们的体脂率处于偏瘦/标准/偏胖/严重肥胖并给出健康建议。",
		// func is a registered callback function, Run is the main body of cmd, custom-made as desired
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("name: ", name)
			fmt.Println("sex: ", sex)
			fmt.Println("height: ", height)
			fmt.Println("weight: ", weight)
			fmt.Println("age: ", age)

			// calculate bmi & calculator...
			bmi := calc.CalcBMI(height, weight)
			bfr := calc.CalcBFRUpgrade(bmi, age, sex)
			fmt.Println("Body Fat Rate:", bfr)

			// healthiness assessment & suggestions...

		},
	}
	/*
		func (f *FlagSet) StringVar(p *string, name string, value string, usage string)
		StringVar defines a string flag with specified name, default value, and usage string.
		The argument p points to a string variable in which to store the value of the flag.
	*/
	cmd.Flags().StringVar(&name, "name", "", "姓名")
	// this means: when we type in command line, what comes after "name" will be saved in variable name as string.
	// if we give nothing after "name", the variable name will be the default value "".
	cmd.Flags().StringVar(&sex, "sex", "", "性别")
	cmd.Flags().Float64Var(&height, "height", 0, "身高")
	cmd.Flags().Float64Var(&weight, "weight", 0, "体重")
	cmd.Flags().IntVar(&age, "age", 0, "年龄")

	// run the defined cmd object
	cmd.Execute()
}
