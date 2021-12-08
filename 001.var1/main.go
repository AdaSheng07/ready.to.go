package main

import (
	"fmt"
)

/*

Go语言变量
- 由字母、数字、下划线组成，其中首个字符不能为数字
- 声明变量的一般形式是使用 var 关键字
- 变量声明的不同方式：
  1. var 变量名 (变量数据类型) = value，如果没有初始化value，默认值为0或""或false
  2. 一次声明多个变量：
     var 变量名1, 变量名2, ... （变量数据类型） = value1, value2, ...
  3. 不声明变量数据类型，依靠 golang 根据值类型自行做变量类型推断
  4. 利用 := 符号声明变量并赋值：
     变量名 := value，类型由 golang 自行判断
     := 左侧的变量不能是已经被 var 声明过的，否则会导致编译错误
     := 左侧可以出现未被 var 声明过的同一变量，相当于对该变量重新赋值，变量的地址不变
     不带声明格式 var 的只能在函数体中出现
  5. 因式分解关键字的写法用于声明全局变量
     var (
         变量名1 变量数据类型
         变量名2 变量数据类型
         ...
     )

*/
func main() {

	// 变量声明方式1
	var hello string = "Hello, golang!"
	fmt.Println(hello)
	var int1 int
	fmt.Println(int1) // int1没有初始化，默认值为0
	var int2 int = 33
	fmt.Println(int2)
	var float3 float64 = 1.234
	fmt.Println(float3)

	// 变量声明方式2
	var s1, s2, s3 string = "twinkle", "twinkle", "little star"
	fmt.Println(s1, s2, s3)
	var num1, num2, num3 = 6, 36, 216 // 可省略数据类型，golang自行判断
	fmt.Println(num1, num2, num3)
	var f1, f2, f3 = 1.01, 42, "item" // 不同数据类型可同时声明赋值
	fmt.Println(f1, f2, f3)

	// 变量声明方式3
	var a = 42
	fmt.Println(a)
	fmt.Printf("%T\n", a) // 打印变量a的数据类型
	var b = 1.355
	fmt.Println(b)
	fmt.Printf("%T\n", b) // 打印变量b的数据类型
	var c = "3652"
	fmt.Println(c)
	fmt.Printf("%T\n", c) // 打印变量c的数据类型
	var d = false
	fmt.Println(d)
	fmt.Printf("%T\n", d) // 打印变量d的数据类型

	// 变量声明方式4
	/* var e, f, g int // 报错：no new variables on the left side of ':=' */
	e, f, g := 1, 2, 3
	fmt.Println(e, f, g)
	fmt.Println("e addr:", &e)
	// 对变量 e 重新赋值，变量 e 的地址不变
	e, j := 4, "woo"
	fmt.Println(e, j)
	fmt.Println("e addr:", &e)

	// 变量声明方式5
	var (
		h int
		i bool
	)
	fmt.Println(h, i)

}
