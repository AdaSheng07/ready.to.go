package main

import "fmt"

func main() {
	var a, b int = 1, 2
	//var op string = "+"

	c := Calculator{
		left:  a,
		right: b,
		//operator: op,
	}
	fmt.Printf("&c @main() is: %p\n", &c)
	fmt.Println(c.Add())
	fmt.Println("c.result =", c.result) // this gives out: c.result = 0, WHY?
	fmt.Println(c.Sub())
	fmt.Println(c.Multiply())
	fmt.Println(c.Divide())
	fmt.Println(c.Remainder())
}
