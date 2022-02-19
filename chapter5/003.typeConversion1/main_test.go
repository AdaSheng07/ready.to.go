package main

import (
	"fmt"
	"testing"
)

func TestAssertion(t *testing.T) {
	r := Refrigerator{Size: 5}
	// var b Box = r   // 从特殊类型Refrigerator到一般类型接口的转换，可以直接赋值
	var c Close = r // 从特殊类型Refrigerator到一般类型接口的转换，可以直接赋值

	// 从特殊类型Refrigerator到一般类型接口的转换？
	r1 := c.(Refrigerator) // 强制转换类型：Refrigerator --> Close接口，不考虑转换的成功或失败，带来很大的风险
	fmt.Println(r1.Size)

	r2 := TestBox{}
	var c2 Close = r2 // 从特殊类型TestBox到一般类型接口的转换
	//fmt.Println(c2.(Refrigerator)) // panic: interface conversion: main.Close is main.TestBox, not main.Refrigerator

	// 运用类型断言判断类型转换是否成功
	// 1. switch-case
	switch cType := c2.(type) {
	case Refrigerator:
		fmt.Println("expected refrigerator, get refrigerator")
		fmt.Println(cType.Size)
	case TestBox:
		fmt.Println("this is a Box")
	default:
		fmt.Println("no type matches, type conversion fails")
	}

	// 2. variable, ok := interface_name.(<type>)
	fridge, ok := checkIfRefrigerator(c2)
	if ok {
		fmt.Println("this is a refrigerator", fridge)
	} else {
		fmt.Println("this is not a refrigerator")
	}
}

// 2. variable, ok := interface_name.(<type>)
func checkIfRefrigerator(c Close) (Refrigerator, bool) {
	r, ok := c.(Refrigerator)
	return r, ok
}

type TestBox struct {
}

func (tb TestBox) Close() error {
	// todo
	return nil
}
