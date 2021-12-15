package main

import "fmt"

func main() {
	// Golang supports: convert a string array into a byte slice, and make amends based on the byte slice
	var a string = "hello"
	fmt.Println(a)
	fmt.Println(a[0])
	// a[0] = 'H' // Error: Cannot assign to a[0]
	aBytes := []byte(a)
	fmt.Println(aBytes)
	fmt.Println("Edit content in the slice! ")
	aBytes[0] = 'H'
	a = string(aBytes)
	fmt.Println(a)

	// To represent more characters, Golang encodes based on UTF-8 and supports accessing ASCII as well
	// Use type rune to gain access to every UTF-8 character
	var b string = "企鹅"
	fmt.Println(b)
	fmt.Println(b[0])
	bBytes1 := []byte(b)
	bBytes2 := []rune(b)
	fmt.Println("byte b is: ", bBytes1, len(bBytes1))
	fmt.Println("rune b is: ", bBytes2, len(bBytes2))
	// bBytes1[0] = '你' // Error: ''你'' (type rune) cannot be represented by the type byte
	bBytes1[0] = 'Q'
	b1 := string(bBytes1)
	fmt.Println("b1 is: ", b1)
	bBytes2[0] = '天'
	b2 := string(bBytes2)
	fmt.Println("b2 is: ", b2)
}
