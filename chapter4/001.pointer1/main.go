package main

import "fmt"

func main() {
	a, b := 1, 2
	fmt.Printf("变量a的地址是：%x\n", &a)
	fmt.Printf("变量b的地址是：%x\n", &b)
	add(a, b)
	fmt.Println(a)
	// in this case, the printed result is still 1
}

// how to change value of a without returning any value in function add()
func add(a, b int) {
	// in this function, values of a and b are copied for calculation, they are not the same a and b in main()
	a = a + b
}
