package main

import "fmt"

// tall and weight are global variables
var tall, weight float64

func main() {
	fmt.Println("before assigning values to tall and weight: ")
	add() // 0

	// tall and weight are global variables
	tall, weight = 1.70, 70
	fmt.Println("after assigning values to tall and weight: ")
	add() // 71.7

	tall, weight := 1.80, 100 // they are not global variables anymore, now they are local variables
	fmt.Println("after redeclare and assign new values to tall and weight: ", tall, weight)
	add() // 71.7

	tall, weight = 100, 200
	fmt.Println("after reassign new values to tall and weight: ", tall, weight)
	add() // 71.7

}

func add() float64 {
	fmt.Println("tall + weight =", tall+weight)
	return tall + weight
}
func sampleSubdomainIf() {
	// v is only valid in if-else{}
	if v := add(); v > 100 {
		fmt.Println(v)
	} else {
		fmt.Println(v)
	}
	// fmt.Printf(v) // error: Unresolved reference 'v'
}

func sampleSubdomainFor() {
	// i is only valid in for-loop {}, same for switch-case
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	// fmt.Println(i) // error: Unresolved reference 'i'
}
