package main

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
)

/*

Go语言变量类型

虽然Go语言自带变量类型判断，但对于在代码中高频使用的变量尽量定义类型，因为自动类型判断有代价，可能出错。
当 Golang 自行判断变量类型出错时，对于数字类型的变量可以强制转换类型，e.g. int <-> float64，但数据本身会有损失，丢失了精度。

*/

func main() {

	// 强制转换字符串类型
	name := "123"
	// 如果name的字符串内容为"take"，输出结果值为0，并err报错：strconv.Atoi: parsing "take": invalid syntax
	val, err := strconv.Atoi(name) // 将string转换为int
	fmt.Printf("val type: %T\n", val)
	fmt.Println("val type:", reflect.TypeOf(val))
	fmt.Println(val, err)

	// 强制转换数字类型
	var a, b = 2, 2.5
	// var c = a * b 此语句报错 Invalid operation: a * b (mismatched types int and float64)
	var c = a * int(b) // 强制转换类型，观察结果
	fmt.Println(c)     // 发现结果不是5，而是4，因为 int(b) 截断了小数位，使b = 2，造成数值的精度损失

	var f1 float64 = 1.234
	var i1 int = int(f1)
	fmt.Println("f1:", f1, "i1:", i1)

	var a1 uint = math.MaxUint64 // unit 无符号整数中的最大值：const MaxUint64 int = 1<<64 - 1，1左移64次后减一
	var a2 int = int(a1)
	fmt.Println(a1) // 输出：18446744073709551615
	fmt.Println(a2) // 输出：-1

	/*
		这与计算机使用存储空间时数字的存储方式有关：
		1. 每种数据类型都有大小限制：math/const.go
		2. 每种数据类型的存储方式不同
		3. 不是每种数据类型之间都可以转换
		4. 可转换数据类型之间的转换是有代价的
	*/

}
