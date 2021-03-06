# Practice Summary 2

基于BMI的体脂计算器 [>> link](https://github.com/AdaSheng07/ready.to.go/blob/8a229c059febb49f2860c2587e7205562ca61b7f/000.homework/0003.bfrCalculator/main.go)

本地拓展后的`gobmi`以及单元测试：[>> link](https://github.com/AdaSheng07/ready.to.go/blob/8a229c059febb49f2860c2587e7205562ca61b7f/staging/src/github.com/armstrongli/go-bmi)

## Task 1: Go Module 的`replace`、`vendor`以及本地扩展

我们需要使用 GitHub 上的[`lib`: github.com/armstrongli/go-bmi](https://github.com/armstrongli/go-bmi) 完成体脂计算器：
1. 在`main.go`的`import`中输入`_"github.com/armstrongli/go-bmi"`并在文件所处`path`主目录下的`go.mod`文件中运行`go mod tidy`和`go mod vendor`，`go.mod`文件更新后：
   ```
   module learn.go
   
   go 1.17
   
   require github.com/armstrongli/go-bmi v0.0.1
   ```
2. 在主目录下新建`directory`命名为`staging`，模仿`"github.com/armstrongli/go-bmi"`的格式建立`../staging/src/github.com/armstrongli/go-bmi`的文件路径，并把我们所需的`bmi.go`的代码拷贝进去：

   ![image](https://github.com/AdaSheng07/ready.to.go/blob/a6e9c0f3f82528562156da10d577a6ce61ca6a9d/pics/homework2_1.png)

3. 利用扩展后的`staging`下目录去`replace`GitHub 上的`lib`，实现根据需求的定制化，运行`go mod tidy`和`go mod vendor`更新`go.mod`文件:
   ```
   // replace github.com/armstrongli/go-bmi with what we implement locally
   replace github.com/armstrongli/go-bmi => ./staging/src/github.com/armstrongli/go-bmi
   ```
完整的`go.mod`文件 [>> go.mod](https://github.com/AdaSheng07/ready.to.go/blob/5f9d67f8344535e2f8bc466ac32ce0486391e60d/go.mod)
```
module learn.go

go 1.17

require (
	github.com/armstrongli/go-bmi v0.0.1
	github.com/shopspring/decimal v1.3.1
	github.com/spf13/cobra v1.3.0
)

require (
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
)

// replace local module
// replace learn.go.tools => ../learn.go.tools

// replace github.com/armstrongli/go-bmi with what we implement locally
replace github.com/armstrongli/go-bmi => ./staging/src/github.com/armstrongli/go-bmi

```

### 最终实现

利用`cobra`命令行运行，依赖 GitHub 上的 [`lib`](https://github.com/armstrongli/go-bmi) 以及在此基础上本地扩展后的依赖实现体脂计算器

[>> Body Fat Rate Calculator](https://github.com/AdaSheng07/ready.to.go/blob/8a229c059febb49f2860c2587e7205562ca61b7f/000.homework/0003.bfrCalculator/main.go) 
```go
    package main
    
    import (
        "fmt"
        gobmi "github.com/armstrongli/go-bmi"
        "github.com/spf13/cobra"
    )
    
    func main() {
        var (
            name   string
            gender string
            height float64
            weight float64
            age    int
        )
        cmd := cobra.Command{
            Use:   "CheckHealthyStatus",
            Short: "A body fat rate calculator based on BMI value.",
            Long: "Given a person's name, gender, height, weight and age, calculate this person's BMI and BFR. \n" +
                "Based on this person's gender and age, set up the BFR benchmark for this person. \n" +
                "Check for this peron's healthy status and give out specialized suggestions.",
            Run: func(cmd *cobra.Command, args []string) {
                fmt.Println("Name:", name)
                fmt.Println("Gender(male/female):", gender)
                fmt.Println("Height(m):", height)
                fmt.Println("Weight(kg):", weight)
                fmt.Println("Age:", age)
    
                // calculate bmi & calculator and print them out
                bmi, _ := gobmi.CalcBMI(weight, height)
                bfr, _ := gobmi.CalcBFR(bmi, age, gender)
    
                // healthiness assessment & suggestions and print them out
                _ = gobmi.GiveOutSuggestions(bfr, gender, age)
            },
        }
        cmd.Flags().StringVar(&name, "name", "", "input name")
        cmd.Flags().StringVar(&gender, "gender", "", "input gender")
        cmd.Flags().Float64Var(&height, "height", 0, "input height")
        cmd.Flags().Float64Var(&weight, "weight", 0, "input weight")
        cmd.Flags().IntVar(&age, "age", 0, "input age")
    
        _ = cmd.Execute()
    }
```
在`staging`下目录本地拓展后实现的几个模块分别有 [>> go-bmi 本地拓展](https://github.com/AdaSheng07/ready.to.go/blob/8a229c059febb49f2860c2587e7205562ca61b7f/staging/src/github.com/armstrongli/go-bmi)
- 性别权重计算 `get_generWeight.go` [>> link](https://github.com/AdaSheng07/ready.to.go/blob/8a229c059febb49f2860c2587e7205562ca61b7f/staging/src/github.com/armstrongli/go-bmi/get_genderWeight.go)
- BMI值的计算 `calc_bmi.go` [>> link](https://github.com/AdaSheng07/ready.to.go/blob/8a229c059febb49f2860c2587e7205562ca61b7f/staging/src/github.com/armstrongli/go-bmi/calc_bmi.go)
- 体脂率的计算 `calc_bfr.go` [>> link](https://github.com/AdaSheng07/ready.to.go/blob/8a229c059febb49f2860c2587e7205562ca61b7f/staging/src/github.com/armstrongli/go-bmi/calc_bfr.go)
- 根据性别和年龄确定评判标准 `get_benchmarkParameters.go` [>> link](https://github.com/AdaSheng07/ready.to.go/blob/8a229c059febb49f2860c2587e7205562ca61b7f/staging/src/github.com/armstrongli/go-bmi/get_benchmarkParameters.go)
- 根据体脂率和评判标准给出建议 `give_suggestions.go` [>> link](https://github.com/AdaSheng07/ready.to.go/blob/8a229c059febb49f2860c2587e7205562ca61b7f/staging/src/github.com/armstrongli/go-bmi/give_suggestions.go)

在本地`Terminal`运行后的结果为：
```
   $ go run ./main.go --name 小强 --gender male --age 35 --height 1.70 --weight 65
   Name: 小强
   Gender(male/female): male
   Height(m): 1.7
   Weight(kg): 65
   Age: 35
   BMI: 22.49%
   BFR: 18.84%
   目前的体脂率水平处于【偏胖】，建议：少坐多动，注意健康饮食
```

## Task 2: 为体脂计算器编写单元测试并完善体脂计算器的验证逻辑

#### BMI值计算：

- 录入0或负数身高，返回错误
- 录入0或负数体重，返回错误
```go
   package gobmi
   
   import "fmt"
   
   func CalcBMI(weightKG, heightM float64) (bmi float64, err error) {
       if weightKG <= 0 {
           err = fmt.Errorf("weight cannot be negative")
           return
       }
       if heightM < 0 {
           err = fmt.Errorf("height cannot be negative")
           return
       }
       if heightM == 0 {
           err = fmt.Errorf("height cannot be 0")
           return
       }
       bmi = weightKG / (heightM * heightM)
       return
   }
```
- 编写单元测试，确保在录入正常身高、体重时，计算结果符合预期 

   [>> 计算 BMI 的单元测试](https://github.com/AdaSheng07/ready.to.go/blob/8a229c059febb49f2860c2587e7205562ca61b7f/staging/src/github.com/armstrongli/go-bmi/calc_bmi_test.go)

```go
   package gobmi
   
   import (
       "github.com/shopspring/decimal"
       "testing"
   )
   
   func TestCalcBMI(t *testing.T) {
       inputHeight, inputWeight := 1.70, 65.0
       expectedOutput := 22.49
       t.Logf("start calculating bmi: \nheight: %.2f \nweight: %.2f \nexpecting bmi = %.2f.\n", inputHeight, inputWeight, expectedOutput)
       actualOutput, err := CalcBMI(inputWeight, inputHeight)
       if err != nil {
           t.Fatalf("major error emerges while calculating bmi: %v\n", err)
       }
       actualFormatOutput, _ := decimal.NewFromFloat(actualOutput).RoundFloor(2).Float64()
       if expectedOutput != actualFormatOutput {
           t.Logf("expecting %f, but got %f.\n", expectedOutput, actualFormatOutput)
           t.Failed()
       }
   }

```
单元测试的运行结果为：
```
   /usr/local/go/bin/go tool test2json -t /private/var/folders/7l/pnr12b8962l7dwhrxxzjd9500000gn/T/GoLand/___TestCalcBMI_in_github_com_armstrongli_go_bmi.test -test.v -test.paniconexit0 -test.run ^\QTestCalcBMI\E$
   === RUN   TestCalcBMI
       calc_bmi_test.go:11: start calculating bmi: 
           height: 1.70 
           weight: 65.00 
           expecting bmi = 22.49.
   --- PASS: TestCalcBMI (0.00s)
   PASS
   
   Process finished with the exit code 0
```

#### 体脂率计算：

- 录入非法BMI、年龄、性别（0、负数、超过150的年龄、非男女的性别输入），返回错误

```go
   package gobmi
   
   import "fmt"
   
   func CalcBFR(bmi float64, age int, gender string) (bfr float64, err error) {
       if bmi <= 0 {
           err = fmt.Errorf("bmi must be positive")
           return
       }
       if age <= 0 {
           err = fmt.Errorf("age must be a positive integer")
           return
       }
       if age > 150 {
           err = fmt.Errorf("age must be reasonable")
           return
       }
       if gender != "male" && gender != "female" {
           err = fmt.Errorf("gender must be male or female")
           return
       }
   
       bfr = (1.2*bmi + 0.23*float64(age) - 5.4 - 10.8*getGenderWeight(gender)) / 100
       fmt.Printf("BMI: %.2f%%\nBFR: %.2f%%\n", bmi, bfr*100)
       return
   }
```
- 编写单元测试，确保在录入正常BMI、年龄、性别时，计算结果符合预期  

  [>> 计算 Body Fat Rate 的单元测试](https://github.com/AdaSheng07/ready.to.go/blob/8a229c059febb49f2860c2587e7205562ca61b7f/staging/src/github.com/armstrongli/go-bmi/calc_bfr_test.go)

```go
   package gobmi
   
   import (
       "github.com/shopspring/decimal"
       "testing"
   )
   
   func TestCalcBFR(t *testing.T) {
       inputBmi := 22.49
       inputAge := 35
       inputGender := "male"
       expectedOutput := 0.1883
       t.Logf("start calculating bfr(body fat rate): \ngender: %s \nage: %d \nbmi: %.2f \nexpectedOutput: %.4f ", inputGender, inputAge, inputBmi, expectedOutput)
       actualOutput, err := CalcBFR(inputBmi, inputAge, inputGender)
       if err != nil {
           t.Fatalf("major error while calculating bfr: %v\n", err)
       }
       actualFormatOutput, _ := decimal.NewFromFloat(actualOutput).RoundFloor(4).Float64()
       if actualFormatOutput != expectedOutput {
           t.Logf("expecting %f, but got %f.\n", expectedOutput, actualFormatOutput)
           t.Failed()
   
       }
   
   }

```
单元测试的运行结果为：
```
   /usr/local/go/bin/go tool test2json -t /private/var/folders/7l/pnr12b8962l7dwhrxxzjd9500000gn/T/GoLand
   /___TestCalcBFR_in_github_com_armstrongli_go_bmi.test -test.v -test.paniconexit0 -test.run ^\QTestCalcBFR\E$
   === RUN   TestCalcBFR
       calc_bfr_test.go:13: start calculating bfr(body fat rate): 
           gender: male 
           age: 35 
           bmi: 22.49 
           expectedOutput: 0.1883 
   BMI: 22.49%
   BFR: 18.84%
   --- PASS: TestCalcBFR (0.00s)
   PASS
   
   Process finished with the exit code 0
```

#### 健康状态评判与建议

编写单元测试，确保录入完整的性别、年龄、身高、体重时，最终获得的健康建议符合预期

[>> 健康状态评判与建议](https://github.com/AdaSheng07/ready.to.go/blob/8a229c059febb49f2860c2587e7205562ca61b7f/staging/src/github.com/armstrongli/go-bmi/give_suggestions_test.go)

```go
    package gobmi
    
    import (
        "github.com/shopspring/decimal"
        "testing"
    )
    
    func TestGiveOutSuggestions(t *testing.T) {
        inputHeight, inputWeight := 1.70, 65.0
        inputGender := "male"
        inputAge := 35
        expectedBMI := 22.49
        expectedBFR := 0.1883
        expectedSuggestion := "目前的体脂率水平处于【偏胖】，建议：少坐多动，注意健康饮食"
    
        t.Logf("start giving out suggestions: \ngender: %s \nage: %d \nheight: %.2f \nweight: %.2f \n", inputGender, inputAge, inputHeight, inputWeight)
        actualBMI, err1 := CalcBMI(inputWeight, inputHeight)
        if err1 != nil {
            t.Fatalf("major error emerges while calculating bmi: %v\n", err1)
        }
        actualFormatBMI, _ := decimal.NewFromFloat(actualBMI).RoundFloor(2).Float64()
        if expectedBMI != actualFormatBMI {
            t.Logf("expecting %f, but got %f.\n", expectedBMI, actualFormatBMI)
            t.Failed()
        }
    
        actualBFR, err2 := CalcBFR(actualBMI, inputAge, inputGender)
        if err2 != nil {
            t.Fatalf("major error emerges while calculating body fat rate: %v\n", err2)
        }
        actualFormatBFR, _ := decimal.NewFromFloat(actualBFR).RoundFloor(4).Float64()
        if expectedBFR != actualFormatBFR {
            t.Logf("expecting %f, but got %f.\n", expectedBFR, actualFormatBFR)
            t.Failed()
    
        }
    
        actualSuggestion := GiveOutSuggestions(actualBFR, inputGender, inputAge)
        if expectedSuggestion != actualSuggestion {
            t.Logf("expecting %s, but got %s.\n", expectedSuggestion, actualSuggestion)
            t.Failed()
        }
    }
```
单元测试的运行结果为：
```
    /usr/local/go/bin/go tool test2json -t /private/var/folders/7l/pnr12b8962l7dwhrxxzjd9500000gn/T/GoLand
    /___TestGiveOutSuggestions_in_github_com_armstrongli_go_bmi.test -test.v -test.paniconexit0 -test.run ^\QTestGiveOutSuggestions\E$
    === RUN   TestGiveOutSuggestions
        give_suggestions_test.go:16: start giving out suggestions: 
            gender: male 
            age: 35 
            height: 1.70 
            weight: 65.00 
    BMI: 22.49%
    BFR: 18.84%
    目前的体脂率水平处于【偏胖】，建议：少坐多动，注意健康饮食
    --- PASS: TestGiveOutSuggestions (0.00s)
    PASS
    
    Process finished with the exit code 0
```
还可以在`Terminal`对`staging`目录下所有的`_test.go`测试文件运行查看，结果为：
```
    $ go test
    BMI: 22.49%
    BFR: 18.84%
    BMI: 22.49%
    BFR: 18.84%
    目前的体脂率水平处于【偏胖】，建议：少坐多动，注意健康饮食
    PASS
    ok      github.com/armstrongli/go-bmi   0.881s
```





