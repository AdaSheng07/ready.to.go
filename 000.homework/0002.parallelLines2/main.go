package main

import "fmt"

func main() {
	for {
		var line1, line2 [2][2]float64 // 声明两个多维的数组
		var ans string                 // 回答是否继续判定直线关系
		// 输入坐标，存入2*2的多维数组，并对输入数据的合法性进行确认
		line1, line2 = InputCoordinates(line1, line2)
		/*
			line1和line2中元素与四组点坐标的对应关系
			line1:             line2:
			    [0] [1]            [0] [1]
			[0]{x1, y1}        [0]{x3, y3}
			[1]{x2, y2}        [1]{x4, y4}
		*/
		// 判断两条直线的关系：平行/相交/重合
		VerdictLinesRelation(line1, line2)

		// 询问是否继续判定直线关系
		fmt.Print("是否继续录入？（继续录入请键入y，其他任意键停止录入）：")
		fmt.Scan(&ans)
		if ans != "y" {
			break // 不继续录入，跳出循环，结束
		}
	}
	// 所有流程结束
	fmt.Println("-------END-------")
}

// VerdictLinesRelation 判断两条直线的关系：平行/相交/重合
func VerdictLinesRelation(arr1, arr2 [2][2]float64) {
	// 先排除所有斜率不存在情况：
	switch {
	// 如果直线1与直线2的斜率都是+Inf，即它们都是垂直于x轴的直线，需要判定是【两线平行】还是【两线重合】
	case arr1[0][0] == arr1[1][0] && arr2[0][0] == arr2[1][0]:
		// 如果两条垂直于x轴的直线在x轴上的截距相等，那么两线重合，否则两线平行
		if arr1[0][0] == arr2[0][0] {
			fmt.Println("两条直线重合")
		} else {
			fmt.Println("两条直线平行")
		}
		return // 判定结束，返回空
	// 如果直线1与直线2中只有一条的斜率是+Inf，那么两条直线一定相交
	case (arr1[0][0] == arr1[1][0]) != (arr2[0][0] == arr2[1][0]):
		fmt.Println("两条直线相交")
		return // 判定结束，返回空
	}
	// 两条直线的斜率都存在的情形：
	// 分别计算两线斜率
	slope1 := CalcSlope(arr1)
	slope2 := CalcSlope(arr2)
	if slope1 != slope2 { // 如果两条直线的斜率不相等，两线一定相交
		fmt.Println("两条直线相交")
		return // 判定结束，返回空
	} else if IsTheSameLine(arr1[0][:], arr2[0][:], slope1) { // 如果两条直线的斜率相等，计算截距判断平行/重合
		fmt.Println("两条直线重合")
		return // 判定结束，返回空
	} else {
		fmt.Println("两条直线平行")
		return // 判定结束，返回空
	}
}

// InputCoordinates 输入坐标，存入2*2的多维数组，并对输入合法性进行判断
func InputCoordinates(arr1, arr2 [2][2]float64) ([2][2]float64, [2][2]float64) {
	// 输入直线1的两组坐标，存入一个2*2的多维数组arr1
	arr1 = [2][2]float64{}
	fmt.Println("输入直线 1 上第一个点的坐标(x1, y1)，中间用空格分开：")
	fmt.Scan(&arr1[0][0], &arr1[0][1])
	fmt.Println("输入直线 1 上第二个点的坐标(x2, y2)，中间用空格分开：")
	fmt.Scan(&arr1[1][0], &arr1[1][1])

	// 输入直线2的两组坐标，存入一个2*2的多维数组arr2
	arr2 = [2][2]float64{}
	fmt.Println("输入直线 2 上第一个点的坐标(x3, y3)，中间用空格分开：")
	fmt.Scan(&arr2[0][0], &arr2[0][1])
	fmt.Println("输入直线 2 上第二个点的坐标(x4, y4)，中间用空格分开：")
	fmt.Scan(&arr2[1][0], &arr2[1][1])

	// 输入坐标合法性判断：不能重复
	// 如果【x1=x2 且 y1=y2】或者【x3=x4 且 y3=y4】，arr1或arr2中的两组坐标是重复的
	if (arr1[0][0] == arr1[1][0] && arr1[0][1] == arr1[1][1]) || (arr2[0][0] == arr2[1][0] && arr2[0][1] == arr2[1][1]) {
		fmt.Println("已输入的坐标重复，不足以确定一条直线，重新输入直线上的点坐标！")
		arr1, arr2 = InputCoordinates(arr1, arr2)
	}
	return arr1, arr2
}

// CalcSlope 根据两点确定一条直线的定理计算直线斜率
func CalcSlope(arr1 [2][2]float64) float64 {
	// 计算斜率
	slope := (arr1[1][1] - arr1[0][1]) / (arr1[1][0] - arr1[0][0])
	return slope
}

// IsTheSameLine 计算两条直线的截距，判断两条斜率相等的直线是【重合】还是【平行】
func IsTheSameLine(arr1 []float64, arr2 []float64, slope float64) bool {
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
