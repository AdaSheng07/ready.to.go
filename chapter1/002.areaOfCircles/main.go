package main

import (
	"fmt"
	"math"
	"strconv"
)

// 计算两个圆的面积之和并输出，精确到小数点后 3 位。

func main() {
	var radius1, radius2 string
	var area float64
	for {
		fmt.Printf("分别输入两个圆的半径，输入完毕后回车: ")
		fmt.Scanln(&radius1, &radius2)
		r1, err1 := strconv.ParseFloat(radius1, 64)
		r2, err2 := strconv.ParseFloat(radius2, 64)
		if r1 >= 0 && r2 >= 0 && err1 == nil && err2 == nil {
			area = math.Pi * (math.Pow(r1, 2) + math.Pow(r2, 2))
			break
		}
		fmt.Printf("Warning: 无效输入！")
	}
	fmt.Printf("两圆面积之和为：%.3f\n", area)
}
