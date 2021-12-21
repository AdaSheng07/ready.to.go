package main

import "fmt"

// tall and weight are global variables
var tall, weight float64

func main() {
	fmt.Println("before assigning values to tall and weight: ")
	add()

	// tall and weight are global variables
	tall, weight = 1.70, 70
	fmt.Println("after assigning values to tall and weight: ")
	add()

	tall, weight := 1.80, 100 // they are not global variables anymore, now they are local variables
	fmt.Println("after redeclare and assign new values to tall and weight: ", tall, weight)
	add()

	tall, weight = 100, 200
	fmt.Println("after reassign new values to tall and weight: ", tall, weight)
	add()

}
func add() {
	fmt.Println("tall + weight =", tall+weight)
}
