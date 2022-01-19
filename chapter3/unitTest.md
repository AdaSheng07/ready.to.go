## 单元测试

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
- 对于预期失败的`case`，需要调用`t.Fail()`, `t.FailNow()`, `t.Error(...)`, `t.Errorf()`, `t.Fatal(...)`, `t.FatalIf(...)`...等来声明失败，在测试时如果代码出现错误或者不符合预期，一定要`fail fast`，尽快结束运行跳出。
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
> $ go test ./...
> 或者
> $ go test ./calculator
> 
> ok      learn.go/chapter3/016.unitTest/calculator       1.724s
> ```
> 或者，在 IDE 中直接运行也可以看到最终结果：  
> ![image](https://github.com/AdaSheng07/ready.to.go/blob/53fb37d8a2b8051a7d12d39f95d4393ba259c161/pics/unit_test_2.png)

在 VS Code 中运行单元测试 //todo lecture 8 02：58：06