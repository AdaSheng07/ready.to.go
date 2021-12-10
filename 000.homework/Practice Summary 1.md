# Practice Summary 1

## 基于BMI的体脂计算器
### 需求
连续输入多人信息，输出每个人的姓名、BMI、体脂率、建议，最终打印总人数和他们的平均体脂率。
### 主体思路
1. 输入单人信息：姓名、年龄、性别、身高、体重
2. 根据每个人的性别和年龄，确定`sexWeight`值，以及对应的体脂率BFR与健康水平判定标准
3. 根据输入的信息计算BMI值，再计算体脂率BFR
4. 根据计算所得的体脂率BFR判定此人的状态：偏瘦/标准/偏胖/肥胖/严重肥胖，并给出对应的建议 
5. 询问是否继续录入信息
6. 信息录入结束后，按照输入顺序，输出每个人的姓名、BMI、体脂率BFR、对应建议
7. 计算并打印所有输入人员的总人数和平均体脂率BFR
### 细节展开与补充
1. 输入单人信息：姓名、年龄、性别、身高、体重：  
   - 明确五个变量的类型并声明：`str`/`float64`/`int`  
   - 利用正则表达式和`regexp`包对变量输入值进行合法性判断：
     - 合法：继续录入下一个变量值
     - 不合法：重新录入该变量，可以设置`label`用`goto`跳转实现  
2. 根据每个人的性别和年龄，确定`sexWeight`值，以及体脂率BFR与健康水平判定标准：  
   - `sexWeight`值：男为`1`，女为`0`
   - 根据每个人的性别和年龄确定体脂率BFR与健康水平判定标准：
     - 单独增加此模块，是为了**简化后续根据体脂率BFR值判断状态的过程，减少if-else嵌套**
     - 根据表格总结规律，其实是简单的数学问题
3. 根据输入的信息计算BMI值，再计算体脂率BFR：
   - BMI计算公式：`BMI = 体重(kg) ÷ (身高(m) × 身高(m))`
   - BFR计算公式：`BFR = 1.2 × BMI + 0.23 × 年龄 - 5.4 - 10.8 × sexWeight`
   - 计算完成后，为了便于最终打印结果，记录当前参与者的姓名、BMI和体脂率BFR，并分别推入切片`anticipates`, `bmiSet`和`bfrSet`中
4. 根据计算所得的体脂率BFR判定此人的状态：偏瘦/标准/偏胖/肥胖/严重肥胖，并给出对应的建议：
   - `if-else`条件语句判断偏瘦/标准/偏胖/肥胖/严重肥胖
   - 判断完成后，将判断结果和对应建议存入切片`suggestions`中
5. 询问是否继续录入信息：
   - 回答`y`：继续录入信息，总人数`num`加一，再回到程序开头输入单人信息
   - 回答其他：结束录入，进入输出环节
6. 信息录入结束后，按照输入顺序，输出每个人的姓名、BMI、体脂率BFR、对应建议：
   - `for-loop`从`0-(num-1)`遍历切片`anticipates`, `bmiSet`, `bfrSet`和`suggestions`
   - 格式化输出每个人的姓名、BMI、体脂率BFR和对应建议
     - 同时累加`bfrSet`中元素之和，记为`sum`
7. 计算并打印所有输入人员的总人数和平均体脂率BFR
   - 利用上一步计算所得`sum`计算平均体脂率：`sum / num`
   - 格式化输出最终的总人数与他们的平均体脂率

### 代码实现与优化
[Version1.0 初级实现](https://github.com/AdaSheng07/ready.to.go/blob/69d48f3fe29d5566806013519d09fb50a516c6db/000.homework/0001.bmiCalculator1/main.go)
</br>
#### 思考
在 Version1.0 代码中，整个代码模块分为：
- 信息录入与数据合法性验证
- 当前人员的健康水平判定标准
- BMI与体脂率BFR的计算与存储
- 基于体脂率BFR的当前人员的健康水平判断与结果存储
- 录入信息流程继续/结束询问
- 打印最终结果

在录入数据合法性验证和询问是否继续录入信息时，为了实现条件转移，复用代码，构成循环，使用了较多的`goto`语句跳转到指定标签。

☞ [为什么说 goto 是一种不好的用法？](https://www.zhihu.com/question/20259336)  
☞ [Why does Go have a "goto" statement?](https://stackoverflow.com/questions/11064981/why-does-go-have-a-goto-statement)

**有没有更优美的写法呢？**

#### 优化
将Version 1中的代码模块进行重构`refactor`成为函数方法，可在`main`函数中直接调用，大大简化`main`函数，使整个操作流程一目了然，还可以在此基础上解决更多需求。  

[Version2.0 重构优化](https://github.com/AdaSheng07/ready.to.go/blob/d1483b82411414f66b69b75ba994f174f14490d0/000.homework/0001.bmiCalculator2/main.go)

#### Furthermore...
在此基础上实现一些新需求：
- 不要求输入性别，同时计算两个性别的体脂
- 不要求输入年龄，计算在当前身高、体重下不同年龄的体脂
- 亚洲、欧洲、美洲的体质不同，年龄所占体脂比重不同，计算在相同身高、体重不变的情况下各洲的体脂......

#### Appendix 参考资料
❖ ︎[在线正则表达式测试](https://tool.oschina.net/regex/)  
❖ [常用正则表达式整理](https://xie.infoq.cn/article/7bf17ad93009c4a1f3045ea26)  
❖ [基础知识-Golang 中的正则表达式](https://www.cnblogs.com/williamjie/p/9686311.html)  
❖ [Golang 正则表达式（regexp）](https://cloud.tencent.com/developer/article/1706173)  
❖ [Golang如何去判断字符串是数字还是字符](https://studygolang.com/topics/8696)  
❖ [Golang之数据验证库validator的使用](https://juejin.cn/post/6990918041395544077)  
❖ [Go的fmt效率问题](http://z-rui.github.io/post/2017/03/go-scanf/)

---
## 两条直线的关系判定
[Version 1 初级实现](https://github.com/AdaSheng07/ready.to.go/blob/69d48f3fe29d5566806013519d09fb50a516c6db/000.homework/0002.parallelLines1/main.go)