package gobmi

import "fmt"

func GiveOutSuggestions(bfr float64, gender string, age int) (suggestion string) {
	suggestion = ""
	delta1, delta2, delta3, s1 := getParameters(gender, age)
	if bfr <= s1 {
		// 【偏瘦】区间为：[0, s1]
		suggestion = "目前的体脂率水平处于【偏瘦】，建议：多吃多运动，强身健体"
		fmt.Println("目前的体脂率水平处于【偏瘦】，建议：多吃多运动，强身健体")
		//suggestions = append(suggestions, "目前的体脂率水平处于【偏瘦】，建议：多吃多运动，强身健体")
	} else if bfr > s1 && bfr <= s1+delta1 {
		// 【标准】区间为：(s1, s1+delta1]
		suggestion = "目前的体脂率水平处于【标准】，建议：继续保持"
		fmt.Println("目前的体脂率水平处于【标准】，建议：继续保持")
		//suggestions = append(suggestions, "目前的体脂率水平处于【标准】，建议：继续保持")
	} else if bfr > s1+delta1 && bfr <= s1+delta2 {
		// 【偏胖】区间为：(s1+delta1, s1+delta2]
		suggestion = "目前的体脂率水平处于【偏胖】，建议：少坐多动，注意健康饮食"
		fmt.Println("目前的体脂率水平处于【偏胖】，建议：少坐多动，注意健康饮食")
		//suggestions = append(suggestions, "目前的体脂率水平处于【偏胖】，建议：少坐多动，注意健康饮食")
	} else if bfr > s1+delta2 && bfr <= s1+delta3 {
		// 【肥胖】区间为：(s1+delta2, s1+delta3]
		suggestion = "目前的体脂率水平处于【肥胖】，建议：多多运动，控制热量摄入"
		fmt.Println("目前的体脂率水平处于【肥胖】，建议：多多运动，控制热量摄入")
		//suggestions = append(suggestions, "目前的体脂率水平处于【肥胖】，建议：多多运动，控制热量摄入")
	} else {
		// 【严重肥胖】区间为：(s1+delta3, +Inf）
		suggestion = "目前的体脂率水平处于【严重肥胖】，建议：少吃饭多喝水，规律作息，必须运动"
		fmt.Println("目前的体脂率水平处于【严重肥胖】，建议：少吃饭多喝水，规律作息，必须运动")
		//suggestions = append(suggestions, "目前的体脂率水平处于【严重肥胖】，建议：少吃饭多喝水，规律作息，必须运动")
	}
	return
}
