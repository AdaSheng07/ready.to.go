package main

import "fmt"

func main() {
	/* 整体思路：
	输入直线1的两组坐标
	输入直线2的两组坐标
	输入合法性判断：有无重合点
	垂直于x轴的斜率不存在情况：都不存在/一个不存在；平行/重合/相交
	斜率存在，斜率相同，截距不同：平行
	斜率存在，斜率相同，截距也相同：重合
	斜率存在，斜率不相同：相交
	*/

	var slope1, slope2 float64 // 声明两条直线的斜率

restart:
	// 输入直线1的两组坐标，存入一个2*2的多维数组line1
	line1 := [2][2]float64{}
	fmt.Println("输入直线 1 上第一个点的坐标(x1, y1)，中间用空格分开：")
	fmt.Scan(&line1[0][0], &line1[0][1])
	fmt.Println("输入直线 1 上第二个点的坐标(x2, y2)，中间用空格分开：")
	fmt.Scan(&line1[1][0], &line1[1][1])

	// 输入直线2的两组坐标，存入一个2*2的多维数组line2
	line2 := [2][2]float64{}
	fmt.Println("输入直线 2 上第一个点的坐标(x3, y3)，中间用空格分开：")
	fmt.Scan(&line2[0][0], &line2[0][1])
	fmt.Println("输入直线 2 上第二个点的坐标(x4, y4)，中间用空格分开：")
	fmt.Scan(&line2[1][0], &line2[1][1])

	/*
		line1和line2中元素与四组点坐标的对应关系
		line1:             line2:
		    [0] [1]            [0] [1]
		[0]{x1, y1}        [0]{x3, y3}
		[1]{x2, y2}        [1]{x4, y4}
	*/

	// 特殊情形判断：1. 输入的点坐标重合；2. 直线斜率不存在
	switch {
	// 1. 输入不合法，两点重合，不足以确定直线，重新输入
	case (line1[0][0] == line1[1][0] && line1[0][1] == line1[1][1]) ||
		(line2[0][0] == line2[1][0] && line2[0][1] == line2[1][1]):
		// 如果【x1=x2 且 y1=y2】或者【x3=x4 且 y3=y4】，line1或line2中的两组坐标是重复的
		fmt.Println("已输入的坐标重复，不足以确定一条直线，重新输入直线上的点坐标！")
		goto restart
	// 2. 如果两条直线中有斜率+Inf的情况：
	// 2.1 如果直线1与直线2的斜率都是+Inf，即它们都是垂直于x轴的直线，需要判定是【两线平行】还是【两线重合】
	case line1[0][0] == line1[1][0] && line2[0][0] == line2[1][0]: // x1 = x2 且 x3 = x4
		fmt.Println("两条直线的斜率都不存在，判断两线是平行or重合")
		// 如果两条垂直于x轴的直线在x轴上的截距相等，那么两线重合，否则两线平行
		if line1[0][0] == line2[0][0] {
			fmt.Println("两条直线重合")
		} else {
			fmt.Println("两条直线平行")
		}
		goto end // 判定结束
	// 2.2 如果直线1与直线2中只有一条的斜率是+Inf，那么两条直线一定相交
	case (line1[0][0] == line1[1][0]) != (line2[0][0] == line2[1][0]): // 【x1 = x2】和【x3 = x4】只满足其中一个
		fmt.Println("两条直线相交")
		goto end // 判定结束
	}

	// 两条直线的斜率都存在的情形：
	// 分别计算两线斜率
	slope1 = calcSlope(line1[0][0], line1[0][1], line1[1][0], line1[1][1])
	slope2 = calcSlope(line2[0][0], line2[0][1], line2[1][0], line2[1][1])
	if slope1 != slope2 { // 如果两条直线的斜率不相等，两线一定相交
		fmt.Println("两条直线相交")
	} else if isTheSameLine(line1[0][:], line2[0][:], slope1) { // 如果两条直线的斜率相等，计算截距判断平行/重合
		fmt.Println("两条直线重合")
	} else {
		fmt.Println("两条直线平行")
	}

end:
	fmt.Println("-------END-------") // 所有流程结束
}

// 函数 calcSlope 根据两点确定一条直线的定理计算直线斜率
func calcSlope(x1, y1, x2, y2 float64) float64 {
	// 计算斜率
	slope := (y2 - y1) / (x2 - x1)
	return slope
}

// 函数 isTheSameLine 计算两条直线的截距，判断两条斜率相等的直线是【重合】还是【平行】
func isTheSameLine(arr1 []float64, arr2 []float64, slope float64) bool {
	// 计算截距
	intercept1 := arr1[1] - slope*arr1[0]
	intercept2 := arr2[1] - slope*arr2[0]
	// 两条斜率相等的直线：截距相等则【重合】，截距不等则【平行】
	if intercept1 == intercept2 { // 如果两条直线的斜率相等，截距也相等，两条直线重合
		return true
	} else { // 否则两条直线平行
		return false
	}
}
