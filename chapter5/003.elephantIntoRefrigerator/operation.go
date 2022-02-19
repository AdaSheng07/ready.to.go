package main

import "fmt"

func operation(elephantIntoRefrigerator PutElephantIntoRefrigerator, r Refrigerator, e Elephant) {
	fmt.Println("To put an elephant into a refrigerator:")

	// 如果使用的是人力，给警告 warning: you are using man power, not efficient
	if _, ok := elephantIntoRefrigerator.(*manpowerpush); ok {
		fmt.Println("warning: you are using man power, not efficient")
	}

	switch elephantIntoRefrigerator.(type) {
	case *manpowerpush:
		{
			fmt.Println("warning: you are using man power, not efficient")
		}
	case *truck:
		{
			fmt.Println("note: you are using truck")
		}
	case *shipElephant:
		{
			fmt.Println("note: you are shipping the elephant")
		}
	}

	elephantIntoRefrigerator.OpenTheDoorOfRefrigerator(r)
	elephantIntoRefrigerator.PutElephantIntoRefrigerator(e, r)
	elephantIntoRefrigerator.CloseTheDoorOfRefrigerator(r)

	fmt.Println("This is how we put an elephant into a refrigerator.")
}
