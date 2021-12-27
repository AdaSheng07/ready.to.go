package main

import (
	"fmt"
	_ "github.com/spf13/cobra"
	"os"
)

func main() {
	var (
		name   string
		sex    string
		height string // float64
		weight string // float64
		age    string // int
	)
	arguments := os.Args
	fmt.Println(arguments)
	// [/private/var/folders/7l/pnr12b8962l7dwhrxxzjd9500000gn/T/GoLand/___go_build_learn_go_chapter2_014_bfrCmdCalculator 小强 男 1.7 70 35]
	// this is a slice, arguments[0] is the program itself, so index of variables starts from 1
	name = arguments[1]
	sex = arguments[2]
	height = arguments[3]
	weight = arguments[4]
	age = arguments[5]

	fmt.Println("name: ", name)
	fmt.Println("sex: ", sex)
	fmt.Println("height: ", height)
	fmt.Println("weight: ", weight)
	fmt.Println("age: ", age)
}
