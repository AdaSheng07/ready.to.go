# 指针 pointer

## 什么是指针

变量是一种用于引用计算机内存地址的占位符。在 Golang 中，我们用`&`来取变量的地址，将`&`放在一个变量名前使用，就会调用返回相应变量的内存地址。一个指针变量指向了一个值的内存地址。比如：
```go
package main

import "fmt"

func main() {
	a, b := 1, 2
	fmt.Printf("变量a的地址是：%x\n", &a)
	fmt.Printf("变量b的地址是：%x\n", &b)
	add(a, b)
	fmt.Println(a)
	// in this case, the printed result is still 1
}

// how to change value of a without returning any value in function add()
func add(a, b int) {
	// in this function, values of a and b are copied for calculation, they are not the same a and b in main()
	a = a + b
}
```
运行上面这段代码得到的结果是：
```
变量a的地址是：c00012a008
变量b的地址是：c00012a010
1
```
定义的`add()`函数没有返回值，其中的`a`与`b`值的改变无法传导至`main()`函数中，如何在没有返回值的条件下，计算`a + b`并用结果覆盖`a`的值？
```go
package main

import "fmt"

func main() {
	a, b := 1, 2
	// instead of copying values in variables a and b, here we use addresses of variables a and b
	add(&a, &b)
	fmt.Println(a) // this time, a will be 3
}

// *int means add() has parameters that are addresses for integer-type variables
func add(x, y *int) {
	// x and y are pointers, *x and *y indicate values of variables x and y, respectively
	// extract values of x and y, add them up and put the result back into variable x
	*x = *x + *y
}
```

![image](https://github.com/AdaSheng07/ready.to.go/blob/b952c0c6d1b68b760ca0431308ce45a9c01adb47/pics/pointer_1.png)

- `&a`,`&b`相当于"编号1"、"编号2"
- `&a`,`&b`的类型为**指针**
- `&`的作用为获取变量的地址
- `*int`表示变量的类型为**整数的指针**，即`&a`和`&b`的类型都是`*int`
- 在 Golang 中，地址（如上例中的`c00012a008`）不能直接作为指针传入，在 C/C++ 中可以
- 在 Golang 中，不能直接对整数进行强制类型转换，如`x(*int)`，是不可行的

## 如何使用指针

- 通过`&<XX>`来获取变量的地址
- 通过`*<XX>`来获取指针所指的目标
- 指针可以指向任何变量（但不能指向常量）
- 指针不支持运算
- 当我们需要对参数变量有互动地赋值传递时，可以用指针

```go
package main

import "fmt"

func main() {
	a, b := 1, 2
	add(&a, &b)
	fmt.Println(a)

	c := &a          // c is a pointer variable, its type is *int, it saves the address of a
	d := &c          // d is also a pointer variable, its type is **int, it saves the address of c
	fmt.Println(*c)  // this will give out value of a
	fmt.Println(*d)  // this will give out c, that is &a, the address of a
	fmt.Println(**d) // this is *(c), which will give out value of a
	fmt.Println("d =", d, "*d =", *d, "**d =", **d)
	// the passing process can be infinite:
	// d is the address of c
	// *d is the value of c, which is the address of a
	// **d is equal to *c, which is the value of a

	// pointers can point to any variables
	m := map[string]string{}
	mp1 := &m // mp1 is a pointer variable, its type is *map[string]string, it saves the address of map m
	fmt.Println(mp1)
	put(m)
	fmt.Println("*mp1 =", mp1)

	f1 := add // f1 is a function
	f1(&a, &b)
	fmt.Println("f1, add =", a) // this gives out 5 from 3 + 2
	f1p1 := &f1                 // f1p1 is a pointer variable, its type is *func(*int, *int), it saves the address of function f1
	// if we use *f1p1(&a, &b)... Error: it does not meet the priority rule, golang calculates first, gives out the result and then use *
	// we want to use *f1p1 first to get the function f1 itself, and then do the calculation, we MUST use brackets
	(*f1p1)(&a, &b)
	fmt.Println("f1p1, add =", a) // this gives out 7 from 5 + 2
}

func put(m map[string]string) {
	m["a"] = "AAA"
}

func add(x, y *int) {
	*x = *x + *y
}
```

## 指针与引用类型的对比理解

```go
package main

import "fmt"

func main() {
	{
		var nothing *int                 // notice: nothing is a pointer that points to nil
		fmt.Println("nothing:", nothing) // this is <nil>
		*nothing = 3                     // *nothing does not exist, so no valid assignment operations --> panic: runtime error: invalid memory address or nil pointer dereference
	}
	// this is also why we must instantiate a map, because we cannot assign any values to a nil map:
	{
		var nothingMap map[string]string
		nothingMap["nothingMap"] = "empty" // panic: assignment to entry in nil map
		fmt.Println("nothingMap:", nothingMap) // this is map[]
	}
	{
		var nothingSlice []int
		fmt.Println(nothingSlice) // this is []
		nothingSlice[0] = 100     // panic: runtime error: index out of range [0] with length 0
		// we cannot assign value to an uninstantiated nil slice, but we can append it w/o instantiation:
		nothingSlice = append(nothingSlice, 100)
		fmt.Println(nothingSlice) // this is [100]
	}
}
```
