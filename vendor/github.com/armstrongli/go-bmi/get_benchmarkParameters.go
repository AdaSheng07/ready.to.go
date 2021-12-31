package gobmi

func getParameters(gender string, age int) (delta1 float64, delta2 float64, delta3 float64, s1 float64) {
	switch getGenderWeight(gender) {
	case 1.0:
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
	case 0:
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
	return
}
