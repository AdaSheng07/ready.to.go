package main

import "fmt"

type manpowerpush struct {
}

func (mpp *manpowerpush) OpenTheDoorOfRefrigerator(Refrigerator) error {
	fmt.Println("用 manpowerpush 做 OpenTheDoorOfRefrigerator(Refrigerator) error")
	return nil
}

func (mpp *manpowerpush) PutElephantIntoRefrigerator(Elephant, Refrigerator) error {
	fmt.Println("用 manpowerpush 做 PutElephantIntoRefrigerator(Elephant, Refrigerator) error")
	return nil
}

func (mpp *manpowerpush) CloseTheDoorOfRefrigerator(Refrigerator) error {
	fmt.Println("用 manpowerpush 做 CloseTheDoorOfRefrigerator(Refrigerator) error")
	return nil
}
