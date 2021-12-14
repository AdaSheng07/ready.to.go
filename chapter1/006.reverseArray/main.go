package main

import "fmt"

func main() {
	var array1 [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	fmt.Println(array1)
	var temp int
	for i, val := range array1 {
		if i >= len(array1)/2 {
			break
		}
		temp = val
		array1[i] = array1[len(array1)-1-i]
		array1[len(array1)-1-i] = temp
	}
	fmt.Println(array1)

	// 尝试用双指针
	var array2 [5]string = [5]string{"love", "me", "like", "you", "do"}
	fmt.Println(array2)
	for left, right := 0, len(array2)-1; left < right; left++ {
		array2[left], array2[right] = array2[right], array2[left]
		right--
	}
	fmt.Println(array2)
}
