package main

import (
	"fmt"
)

func main() {
	fmt.Println("Round 1")
	for i := 0; i < 5; i++ {
		fmt.Println("你好，Golang!")
	}

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
	m := 0
	for {
		fmt.Println("Yo, Golang!")
		m++
		if m >= 5 {
			break // 打断跳出for循环
		}
	}

	fmt.Println("Round 5")
	m = 0
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
