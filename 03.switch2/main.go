package main

import (
	"fmt"
	"strconv"
)

/*
switch与if-else的比较
1. switch可以完全替代if-else条件表达式，且分支表达式可以做不一样的事，switch的分支表达式可以随时终止
2. switch可以并入其他条件分支继续执行，e.g. case 1 执行完毕后，直接调入 case 2 执行，并且不需要符合 case 2 条件：fallthrough
3. fallthrough的情况在状态机中经常存在，用于状态的流转
*/

func main() {
	var money float64
	var busy bool
	var num1, num2 string
	fmt.Println("兜里有多少钱？")
	fmt.Scanln(&num1)
	money, _ = strconv.ParseFloat(num1, 64)
	fmt.Println("忙不忙？")
	fmt.Scanln(&num2)

	switch {
	case num2 == "Y" || num2 == "y" || num2 == "忙":
		busy = true
	case num2 == "N" || num2 == "n" || num2 == "不忙":
		busy = false
	}

	switch {
	case money == 0:
		fmt.Println("饿着吧")
	case money > 0 && money <= 20:
		fmt.Println("点个外卖")
		if !busy {
			fmt.Println("不忙，吃点零食")
			break // 分支表达式随时终止程序
		} else {
			fmt.Println("忙，差不多吃点得了")
		}
	case !busy && money > 20 && money <= 200:
		fmt.Println("下个馆子")
	case !busy && money > 200 && money <= 2000:
		fmt.Println("去米其林")
	case !busy && money > 2000 && money <= 20000:
		fmt.Println("出去玩一圈")
		fallthrough // 直接进入下一个处理分支而无需判断条件
	case !busy && money > 20000 && money <= 50000:
		fmt.Println("嗨起来！") // 因fallthrough此条也将执行
	default:
		fmt.Println("容我想想")
	}
}
