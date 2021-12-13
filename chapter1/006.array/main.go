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
}
