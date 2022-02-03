# 对象 object

## 面向对象的必要性

以之前的体脂计算器为例，如果班上一共有100个人，要把每个人体脂相关的信息都录入进程序中，在输入一个人的姓名时，就能输出这个人的身高、体重、年龄、性别和体脂信息。

按照之前的程序编写方法，我们可能会创建100组信息变量：
```
name1, gender1, height1, weight1, age1 := ...
name2, gender2, height2, weight2, age2 := ...
name3, gender3, height3, weight3, age3 := ...
...
```
或者创建5个数组来分类存储信息，同索引数的一列代表一个人的一组信息：
```
names := []string{}
genders := []string{}
heights := []float64{}
weights := []float64{}
age := []int{}
```
以上方法可行，但当数据体量大幅增加时，对于数据的增删改查都非常麻烦，而这些变量之间是存在一定联系的，所以我们引入**对象**。

利用对象，我们可以将与同一个人相关的所有属性/信息（姓名、性别、身高、体重和年龄）集合在一起。每一个对象（及其所有属性）都是一个**完整、独立**的整体，与其它对象是隔离的。

## Go 的面向对象编程

- Golang 支持面向对象编程（OOP，Object-oriented programming），虽然 Golang 并不是严格的面向对象编程语言，是弱对象的
- Golang 中的结构体`struct`与其它语言（如 Java）中的类`class`在面向对象编程中处于同等地位，在 Golang 中利用`struct`新建一个对象，将相关的内容组装：
    ```go
    package main
    
    func main() {
    
    }
    type Person struct {
            name   string
            gender string
            height float64
            weight float64
            age    int
        }
    ```
  在提取对象、组装相关内容和建立程序架构时，遵循**高内聚、低耦合**的规则：
  - 高内聚：内聚是指内部间聚集、关联的程度，高内聚就是指要高度的聚集和关联。
  - 低耦合：耦合又称为快间联系，指软件系统结构中各模块间相互联系紧密程度，耦合性越高，模块的独立性越差。在低耦合的设计中，在对一个模块进行变更时，对其它模块的变更很少。
- Golang 中没有用哪个的继承、方法重载、构造函数等，但 Golang 也可以做到继承
- Golang 通过接口`interface`完成面向对象编程

## 体脂计算器与对象

```go
package main

import (
	"fmt"
	gobmi "github.com/armstrongli/go-bmi"
)

func main() {

	persons := []Person{
		{
			"小强",
			"男",
			1.7,
			70,
			35,
		},
	}
	for _, item := range persons {
		bmi, err := gobmi.CalcBMI(item.weight, item.height)
		fmt.Println(bmi, err)
	}

}

type Person struct {
	name   string
	gender string
	height float64
	weight float64
	age    int
}
```




