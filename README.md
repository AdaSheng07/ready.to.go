# Ready To Go

## INDEX




****
## 🔷 Module I


## 环境配置

## 变量 常量

## 计算符号 字符串

## 条件 分支表达式 循环

## 🔸 数组 切片

### 数组 [link](https://github.com/AdaSheng07/ready.to.go/blob/56ec88917763f732b33170478d0e1d794ec9bef9/chapter1/006.array1/main.go)  
数组是具有相同唯一类型的一组已编号且长度固定的数据项序列，即由固定长度的特定类型元素组成的序列。  
这种类型可以是任意的原始类型例如整型`int`、字符串`string`、浮点型`float`或者自定义类型。

相对于去声明 `number0, number1, ..., number99` 的变量，使用数组形式 `numbers[0], numbers[1] ..., numbers[99]` 更加方便且易于扩展。

**数组的声明与初始化**
```
var variable_name [SIZE] variable_type = [SIZE]variable_type{......}
var variable_name = [SIZE]variable_type{......}
variable_name := [SIZE]variable_type{......}
```
如果数组的长度不确定，可以用`...`代替数组的长度，编译器将根据数组中的元素个数自行推断数组的长度。
```
e := [...]int{1, 2, 3, 4}
```
设置数组长度以后，还可以通过指定下标来初始化元素：
```
age := [5]int{1:24,4:35}
```

**数组的特性**
- 属于变量，固定长度，固定类型，动态赋值（运行过程中数组的元素值可以改变）
- 数组的类型由数组的元素类型和数组的长度两者决定，长度和元素类型一旦定义，都不可变：
  ```
  ages3 := [5]int{34, 75, 25, 57, 24}
  ages3 = [6]int{34, 75, 25, 57, 24, 99} // 报错：ages3的类型type是[5]int，不能重复赋值为[6]int
  ```

**数组的常规操作**
- 数组是变量，可以对整个数组进行动态赋值
- 数组中的每个元素也是变量，可以对数组中的每个元素进行动态赋值，注意下标从`0`开始，不可越界为元素赋值
- 数组有固定长度，可用`len()`求出数组长度
- 可用`for range`搭配访问每个元素：
  ``` go
  var array2 [3]int = [3]int{1, 2, 3}
  for i := 0; i < len(array2); i++ {
      fmt.Println(array2[i])
  }
  for i, val := range array2 {
      fmt.Printf("%d, array2[%d]: %d\n", array2[i], i, val)
      fmt.Printf("%d\t%d\n", i, val)
  }
      
  /*
  1
  2
  3
  1, array2[0]: 1
  0       1
  2, array2[1]: 2
  1       2
  3, array2[2]: 3
  2       3
  */
  ```
### 多维数组 [link](https://github.com/AdaSheng07/ready.to.go/blob/d5077093f5549509d58f58ed096831d06cffd7aa/chapter1/006.array2/main.go)

**多维数组的声明**

声明一个`n`维数组：
```
    var variable_name [SIZE1][SIZE2]...[SIZEn]variable_type
    
    e.g. var a [3][3]int
```
一组方括号`[]`定义的是一个一维数组，多组方括号`[][]`或者`[][][]`或者任意多`[]`的定义多维数组。

以下方的二维数组为例，二维数组可认为是一个表格，`x`为行，`y`为列，二维数组可以通过`a[i][j]`来表示。

![img.png](pics/img.png)

**初始化二维数组**

```
    // 方式一：数组长度有限
    variable_name := [SIZE1][SIZE2]variable_type{}
        variable_name[0] = [SIZE2]variable_type{......}
        variable_name[1] = [SIZE2]variable_type{......}
        ...
        variable_name[SIZE1 - 1] = [SIZE2]variable_type{......}
    
    variable_name := [SIZE1][SIZE2]variable_type{
        [SIZE2]variable_type{......}
        [SIZE2]variable_type{......}
        ...
        [SIZE2] variable_type{......}
    }
    
    // 方式二：优化数组长度管理，支持动态添加
    variable_name := [...][SIZE2]variable_type{
        [SIZE2]variable_type{......}
        [SIZE2]variable_type{......}
        [SIZE2]variable_type{......}
        ...
    }
```
**多维数组的遍历**

多维数组的遍历通过`for-loop`层层降维，也可以用`for range`访问每个元素。对于一个`n`维的数组：
```
    for d1, d1val := range array{
        for d2, d2val := range d1val{
            for d3, d3val := range d2val{
            ...
            }
        }
    }
```
再以二维数组为例：
``` go
    c := [...][3]string{
            [3]string{"take", "take", "take"},
            [3]string{"go", "go", "go"},
            [3]string{"fine", "fine", "fine"},
        }
    for d1, d1val := range c {
        for d2, d2val := range d1val {
            fmt.Println(d1, d1val, d2, "val =", d2val)
        }
    }
```
***
### 切片 [link]()
****


## Map


## 🟦 Module 1 Practice Collection

### 数组
 
> **Q1**  创建一个一维数组，并反转它  
>
>      [>> How to Reverse an Array w/o Creating a New One?](https://github.com/AdaSheng07/ready.to.go/blob/76668b88b729bcbd51f76fcbb93e07b1997d2155/chapter1/006.reverseArray/main.go)

> **Q2**  用多维数组写一个日历表，需要考虑每个月的天数不同，需要考虑是平年还是闰年来专门处理二月  
>
>

> **Q3**  【提升篇】日历按照一周的宽度显示（第一列为周一），且每个日期匹配到对应的列

****
## Module II
## 函数

❖ [Go 语言向函数传递数组](https://www.runoob.com/go/go-passing-arrays-to-functions.html)

## 包
## 函数方法论



