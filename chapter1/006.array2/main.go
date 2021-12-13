package main

import "fmt"

func main() {
	// 初始化二维数组方式一：难以长期维护
	a1 := [3]int{11, 12, 13}
	a2 := [3]int{21, 22, 23}
	a3 := [3]int{31, 32, 33}
	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)

	// 初始化二维数组方式二：数组长度有限
	var a [3][3]int
	a[0] = [3]int{11, 12, 13}
	a[1] = [3]int{21, 22, 23}
	a[2] = [3]int{31, 32, 33}
	fmt.Println(a)

	var b = [3][3]string{
		[3]string{"take", "take", "take"},
		[3]string{"go", "go", "go"},
		[3]string{"fine", "fine", "fine"},
	}
	for _, val := range b {
		fmt.Println(val)
	}
	// 初始化二维数组方式三：优化数组长度管理，支持动态添加
	c := [...][3]string{
		[3]string{"take", "take", "take"},
		[3]string{"go", "go", "go"},
		[3]string{"fine", "fine", "fine"},
	}
	fmt.Println(c)
	for d1, d1val := range c {
		for d2, d2val := range d1val {
			fmt.Println(d1, d1val, d2, "val =", d2val)
		}
	}

	// 初始化多维数组方式四：定义空数组，利用append添加一维数组
	// 1. 创建数组
	var values [][]int
	// 2. 使用append()函数向空的二维数组添加两行一维数组
	row1 := []int{1, 2, 3}
	row2 := []int{4, 5, 6}
	values = append(values, row1)
	values = append(values, row2)
	// 3. 显示两行数据
	fmt.Println("Row 1")
	fmt.Println(values[0])
	fmt.Println("Row 2")
	fmt.Println(values[1])
	// fmt.Println("Row 3")
	// fmt.Println(values[2])// 报错：panic: runtime error: index out of range [2] with length 2

	// 遍历二维数组
	fmt.Println("遍历一层")
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}
	fmt.Println("遍历两层")
	for i, val := range a {
		for j := 0; j < len(val); j++ {
			fmt.Printf("%d\t%d\t%d\t", i, j, val[j])
		}
		fmt.Println()
	}
}
