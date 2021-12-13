package main

import "fmt"

func main() {
	// 存储单个变量
	var age int = 35
	fmt.Println(age)
	// 存储一组相同类型的变量，数组声明的不同方式
	var ages1 [5]int = [5]int{35, 23, 45, 46, 25}
	fmt.Println(ages1)
	var ages2 = [5]int{32, 45, 53, 23, 34}
	fmt.Println(ages2)
	ages3 := [5]int{34, 75, 25, 57, 24}
	fmt.Println(ages3)
	// ages3 = [6]int{34, 75, 25, 57, 24, 99} // 报错：ages3的类型是[5]int，不能重复赋值为[6]int

	a := [3]int{1, 2, 3}
	var b [3]int = [3]int{}
	// var c [3]int = [4]int{} // 编译不通过，类型不匹配错误
	d := [5]int{}
	e := [...]int{1, 2, 3, 4}
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(len(e)) // 数组有固定长度

	// 数组赋值
	var array1 [3]int
	array1 = [3]int{1, 1, 1}
	fmt.Println(array1)
	array1 = [3]int{2, 2, 2}
	fmt.Println(array1)

	// 对数组中的元素赋值
	var array2 [3]int
	array2[0] = 1
	array2[1] = 2
	array2[2] = 3
	fmt.Println(array2)
	// array2[-1] = 1 // 报错：越界赋值
	// array2[5] = 3 // 报错：越界赋值

	for i := 0; i < len(array2); i++ {
		fmt.Println(array2[i])
	}

	// for range 遍历访问数组
	for i, val := range array2 {
		fmt.Printf("%d, array2[%d]: %d\n", array2[i], i, val)
		fmt.Printf("%d\t%d\n", i, val)
	}

}
