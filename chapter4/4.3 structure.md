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

![image](https://github.com/AdaSheng07/ready.to.go/blob/c92563c53948bf266ed80f8d757221108e9eb743/pics/structure_2.png)

[>> 完整代码实现]()

运行结果为：
```
3
-1
2
0
1
```

## 结构体的操作

- 结构体的属性在**同一个包内**均可见
- **跨包的私有**结构体、成员变量和成员函数在包外部不可见
- 只有**公有**的结构体、成员变量、成员函数在包外部可见
- 结构体的成员函数执行时，只能通过**结构体指针的成员函数**进行更改

> 【注意】如果在执行结构体的成员函数时，需要对结构体的成员变量进行更改，只能通过**结构体指针的成员函数**。比如，在结构体`Calculator`中新增属性`result`，并且在调用成员函数`Add()`之后将计算所得的结果存入`result`中：
>    ``` go
>    package main
>    
>    import "fmt"
>    
>    type Calculator struct {
>    left, right int
>    operator    string
>    result int
>    }
>    
>    func (c Calculator) Add() int {
>    	// structure variable c is a receiver, by receiver we get access to member variables and functions
>    	tempResult := c.left + c.right
>    	c.result = tempResult
>    	fmt.Println("调用Add()函数，c的result =", c.result)
>    	return tempResult
>    }
>    
>    
>    func main() {
>    var a, b int = 1, 2
>    //var op string = "+"
>    
>    	c := Calculator{
>    		left:  a,
>    		right: b,
>    		//operator: op,
>    	}
>    	fmt.Println(c.Add())
>    	fmt.Println("c.result =", c.result) // this gives out: c.result = 0, WHY?
>    }
>    ```
> 运行结果为：
> ```
> 调用Add()函数，c的result = 3
> 3
> c.result = 0
> ```
> 这是因为`Add()`中的`c`与`main()`中的`c`并不是同一个，`Add()`建立在结构体对象`Calculator`上，而不是结构体`Calculator`的指针上，我们可以分别打印两处`c`的地址来比较，得到：
> ```
> &c @main() is: 0xc000098180
> &c @c.Add() is: 0xc0000981b0
> ```
> 两处的`c`地址不同，`tempResult`的值仅覆盖了地址为`0xc0000981b0`的`c`中，`main()`中的`c`无法变更。

**如何解决这个问题？** >> 将`Add()`更改为定义在**结构体指针的成员函数**：
```
func (c *Calculator) Add() int {
	// structure variable c is a receiver, by receiver we get access to member variables and functions
	tempResult := c.left + c.right
	c.result = tempResult
	fmt.Printf("&c @c.Add() is: %p\n", &c)
	fmt.Println("调用Add()函数，c的result =", c.result)
	return tempResult
}
```
运行结果为：
```
&c @main() is: 0xc00006c180
&c @c.Add() is: 0xc00000e030
调用Add()函数，c的result = 3
3
c.result = 3
```
> 【观察】程序更新后的`c.Add()`中的`c`的地址不同于更改之前的地址`0xc0000981b0`，此时的`c`是一个新的对象。Golang 内部在成员函数进行计算的过程中，会出现一个新的拷贝，这个拷贝可能是一个引用。
> 
> 定义在结构体对象本身上的成员函数，无论怎样操作，都无法触碰到结构体对象实体内的内容，无法更改内部属性。
> 
> 如果一个对象含有非常多的属性、数据和值，而成员函数有建立在结构体对象本身，每一次调用成员函数都会拷贝一次结构体对象，产生一次副本，最终会产生非常多的垃圾数据，依赖 Golang 的自动垃圾回收功能，内存会在瞬时不断地上升和下降。对象越多，消耗越大，程序性能的下降越明显。
> 
> 在使用场景中，如果我们需要保护对实体内部的值，保证它在函数被调用的过程中不被更改，或者确保运算的正确性时，我们需要对其进行**显式声明**，比如，当其它成员函数都是建立在对象本身上时，增加建立在对象指针上的成员函数`SetResult`，只有`SetResult`能够触及结构体对象实体的内容：
> ```
> func (c *Calculator) SetResult(result int) {
> 	c.result = result
> }
> ```
> 在将成员函数定义在结构体的指针上时，本身就完成了一个很好的封装。可以尽量将成员函数定义在结构体的指针上，但这样的操作也是有风险的。

## 结构体的嵌套

在 Golang 中，**继承**是通过结构体的嵌套来完成的。结构体的嵌套是指**直接嵌入其它结构体**完成结构体的定义。被嵌入的结构体的所有成员变量、成员方法都可以直接被沿用继承。

- 被嵌入的结构体的成员变量、成员函数也遵循公有、私有访问限制（同包/跨包）
- 嵌套结构体的语法与定义成员变量的语法是不同的，区别在于被嵌入的结构体是匿名的，没有变量名。如果一个结构体`struct`嵌套了另一个**有变量名的结构体**，这个模式叫做**组合**
    ```
    type calculator struct {
        a, b int
        operator string
    }
    
    func (p *calculator) Add() {
        return p.a + p.b 
    }
    // this function cannot be used in NewCalculator, c.Add() does not exist
    
    type NewCalculator struct {
        old calculator // this is a combination of two structures, the structure calculator was quoted by a memeber variable old
    }
    
    // if new funcitons, e.g. Add() and QuadraticSum(), are added into structure calculator, old cannot have then and we need to copy and maintein new functions ourselves, as below: 
    
    func (c *NewCalculator) Add() int {
        return c.old.Add()
    }
    // ...
    func (c *NewCalculator) QuadraticSum() int {
        return c.old.a * c.old.a + c.old.b * c.old.b
    }
    ```
  具体实现和练习如下，观察结构体不同嵌套方式的最终结果：
  ```go
      package main
    
      import "fmt"
    
      func main() {
        var num1, num2 int = 2, 3
    
        p := calculator{
          a: 0,
          b: 1,
        }
        fmt.Printf("%d + %d = %d\n", p.a, p.b, p.Add())
    
        old := calculator{
          a: num1,
          b: num2,
        }
        fmt.Printf("%d + %d = %d\n", old.a, old.b, old.Add())
    
        var c1 NewCalculator1
        c1.old.a = 3
        c1.old.b = 4
        fmt.Println("Nest structure calculator into structure NewCalculator1 w name 'old':")
        fmt.Printf("%d + %d = %d\n", c1.old.a, c1.old.b, c1.old.Add())
        fmt.Printf("%d + %d = %d\n", c1.old.a, c1.old.b, c1.Add())
        // c1.Add(): unsolved reference error if Add() is not redefined on *NewCalculator1
    
        c2 := NewCalculator2{}
        c2.a = 5
        c2.b = 6
        fmt.Println("Nest structure calculator into structure NewCalculator2 w/o name:")
        fmt.Printf("%d + %d = %d\n", c2.a, c2.b, c2.Add())
        c2.calculator.a = 7
        c2.calculator.b = 8
        fmt.Printf("%d + %d = %d\n", c2.calculator.a, c2.calculator.b, c2.calculator.Add())
    
        c3 := NewCalculator2{}
        c3.a = 9
        c3.b = 10
        fmt.Printf("%d + %d = %d\n", c3.a, c3.b, c3.Add())
        fmt.Printf("%d^2 + %d^2 = %d\n", c3.a, c3.b, c3.QuadraticSum())
      }
    
      type calculator struct {
        a, b int
      }
    
      func (p *calculator) Add() int {
        return p.a + p.b
      }
    
      func (p *calculator) QuadraticSum() int {
        return p.a*p.a + p.b*p.b
      }
    
      type NewCalculator1 struct {
        old calculator // a combination of two structures
      }
    
      func (c1 *NewCalculator1) Add() int {
        return c1.old.Add() // redefine Add() on *NewCalculator, otherwise there is no access to c1.Add()
      }
    
      type NewCalculator2 struct {
        calculator // nesting a anonymous structure calculator
      }
    
  ```
- 结构体也可以嵌套其它外部项目，并进行扩展定制，比如：
  ``` go
    type MyCommand struct {
        cobra.Command
    }
  ```
- 如果一个结构体`struct`嵌套了多个匿名结构体，那么这个结构可以直接访问多个匿名结构体的方法，从而实现**多重继承**
- 在创建变量时，嵌入的对象也需要实例化，尤其是引用类型，如：指针和`map`等。如果结构体`StructA`嵌套的是另一个结构体`StructB`的指针，创建结构体`StructA`的变量时，结构体`StructB`的指针需要实例化：
  ```
  type StructA struct {
      *StructB
  }
  // ...
  func main() {
      // ...
      newA := StructA{&StructB}
  }
  ```
  再比如，如果结构体`StructB`中有`map`属性的成员变量，在新建`StructB`变量时，`map`属性的成员变量也需要实例化，否则会`panic: assignment to entry in nil map`：
  ```go
      c5 := StructB{}
      c5.options["key"] = "value" // panic: assignment to entry in nil map
  
      c6 := StructB{
          options: map[string]string{},
      }
      c6.options["key"] = "value"
      fmt.Println(c6.c)
      fmt.Println(c6.options)
  ```
  非引用类型的结构体对象在嵌套时能够自动被实例化，但引用类型，如指针和`map`必须手动实例化，否则会有空指针或空`map`报错。
















