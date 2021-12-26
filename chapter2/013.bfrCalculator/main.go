package main

import (
	"fmt"
	"learn.go/chapter2/013.bfrCalculator/calc"
	"regexp"
	"strconv"
)

func main() {

	// 全局变量声明
	var name string = ""   // 参与者姓名
	var age int = 0        // 初始化年龄age
	var sex string = ""    // 初始化性别sex
	var height float64 = 0 // 初始化身高height
	var weight float64 = 0 // 初始化体重weight

	var num int = 1                    // 参与者数量统计，初始化为1
	var sum float64                    // 计算所有参与者bmi值的和
	var ans string                     // 回答是否继续录入数据
	var sexWeight int                  // 性别识别
	var bmi float64                    // BMI(Body Mass Index) 身体质量指数
	var bfr float64                    // BFR(Body Fat Rate) 体脂率
	var s1 float64                     // 体脂率【标准】范围的下限
	var delta1, delta2, delta3 float64 // 体脂率【标准】【标准-偏重】【标准-肥胖】范围的长度
	var anticipates []string           // 切片anticipates存放所有参与者的姓名
	var bmiSet []float64               // 切片bmiSet存放所有参与者的姓名
	var bfrSet []float64               // 切片bfrSet存放所有参与者的姓名
	var suggestions []string           // 切片suggestions存放所有参与者的姓名

	for {
		// 开始录入参与者信息，并保证所有输入数据的合法性
		// 录入姓名并验证数据合法性
		name = InputName(num)
		// 录入年龄并验证数据合法性
		age = InputAge(name, num)
		// 录入性别并验证数据合法性
		sex = InputSex(name)
		// 录入身高并验证数据合法性
		height = InputHeight(name)
		// 录入体重并验证数据合法性
		weight = InputWeight(name)

		/* 男性与女性不同性别区间的体脂率范围不同，但整体来看可以依据节点划分，各性别的各判断标准对应的体脂率范围长度分别相同：
		      sex   age    |     偏瘦       |       标准        |      偏重       |       肥胖      |   严重肥胖
		【1】  男   18-39  0%              10%      +6%        16%     +5%      21%      +5%      26%
		【2】  男   40-59  0%              11%      +6%        17%     +5%      22%      +5%      27%
		【3】  男    60+   0%              13%      +6%        19%     +5%      24%      +5%      29%
		       男                          s1                s1+6%            s1+6%+5%        s1+6%+5%+5%
		【4】  女   18-39  0%              20%      +7%        27%     +7%      34%      +5%      39%
		【5】  女   40-59  0%              21%      +7%        28%     +7%      35%      +5%      40%
		【6】  女    60+   0%              22%      +7%        29%     +7%      36%      +5%      41%
		       女                          s1                s1+7%            s1+7%+7%        s1+7%+7%+5%
		*/

		// 依据性别与年龄，确定当前参与者适用的体脂状态判定标准
		sexWeight, delta1, delta2, delta3, s1 = PersonalStandard(sex, sexWeight, delta1, delta2, delta3, age, s1)

		// 计算BMI与体脂率BFR
		bmi, bfr = Calculate(bmi, weight, height, bfr, age, sexWeight)
		// 记录当前参与者的姓名、BMI和体脂率BFR，并分别推入切片anticipates, bmiSet和bfrSet中
		anticipates = append(anticipates, name)
		bmiSet = append(bmiSet, bmi)
		bfrSet = append(bfrSet, bfr)

		// 判断当前参与者体脂率所处体脂状态并给出建议，将建议推入切片suggestions中
		suggestions = GiveOutSuggestions(bfr, s1, suggestions, delta1, delta2, delta3)

		// 询问是否继续录入信息
		fmt.Print("是否继续录入？（继续录入请键入y，其他任意键停止录入）：")
		fmt.Scanln(&ans)
		if ans != "y" {
			break // 不继续录入，跳出循环，输出结果
		}
		num++ // 如果继续录入，参与者人数num加一
	}

	// 打印每个参与者的姓名、BMI、体脂率以及对应建议
	PrintOutResults(num, anticipates, bmiSet, bfrSet, suggestions, sum)

}

// PrintOutResults 打印每个参与者的姓名、BMI、体脂率以及对应建议
func PrintOutResults(num int, anticipates []string, bmiSet []float64, bfrSet []float64, suggestions []string, sum float64) {
	for i := 0; i < num; i++ {
		fmt.Printf("------- %d号参与者 -------\n", i+1)
		fmt.Printf("姓名: %s\nBMI: %.2f\n体脂率: %.2f%%\n%s\n", anticipates[i], bmiSet[i], bfrSet[i]*100, suggestions[i])
		sum += bfrSet[i]
	}
	fmt.Printf("\n一共有 %d 名参与者，他们的平均体脂率是%.2f%%\n", num, sum/float64(num)*100)
	fmt.Println("\n-------THE END-------")
}

// GiveOutSuggestions 判断当前参与者体脂率所处体脂状态并给出建议，将建议推入切片suggestions中
func GiveOutSuggestions(bfr float64, s1 float64, suggestions []string, delta1 float64, delta2 float64, delta3 float64) []string {
	if bfr <= s1 {
		// 【偏瘦】区间为：[0, s1]
		suggestions = append(suggestions, "目前的体脂率水平处于【偏瘦】，建议：多吃多运动，强身健体")
	} else if bfr > s1 && bfr <= s1+delta1 {
		// 【标准】区间为：(s1, s1+delta1]
		suggestions = append(suggestions, "目前的体脂率水平处于【标准】，建议：继续保持")
	} else if bfr > s1+delta1 && bfr <= s1+delta2 {
		// 【偏胖】区间为：(s1+delta1, s1+delta2]
		suggestions = append(suggestions, "目前的体脂率水平处于【偏胖】，建议：少坐多动，注意健康饮食")
	} else if bfr > s1+delta2 && bfr <= s1+delta3 {
		// 【肥胖】区间为：(s1+delta2, s1+delta3]
		suggestions = append(suggestions, "目前的体脂率水平处于【肥胖】，建议：多多运动，控制热量摄入")
	} else {
		// 【严重肥胖】区间为：(s1+delta3, +Inf）
		suggestions = append(suggestions, "目前的体脂率水平处于【严重肥胖】，建议：少吃饭多喝水，规律作息，必须运动")
	}
	return suggestions
}

// Calculate 计算BMI与体脂率BFR
func Calculate(bmi float64, weight float64, height float64, bfr float64, age int, sexWeight int) (float64, float64) {
	bmi = calc.CalcBMI(height, weight)
	bfr = calc.CalcBFR(bmi, age, sexWeight)
	return bmi, bfr
}

// PersonalStandard 依据性别与年龄，确定当前参与者适用的体脂状态判定标准
func PersonalStandard(sex string, sexWeight int, delta1 float64, delta2 float64, delta3 float64, age int, s1 float64) (int, float64, float64, float64, float64) {
	switch {
	case sex == "男":
		sexWeight = 1          // 男性的sexWeight为1
		delta1 = 0.06          // 男性的【标准】范围区间长度为6%
		delta2 = 0.05 + delta1 // 男性的【偏重】范围区间长度为5%
		delta3 = 0.05 + delta2 // 男性的【肥胖】范围区间长度为5%
		// 根据不同年龄段，确定【标准】范围的下限作为起始点
		if age >= 60 {
			s1 = 0.13 // 年龄60+的男性【标准】范围从 13% 开始，对应上表【3】
		} else if age >= 40 {
			s1 = 0.11 // 年龄40-59的男性【标准】范围从 11% 开始，对应上表【2】
		} else {
			s1 = 0.10 // 年龄18-39的男性【标准】范围从 10% 开始，对应上表【1】
		}
	case sex == "女":
		sexWeight = 0          // 女性的sexWeight为0
		delta1 = 0.07          // 女性的【标准】范围区间长度为7%
		delta2 = 0.07 + delta1 // 女性的【偏重】范围区间长度为7%
		delta3 = 0.05 + delta2 // 女性的【肥胖】范围区间长度为5%
		// 根据不同年龄段，确定【标准】范围的下限作为起始点
		if age >= 60 {
			s1 = 0.22 // 年龄60+的女性【标准】范围从 22% 开始，对应上表【6】
		} else if age >= 40 {
			s1 = 0.21 // 年龄40-59的女性【标准】范围从 21% 开始，对应上表【5】
		} else {
			s1 = 0.20 // 年龄18-39的男性【标准】范围从 20% 开始，对应上表【4】
		}
	}
	return sexWeight, delta1, delta2, delta3, s1
}

// InputWeight 录入体重并做数据合法性验证
func InputWeight(name string) float64 {
	var weight float64 = 0 // 初始化体重weight
	fmt.Printf("请输入%s的体重(kg)：", name)
	fmt.Scan(&weight)
	// 输入合法性验证：必须是正浮点数，且值在合理区间内
	result5, _ := regexp.MatchString("^[1-9]\\d*\\.\\d*|0\\.\\d*[1-9]\\d*$", strconv.FormatFloat(weight, 'f', 2, 64))
	if !result5 {
		fmt.Print("体重为正小数，小数点后最多保留2位，请重新输入！\n")
		weight = InputWeight(name) // 重新录入体重
	} else if weight > 635 || weight < 10 {
		fmt.Print("请输入真实体重！\n")
		weight = InputWeight(name) // 重新录入体重
	}
	return weight
}

// InputHeight 录入身高并做数据合法性验证
func InputHeight(name string) float64 {
	var height float64 = 0 // 初始化身高height
	fmt.Printf("请输入%s的身高(m)：", name)
	fmt.Scan(&height)
	// 输入数据合法性验证：必须是正浮点数，且值在合理区间
	result4, _ := regexp.MatchString("^[1-9]\\d*\\.\\d*|0\\.\\d*[1-9]\\d*$", strconv.FormatFloat(height, 'f', 2, 64))
	if !result4 {
		fmt.Print("身高为正小数，小数点后最多保留2位，请重新输入！\n")
		height = InputHeight(name) // 重新录入身高
	} else if height < 0.7 || height > 2.455 {
		fmt.Print("身高单位为米，请重新输入正确身高！\n")
		height = InputHeight(name) // 重新录入身高
	}
	return height
}

// InputSex 录入性别并做数据合法性验证
func InputSex(name string) string {
	var sex string = "" // 初始化性别sex
	fmt.Printf("请输入%s的性别：", name)
	fmt.Scan(&sex)
	// 输入合法性验证：必须为"男"或"女"
	if sex != "男" && sex != "女" {
		fmt.Print("性别必须为 男 或者 女，请重新录入性别！\n")
		sex = InputSex(name) // 重新录入性别
	}
	return sex
}

// InputAge 录入年龄并做数据合法性验证
func InputAge(name string, num int) int {
	var age int
	fmt.Printf("请输入%s的年龄：", name)
	fmt.Scan(&age)
	if result2, _ := regexp.MatchString("[1-9][0-9]+", strconv.Itoa(age)); !result2 { // 保证输入年龄都是正整数
		fmt.Print("年龄必须是正整数，请重新输入！\n")
		age = InputAge(name, num) // 重新录入年龄
	} else if age < 18 { // 剔除年龄不到18岁的参与者
		fmt.Print("未成年人体脂率不具参考意义，请录入成年人数据！\n")
		name = InputName(num) // 重新录入参与者，覆盖
		age = InputAge(name, num)
	} else if age > 122 { // 世界吉尼斯纪录最长寿者为122岁
		fmt.Print("请输入真实年龄！\n")
		age = InputAge(name, num) // 重新录入年龄
	}
	return age
}

// InputName 录入姓名并做数据合法性验证
func InputName(num int) string {
	var name string // 初始化姓名name
	fmt.Printf("请输入第 %d 位参与者的姓名：", num)
	fmt.Scan(&name)
	if result1, _ := regexp.MatchString("[\u4e00-\u9fa5a-zA-Z]+", name); !result1 {
		// 输入合法性验证：姓名必须由字母和汉字组成
		fmt.Print("姓名必须由字母或汉字组成，请重新输入！\n")
		name = InputName(num)
	}
	return name
}
