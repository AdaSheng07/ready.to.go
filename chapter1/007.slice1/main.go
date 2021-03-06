package main

import "fmt"

func main() {
	// slice declaration
	a := []int{}
	fmt.Println(a)
	b := [0]int{}
	fmt.Println(b)
	// a and b are both [], but they are different types
	// a is a slice, while b is an array

	// compare slice with array
	array := [3]int{111, 222, 333} // type: [3]int
	fmt.Println(array)
	slice := []int{111, 222, 333} // type: []int
	fmt.Println(slice)
	for i, val := range slice {
		fmt.Println(i, val)
	}
	// slice supports ALL operations which are supported by array

	// In addition, slice also supports...
	// 1. dynamically add the same declared slice type of elements in the slice
	var slice1 []int = []int{111, 222, 333}
	var slice2 []int = []int{222, 333, 444}
	var slice3 []int
	fmt.Println("Use append to add elements: ")
	slice1 = append(slice1, 444)
	slice2 = append(slice2, 555, 666)
	slice3 = append(slice2, 777)
	fmt.Printf("slice1: %v\n", slice1)
	fmt.Printf("slice2: %v\n", slice2)
	fmt.Printf("slice3: %v\n", slice3)

	// 2. dynamically delete elements from the slice
	fmt.Println("Use append to delete elements: ")
	// NOTICE: slice[2 : 5] = {slice[2], slice[3], slice[4]}
	fmt.Printf("Before deleting elements: slice1 = %v\n", slice1)
	fmt.Printf("Before deleting elements: slice2 = %v\n", slice2)
	fmt.Printf("Before deleting elements: slice3 = %v\n", slice3)
	slice2 = append(slice2[:1], slice2[3:]...)
	fmt.Printf("After deleting elements: slice1 = %v\n", slice1)
	fmt.Printf("After deleting elements: slice2 = %v\n", slice2)
	fmt.Printf("After deleting elements: slice3 = %v\n", slice3)
	// NOTICE: how and why does slice3 change as well?

	// Before invert a new element, make a backup
	a = []int{6, 4, 2, 0}
	backup := a[1:]
	fmt.Println("backup is: ", backup) // [4 2 0]
	a = append(a[:1], 120)
	fmt.Println("a is: ", a)           // [6 120]
	fmt.Println("backup is: ", backup) // [120 2 0]
	a = append(a, backup...)
	fmt.Println("a is: ", a) // [6 120 120 2 0], not [6 120 4 2 0], why?
	// `backup` updates as `a` updates, they share the same memory space and change synchronously.
	// How can we solve it?
	a2 := []int{6, 4, 2, 0}
	backup2 := append([]int{}, a2[1:]...)
	fmt.Println("backup2 is: ", backup2) // [4 2 0]
	a2 = append(a2[:1], 220)
	fmt.Println("a2 is: ", a2)           // [6 220]
	fmt.Println("backup2 is: ", backup2) // [4 2 0]
	a2 = append(a2, backup2...)
	fmt.Println("a2 is: ", a2) // [6 220 4 2 0]
	// We can use copy as well to make an effective backup
	a3 := []int{6, 4, 2, 0}
	var backup3 []int = make([]int, len(a3[1:]), cap(a3[1:])*2)
	copy(backup3, a3[1:])
	fmt.Println("backup3 is: ", backup3) // [4 2 0]
	a3 = append(a3[:1], 320)
	fmt.Println("a3 is: ", a3)           // [6 320]
	fmt.Println("backup3 is: ", backup3) // [4 2 0]
	a3 = append(a3, backup3...)
	fmt.Println("a3 is: ", a3) // [6 320 4 2 0]

	// 3. use two slices to create a new slice
	fmt.Println("Put slice1 and slice2 together to create slice4 ")
	slice4 := append(slice1[:], slice2[:]...)
	fmt.Printf("slice4 is %v", slice4)

	// 4. use make(slice_name type, len(slice_name), cap(slice_name)) to expand slice capacity
	slice5 := make([]int, len(slice4), cap(slice4)*3)
	fmt.Printf("slice4 is %v, its length is %d, its capacity is %d\n", slice4, len(slice4), cap(slice4))
	fmt.Printf("slice5 is %v, its length is %d, its capacity is %d\n", slice5, len(slice5), cap(slice5))

	// 5. use copy(dst []Type, src []Type) to duplicate the content of a slice
	// before copying, use make() to declare and initialize the destination slice with capacity 2 times that of the source slice
	var slice6 []int = make([]int, len(slice4), cap(slice4)*2)
	fmt.Println("Copy the content of slice4 to slice6:")
	copy(slice6, slice4)
	fmt.Printf("slice4 is %v, slice6 is %v\n", slice4, slice6)
}
