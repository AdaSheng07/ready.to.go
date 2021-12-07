package main

import "fmt"

/*
if语句语法：
	if 布尔表达式 {
   		// 在布尔表达式为 true 时执行
	}

if-else语句语法：
	if 布尔表达式 {
   		// 在布尔表达式为 true 时执行
	} else {
		// 在布尔表达式为 false 时执行
	}

if语句嵌套语法：
	if 布尔表达式 1 {
   		// 在布尔表达式 1 为 true 时执行
		if 布尔表达式 2 {
			// 在布尔表达式 2 为 true 时执行
		}
	}

else-if语句嵌套语法：
	if 布尔表达式 1 {
   		// 在布尔表达式 1 为 true 时执行
	} else if 布尔表达式 2 {
		// 在布尔表达式 1 为 false 且布尔表达式 2 为 true 时执行
	} else if 布尔表达式 3 {
		// 在布尔表达式 1 为 false, 布尔表达式 2 为 false, 且布尔表达式 3 为 true 时执行
	} else {
		// 在布尔表达式 1,2,3 均为 false 时执行
	}

*/

func main() {

	// round 1
	fmt.Println("Round 1")
	var fruit1 = "six apples"
	watermelon := true
	if watermelon {
		fruit1 = "an apple"
	}
	fmt.Printf("I bought %s.\n", fruit1)

	// round 2
	fmt.Println("Round 2")
	var fruit2 = "six apples"
	banana := false
	if !banana {
		fruit2 = "a watermelon and two bananas"
	} else {
		fruit2 = "four apples"
	}
	fmt.Printf("I bought %s.\n", fruit2)

	// round 3
	fmt.Println("Round 3")
	var fruit3 = "strawberries"
	var applePrice = 4
	if applePrice >= 7 {
		fmt.Printf("I only bought %s.\n", fruit3)
	} else if applePrice > 5 {
		fmt.Printf("I bought %s and two apples.\n", fruit3)
	} else {
		fmt.Printf("I bought %s and four apples.\n", fruit3)
	}
}
