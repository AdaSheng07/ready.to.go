## 包 `package`

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