package main

import "fmt"

func main() {
	// iteration: fibonacci sequence problem
	fmt.Println(fibonacci(10))
	// if n = 100, this go procedure will run really slowly and take all CPU
	// to avoid this disadvantage, we replace iteration with for-loop
}
func fibonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

// 1 1 2 3 5
// 1 1 2 3 5 8 13 21 34 55
