package main

import (
	"fmt"
)

// 用多维数组写一个日历表，需要考虑每个月的天数不同，以及是平年还是闰年来专门处理二月
//【提升篇】日历按照一周的宽度显示（第一列为周一），且每个日期匹配到对应的列

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
	var month [12]string = [12]string{"January", "February", "March", "April", "May", "June",
		"July", "August", "September", "October", "November", "December"}
	// 星期数组
	var weekday [7]string = [7]string{"M", "T", "W", "T", "F", "S", "S"}

	// 生成年日历
	var calendar [][]int
	for _, val := range month {
		switch val {
		case "January", "March", "May", "July", "August", "October", "December":
			calendar = append(calendar, bigMonth)
		case "April", "June", "September", "November":
			calendar = append(calendar, smallMonth)
		case "February":
			calendar = append(calendar, feb)
		}
	}

	// days是这一天在这一年中是第几天
	var days int = 0

	// 按月打印年日历
	for i := 0; i <= 11; i++ {
		fmt.Printf("\n%s\n", month[i])
		printWeekday(weekday)
		// 计算每个月的第一天是这一年的第几天，即days的值
		if i == 0 {
			days = 1 // 一月的第一天也是是这一年的第一天
		} else {
			days += len(calendar[i-1]) // 第j个月的第一天是一月到(j-1)月的所有天数+1，e.g.12月的第一天是1-11月所有天数+1
		}
		// 计算并记录每个月的第一天是星期几，确定起始位置
		startWeekday := startWeekday(year, days)
		// 打印每个月的月历
		printMonthly(calendar, i, startWeekday)
	}
}

// 打印每个月的月历
func printMonthly(calendar [][]int, i int, startWeekday int) {
	for j, val := range calendar[i] {
		fmt.Printf("%d\t", val)
		startWeekday++
		if startWeekday%7 == 1 && j < (len(calendar[i])-1) { // 一周已经打印到周日，换行
			fmt.Println()
		}
	}
	fmt.Println()
}

// 打印星期栏
func printWeekday(weekday [7]string) {
	for i := 0; i < 7; i++ {
		fmt.Printf("%s\t", weekday[i])
	}
	fmt.Println()
}

// 计算并记录每个月的第一天是星期几，确定起始位置
func startWeekday(year int, days int) int {
	// 计算每年每个月的第一天是星期几
	// W = ([Y-1] + [(Y-1)/4] - [(Y-1)/100] + [(Y-1)/400] + days) % 7
	var startWeekday int = ((year - 1) + (year-1)/4 - (year-1)/100 + (year-1)/400 + days) % 7
	// 根据是星期几确定输出时对准的位置
	switch startWeekday {
	case 0: // 周日
		fmt.Printf("\t\t\t\t\t\t")
	case 1: // 周一
		break
	case 2: // 周二
		fmt.Printf("\t")
	case 3: // 周三
		fmt.Printf("\t\t")
	case 4: // 周四
		fmt.Printf("\t\t\t")
	case 5: // 周五
		fmt.Printf("\t\t\t\t")
	case 6: // 周六
		fmt.Printf("\t\t\t\t\t")
	}
	return startWeekday
}

/*
	$ go run ./main.go
	输入需要日历的年份：
	2021

	January
	M       T       W       T       F       S       S
									1       2       3
	4       5       6       7       8       9       10
	11      12      13      14      15      16      17
	18      19      20      21      22      23      24
	25      26      27      28      29      30      31

	February
	M       T       W       T       F       S       S
	1       2       3       4       5       6       7
	8       9       10      11      12      13      14
	15      16      17      18      19      20      21
	22      23      24      25      26      27      28

	March
	M       T       W       T       F       S       S
	1       2       3       4       5       6       7
	8       9       10      11      12      13      14
	15      16      17      18      19      20      21
	22      23      24      25      26      27      28
	29      30      31

	April
	M       T       W       T       F       S       S
							1       2       3       4
	5       6       7       8       9       10      11
	12      13      14      15      16      17      18
	19      20      21      22      23      24      25
	26      27      28      29      30

	May
	M       T       W       T       F       S       S
											1       2
	3       4       5       6       7       8       9
	10      11      12      13      14      15      16
	17      18      19      20      21      22      23
	24      25      26      27      28      29      30
	31

	June
	M       T       W       T       F       S       S
			1       2       3       4       5       6
	7       8       9       10      11      12      13
	14      15      16      17      18      19      20
	21      22      23      24      25      26      27
	28      29      30

	July
	M       T       W       T       F       S       S
							1       2       3       4
	5       6       7       8       9       10      11
	12      13      14      15      16      17      18
	19      20      21      22      23      24      25
	26      27      28      29      30      31

	August
	M       T       W       T       F       S       S
													1
	2       3       4       5       6       7       8
	9       10      11      12      13      14      15
	16      17      18      19      20      21      22
	23      24      25      26      27      28      29
	30      31

	September
	M       T       W       T       F       S       S
					1       2       3       4       5
	6       7       8       9       10      11      12
	13      14      15      16      17      18      19
	20      21      22      23      24      25      26
	27      28      29      30

	October
	M       T       W       T       F       S       S
									1       2       3
	4       5       6       7       8       9       10
	11      12      13      14      15      16      17
	18      19      20      21      22      23      24
	25      26      27      28      29      30      31

	November
	M       T       W       T       F       S       S
	1       2       3       4       5       6       7
	8       9       10      11      12      13      14
	15      16      17      18      19      20      21
	22      23      24      25      26      27      28
	29      30

	December
	M       T       W       T       F       S       S
					1       2       3       4       5
	6       7       8       9       10      11      12
	13      14      15      16      17      18      19
	20      21      22      23      24      25      26
	27      28      29      30      31

*/
