# ⚪️ <u>Ready To Go</u>

## INDEX

****
# 🟡 <u>Module 1</u>

## 环境配置
## 变量 常量
## 计算符号 字符串
## 条件 分支表达式 循环
****
## 🔶 数组 切片
### 🔸 数组 [link](https://github.com/AdaSheng07/ready.to.go/blob/56ec88917763f732b33170478d0e1d794ec9bef9/chapter1/006.array1/main.go)  
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
### 🔸 多维数组 [link](https://github.com/AdaSheng07/ready.to.go/blob/d5077093f5549509d58f58ed096831d06cffd7aa/chapter1/006.array2/main.go)

**多维数组的声明**

声明一个`n`维数组：
```
    var variable_name [SIZE1][SIZE2]...[SIZEn]variable_type
    
    e.g. var a [3][3]int
```
一组方括号`[]`定义的是一个一维数组，多组方括号`[][]`或者`[][][]`或者任意多`[]`的定义多维数组。

以下方的二维数组为例，二维数组可认为是一个表格，`x`为行，`y`为列，二维数组可以通过`a[i][j]`来表示。

![image](https://github.com/AdaSheng07/ready.to.go/blob/e164bedc9f631503e3a19aa8e6a4fe2113e0add1/pics/img.png)

**初始化二维数组**

```
    方式一：数组长度有限
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
    
    方式二：优化数组长度管理，支持动态添加
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

### 🔸 数组与切片的比较

切片是对数组的抽象。数组的长度是不可改变的，
而相比之下，切片是一种更加灵活，功能更多的内置类型，可以理解为一种变长数组，或者动态数组。
切片的长度是不固定的，可以利用`append`追加元素，在追加时可能使切片的容量增大。

在实际开发运用中，极少情况下才会使用数组，大部分情况下使用切片。
使用数组的特定情况：
- 初始化内容
- 固定数组，固定长度，内容不会做变更
- 节省内存空间，切片会自动扩容与空间释放

### 🔸 切片的声明与初始化 [link](https://github.com/AdaSheng07/ready.to.go/blob/02c813c001594cca665171f452d097302ca9f901/chapter1/007.slice1/main.go)

一个切片在未初始化之前默认为`nil`，长度为`0`。切片的声明方式有：
```
  方式一：声明未指定大小的数组来定义切片
  var slice_name []type 
  * 注意：切片声明时不需要说明长度
  
  方式二：内置函数 make() 初始化切片
  var slice_name []type = make([]tape, length, capacity)
  slice_name := make([]type, length, capacity)
  * 注意：capacity是可选参数，length是数组的长度，也是切片的初始长度
  
  方式三：引用数组
  var array_name [array_size]type
  slice_name := [array_name:]
  slice_name := array_name[startIndex:endIndex]
  
  方式四：引用切片
  var slice1_name []type
  slice2_name := [slice1_name:]
  slice3_name := slice1_name[startIndex:endIndex]
```

### 🔸 切片的操作

**动态增加、删除、截取切片中的元素** [link](https://github.com/AdaSheng07/ready.to.go/blob/02c813c001594cca665171f452d097302ca9f901/chapter1/007.slice1/main.go)
1. 利用`append`对切片进行追加、插入和删除元素的操作

- 切片没有直接的删除操作
- `slice_name[lower-bound:upper-bound]`的区间是**左闭右开**的
- 用`append`插入元素时会覆盖原有元素，需要提前做备份，如何做一份有效备份？
  ```
  a := []int{6, 4, 2, 0}
    
  1. Invalid, backup changes while a changes:
  backup := a[1:] 
    
  2. Valid, backup does not change as a changes: 
  backup := append([]int{}, a[1:]...)
    
  3. Valid, backup does not change as a changes:
  var backup []int = make([]int, len(a[1:]), cap(a[1:])*2)
  copy(backup, a[1:])
  ```
2. 还可以利用`make`对切片进行扩容
```
  make(slice_name type, len(slice_name), cap(slice_name))
```
3. 利用`copy`拷贝切片的内容
- 在拷贝`source`切片到`destination`切片之前，需要先声明`destination`切片并初始化容量为`source`切片的两倍（可用`make`）
- 语法：`copy(dst []Type, src []Type)`

**`Golang`支持`string`与`byte`、`rune`切片之间的转换** [link](https://github.com/AdaSheng07/ready.to.go/blob/02c813c001594cca665171f452d097302ca9f901/chapter1/007.slice2/main.go)
- 只做特定的支持：`string array`与`byte slice`之间的切换，其它的类型如`int`等是无法转换为`byte slice`的
- 为了表示更多字符，`Golang`字符串实现基于`UTF-8`编码，也支持`ASCII`码方式逐字符访问
- 在`Golang`中，通过`rune`类型,可以方便地对每个`UTF-8`字符进行访问
- 如果仅使用`byte`来读取、转换、写入字符，一般不会出现，但如果继续使用`byte`对字符进行修改，则会出现字符非`ASCII`码的问题

**Appendix**

☞  [Go 字符串编码，Unicode 和UTF-8](https://segmentfault.com/a/1190000019361462)  
☞  [GO操作切片数组时不当，数据被覆盖](https://blog.csdn.net/weixin_44145242/article/details/111299356)  
☞  [Go 字符串编码，Unicode 和UTF-8](https://segmentfault.com/a/1190000019361462)
****

## 🔶 Map

`Map`是一个`key-value`组合的结构体，一种无序的键值对的集合，常用实现方式是二叉树和哈希表（散列表）。
`Map`通过`key`来快速检索数据，`key`的作用类似于数组与切片中的索引，指向数据的值`value`。
`key`总是唯一的，相同的`key`拿到相同的`value`，更新`value`时，会覆盖相同`key`的原有`value`值。

### 🔸 使用`Map`的优势：快速查找，从`key`定位到`value`

当数据量很大时，e.g. >10000，仍然用数组或者切片来存储，会出现什么问题？

如果我们需要查找数据集合中的某一组数据，需要`for-loop`遍历整个数组/切片，再比较值是否相等来锁定目标数据，时间复杂度是`O(n)`，很大。

如果在总共` 10000 `人中找第` 5000 `个人` 2000 `次，一共在切片上比较`5000 * 2000`次，效率太低。

`Map`是为了解决这样的问题而存在的，只要数据组中的每一组数据都是是唯一的，我们就可以用`key-value`构造`Map`来进行存储。

![img.png](pics/map.png)

### 🔸 `Map`的定义与初始化 [link](https://github.com/AdaSheng07/ready.to.go/blob/0bc031aec6339e4f13d7ab4705546030a8ec0dc6/chapter1/008.map1/main.go)

定义`map`时注意`key`与`value`的类型都需要声明。
`map`定义可以是符合类型，它的`key`与`value`都可以是任意类型，e.g. `int`,`float64`, `array`, `slice`, etc.

主要有两种定义方式：
```
  1. use keyword `map` to declare and initialize
  var map_variable_name map[key_data_type]value_data_type
  map_variable_name := map[key_data_type]value_data_type{}
  map_variable_name := map[key_data_type]value_data_type{key1: value1, key2: value2, key3: value3, ...}
  
  2. use keyword `map` to declare, then use built-in function make() to initialize
  var map_variable_name map[key_data_type]value_data_type
  map_variable_name = make(map[key_data_type]value_data_type, length)
  
  3. use built-in function make() to declare, then add in key-value elements
  map_variable_name := make(map[key_data_type]value_data_type, length)
  map_variable_name[key1] = value1
  map_variable_name[key2] = value2
  ...
 ```
`Map`定义的`key`与`value`类型可以**嵌套使用**，但要注意多层嵌套后代码意义是否会混淆，影响可读性，如：
```
  map_variable_name := map[string]map[string]map[int]float64{}
```

### 🔸 对`map`的操作 [link](https://github.com/AdaSheng07/ready.to.go/blob/0bc031aec6339e4f13d7ab4705546030a8ec0dc6/chapter1/008.map2/main.go)

如果不初始化/实例化`map`，就会默认初始化为`nil map`。`Map`不用实例化就可以读取和删除，但`nil map`**不能**用来写入/存放键值对：
```
  panic: assignment to entry in nil map
```

**`Map`的增删改查**

`Map`属于`Golang`的变量范畴，也是强类型的。当定义好`map`后，它只能容纳对应类型的数据。

>【回顾】如何在切片中间插入一个元素?
> 
> 备份切片扩容 >> 拷贝备份 >> 用`append`插入值 >> 用`append`结合备份与追加元素后的新切片

此操作的**风险高，代价大**，而`Map`的增删改查是极其方便的：

```
  1. add/change key-value
  map_variable_name[key_name] = key_value
  * If this key does not exist in map, add this key-value pair in map; 
    else, change its value in the map.
  
  2. delete key-value
  delete(map_variable_name, key)
  * We can delete the same key repeatedly from map.
  
  3. lookup key-value：take key as index of map
  map_variable_name[key]
```
【注意】如果此时`map`中不存在这个`key`，会自动加入此`key`，但返回的`value`是**假值**。  

> 如何判断验证真假值（`key`在`map`中是否存在）呢？
> ```
> value, ok := map_variable_name[key]
> * If ok is true, value is a true value;
>   else if ok is false, value is a false value.
> ```


**`Map`的遍历**
```
  for key, value := range map_variable_name {
      fmt.Printf("%v\t%v\n", key, value)
  }
  
  for key := range map_variable_name{
      fmt.Println(map_variable_name[key])
  }
```

**`Map`的合并**
```
  map_variable_name1, map_variable_name2 := map[key_type]value_type, map[key_type]value_type
  map_variable_name1[key1] = value1
  map_variable_name2[key2] = value2
  for k, v := range map_variable_name1{
      map_variable_name2[k] = v
  }
```

**Appendix**

☞   [Go 语言Map(集合)](https://www.runoob.com/go/go-map.html)  
☞   [Go语言map的创建](https://haicoder.net/golang/golang-map-make.html)

## 🔵 Module 1 Practice Collection

### 🔹 数组

> **Q1**  创建一个一维数组，并反转它  
>
>       >>  [How to Reverse an Array w/o Creating a New One?](https://github.com/AdaSheng07/ready.to.go/blob/76668b88b729bcbd51f76fcbb93e07b1997d2155/chapter1/006.reverseArray/main.go)

> **Q2**  用多维数组写一个日历表，需要考虑每个月的天数不同，以及是平年还是闰年来专门处理二月  
>
>       >>  [Print out Calendar of a Given Year](https://github.com/AdaSheng07/ready.to.go/blob/b21fd48ba4780bb7b5fc1dc8a919e1cd3ef14111/chapter1/006.calendar1/main.go)

> **Q3**  【提升篇】日历按照一周的宽度显示（第一列为周一），且每个日期匹配到对应的列
>
>       >>  [Print out Calendar w Weekdays of a Given Year (Advanced)](https://github.com/AdaSheng07/ready.to.go/blob/4f419675f04290dfedfaed716de0d752b912bd1f/chapter1/006.calendar2/main.go)  
>
>      ☞  [计算任何一天是星期几的几种算法](https://blog.csdn.net/luoyayun361/article/details/54881835)  
>      ☞  [golang向上取整、向下取整和四舍五入](https://studygolang.com/articles/12965)
### 🔹 切片
> **Q1**  创建一个一维整数切片，并用循环对它由从小到大排序
>       
> 
> **Q2**  对一副新扑克牌打乱顺序
>       
> 
> **Q3**  有一个包含中英文的切片，如果是英文的，转换它们的大小写
>       
>
> **Q4**  切片去重的实现
>       
>
### 🔹 Map 
> **Q1**  用 Map 管理 20人 的分数，并做如下操作：
> 1. 算出所有人的平均分
> 2. 根据分数高低对这 20 分排名，高分在前 
> 3. 相同分数的在同一行
>       []()

****

# 🟡 <u>Module 2</u>

## 🔶 函数

函数是基本的代码块，用于执行一个任务，以达成预期目的。
在编程中，我们使用不同的函数，划分它们各自的功能，来完成不同的任务。
每个程序必须有命名为`main()`的主函数。

我们为什么需要函数？

- 将一段经常需要使用的代码封装起来，在需要使用时可以直接调用，提高代码的复用率
- 简化代码逻辑，提高代码效率
- 节省代码阅读成本，提高代码的可读性
- 当代码出现问题时，可以更快地锁定`Error`位置，提高可维护性
- 更新编写代码的思路：从顶层展开伪代码结构（目的、投入、预期、产出）再进行模块化落实  
  ...

### 🔸 函数的定义与作用域 [link](https://github.com/AdaSheng07/ready.to.go/blob/f03ac8deb4c07d421624d5c91d1efccbbf8b95b6/chapter2/009.function5/main.go)

完整的函数的定义格式如下：
```
  func function_name([parameter list])(return-value_list){
      // function body: executable statements
  }
```
`func`、函数名`function_name`、参数列表`[parameter list]`的括号`()`和函数体的大括号`{}`是必须的。
由`func`开始函数的声明，函数名`function_name`、参数列表`[parameterlist]`和返回值列表`return-value_list`构成了**函数签名**。

**函数名`function_name`**
- 函数的名称，在同一个包`(package)`里必须唯一
- 函数是通过名称来调用的，名称一定要有意义

**参数列表`[parameter list]`**
- 形式参数，函数的输入项，是可选项
- 参数列表如果为空，调用此函数时也不能有输入项
- 如果参数列表不为空，当函数被调用时，可以将值传递给参数，这个值就是实际参数
- 参数列表指定的内容：参数类型、顺序以及参数的个数
 
**返回值列表`return-value_list`**
- 调用函数后的产出结果通过返回值列表返回
- 当有返回值列表时，函数体中必须有`return`来呼应与之匹配，否则报错
- 有些函数的功能实现不需要返回值，返回值列表也是可选项，可以为空
- 参数和返回值可以是列表，表示可以传入多个参数，也可以返回多个参数

**函数体`function body`**：函数定义的可执行代码的集合

函数在定义时，根据预期目的的不同，会有很多**变种**：
- 只有函数名`function_name`（目的），没有参数列表`[parameter list]`（投入）、返回值列表`return-value_list`（预期和产出），是单纯的函数调用，可用于：
  - 内容的输入输出
  - 加载全局变量
- 有函数名`function_name`（目的）和参数列表`[parameter list]`（投入），
  没有返回值列表`return-value_list`（预期和产出）

**函数的作用域** 
- 函数体的大括号`{}`定义函数的作用域
- 作用域是指特定实体的有效范围，可包含变量、常量、函数、接口、对象等，它们互相可见、可操作
- 作用域可嵌套，在嵌套时，子作用域可见母作用域的所有元素，且子作用域可定义与母作用域同名的变量、常量等，在操作时遵守就近原则
- 已经在`main()`母作用域定义过的变量可以在另一个作用域`block`重定义，覆盖母作用域的变量值
- 一个变量在作用域外部已经定义赋值为`a`而在作用域内部需要使用时：
  - 如果在作用域内部**重新赋值为`b`**，那么变量值将被覆盖更新为`b`，在作用域之外也**保持更新后的值`b`**
  - 如果在作用域内部**重新定义并赋值为`b`**，相当于新定义了一个重名的变量，值为`b`且**仅限于作用域内**，作用域外的同名变量值仍然是`a`
- 在作用域之外定义的变量在所有作用域内外都可以重新赋值使用
- 在处理相似逻辑的代码时，为防止重用共享变量而不自知，可以用`{}`对作用域进行隔离

```go
    package main
    
    import "fmt"
    
    // quantity is outside all functions, so it can be used in all functions
    var quantity int
    
    func main() {
        quantity = 2
        {
            fruit := "banana"
            fmt.Println(fruit, quantity)
        }
        {
            fruit := "banana"
            fmt.Println(fruit, quantity)
        }
        var fruit string = "apple"
        fmt.Printf("在main函数作用域内，fruit is %s, 一共%d个", fruit, quantity)
        newFruit()
    } 
    func newFruit(){
        quantity = 1
        // fmt.Println("newFruit作用域内, 未重定义时，fruit is", fruit) // 报错：Unresolved reference 'fruit'
        var fruit = "watermelon"
        fmt.Printf("在newFruit函数作用域内，fruit is %s, 一共%d个", fruit, quantity)
    }

```

### 🔸 函数的参数与返回值 [link](https://github.com/AdaSheng07/ready.to.go/blob/f03ac8deb4c07d421624d5c91d1efccbbf8b95b6/chapter2/009.function1/main.go)

**<u>定长参数与不定长参数</u>**

- 定长参数：指定具体名称、类型，使用函数的时候必须一一对应，否则报错：
```
  func constructHello(name string) string {
      return fmt.Sprintf("hello, %s", name)
  }
  func main(){
      constructHello("Tom")
  }
```
- 不定长参数：除了指定名称、类型外，需要额外指定特殊符号`...`：
```
  func constructHello(name ...string) string {
      return fmt.Sprintf("hello, %s", name)
  }
  func main(){
      constructHello("Tom", "Jerry")
  }
```

**<u>函数调用与传递参数的方式</u>**

默认情况下，Go 语言使用的是**值传递**，即在调用过程中不会影响到实际参数。

函数使用的参数可称为函数的形式参数，可理解为**预期投入**，定义在函数体内的局部变量。

调用参数，可以通过两种方式来传递参数：
- **值传递**：在调用函数时将实际参数复制一份传递到函数中，这样在函数中如果对参数进行修改，将不会影响到实际参数。
- **引用传递**：在调用函数时将实际参数的地址传递到函数中，那么在函数中对参数所进行的修改，将影响到实际参数。

一个函数的参数用什么方式传递，要看这个函数传递的参数是什么类型：
- `array`, `map`和`interface`等本身就是引用，在作为参数传递时，拷贝的是变量本身，变量变更时实际参数也跟着变更 
- `int`, `string`是值传递，传递的是该参数的一个副本
- 传递一个指针类型的参数，其实传递的是这个该指针的一份拷贝，而不是这个指针指向的值

**<u>递归</u>** [link](https://github.com/AdaSheng07/ready.to.go/blob/f03ac8deb4c07d421624d5c91d1efccbbf8b95b6/chapter2/010.iteration1/main.go) 

如果一个函数的调用链中存在自己调用自己，则将这种调用方式称为递归。以斐波那契数列计算核心为例：
```go
    package main
    func fibonacci(n int) int{
        if n==0 {
            return 0
        }
        if n==1 {
            return 1
        }
        return fibonacci(n-1)+fibonacci(n-2)
    }
```
运用递归时，注意要整个函数需要退出条件与机制，否则函数将无止境地递归调用运行下去。先分析好**终止条件**，再判断处理细节。

**<u>返回值</u>**

在编写代码时，强烈建议**命名返回值**：
```
  func function_name([parameter list])(return_value_name return_value_type){
      function body
  }
```
- 命名返回值在计算逻辑过程中可以直接使用，且可以参与计算
- 当返回值被命名时，在函数体中对返回值进行复制之后，`return`可以为空
- 命名返回值还可以帮组合使用人员精确了解每个返回值的意义
  - 使用方法的人必须知道每个参数的意义，查看函数的实现过程，否则很难理解返回值代表什么。  
  - 对返回值进行命名就可以简化理解过程，快速知道函数的目的和预期产出。
  - 代码提示信息即可告知具体参数及其意义。

### 🔸 将函数作为参数与返回值 [link](https://github.com/AdaSheng07/ready.to.go/blob/f03ac8deb4c07d421624d5c91d1efccbbf8b95b6/chapter2/009.function2/main.go)

**<u>提取函数并重构</u>**
- 选取函数片段 -> Refactor -> Extract Method...
- 函数重命名：Refactor -> Rename

**<u>使用函数</u>**

- 一个函数在定义后，`Golang`支持将该函数作为形式参数传入另一个函数。
- 被传入函数有时也称作**回调函数(callback function)**。
- 形式参数同时也是作为变量存在。
- 以**函数签名**（包括`func`，函数名，参数列表和返回值列表）作为主调函数的形式参数的类型，即传递一个指向函数的指针。

<br>以基于BMI的体脂计算器为例，如果提取重构两个函数分别给出男性和女性的健康程度判断及建议，并在主函数中将这两个函数分别作为形式参数传入，相当于在告诉主调函数中应当使用的方法是什么。

```go
  package main
  
  import "fmt"
  
  func main() {
      var age int
      var fatRate float64
      var isMale bool
      var result string
      // ...
      if isMale {
          result = getFinalFatState(age, fatRate, getFinalFatStateForMale)
      } else {
          result = getFinalFatState(age, fatRate, getFinalFatStateForFemale)
      }
      fmt.Println(result)
  }
  
  func getFinalFatState(age int, fatRate float64, getSuggestion func(age int, fatRate float64) string) string {
      return getSuggestion(age, fatRate)
  }
  
  func getFinalFatStateForMale(age int, fatRate float64) string { 
      // ... 
      return "This man's healthiness state is ..."
  }
  
  func getFinalFatStateForFemale(age int, fatRate float64) string { 
      // ... 
      return "This woman's healthiness state is ..."
  }
```

**<u>将函数作为返回值（方法）</u>**

用一个函数来返回另一个函数，可以生成一个工具去做计算或是加工已有的计算，根据需求进行定制：
```
  func generateNewFunction([parameter list 1]) func([parameter list 2])(return value list 2) {
      // ...
      return func([parameter list 2])(return value list 2) {
          // custom-made function
      }
  }
```

更多使用方法请参考闭包(`closure`) 。

**<u>闭包`closure`</u>** [link](https://github.com/AdaSheng07/ready.to.go/blob/f03ac8deb4c07d421624d5c91d1efccbbf8b95b6/chapter2/011.closure)

闭包是函数与其相关的引用环境组成的实体。一个函数和对其**周围状态（又称上下文）的引用捆绑**在一起，这样的组合成为闭包（`closure`）。闭包可以让我们在一个**内层函数中访问到其外层函数的作用域**。

在操作上，闭包是一种用于保存函数和环境的记录。环境记录关联性的映射，将函数的每个自由变量与创建闭包时所绑定名称的值或引用相关联。环境决定了函数的特殊性与闭包的特性。

分析函数运行时，重点关注的函数本身及其上下文，比如使用的变量、调用的方法、`golang`的值传递等。闭包函数变量在被定义的时候，与哪些变量产生了关联，在闭包方法被调用运行时，闭包方法会回到当初被定义的位置，与原来的环境/周围状态/上下文发生互动，得到执行的最终结果。

```go
    package main
    
    import (
      "fmt"
    )
    
    var times int
    
    // what if we return a function defined in another function body 
    // and using variables outside its scope?
    func calcSumFunc() func(nums ...int) {
        var sum int
        var weightedSum float64
        weight := 0.5
        return func(nums ...int) {
            for _, value := range nums {
                sum += value
            }
            weightedSum = float64(sum) * weight
            times++
            fmt.Println(sum, weightedSum)
        }
    }
    
    func main() { 
        // a function variable is declared to calcWeightedSum, calcWeightedSum is a closure
        // calcWeightedSum keeps its citation to weight, sum and times.
        // weight, sum and times seems to escape, but their life cycles have not ended yet
        calcWeightedSum := calcSumFunc()
        calcWeightedSum(1, 2, 3, 4, 5, 6, 7, 8, 9) // 45 22.5
        calcWeightedSum(1, 2, 3, 4, 5, 6, 7, 8, 9) // 90 45
        calcWeightedSum(1, 2, 3, 4, 5, 6, 7, 8, 9) // 135 67.5
        calcWeightedSum(1, 2, 3, 4, 5, 6, 7, 8, 9) // 180 90
        calcWeightedSum(1, 2, 3, 4, 5, 6, 7, 8, 9) // 225 112.5
        calcWeightedSum(1, 2, 3, 4, 5, 6, 7, 8, 9) // 270 135
        fmt.Println(times)                         // 6
        // here we call calcSumFunc for another five times
        // they return five closures, they cite different sums and weights (even though they have the same values)
        // times still accumulates because it is a global variable
        calcSumFunc()(1, 2, 3, 4, 5, 6, 7, 8, 9) // 45 22.5
        calcSumFunc()(1, 2, 3, 4, 5, 6, 7, 8, 9) // 45 22.5
        calcSumFunc()(1, 2, 3, 4, 5, 6, 7, 8, 9) // 45 22.5
        calcSumFunc()(1, 2, 3, 4, 5, 6, 7, 8, 9) // 45 22.5
        calcSumFunc()(1, 2, 3, 4, 5, 6, 7, 8, 9) // 45 22.5
        fmt.Println(times)                       // 11
    }
```

闭包最主要的意义在于缩小变量的作用域，减少对全局变量的污染，同时可以增加方法的灵活性和自由度。

### 🔸 `Golang`的内置函数

`Go`语言标准库提供了多种可动用的内置的函数。

| built-in functions | applications      | specification                                                             |
|--------------------|-------------------|---------------------------------------------------------------------------|
| `close()`          | 管道关闭              |                                                                           |
| `len()`            | 接受不同类型参数并返回该类型的长度 | 例如字符串、数组、切片、`map`和管道。<br/>如果我们传入的是字符串则返回字符串的长度，如果传入的是数组，则返回数组中包含的元素个数。    |
| `cap()`            | 接受不同类型参数并返回该类型的长度 | 用于返回某个类型的最大容量，例如切片和`map`，根据不同类型，返回意义不同。                                   |
| `new()`            | 内存分配              | 用于值类型和用户定义的类型，如自定义结构，将类型作为参数。                                             |
| `make()`           | 内存分配              | 用于内置引用类型（切片、map 和管道），将类型作为参数。<br/>`new(T)`分配类型`T`的零值并返回其地址，也就是指向类型`T`的指针。 |
| `copy()`           | 操作切片              | 复制切片，`copy(dst []Type, src []Type)`                                       |
| `append()`         | 操作切片              | 编辑切片，增删改，`append(slice, elem1, elem2)`，`append(slice, anotherSlice...)`   |
| `panic()`          | 错误处理              | 用于处理严重错误，使当前运行函数直接异常退出，如果异常退出没有被捕获，则会持续向上层递进，直到有捕获的地方，或`main`函数退出。        |
| `recover()`        | 错误处理              | 用于捕获严重错误。                                                                 |
| `print()`          | 打印                |                                                                           |
| `println()`        | 打印                |                                                                           |
| `complex()`        | 操作复数              |                                                                           |
| `real()`           | 操作复数              |                                                                           |
| `imag()`           | 操作复数              |                                                                           |
| `defer()`          | 注册延迟调用机制          | 常用于关闭文件、重置共享变量等。常与`defer`结合使用。                                            |

**<u>`defer`函数</u>** [link](https://github.com/AdaSheng07/ready.to.go/blob/f03ac8deb4c07d421624d5c91d1efccbbf8b95b6/chapter2/012.defer/main.go)

`defer`是`golang`内置的函数，是Go语言提供的一种用于注册延迟调用的机制，以用来保证一些资源被回收和释放。`defer`注册的延迟调用可以在当前函数执行完毕后执行（包括通过`return`正常结束或者`panic`导致的异常结束），后定义的`defer`先执行。

在定义`defer`函数时，注意它与周围环境有哪些关联关系，与我们使用的方法和作用域严格绑定分析。

> 【重点】`defer`函数的陷阱
>
> ```go
> package main
> import ( 
> "fmt"
> "time"
> )
> // why is duration not 5 seconds?
> func deferGuess() {
> 	startTime := time.Now()
> 	fmt.Println("start time:", startTime)
> 	defer fmt.Println("duration: ", time.Now().Sub(startTime)) // in nanosecond level
> 	time.Sleep(time.Second * 5)                                // 5 seconds
> 	fmt.Println("finish time:", time.Now())
> }
> ```
`defer`注册的函数是逆序执行的，即**先注册后执行**，先注册进内存栈中，得到信号之后从栈内弹出，原则是先入后出。

在本例中，`defer`注册的函数里的部分，即`time.Now().Sub(startTime)`会预先运行完毕（纳秒级别），准备好被打印（运行最后一层）。

> 如何解决这个问题？>>> 利用闭包`closure`
>
> 在defer函数定义时，对外部变量的引用是有两种方式的，分别是作为**函数参数**和作为**闭包**引用：
> - 作为函数参数，则在`defer`定义时就把值传递给`defer`，并被`cache`起来
> - 作为闭包引用的话，则会在`defer`函数真正调用时根据整个上下文确定当前的值
>
> ```
> defer func() {
>   fmt.Println("use closure to calculate duration: ", time.Now().Sub(startTime)) // about 5 seconds
> }()
> ```
>

**<u>`panic`函数</u>** [link](https://github.com/AdaSheng07/ready.to.go/blob/f03ac8deb4c07d421624d5c91d1efccbbf8b95b6/chapter2/012.panic/main.go)

`golang`内置了多种`panic`，如`nil pointer`, `index out of range`, `concurrent read/write map`等。

`panic`也可以主动通过调用`panic`函数抛出。即使程序`panic`，`defer`函数仍然会照常调用。

观察`goroutine`的调用栈（调用链条），从内层到外层，由近及远报错：
```
  panic: assignment to entry in nil map
  
  goroutine 1 [running]:
  main.recoverSample(...)
          /Users/.../gopath/src/learn.go/chapter2/012.recover/recover.go:5
  main.main()
          /Users/.../gopath/src/learn.go/chapter2/012.recover/main.go:4 +0x2e

```

**<u>`recover`函数</u>** [link](https://github.com/AdaSheng07/ready.to.go/blob/f03ac8deb4c07d421624d5c91d1efccbbf8b95b6/chapter2/012.recover)

`recover`用于捕获严重错误，它通常位于`defer`引入的函数体中，用于捕获正在运行的函数中出现的严重错误。

`defer`的`recover`只能出于当前函数的调用栈中，如果脱离了当前函数的调用栈，`recover`将无法捕获：
```go
    package main
    import "fmt"
    
    func main() {
        recoverSample()
    }
    
    func recoverSample() { 
        // adding defer with recover will catch panic in the program
        // with this fragment, program will not drop out automatically
        defer func() {
          if r := recover(); r != nil {
              fmt.Println("fatal error discovered here:", r)
          }
        }()
    
        defer catchPanicUpgraded()
    
        defer catchPanic()
    
        var nameScore map[string]int
        nameScore["lisa"] = 100 // panic: assignment to entry in nil map
    
    }
    // catchPanic: refactor func() to get it
    // in this case, the panic will not be caught, why?
    // this is because when we use catchPanic, the call process of func() has escaped the running of recoverSample
    // they are not in the same stack anymore, therefore panic error cannot be caught
    func catchPanic() {
        func() {
            if r := recover(); r != nil {
                fmt.Println("fatal error discovered:", r)
            }
        }()
    }
    
    // upgrade it: this one will catch panic
    func catchPanicUpgraded() {
        if r := recover(); r != nil {
            fmt.Println("fatal error discovered finally:", r)
        }
    }

```

> 【小结】在编写代码时，将`defer`, `panic`和`recover`相结合使用，可以写出健壮的程序。当工作中出现错误，常用它们解决。善用`debug.PrintStack()`打印调用栈，确定错误位置，在大型项目中非常实用。


### 🔸 函数作为特殊变量 [link](https://github.com/AdaSheng07/ready.to.go/blob/f03ac8deb4c07d421624d5c91d1efccbbf8b95b6/chapter2/009.function3/main.go)

函数除了单独定义外，还可以作为变量使用，变量类型是**方法**。该变量可以作为方法调用来使用。函数变量可以这样定义：
```
    var function_name func([parameter list]) (return value list)
```
函数变量在赋值时有条件：
- 变量类型不能变，函数**只能作为函数**来使用，
- 变量定义后必须使用。普通函数可以单独存在而不必使用，但是**函数变量定义后必须使用**
- 强类型，定义的函数变量的**参数类型和返回值类型必须一致**，才能够进行赋值

**匿名函数**

匿名函数没有函数名，只有函数逻辑体，定义格式为：
```
    func([parameter list])(return_value_list) {
        function body: executive statements
    }
```
> 【应用】什么情况下使用匿名函数？
> 
> 对于只用到一次，不会重复使用的函数，不需要命名，在定义函数之后立即使用；还可以作为回调函数使用：
> 
> ```
> 1. use anonymous functions just after declaration
> 
> func([parameter list]) (return_value_list) {
> function body: executive statements
> }([parameter list])
> 
> 2. use anonymous functions as callback functions
> 
> func function_name_1([parameter list 1], function_name_2 func([parameter list 2])(return_value_list_2))(return_value_list_1){
> function body: executive statements
> }
> func main() {
> function_name_1([parameter list 1], func([parameter list 2])(return_value_list_2){
> function body: executive statements
> })
> }
> ```
>

### 🔸 特殊函数 [link](https://github.com/AdaSheng07/ready.to.go/blob/f03ac8deb4c07d421624d5c91d1efccbbf8b95b6/chapter2/009.function4/main.go)

**`init`函数**

`init`函数是在包被引用时**用于包初始化**的函数。
特殊点在于：
- 函数参数为空
- 不需要、也不可以被调用，由`Golang`默认自动执行
- 一个`Go`文件、包中可以有多个`init`函数，依照定义顺序在`main`主函数运行前自动运行


**Appendix**

☞   [Go 语言向函数传递数组](https://www.runoob.com/go/go-passing-arrays-to-functions.html)  
☞   [Go语言参数传递是传值还是传引用](https://www.flysnow.org/2018/02/24/golang-function-parameters-passed-by-value.html)  
☞   [Go 语言函数](https://www.runoob.com/go/go-functions.html)  
☞   [Go语言将函数作为返回值](http://c.biancheng.net/view/4781.html)  
☞   [回调函数和闭包](https://www.cnblogs.com/f-ck-need-u/p/9878898.html)  
☞   [Go 语言闭包详解](https://www.sulinehk.com/post/golang-closure-details/)  
☞   [golang中的闭包的意义和用法](https://blog.csdn.net/jason_cuijiahui/article/details/84720411)  
☞   [Golang之轻松化解defer的温柔陷阱](https://segmentfault.com/a/1190000018169295)

## 🔶 包 `package`

像之前我们常用的`fmt`与`math`都是 Go 语言中的内置包，它们被称为标准库。`fmt`提供了格式化输入输出功能，`math`提供了基本的数学函数。

包是在同一个目录（或文件夹）下的、总是一起编译、使用的一组源文件的集合。  

为什么要用包？

- 更好地管理项目模块、功能：分工负责处理外部调用/做业务逻辑计算、转换等/存储...
- 实现同名方法、变量等：同名不同包  
  ![image](https://github.com/AdaSheng07/ready.to.go/blob/486e5b1d933ec33a70e5c483a9ccefc05a72e11a/pics/package_1.png)
- 控制访问范围（又称"作用域"）：比如处于整个包中的全局变量，在其它包中还是不可见的

### 🔸 包的可访问性

在同一个包中，所有的函数、变量、常量、对象等都是可见的（visible）。e.g. 在一个文件夹`011.closure`中，定义在`go`文件`closure.go`中的变量、函数等是可以在`main.go`中可见并使用的。在`main.go`中定义的变量、函数等，在同文件夹下的其它`go`文件中可以直接使用。

在不同的的包中，只有公开的（public）函数、变量、常量、对象等是可见的（visible）。在 Golang 中凡是以大写字母开头的，就是公开的，以非大写字母开头的，就是未导出的函数功能`unexported function usage`，非公开/私有的/不可用的（锁）。

根据函数是否可以访问，分为公有（public）和私有（private）两种：
- **公有（public）**：可以被其它包访问的函数、变量。以大写字母开头的函数是公有函数、公有变量。
- **私有（private）**：只能在当前包访问的函数、变量。不是大写字母开头的函数均为私有函数、私有变量。

可以用外接插口来类比理解公有/私有：

PC外接插口有USB口、电源口、HDMI口（显示器口）、SD Card口，而PC内部的插槽可能还有内存插槽、PCIe、SATA、CPU插槽等，对于内部的组装不可见，我们并不感兴趣，我们只需要使用可见的外接插口。

再比如MacBook的外接插口是USB-C，而内部插槽（或焊接位）也有内存插槽、PCIe、CPU插槽（或焊接位）等，对内的组装不可见，对外可见的接口统一为USB Type-C。

### 🔸 包函数的使用

编辑器`GoLand`会自动扫描`module`路径名称下的所有文件、函数和方法等，在使用时会自动提示导入包。如果要使用非本package（包）中的函数，必须：
- 导入要使用的包
  ```
  import "<package name>"
    package name=<module path>+<directory>
  ```
- 调用包内的函数
  ```
  <package name>.<function name>
  ```

### 🔸 包的初始化

`init`函数在包`package`被引用（import）时会被调用，`init`函数可以做包的初始化。比如：全局变量的初始化、注册、预先计算等。

通过使用特殊符号——`_`（下划线），可以做到纯初始化调用。

在 Golang 中，`_`是一个特殊的符号：
- 作为变量时，表示不定义具体对象的变量
- 作为形式参数时，表示一个不存在的形式参数，无法使用
- 作为包引用时，表示这个包不需要被调用，但它是项目的需要，在import时可以用`_`，不至于因为未使用此包在保存go文件时删除这条import  
  ![image](https://github.com/AdaSheng07/ready.to.go/blob/70dc6ec01604a13625abf6d7fb941619a1899b1e/pics/package_2.png)

### 🔸 扩展当前包

有时候为了简化调用，可以将要引用的包作为当前包的扩展包引用进来，从来可以像使用当前包的函数一样使用引用包的函数：

  ```go
    package main
    import (
		"fmt"
		."fmt"
    )
    func main() {
		Println("asdf")
		fmt.Println("1234")
    }
  ```
>【注意】
> - 同一个包中函数名称、变量名称不能重名 
> - 扩展包中的函数不可控

### 🔸 包别名

在引用包的时候，有时会遇到引用的多个包名称相同（哪怕是不同文件夹名，以包名称为准）的情况，这时需要使用别名来进行区分。有时也通过别名来简化引用：
```
import (
    "fmt"
    "learn.go/utils"
    c02 "learn.go/chapter02/utils"
    c03 "learn.go/chapter03/utils"
) 
func main() {
    c02.Cacl()
    c03.Calc()
    fmt.Println(...)
}
```

## 🔶 Go Module

### 🔸 初识 Go Module

Go Module是 Golang 官方提供的依赖管理方案。一个 Go Module 代表一个独立的、可使用的模块。

``` 
  module xxxx   // 定义module的名称
  go 1.17       // 定义golang的版本，不同版本的module行为上略有区别，但核心规则是相同的
```
**初始化Go Module**

Golang 官方提供初始化 Go Module 命令行`$ go mod init[module path]`:

[!image](https://github.com/AdaSheng07/ready.to.go/blob/3e4f5dd099b31633b4473288dc1b2b5c913879b9/pics/module_1.png)

- `[module path]`为可选参数，但当其不是在`$GOPATH/src`文件夹中时，必须指定。`[module path]`会被默认为`module`的名称。
- 在`$GOPATH/src`文件夹中时，可以不指定，默认会以相对于`src`文件夹的路径。
- 生成的`go.mod`文件中不包含 Golang 官方包。

### 🔸 Go Module 实战例题

如果需要利用从命令行传入的数据进行计算，可以使用`os.Args`：
  ```
  var (
      name   string
      sex    string
      height string // float64
      weight string // float64
      age    string // int
  )
  arguments := os.Args
  fmt.Println(arguments)
  ```
在这里，`arguments`是一个切片`slice`，在编辑配置给出数据后，打印的`arguments`如下：
  ```
  [/private/var/folders/7l/pnr12b8962l7dwhrxxzjd9500000gn/T/GoLand/___go_build_learn_go_chapter2_014_bfrCmdCalculator 小强 男 1.7 70 35]
  name:  小强
  sex:  男
  height:  1.7
  weight:  70
  age:  35
  ```
由此可知，切片`arguments`的第一个元素是 go 文件本身，传入的数据索引从`1`开始，而且数据必须严格按照顺序传入，且所有录入数据的格式都是`string`，需要手动转换，这样的程序明显健壮性不足。

如何改进？

我们当然可以写一个说明文档，解释传入数据的方法和要求，但更简洁的方法是参照`git`的命令行语句，用不同的参数执行不同的功能。

**引用第三方包**

在引入第三方包时，如果此包本地从未缓存过，可以先定义一个别名为"`_`"的引用，然后运行`go mod tidy`来让系统初始化引入的第三方的包，然后再使用。

比如，这里我们引入`cobra`：
  ```
  import (
      _ "github.com/spf13/cobra"
  )
  ```
再在`Terminal`的`go.mod`文件路径下执行`go mod tidy`完成配置：
  ```
  $ go mod tidy
  go: finding module for package github.com/spf13/cobra
  go: downloading github.com/spf13/cobra v1.3.0
  go: found github.com/spf13/cobra in github.com/spf13/cobra v1.3.0
  go: downloading github.com/inconshreveable/mousetrap v1.0.0
  go: downloading github.com/spf13/pflag v1.0.5
  ```

> 【注意】这里`github.com/spf13/cobra`是第三方包的路径，`cobra`是目录下的文件名，并不是包的名称。`control/command` + 左键点击可以`goto`自动跳转至文件目录下进行查看项目正在依赖的包，打开其中的 go 文件可以查看包的名称，大部分开源社区的 module 文件名与包名都是一致的。
>
> ![image](https://github.com/AdaSheng07/ready.to.go/blob/ab8d28ec0c11cd5b153f96f0baba3f7e133aed07/pics/package_3.png)

**使用第三方依赖包**

利用`cobra`包中的`Command`功能传入计算，替代原代码中的`os.Args`。定义其中需要的变量，以及命令行参数需要的命令行对象：
  ```
  func main() {
      var (
          name   string
          sex    string
          height float64
          weight float64
          age    int
      )
      cmd := cobra.Command{
          Use:   "bfrCheck",    // use: set a name for this command
          Short: "基于BMI的体脂计算器", // short: a short description for this command
          // long: a detailed explanation for this command
          Long: "录入姓名、性别、身高、体重和年龄，计算他们的BMI值，基于他们的性别和年龄生成他们的体脂率标准，判断他们的体脂率处于偏瘦/标准/偏胖/严重肥胖并给出健康建议。",
          // func is a registered callback function, Run is the main body of cmd, custom-made as desired
          Run: func(cmd *cobra.Command, args []string) {
              fmt.Println("name: ", name)
              fmt.Println("sex: ", sex)
              fmt.Println("height: ", height)
              fmt.Println("weight: ", weight)
              fmt.Println("age: ", age)
              // calculate bmi & bfr...
              
              // healthiness assessment & suggestions
              
          },
      }
  ```
利用`Flags()`配置需要的各参数：
  ```
      /*
          func (f *FlagSet) StringVar(p *string, name string, value string, usage string)
          StringVar defines a string flag with specified name, default value, and usage string.
          The argument p points to a string variable in which to store the value of the flag.
      */
      cmd.Flags().StringVar(&name, "name", "", "姓名")
      // this means: when we type in command line, what comes after "name" will be saved in variable name as string.
      // if we give nothing after "name", the variable name will be the default value "".
      cmd.Flags().StringVar(&sex, "sex", "", "性别")
      cmd.Flags().Float64Var(&height, "height", 0, "身高")
      cmd.Flags().Float64Var(&weight, "weight", 0, "体重")
      cmd.Flags().IntVar(&age, "age", 0, "年龄")
  ```
在`main()`增加语句运行已定义的命令行对象：
  ```
  cmd.Execute()
  ```
在`Terminal`运行程序得到：
  ```
  $ go run ./main.go
  name:  
  sex:  
  height:  0
  weight:  0
  age:  0
  ```
通过命令行的帮助选项查看配置与使用方法：
  ```
  $ go run ./main.go --help
  录入姓名、性别、身高、体重和年龄，计算他们的BMI值，基于他们的性别和年龄生成他们的体脂率标准，判断他们的体脂率处于偏瘦/标准/偏胖/严重肥胖并给出健康建议。
  
  Usage:
    bfrCheck [flags]
  
  Flags:
        --age int        年龄
        --height float   身高
    -h, --help           help for bfrCheck
        --name string    姓名
        --sex string     性别
        --weight float   体重
  ```
通过命令行对象，我们可以实现乱序输入，并且数据格式会自动转换。实际使用方法为：
  ```
  $ go run ./main.go --age 35 --name 小强 --weight 70 --height 1.70 --sex 男
  name:  小强
  sex:  男
  height:  1.7
  weight:  70
  age:  35
  ```
在命令行对象的`Run`的主体函数中，我们也可以调用其它包的函数完善其它功能。比如：
  ```
  // calculate bmi & bfr...
  bmi := calc.CalcBMI(height, weight)
  bfr := calc.CalcBFRUpgrade(bmi, age, sex)
  fmt.Println("Body Fat Rate:", bfr)
  ```
在`Terminal`中运行后得到的结果：
  ```
  $ go run ./main.go --age 35 --name 小强 --weight 65 --height 1.70 --sex 男
  name:  小强
  sex:  男
  height:  1.7
  weight:  65
  age:  35
  Body Fat Rate: 0.18489619377162636
  ```

> 【总结】在使用第三方依赖包时，先定位需要使用的某个工具，导入（巧用"`_`"技巧）后`go mod tidy`拉取需要的缺少的模块，移除不用的模块。
> ```
> Usage:
>
>    go mod <command> [arguments]
>
> The commands are:
>
>    download    download modules to local cache
>    edit        edit go.mod from tools or scripts
>    graph       print module requirement graph
>    init        initialize new module in current directory
>    tidy        add missing and remove unused modules
>    vendor      make vendored copy of dependencies
>    verify      verify dependencies have expected content
>    why         explain why packages or modules are needed
>
> Use "go help mod <command>" for more information about a command.
> ```

### 🔸 Go Module 进阶

**Module 的使用**

- 像普通的包一样直接使用函数
- 编写代码时，会遇到找不到`package`、`function`的错误，一些较老的项目无法使用 Go Module，这时需要运行：
  ```
  go get <package> // as updated in golang 1.17
  go mod tidy
  ```
- 像正常一样使用代码
- 如果在 VS Code 中打开项目，注意`go.mod`文件务必位于 VS Code 锁打开文件夹的位置（golang 1.17），GoLand 中不存在这个问题
- 运行`go mod tidy`来保证项目中使用到的 go modules 被包含在项目依赖中

**Go Module 替换（replace）**

Go Module 替换（replace）是用另一个实现替换默认要使用的实现，类似于"狸猫换太子"。一般使用的场景有：
- 默认`go mod tidy`出来的都是最新版本，如果最新版本不合适，或不能使用等，需要将其替换为相对旧一些的版本。比如：
  ```
  replace github.com/spf13/cobra => github.com/spf13/cobra v1.2.0
  ```
- upstream代码（上游代码库）不太符合我们的需求，我们可以在此基础上做自身的定制，并开源在`git`服务上，需要将`module`重定向。在 Golang 中永远都是从源码出发的，可以直接从实现出发来使用。
- 本地有些项目的代码，没有`git`支持，需要作为独立的`module`使用。比如：
  ```
  replace learn.go.tools => ../learn.go.tools
  ```
> 【注意】本地项目代码作为`module`替换：
> 1. 替换的是`module`，寻找的是替换后的`go.mod`所在的路径位置**相对**于原来的`go.mod`所在文件夹的位置。
> 2. 可以通过`Terminal`验证`replace`时的重定向路径是否存在和正确，保证`module`能够成功`replace`。
> 3. `replace`之后的版本一定被使用。

> 【常见问题】使用`go mod`时报错`i/o timeout`
>
>      解决办法是在执行前在`Terminal`中设置环境变量：
>   ```
>   $ export GO111MODULE=on
>   $ export GOPROXY=https://goproxy.io 
>   或
>   $ export GO111MODULE=on
>   $ export GOPROXY=https://goproxy.cn
> 
>   对于Golang 1.13及以上版本，还可以：
>   $ go env -w GOPROXY=https://goproxy.cn
>   ```

### 🔸 Vendor

Vendor 是把项目`module`定义的所有依赖制作一个副本保存在项目的`vendor`目录。可以理解为：将项目定义的依赖做一个**快照**并保存下来，避免项目依赖的变更影响项目的一致性。Vendor 是几乎所有大型项目中都会使用的，稳定性强，保证项目的一致性。

Go Module 深度支持`vendor`，通过在`go.mod`所在目录下运行命令行`go mod vendor`将项目的依赖保存到`vendor`中：
  ```
  $ go mod vendor
  $ go mod vendor -v
  # github.com/inconshreveable/mousetrap v1.0.0
  ## explicit
  github.com/inconshreveable/mousetrap
  # github.com/spf13/cobra v1.3.0
  ## explicit; go 1.15
  github.com/spf13/cobra
  # github.com/spf13/pflag v1.0.5
  ## explicit; go 1.12
  github.com/spf13/pflag
  ```
Go Module 的定义与 `vendor`本身就是**共存**的，通过`go build/run`的命令行选项来控制`-mod=<vendor|mod|readonly>`：
- `vendor`——使用`vendor`中的依赖编译项目
- `mod`——使用`go module`定义的依赖编译项目，并且会自动更新`go.mod`定义
- `readonly`——使用`go module`定义的依赖编译项目，并且不做任何依赖的升级

> 【注意】`vendor`目录存在/不存在：
> 1. 项目中如果有`vendor`目录，在编译时默认使用`vendor`提供的依赖，如果在`Terminal`中运行了`go mod tidy`后，必须要再运行`go mod vendor`才能使依赖生效
> 2. 如果没有`vendor`目录，在编译时默认使用`readonly`
> 3. 在开发时最好先将目录中的`vendor`目录删除，在最后提交代码前再对`go.mod`文件执行`go mod vendor`

## 🔶 函数方法论

### 🔸 整洁代码的价值

### 🔸 方法、函数内部

### 🔸 方法、函数之间

## 🔵 Module 2 Practice Collection

### 🔹 函数

> **Q1** 在同一个 Go 文件里能创建同名的函数吗？如果能，会发生什么问题？如果不能，怎么在同一个大项目中创建相同名字的函数？能否用一个文件包含整个项目？
> 
>> 除了`init`函数，其他函数均不能同名，不能在同一个文件或包中定义同名函数。  
>> 从理论上，可以用一个文件夹包含整个项目，但非常困难。如果在一个包中包含所有的项目文件，在做检测或关注某个功能时，会迷失在每个细节中。

> **Q2** 利用递归实现一个汉诺塔问题的求解。

> **Q3**  从 1~100 中心里想一个数字，然后让程序去猜。程序问：是xx吗？你只能回答高了/低了/对了。若没有猜中，程序继续猜，直到猜中为止。 
> <br>方法 1：逐个数字猜 
> <br>方法 2：每次排除一半的数字
> 
>      [>>  How to Guess a Number?](https://github.com/AdaSheng07/ready.to.go/blob/53b015ce07ec17420d0ef971174a67fb5df70657/chapter2/010.iteration2/main.go)


### 🔹 包

> **Q1** 我们已经学会引用我们自己的 module 中的包、方法，怎么使用其他人开发的，或者说开源社区
的 module 呢？

### 🔹 Go Module

> **Q1** Go Module 实战：让体脂计算器接受从命令行传入的姓名、性别、身高、体重、年龄，直接计算其体脂并给出结果。

> **Q2** Go Module 管理依赖已经很完善，是不是只要使用了 Go Module 就万无一失？
>> 不是万无一失：
>> - Go Module 完全依赖`git`服务，与`git`上的`branch`、`release tag`紧密相关。`Git`服务不属于 Go Module 的管理范畴。
>> - `Git`上的`branch`、`release tag`可以随时被删除、重写
>> - Go Module在依赖不改变时是可靠的，但在依赖改变时就不再可靠了
>>
>> 对项目的影响：
>> - 项目可能因为`git`服务的变更而无法编译
>> - 编译出来的程序行为不一致
> 
> 如何用一种更可靠、稳定的方式来管理依赖项目呢？如何保证不同使用者编译、运行相同程序的行为一模一样？
> -> 引入`Vendor`

> **Q3** 使用GitHub上的公有函数来实现功能，重写BMI和体脂计算器，用`module replace`替换GitHub上的，并使用`vendor`为项目依赖提供保障。

****

# 🟡 <u>Module 3</u>

## 🔶 异常处理

异常`Exception`是程序没有按照预期运行的统称，可能是因为输入错误，也有可能是程序本身设计上的缺陷、代码上的漏洞导致的。

常见的异常有：
- 除0
- 索引`index`错误：切片、数组等的更新编辑操作
- 溢出：e.g. 永久调用的递归，内存溢出
- 空指针：e.g. `*int`未赋值；未实例化的`map(nil)`
- 类型转换失败：
    ```go
      package main
      
      import "fmt"
      
      func main() {
          convertTypes()
      }
      
      func convertTypes() {
          var a interface{}
          a = "convert a"
          b := a.(int)
          fmt.Println(b)
      }
      
      // panic: interface conversion: interface {} is string, not int
    ```
- 并发读写`bug`
- 管道相关  
...

### 🔸 处理异常

**1. 暴露错误**

   在编程时，要及时暴露错误，而不是忽略它们。错误可能的来源有：
   - 主动暴露的错误：`panic`,`error`
   - 调用其它函数（非当前函数）返回的错误：这种错误是可预测的，需要尽快暴露错误并结束当前函数运行，而非忽略它们。

**2. 处理错误**

   当调用其它函数（非本函数）时，如果返回有错误`error`，务必处理错误，不可轻易忽略。常见做法是返回错误信息：
  ```go
  package main
  
  import "fmt"
  
  func main() {
      gender, weight, err := inputInfo()
      fmt.Printf("gender: %s\n", gender)
      fmt.Printf("weight: %.2f\n", weight)
      fmt.Println(err)
  }
  
  func inputInfo() (sex string, weight float64, err error) {
      sex = ""
      weight = 0
      fmt.Print("input your gender: ")
      _, _ = fmt.Scanln(&sex)
      if sex != "male" && sex != "female" {
          err = fmt.Errorf("gender: %s is neither male nor female", sex)
          return
      }
      fmt.Print("input your weight: ")
      _, _ = fmt.Scanln(&weight)
      if weight <= 10 || weight > 200 {
          err = fmt.Errorf("weight: %.2f is out of range", weight)
          return
      }
      return
  }
  ```

**3. 抓住错误**

在程序运行时，有一类错误是无法预测的错误，它们不是主动返回的`error`，而是直接以`panic`的当时出现，直接熬制应用程序崩溃，尤其是在使用了不可控的第三方函数，或使用没有进行完整测试、校验的代码。这种错误**必须**被抓组，避免整个程序崩溃退出。  

像除0、索引`index`错误、溢出、空指针、类型转换失败、并发读写、管道操作错误等都属于这种类型的错误。  

抓住这种错误的方式是在函数题重新定义`defer`方法，并在方法中使用`recover`来捕获这类异常，比如：
  ```
  defer func() {
      r := recover()
          if r != nil {
          // detection: handle exception information
      }
  }
  ```
抓住错误意味着本次程序运行出现一些问题，比如输入的数据不合法/不合要求等，并不代表程序必须要崩溃退出。

### 🔸 如何写出健壮的代码？

**在编写代码时，任何担心出现的问题必定会出现（墨菲定律）。**

为了提高编写代码的质量，需要精确控制每段代码的行为，并且做好足够的安全性保障：
- 控制定定义域
- 空指针检测
- 使用`for-range`遍历
- 谨慎使用共享变量
- 控制并抓住异常


## 🔶 单步调试 `debug`

单步调试是指程序开发中，为了找到程序的`bug`，通常采用的一种调试手段。即：一步一步跟踪程序执行的流程，根据变量的值，找到错误的原因。

### 🔸 断点

断点是单步调试过程中添加的特定的检查点，在 GoLand 的 Bookmarks 中查看。在单步调试时，会直接运行到断点并停下。

如果没有断点，程序则会正常运行，直至遇到阻塞才会停下，否则直接运行直至退出。

### 🔸 `main()`函数单步调试练习

## 🔶 单元测试

单元测试是指在计算机编程中针对一块特定的模块、组建、方法进行测试以验证其是否满足业务、质量需求的测试方法。

### 🔸 如何编写 Golang 单元测试？

单元测试的关键组成部分：
- 预备案例
- 预期结果
- 组件调用
- 衡量预期

Golang 针对单元测试，有两种测试方法： 
1. 黑盒测试  
测试案例与源码在同一个目录下，但是不在同一个包中。测试案例仅针对目标包里暴露出的公有方法进行测试。
2. 白盒测试  
测试案例与源码在同一个目录、同一个包中。测试案例可以看到包中的私有变量、方法，可以针对每一项进行测试。

Golang 的单元测试编写规则：
- 必须包含在以`_test.go`结尾的文件中
- 必须符合命名规则`func TestXxxxx(t *testing.T){...}`或者`func Test_Xxxxx(t *testing.T){...}`
- 对于预期失败的`case`，需要调用`t.Fail()`, `t.FailNow()`, `t.Error(...)`, `t.Errorf()`, `t.Fatal(...)`, `t.FatalIf(...)`...等来声明失败
  ```go
  package main
  import (
  
  "math"
  "testing"
  )
  func TestMain(t *testing.T) {
      var expectedBMI float64 = 18.9
      actualOutput := calcBMI(1.70, 69)
      if expectedBMI != actualOutput {
          // ...
          t.Fail()
      }
  }
  func calcBMI(height, weight float64) (bmi float64){
      bmi = weight / math.Pow(height, 2)
      return
  }
  ```
  在以上声明语句中：
  - `t.Fail()`和 `t.FailNow()`标记`test case`已经失败，程序继续往下执行，但没有任何输出
  - `t.Error(...)`和 `t.Errorf()`不仅标记`test case`是失败的，而且会抛出错误信息，程序会继续往下执行
  - `t.Fatal(...)`和`t.FatalIf(...)`在`test case`失败后直接结束程序运行，跳过之后的所有代码，相当于调用了`Exit()`，结束单元测试

### 🔸 如何运行 Golang 单元测试？

单元测试可以通过命令行运行，也可通过 VS Code、GoLand 运行。同样地，也可以对单元测试进行单步调试`debug`。

命令行运行：
```
go test <package path>[/...][-run<UT function name regex rule>]

e.g.
go test <package path>
go test <package path>/...
go test <package path>/...-run^TestMain$
```
> 【注意】  
> ![image](https://github.com/AdaSheng07/ready.to.go/blob/29f947b1983e92ed97e952ce3795a3ef5c96d788/pics/unit_test_1.png)
> 
> 当在此文件目录下运行单元测试`calc_bmi_test.go`时，会报错：
> ```
> $ go test ./calculator/calc_bmi_test.go
> # command-line-arguments [command-line-arguments.test]
> calculator/calc_bmi_test.go:12:24: undefined: CalcBMI
> FAIL    command-line-arguments [build failed]
> FAIL
> ```
> 这是因为在`calc_bmi_test.go`中我们使用了`CalcBMI`这个同包不同文件的公有函数，`CalcBMI`并没有在`calc_bmi_test.go`文件中定义。解决办法是在**整个文件夹目录下运行**单元测试，如：
> ```
> $ go test ./calculator
> ok      learn.go/chapter3/016.unitTest/calculator       1.724s
> ```
> 或者，在 IDE 中直接运行也可以看到最终结果：  
> ![image](https://github.com/AdaSheng07/ready.to.go/blob/53fb37d8a2b8051a7d12d39f95d4393ba259c161/pics/unit_test_2.png)

在 VS Code 中运行单元测试 //todo lecture 8 02：58：06




## 🔵 Module 3 Practice Collection

### 🔹 异常处理

### 🔹 单步调试

> **Q1** 单步调试这么强大，是不是任何问题都可以通过单步调试来完成？除了单步调试之外，还有什么方式能够辅助发现异常点？
>> 
>>

> **Q2**
> 

> **Q3**  
>
>      [>>  ]()


### 🔹 单元测试

> **Q1** 单元测试是必须的吗？怎么保证编写的代码符合预期？是先有主代码，再有测试案例；还是先有测试案例，后有主代码？单元测试会拖慢整个项目的进度吗

### 🔹 


****














