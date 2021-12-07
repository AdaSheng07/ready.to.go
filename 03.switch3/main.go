package main

import (
	"fmt"
	"reflect"
)

/*
switch语句原生支持类型判断
1. 可以利用 type-switch 来判断某个 interface 变量中实际存储的变量类型
2. interface{} 是一个空的 interface 类型，是抽象的，其内容可以【接受】任何具体的类型如 int, float64, string 等
3. 用 type-switch 类型判断的同时可以赋值新的变量，且类型自动匹配，新的变量的类型是万用的
4. case 也可以是多分支条件的，变量仍是空接口类型interface{}，可对它直接赋值，但是获取它的值进行运算时需要断言，e.g.newMoney.(int), newMoney.(float64)

p.s. verb 通用输出：
	%v	值的默认格式表示
	%+v	类似%v，但输出结构体时会添加字段名
	%#v	值的Go语法表示
	%T	值的类型的Go语法表示
	%%	百分号
*/

func main() {
	var money interface{} = 100.5
	// var money interface{} = "100块"
	// var money interface{} = 43.5

	switch newMoney := money.(type) { // 类型判断的同时赋值新的变量 newMoney
	case int, int8, int16, int32, int64: // case 可以多条件
		newMoney = 10
		// newMoney = newMoney + 12.34 // Invalid operation: newMoney + 12.34 (cannot convert the constant 12.34 to the type int)
		fmt.Printf("newMoney %#v 是 %v\n", newMoney.(int)+10, reflect.TypeOf(newMoney))
	case float64, float32: // case 可以多条件
		newMoney = 10.5
		fmt.Printf("newMoney %#v 是 %v\n", newMoney.(float64)+10.5, reflect.TypeOf(newMoney))
	case string:
		newMoney = newMoney + "哈"
		fmt.Printf("newMoney %#v 是 %v\n", newMoney, reflect.TypeOf(newMoney))
	default:
		fmt.Printf("newMoney %#v 是 %v\n", newMoney, reflect.TypeOf(newMoney))
		fmt.Printf("money %#v 类型未知\n", money)
	}
	fmt.Println(money, reflect.TypeOf(money))
	// 可以对interface类型的变量重新赋值，变量类型也可以发生改变
	money = 34.6
	fmt.Println(money, reflect.TypeOf(money))
	fmt.Println("END")

}
