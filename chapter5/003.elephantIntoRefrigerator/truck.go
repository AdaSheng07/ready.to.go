package main

import "fmt"

type truck struct {
}

func (t *truck) OpenTheDoorOfRefrigerator(Refrigerator) error {
	fmt.Println("用 manpowerpush 做 OpenTheDoorOfRefrigerator(Refrigerator) error")
	return nil
}

func (t *truck) PutElephantIntoRefrigerator(Elephant, Refrigerator) error {
	fmt.Println("用 manpowerpush 做 PutElephantIntoRefrigerator(Elephant, Refrigerator) error")
	return nil
}

func (t *truck) CloseTheDoorOfRefrigerator(Refrigerator) error {
	fmt.Println("用 manpowerpush 做 CloseTheDoorOfRefrigerator(Refrigerator) error")
	return nil
}
