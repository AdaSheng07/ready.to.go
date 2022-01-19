## 异常处理

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