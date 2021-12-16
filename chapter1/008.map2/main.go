package main

import "fmt"

func main() {
	// operations on map
	// it is illegal to write in a nil map, in other words, an uninitialized map
	var map1 map[string]string
	fmt.Println(map1)
	// map1["Germany"] = "Berlin" // it can run w/o warnings, but end up with
	// panic: assignment to entry in nil map

	// 1. add/change key-value
	map2 := map[string]string{"France": "Paris", "Japan": "Tokyo", "England": "London", "United States": "New York"}
	fmt.Println("Before alteration, map2 is:", map2)
	map2["Germany"] = "Berlin"
	map2["China"] = "Beijing"
	map2["United States"] = "Washington DC"
	fmt.Println("After alteration, map2 is:", map2)
	// 2. delete key-value
	delete(map2, "China")
	// 3. lookup key-value：take key as index of map
	fmt.Println("The capital of Japan is", map2["Japan"])

	// traverse a map
	for key, value := range map2 {
		fmt.Printf("The capital of %s is %s. \n", key, value)
	}

	// true value or false value?
	check := []string{"England", "China", "France"}
	fmt.Printf("we want to know about capitals of %v.\n", check)
	for i := range check {
		capital, ok := map2[check[i]]
		if ok {
			fmt.Printf("true value: capital of %s is %s.\n", check[i], capital)
		} else {
			fmt.Printf("false value: we don't know about %s.\n", check[i])
		}
	}

}
