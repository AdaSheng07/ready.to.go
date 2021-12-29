package main

import (
	"fmt"
	"runtime/debug"
)

func main() {
	convertTypes()
	fmt.Println("finish!")
}

func convertTypes() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("panic: major error!")
			debug.PrintStack()
		}
	}()
	var a interface{}
	a = "convert a"

	b := a.(int)
	fmt.Println(b)
}

/*
panic: interface conversion: interface {} is string, not int

goroutine 1 [running]:
main.convertTypes()
        /Users/shengyitang/gopath/src/learn.go/chapter3/015.convertTypes/main.go:13 +0x2e
main.main()
        /Users/shengyitang/gopath/src/learn.go/chapter3/015.convertTypes/main.go:6 +0x17

Process finished with the exit code 2

*/
