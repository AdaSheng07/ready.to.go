# 结构体 structure

### 为什么要有结构体？ 

当一组数据总是围绕着某个主体时，表示这些数据与这个主体之间有紧密的关联关系，可将该主体提取为一个对象，围绕着主体的数据为这个对象的**属性**，**围绕着这些数据的操作**即为该主体的**成员函数**。

在 Golang 中，对象的体现方式为**结构体`struct`**。

## 什么是结构体

结构体是由一系列具有**相同类型或不同类型**的数据，比如`int`, `string`, `[]int`, `float64`, `map[xx]yyy`, etc.,
构成了一个数据集合的整体，表示一项记录。

### Example: 写一个支持加减乘除和余数运算的计算器

如果不使用结构体，我们可能的写法是：

```go
package main

import (
	"fmt"
)

// do calculation on integers: if we do not use structure...
func main() {
	var left, right int
	var operator string
	fmt.Scanln(&left)
	fmt.Scanln(&operator)
	fmt.Scanln(&right)
	switch operator {
	case "+":
		fmt.Println(left + right)
	case "-":
		fmt.Println(left - right)
	case "*":
		fmt.Println(left * right)
	case "/":
		fmt.Println(left / right)
	case "%":
		fmt.Println(left % right)
	// if we want to add another type of calculation, we need to add another case
	// switch-case makes a whole, which is not convenient to test
	default:
		fmt.Println("this calculation is not supported.")
	}
}
```
这样写的问题在于，后续如果要在它的基础上增加需求、进行变更和编写测试时，`switch-case`是作为一个整体来被修改和测试，变更范围会很大，提高了风险。

解决方法是将这些计算类型提取出来，定义为一个结构体`structure`：

![image](https://github.com/AdaSheng07/ready.to.go/blob/5527b38186256ef5a1e1984e28cb269beb658a2b/pics/structure_1.png)

## 结构体的定义

定义结构体需要使用`type`和`struct`语句，`type`设定了结构体的名称，`struct`定义一个新的数据类型，结构体中有一个或者多个同类型或不同类型的成员：
```
type strcut_variable_type_name struct{
    <attribute_1> <attribute type>
    <attribute_2> <attribute type>
    ...
    <attribute_n> <attribute type>
}
```
定义结构体类型之后，就能用结构体来声明一些变量：
```
variable_name := structure_variable_type_name {value_1, value_2, ..., value_n} 
// value_1, value_2, ..., value_n match attribute_1, attribute_2, ..., attribute_n respectively

// or we can use map key-value to declare variables
variable_name := structure_variable_type_name {attribute_1: value_1, attribute_2: value_2, ..., attribute_n: value_n}
// neglected attributes in declaration will initialized as 0 or <nil>

// other methods to declare a structure variable:

type Person struct {
    Name string
    Age int
    Height float64
    Weight float64
    Gender string
    calcFatRateFunc func()
    children []*Person
}

var <variable_name> [*]<structure_variable_type_name>
e.g. var tom Person
     var tom *Person

<variable_name> := [&]<structure_variable_type_name>{}
e.g. tom := Person{}
     tom := &Person{}

<variable_name> := [&]<structure_variable_type_name>{[attribute_name: attribute_value, ...]}
e.g. tom := &Person{"Tom", 23}
     tom := &Person{Name: "Tom", Age: 23} // highly recommended

var <variable_name> *<structure_variable_type_name> = new(<structure_variable_type_name>)
e.g. var tom *Person = new(Person)

// new is a built-in function in Golang, when we use this method to declare a structure variable, we get a pointer of this structure_variable Person
// func new(Type) *Type
// tom is a pointer variable of type *Person
// this declaration is equal to: 
//   tom := &Person{}
```

## 结构体的属性

```
type calculator struct {
    a, b int
    sum int
    product int
    quotient int
    remainder int
}

func (c calculatr) Add() int {
    return c.a + c.b
}
```

两个结构体`struct`之间的操作是由某一方发起的，属于某个对象。利用成员函数可以实现结构体之间的操作，实现彼此的互动。**发起操作的对象拥有这个属性。** 在上例中，结构体变量对象`c`拥有成员函数`Add()`，增加了它的属性。

> 【Tips】成员函数可以在同一个包的不同文件中定义。

在此基础上，我们可以优化支持加减乘除和余数运算的计算器的编写：



## 结构体的操作

## 结构体的嵌套










