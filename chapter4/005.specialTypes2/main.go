package main

import (
	"fmt"
	"math/rand"
)

type Math = int
type English = int

func main() {
	var mathScore1 int = 100
	var mathScore2 Math = 99
	mathScore2 = mathScore1 // value of mathScore1 can be assigned to mathScore2, they have the same type
	fmt.Println(mathScore2)

	v := voting{
		students: []*student{
			&student{name: fmt.Sprintf("%s", "a")},
			&student{name: fmt.Sprintf("%s", "b")},
			&student{name: fmt.Sprintf("%s", "c")},
			&student{name: fmt.Sprintf("%s", "d")},
			&student{name: fmt.Sprintf("%s", "e")},
			&student{name: fmt.Sprintf("%s", "f")}, &student{name: fmt.Sprintf("%s", "g")},
			&student{name: fmt.Sprintf("%s", "h")},
			&student{name: fmt.Sprintf("%s", "i")},
			&student{name: fmt.Sprintf("%s", "j")},
			&student{name: fmt.Sprintf("%s", "k")},
			&student{name: fmt.Sprintf("%s", "l")},
			&student{name: fmt.Sprintf("%s", "m")},
			&student{name: fmt.Sprintf("%s", "n")},
			&student{name: fmt.Sprintf("%s", "o")},
			&student{name: fmt.Sprintf("%s", "p")},
			&student{name: fmt.Sprintf("%s", "q")},
			&student{name: fmt.Sprintf("%s", "r")},
			&student{name: fmt.Sprintf("%s", "s")},
			&student{name: fmt.Sprintf("%s", "t")},
			&student{name: fmt.Sprintf("%s", "u")},
			&student{name: fmt.Sprintf("%s", "v")},
			&student{name: fmt.Sprintf("%s", "w")},
			&student{name: fmt.Sprintf("%s", "x")},
			&student{name: fmt.Sprintf("%s", "y")},
			&student{name: fmt.Sprintf("%s", "z")},
		},
	}
	monitor := v.goRun()
	fmt.Println(*monitor)
	monitor.Distribute()

	var stdXQ = &student{name: "xiaoqiang"}
	var stdLH = student{name: "lihua"}
	var ldXQ Monitor = Monitor(*stdXQ)
	var ldLH *Monitor = (*Monitor)(&stdLH)
	ldXQ.Distribute()
	ldLH.Distribute()
	fmt.Println(ldXQ, ldLH)

	// 与byte的强行转换类似：
	//bytesTest1 := []byte{}
	//var str1 string = string(bytesTest1)
}

// renaming types improves readability of codes, for example:
func getMathAndEnglishScoresOfStudents(name string) (Math, English) {
	// ...
	return 0, 0
}

// vote game
type voting struct {
	students []*student
}

//
func (v *voting) goRun() *Monitor {
	for _, item := range v.students {
		randInt := rand.Int()
		if randInt%2 == 0 {
			item.voteAgree(v.students[randInt%len(v.students)])
		} else {
			item.voteDisagree(v.students[randInt%len(v.students)])
		}
	}
	maxScore := -1
	maxScoreIndex := -1
	for i, item := range v.students {
		if maxScore < item.agree {
			maxScore = item.agree
			maxScoreIndex = i
		}
	}
	if maxScoreIndex >= 0 {
		// return v.students[maxScoreIndex] // 此时返回的是student，不是monitor，程序报错
		// 强制转换对象类型：
		return (*Monitor)(v.students[maxScoreIndex])
	}
	return nil
}

type student struct {
	name     string
	agree    int
	disagree int
}

// 使用嵌套对象定义（继承）方式来定义班长
//type Monitor struct {
//	student
//}

// Monitor 使用类型（对象）重定义
type Monitor student

// Distribute 新类型Monitor可以拥有自己的方法
func (m *Monitor) Distribute() {
	fmt.Println("发作业")
}

// 也可以对函数类型添加新的方法
type WhatIfAFunction func()

func (w *WhatIfAFunction) test111() {

}

type WhatIfAnArray []string // 引用类型

func (std *student) voteAgree(voted *student) {
	voted.agree++
}

func (std *student) voteDisagree(voted *student) {
	voted.disagree++
}
