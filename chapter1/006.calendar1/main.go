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

	// 大小月，大月31天，小月30天
	for i := 1; i <= 30; i++ {
		bigMonth = append(bigMonth, i)
		smallMonth = append(smallMonth, i)
	}
	bigMonth = append(bigMonth, 31)

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

	var calendar [][]int
	calendar = append(calendar, bigMonth)   // january
	calendar = append(calendar, feb)        // february
	calendar = append(calendar, bigMonth)   // march
	calendar = append(calendar, smallMonth) // april
	calendar = append(calendar, bigMonth)   // may
	calendar = append(calendar, smallMonth) // june
	calendar = append(calendar, bigMonth)   // july
	calendar = append(calendar, bigMonth)   // august
	calendar = append(calendar, smallMonth) // september
	calendar = append(calendar, bigMonth)   // october
	calendar = append(calendar, smallMonth) // november
	calendar = append(calendar, bigMonth)   // december

	var month [12]string = [12]string{"january", "february", "march", "april", "may", "june",
		"july", "august", "september", "october", "november", "december",
	}

	for i := 0; i <= 11; i++ {
		fmt.Println(month[i])
		for _, val := range calendar[i] {
			fmt.Printf("%d ", val)
		}
		fmt.Println()
	}

}
