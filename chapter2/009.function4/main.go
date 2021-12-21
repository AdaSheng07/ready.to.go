package main

import "fmt"

// special functions

var T int = start()

// 1. init() runs before main()
func init() {
	fmt.Println("count down: 3")
}
func init() {
	fmt.Println("count down: 2")
}
func init() {
	fmt.Println("count down: 1")
}

func main() {
	fmt.Printf("count down: %v\n", T)
	fmt.Println("Boom!")
}
func start() int {
	fmt.Println("Start!")
	return 0
}
