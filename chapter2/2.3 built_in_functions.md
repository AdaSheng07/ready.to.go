## `Golang`的内置函数

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

### 🔸 `defer`函数 [link](https://github.com/AdaSheng07/ready.to.go/blob/f03ac8deb4c07d421624d5c91d1efccbbf8b95b6/chapter2/012.defer/main.go)

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

### 🔸 `panic`函数 [link](https://github.com/AdaSheng07/ready.to.go/blob/f03ac8deb4c07d421624d5c91d1efccbbf8b95b6/chapter2/012.panic/main.go)

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

### 🔸 `recover`函数 [link](https://github.com/AdaSheng07/ready.to.go/blob/f03ac8deb4c07d421624d5c91d1efccbbf8b95b6/chapter2/012.recover)

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