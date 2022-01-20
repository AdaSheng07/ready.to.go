## 函数

函数是基本的代码块，用于执行一个任务，以达成预期目的。

在编程中，我们使用不同的函数，划分它们各自的功能，来完成不同的任务。每个程序必须有命名为`main()`的主函数。

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

**定长参数与不定长参数**

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

**函数调用与传递参数的方式**

默认情况下，Go 语言使用的是**值传递**，即在调用过程中不会影响到实际参数。

函数使用的参数可称为函数的形式参数，可理解为**预期投入**，定义在函数体内的局部变量。

调用参数，可以通过两种方式来传递参数：
- **值传递**：在调用函数时将实际参数复制一份传递到函数中，这样在函数中如果对参数进行修改，将不会影响到实际参数。
- **引用传递**：在调用函数时将实际参数的地址传递到函数中，那么在函数中对参数所进行的修改，将影响到实际参数。

一个函数的参数用什么方式传递，要看这个函数传递的参数是什么类型：
- `array`, `map`和`interface`等本身就是引用，在作为参数传递时，拷贝的是变量本身，变量变更时实际参数也跟着变更
- `int`, `string`是值传递，传递的是该参数的一个副本
- 传递一个指针类型的参数，其实传递的是这个该指针的一份拷贝，而不是这个指针指向的值

**函数的递归** [link](https://github.com/AdaSheng07/ready.to.go/blob/f03ac8deb4c07d421624d5c91d1efccbbf8b95b6/chapter2/010.iteration1/main.go)

如果一个函数的调用链中存在自己调用自己，则将这种调用方式称为**递归**。以斐波那契数列计算核心为例：
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

**命名返回值**

在编写代码时，强烈建议命名返回值：
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