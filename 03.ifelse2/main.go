package main

import "fmt"

func main() {
	var money int
	if money >= 20 {
		fmt.Println("点个外卖")
	} else if money >= 200 {
		fmt.Println("去下馆子")
	} else if money >= 2000 {
		fmt.Println("去米其林")
	} else if money >= 20000 {
		fmt.Println("出国玩一圈")
	} else {
		fmt.Println("饿着吧")
	}
}
