package main

import "fmt"

func main() {
	a, b := 1, 2
	add(&a, &b)
	fmt.Println(a)

	c := &a          // c is a pointer variable, its type is *int, it saves the address of a
	d := &c          // d is also a pointer variable, its type is **int, it saves the address of c
	fmt.Println(*c)  // this will give out value of a
	fmt.Println(*d)  // this will give out c, that is &a, the address of a
	fmt.Println(**d) // this is *(c), which will give out value of a
	fmt.Println("d =", d, "*d =", *d, "**d =", **d)
	// the passing process can be infinite:
	// d is the address of c
	// *d is the value of c, which is the address of a
	// **d is equal to *c, which is the value of a

	// pointers can point to any variables
	m := map[string]string{}
	mp1 := &m // mp1 is a pointer variable, its type is *map[string]string, it saves the address of map m
	fmt.Println(mp1)
	put(m)
	fmt.Println("*mp1 =", mp1)

	f1 := add // f1 is a function
	f1(&a, &b)
	fmt.Println("f1, add =", a) // this gives out 5 from 3 + 2
	f1p1 := &f1                 // f1p1 is a pointer variable, its type is *func(*int, *int), it saves the address of function f1
	// if we use *f1p1(&a, &b)... Error: it does not meet the priority rule, golang calculates first, gives out the result and then use *
	// we want to use *f1p1 first to get the function f1 itself, and then do the calculation, we MUST use brackets
	(*f1p1)(&a, &b)
	fmt.Println("f1p1, add =", a) // this gives out 7 from 5 + 2

	{
		var nothing *int                 // notice: nothing is a pointer that points to nil
		fmt.Println("nothing:", nothing) // this is <nil>
		// *nothing = 3                     // *nothing does not exist, so no valid assignment operations --> panic: runtime error: invalid memory address or nil pointer dereference
	}
	// this is also why we must instantiate a map, because we cannot assign any values to a nil map:
	{
		var nothingMap map[string]string
		// nothingMap["nothingMap"] = "empty" // panic: assignment to entry in nil map
		fmt.Println("nothingMap:", nothingMap) // this is map[]
	}
	{
		var nothingSlice []int
		fmt.Println(nothingSlice) // this is []
		// nothingSlice[0] = 100     // panic: runtime error: index out of range [0] with length 0
		// we cannot assign value to an uninstantiated nil slice, but we can append it w/o instantiation:
		nothingSlice = append(nothingSlice, 100)
		fmt.Println(nothingSlice) // this is [100]
	}
}

func put(m map[string]string) {
	m["a"] = "AAA"
}

func add(x, y *int) {
	*x = *x + *y
}
