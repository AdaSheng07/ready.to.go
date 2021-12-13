package main

import (
	"fmt"
)

/*
Golang中的循环只有一种，就是 for 循环
for-loop形式
	1. for init; condition; post {	}
	2. for ; condition; post {	}
	3. for condition {	}
	4. for {		}
* 无 condition 或者 condition 永远为 true 时在循环体内需要有退出条件

p.s.
	init： 一般为赋值表达式，给控制变量赋初值；
	condition： 关系表达式或逻辑表达式，循环控制条件；
	post： 一般为赋值表达式，给控制变量增量或减量。

for-loop中的循环控制语句
	break: 用于中断当前 for 循环或跳出 switch 语句
	continue: 跳过当前循环的剩余语句，然后继续进行下一轮循环
	goto: 将控制转移到被标记的语句

*/

func main() {
	fmt.Println("Round 1")
	for i := 0; i < 5; i++ {
		fmt.Println("你好，Golang!")
	}
	/* for-loop的工作逻辑
	定义i并初始化为0后：
	1. 判断i是否小于5，是，继续，不是，退出循环
	2. 输出"你好，Golang!"
	3. i++，i=1，再次判断i是否小于5...	*/

	fmt.Println("Round 2")
	j := 0
	for ; j < 5; j++ {
		fmt.Println("Hello, Golang!")
	}

	fmt.Println("Round 3")
	k := 0
	for k < 5 { // 省略;
		fmt.Println("Hi, Golang!")
		k++
	}

	fmt.Println("Round 4")
	l := 0
	for {
		fmt.Println("Yo, Golang!")
		l++
		if l >= 5 {
			break // 打断跳出for循环
		}
	}

	fmt.Println("Round 5")
	m := 0
	for {
		fmt.Println("Hello, Golang!", m)
		m++
		if m >= 10 {
			break
		}
		if m%2 == 0 {
			fmt.Println("Round 5，被continue跳过")
			continue
		}
		fmt.Println("Round 5，练习跳过", m)
	}

}
