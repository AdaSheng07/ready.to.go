package main

import "fmt"

// 用多维数组写一个日历表，需要考虑每个月的天数不同，以及是平年还是闰年来专门处理二月

func main() {
	//   []    []    int
	// month   day
	var year int
	fmt.Println("输入需要日历的年份：")
	_, _ = fmt.Scan(&year)
	var bigMonth []int   // jan, mar, may, jul, aug, oct, dec
	var smallMonth []int // apr, jun, sep, nov
	var feb []int

	// 初始化大小月，大月31天，小月30天
	for i := 1; i <= 30; i++ {
		smallMonth = append(smallMonth, i)
	}
	bigMonth = append(smallMonth, 31)

	// 判断平年与闰年，确定二月的天数
	var febDays int
	if year%4 != 0 {
		febDays = 28
		//fmt.Println("不是闰年")
	} else if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
		febDays = 29
		//fmt.Println("是闰年")
	} else {
		febDays = 28
		//fmt.Println("不是闰年")
	}
	for i := 1; i <= febDays; i++ {
		feb = append(feb, i)
	}

	// 月份数组
	var month [12]string = [12]string{"january", "february", "march", "april", "may", "june",
		"july", "august", "september", "october", "november", "december",
	}

	// 生成年日历
	var calendar [][]int
	for _, val := range month {
		switch val {
		case "january", "march", "may", "july", "august", "october", "december":
			calendar = append(calendar, bigMonth)
		case "april", "june", "september", "november":
			calendar = append(calendar, smallMonth)
		case "february":
			calendar = append(calendar, feb)
		}
	}

	// 按月打印年日历
	for i := 0; i <= 11; i++ {
		fmt.Println(month[i])
		for _, val := range calendar[i] {
			fmt.Printf("%d ", val)
		}
		fmt.Println()
	}

}
