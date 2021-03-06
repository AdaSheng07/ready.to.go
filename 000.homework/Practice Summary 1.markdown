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
![image](https://github.com/AdaSheng07/ready.to.go/blob/b86cbb5fbd925dee911c45aea541a8ad32767a3d/000.homework/img.png)
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

#### 思考
在`Version1.0`代码中，整个代码模块分为：
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
将`Version1.0`中的代码模块进行重构`refactor`成为函数方法，可在`main`函数中直接调用，大大简化`main`函数，使整个操作流程一目了然，还可以在此基础上解决更多需求。

[Version2.0 重构优化](https://github.com/AdaSheng07/ready.to.go/blob/d1483b82411414f66b69b75ba994f174f14490d0/000.homework/0001.bmiCalculator2/main.go)

### Furthermore...
在此基础上实现一些新需求：
- 不要求输入性别，同时计算两个性别的体脂
- 不要求输入年龄，计算在当前身高、体重下不同年龄的体脂
- 亚洲、欧洲、美洲的体质不同，年龄所占体脂比重不同，计算在相同身高、体重不变的情况下各洲的体脂......

### Appendix 参考资料

❖ ︎[在线正则表达式测试](https://tool.oschina.net/regex/)  

❖ [常用正则表达式整理](https://xie.infoq.cn/article/7bf17ad93009c4a1f3045ea26)

❖ [基础知识-Golang 中的正则表达式](https://www.cnblogs.com/williamjie/p/9686311.html)

❖ [Golang 正则表达式（regexp）](https://cloud.tencent.com/developer/article/1706173)

❖ [Golang如何去判断字符串是数字还是字符](https://studygolang.com/topics/8696)

❖ [Golang之数据验证库validator的使用](https://juejin.cn/post/6990918041395544077)

❖ [Go的fmt效率问题](http://z-rui.github.io/post/2017/03/go-scanf/)

---
## 两条直线的关系判定

### 需求
两点决定一条直线，两条线是否平行取决于两条线的斜率是否一样。

### 主体思路

1. 输入直线1的两组坐标
2. 输入直线2的两组坐标 
3. 输入合法性判断：有无重合点 
4. 垂直于x轴的斜率不存在情况：
   - 都不存在/一个不存在
   - 平行/重合/相交 
5. 斜率存在的情况：
   - 斜率相同，截距不同：平行
   - 斜率相同，截距也相同：重合
   - 斜率不相同：相交
6. 输出判定结果

### 细节展开与补充

1. 输入直线1的两组坐标
   - 将坐标存入`2 * 2`数组`line1`
2. 输入直线2的两组坐标
   - 将坐标存入`2 * 2`数组`line2`
   ```
        /*  line1              line2
              [0] [1]            [0] [1]
          [0]{x1, y1}        [0]{x3, y3}
          [1]{x2, y2}        [1]{x4, y4}  */
   ```
3. 输入合法性判断：有无重合点
   - 如果`x1 == x2 && y1 == y2`或者`x3 == x4 && y3 == y4`，`line1`或`line2`中的两组坐标是有重复的
   - 有重复，重新输入两组坐标
4. 斜率不存在情况：
    - 都不存在：两条直线都垂直于x轴，需要判定是两线【平行】还是两线【重合】
    - 只有一条直线的斜率不存在：两条直线一定【相交】
    - 输出判定结果
5. 斜率存在的情况：
    - 计算斜率并比较：
        - 斜率相同，截距不同：两线【平行】
        - 斜率相同，截距也相同：两线【重合】
        - 斜率不相同：两线【相交】
    - 输出判定结果

### 代码实现与优化

[Version1.0 初级实现](https://github.com/AdaSheng07/ready.to.go/blob/69d48f3fe29d5566806013519d09fb50a516c6db/000.homework/0002.parallelLines1/main.go)

#### 优化
在`Version1.0`基础上进行优化，抛弃`goto`语句，将输入合法性判断与输入模块合并重构`refactor`为一个新函数`InputCoordinates`，将斜率存在与不存在的情况的判定合并重构`refactor`为一个新函数`VerdictLinesRelation`，跳出数学思维。

[Version2.0 重构优化](https://github.com/AdaSheng07/ready.to.go/blob/a07834efef9ad9a42339682a36058524a75d7d4c/000.homework/0002.parallelLines2/main.go)

通过重构优化，将`main`主函数代码模块进一步简化抽象为：
1. 输入四组坐标并判断合法性，确认没有出现重复坐标
2. 直接判断两直线的关系，并输出结果
3. 增加了询问再来一次/结束判定直线关系的功能

### Appendix 参考资料
❖ ︎[Go 语言多维数组](https://www.runoob.com/go/go-multi-dimensional-arrays.html)

❖ ︎[Go 语言向函数传递数组](https://www.runoob.com/go/go-passing-arrays-to-functions.html)

❖ ︎[Golang函数的值传递与引用传递，以及如何用指针和切片传递来避免数组传递的缺点](https://blog.csdn.net/benben_2015/article/details/80884537)
