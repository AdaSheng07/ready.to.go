package main

import "fmt"

func main() {

	// 常量相对于变量定义，一旦声明定义了，整个生命周期都不允许修改
	const pi = 3.14159265358979
	fmt.Println(pi)
	// pi = 3.14 报错：cannot assign to pi

	// 定义的常量可以不使用
	const goldenRatio = 6.18
}
