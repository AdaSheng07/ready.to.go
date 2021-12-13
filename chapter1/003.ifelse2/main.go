package main

import "fmt"

/*
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
	var money int
	if money >= 20 {
		fmt.Println("点个外卖")
	} else if money >= 200 {
		fmt.Println("去下馆子")
	} else if money >= 2000 {
		fmt.Println("去米其林")
	} else if money >= 20000 {
		fmt.Println("出国玩一圈")
	} else {
		fmt.Println("饿着吧")
	}
}
