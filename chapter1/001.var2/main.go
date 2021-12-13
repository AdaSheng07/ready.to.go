package main

import (
	"fmt"
	"reflect"
)

/*

Go语言变量类型
1. Golang是强类型语言，在赋值过程中，类型必须保持一致，前后使用如果改变数据类型会不兼容报错
2. 变量必须先定义，后使用，且必须被用到
3. Golang会为每个变量设定默认值
4. 在同一个作用域内，变量不能重名
5. Golang会根据值类型做变量类型判断

*/

func main() {
	// 1. Golang是强类型语言
	var hello string = "hello"
	fmt.Println(hello)
	hello = "hello 2nd time"
	fmt.Println(hello)
	/* 如果赋值：hello = 33，报错：'33' (type untyped int) cannot be represented by the type string
	   hello的类型已经声明为string，不能赋值int */
	hello = "33"
	fmt.Println(hello)

	// 2. 声明定义变量后使用
	var int2 = 42
	fmt.Println(int2)
	float3 := 1.234
	fmt.Println(float3)
	var int4, int5 = 44, 55
	fmt.Println(int4, int5)
	float4, float5 := 4.0, 5.0
	fmt.Println(float4, float5)
	fmt.Printf("float4: %T, float5: %T\n", float4, float5) // 虽然打印显示为4和5，但float4与float5都是float64型变量
	var (
		int6, int7 = 6, 7
	)
	fmt.Println(int6, int7)

	// 使用变量
	var side1 int = 3
	var side2 int = 4
	fmt.Println(side1 * side2)
	var area = side1 * side2
	fmt.Println(area)

	// 3. Golang为每个变量设定默认值，观察下列输出
	var length, width int
	fmt.Println(length)
	fmt.Println(width)
	var string1 string
	fmt.Println(string1)
	var mapA map[string]string
	fmt.Println(mapA)

	// 4. 在同一个作用域内，变量不能重名
	fmt.Println("length addr:", &length)
	fmt.Println("width addr:", &width)
	{
		length := 20
		fmt.Println(length)
		width := 40
		fmt.Println(width)
		fmt.Println("length addr:", &length)
		fmt.Println("width addr:", &width)
	}

	// 5. Golang根据值类型做变量类型判断
	a1 := "hello"
	fmt.Println(reflect.TypeOf(a1))
	a2 := 3
	fmt.Println(reflect.TypeOf(a2))
	a3 := 3.0
	fmt.Println(reflect.TypeOf(a3))
	a4 := &a3 // 地址
	fmt.Println(reflect.TypeOf(a4))
}
