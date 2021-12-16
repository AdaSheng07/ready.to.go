package main

import "fmt"

func main() {

	// map declaration and initialization
	// 1. use keyword `map` to declare and initialize
	var map1 map[string]string
	map2 := map[string]string{"China": "Beijing"}
	map3 := map[string]string{"France": "Paris", "Japan": "Tokyo", "England": "London"}
	fmt.Println(map1)
	fmt.Println(map2)
	fmt.Println(map3)
	// 2. use keyword `map` to declare, then use built-in function make() to initialize
	var map4 map[string]string
	map4 = make(map[string]string, 4) // length = 4
	fmt.Println("map4 is: ", map4, ", with length of", len(map4))
	// 3. use built-in function make() to declare, then add in key-value elements
	map5 := make(map[string]string, 3) // length = 3
	map5["Germany"] = "Berlin"
	map5["Italy"] = "Rome"
	map5["Canada"] = "Ottawa"
	map5["United States"] = "Washington DC"
	fmt.Println("map5 is: ", map5, ", with length of", len(map5))

}
