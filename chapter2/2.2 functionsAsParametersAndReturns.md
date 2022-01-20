## 将函数作为参数与返回值 [link](https://github.com/AdaSheng07/ready.to.go/blob/f03ac8deb4c07d421624d5c91d1efccbbf8b95b6/chapter2/009.function2/main.go)

### 🔸 提取函数并重构
- 选取函数片段 -> Refactor -> Extract Method...
- 函数重命名：Refactor -> Rename

### 🔸 使用函数

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

### 🔸 将函数作为返回值（方法）

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

### 🔸 闭包`closure` [link](https://github.com/AdaSheng07/ready.to.go/blob/f03ac8deb4c07d421624d5c91d1efccbbf8b95b6/chapter2/011.closure)

闭包是函数与其相关的引用环境组成的实体。一个函数和对其**周围状态（又称上下文）的引用捆绑**在一起，这样的组合成为闭包（`closure`）。闭包可以让我们在一个**内层函数中访问到其外层函数的作用域**。

在操作上，闭包是一种用于保存函数和环境的记录。环境记录关联性的映射，将函数的每个自由变量与创建闭包时所绑定名称的值或引用相关联。环境决定了函数的特殊性与闭包的特性。

分析函数运行时，重点关注的函数本身及其上下文，比如使用的变量、调用的方法、`golang`的值传递等。闭包函数变量在被定义的时候，与哪些变量产生了关联，在闭包方法被调用运行时，闭包方法会回到当初被定义的位置，与原来的环境/周围状态/上下文发生互动，得到执行的最终结果。

闭包最主要的意义在于**缩小变量的作用域**，减少对全局变量的污染，同时可以增加方法的灵活性和自由度。

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