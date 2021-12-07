package main

import (
	"fmt"
)

func main() {

	// 常规运算：加、减、乘、除、取余
	var a, b int = 3, 10
	fmt.Println("a + b =", a+b)
	fmt.Println("a - b =", a-b)
	fmt.Println("a * b =", a*b)
	fmt.Println("a / b =", a/b)
	fmt.Println("a % b =", a%b)

	var c, d int8 = 127, 1      // int8的取值范围是-128～127
	fmt.Println("c + d =", c+d) // 溢出
	fmt.Println("c - d =", c-d)
	fmt.Println("c * d =", c*d)
	fmt.Println("c / d =", c/d)
	fmt.Println("c % d =", c%d)

	var e, f uint8 = 255, 5     // uint8是无符号数，取值范围是0～255
	fmt.Println("e + f =", e+f) // 溢出
	fmt.Println("e - f =", e-f)
	fmt.Println("e * f =", e*f)
	fmt.Println("e / f =", e/f)
	fmt.Println("e % f =", e%f)

	// 比较运算符：与&&，或||，非！，相等==，不等!=，大于>，小于<，大于等于>=，小于等于<=
	// 通常用于条件语句if-else
	fmt.Println((a+b < 20) && (a*b < 100)) // true
	fmt.Println((c%d == 0) || (e%f != 0))  // true
	fmt.Println(!(b%a >= 2))               // true

}
