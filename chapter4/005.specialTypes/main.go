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
			&student{name: fmt.Sprintf("%s", "f")},
			&student{name: fmt.Sprintf("%s", "g")},
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
			&student{name: fmt.Sprintf("%s", "aa")},
			&student{name: fmt.Sprintf("%s", "bb")},
			&student{name: fmt.Sprintf("%s", "cc")},
			&student{name: fmt.Sprintf("%s", "dd")},
			&student{name: fmt.Sprintf("%s", "ee")},
			&student{name: fmt.Sprintf("%s", "ff")},
			&student{name: fmt.Sprintf("%s", "gg")},
			&student{name: fmt.Sprintf("%s", "hh")},
			&student{name: fmt.Sprintf("%s", "ii")},
			&student{name: fmt.Sprintf("%s", "jj")},
			&student{name: fmt.Sprintf("%s", "kk")},
			&student{name: fmt.Sprintf("%s", "ll")},
			&student{name: fmt.Sprintf("%s", "mm")},
			&student{name: fmt.Sprintf("%s", "nn")},
			&student{name: fmt.Sprintf("%s", "oo")},
			&student{name: fmt.Sprintf("%s", "pp")},
			&student{name: fmt.Sprintf("%s", "qq")},
			&student{name: fmt.Sprintf("%s", "rr")},
			&student{name: fmt.Sprintf("%s", "ss")},
			&student{name: fmt.Sprintf("%s", "tt")},
			&student{name: fmt.Sprintf("%s", "uu")},
			&student{name: fmt.Sprintf("%s", "vv")},
			&student{name: fmt.Sprintf("%s", "ww")},
			&student{name: fmt.Sprintf("%s", "xx")},
			&student{name: fmt.Sprintf("%s", "yy")},
			&student{name: fmt.Sprintf("%s", "zz")},
		},
	}
	monitor := v.goRun()
	fmt.Println(*monitor)
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
		return v.students[maxScoreIndex]
	}
	return nil
}

type student struct {
	name     string
	agree    int
	disagree int
}

type Monitor = student

func (std *student) voteAgree(voted *student) {
	voted.agree++
}

func (std *student) voteDisagree(voted *student) {
	voted.disagree++
}
