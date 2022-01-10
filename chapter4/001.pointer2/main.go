package main

import "fmt"

func main() {
	a, b := 1, 2
	// instead of copying values in variables a and b, here we use addresses of variables a and b
	add(&a, &b)
	fmt.Println(a)
}

// *int means add() has parameters that are addresses for integer-type variables
func add(x, y *int) {
	// x and y are pointers, *x and *y indicate values of variables x and y, respectively
	// extract values of x and y, add them up and put the result back into variable x
	*x = *x + *y
}
