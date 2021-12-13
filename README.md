# Ready To Go

## INDEX

---

## Module I
### 环境配置

### 变量 常量

### 计算符号 字符串

### 条件 分支表达式 循环


### 基于BMI的体脂计算器


### 数组 切片

#### 数组 [link](https://github.com/AdaSheng07/ready.to.go/blob/62fce9be1b277ad49d400d6257d41182817f5887/chapter1/006.array/main.go)  
数组是具有相同唯一类型的一组已编号且长度固定的数据项序列，即由固定长度的特定类型元素组成的序列。  
这种类型可以是任意的原始类型例如整型`int`、字符串`string`、浮点型`float`或者自定义类型。

相对于去声明 `number0, number1, ..., number99` 的变量，使用数组形式 `numbers[0], numbers[1] ..., numbers[99]` 更加方便且易于扩展。

数组元素可以通过索引（位置）来读取（或者修改），索引从 `0` 开始，第一个元素索引为 `0`，第二个索引为 `1`，以此类推。

- 数组的声明与初始化
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

- 数组的特性：
  - 属于变量，固定长度，固定类型，动态赋值（运行过程中数组的元素值可以改变）
  - 数组的类型由数组的元素类型和数组的长度两者决定，长度和元素类型一旦定义，都不可变：
    ```
    ages3 := [5]int{34, 75, 25, 57, 24}
    ages3 = [6]int{34, 75, 25, 57, 24, 99} // 报错：ages3的类型type是[5]int，不能重复赋值为[6]int
    ```
- 数组的常规操作
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

---
### Map

---
## Module II
### 函数

### 包
### 函数方法论



