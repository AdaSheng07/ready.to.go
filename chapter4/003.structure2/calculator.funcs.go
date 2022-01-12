package main

import "fmt"

func (c *Calculator) Add() int {
	// structure variable c is a receiver, by receiver we get access to member variables and functions
	tempResult := c.left + c.right
	c.result = tempResult
	fmt.Printf("&c @c.Add() is: %p\n", &c)
	fmt.Println("调用Add()函数，c的result =", c.result)
	return tempResult
}

// Sub : function Sub() is based on object Calculator, not the pointer of Calculator
func (c Calculator) Sub() int {
	return c.left - c.right

}
func (c Calculator) Multiply() int {
	return c.left * c.right

}
func (c Calculator) Divide() int {
	return c.left / c.right

}
func (c Calculator) Remainder() int {
	return c.left % c.right

}

//func (c *Calculator) SetResult(result int) {
//	c.result = result
//}
