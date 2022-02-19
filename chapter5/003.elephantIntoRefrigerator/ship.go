package main

import "fmt"

type shipElephant struct {
}

func (se *shipElephant) OpenTheDoorOfRefrigerator(Refrigerator) error {
	fmt.Println("用 ship 做 OpenTheDoorOfRefrigerator(Refrigerator) error")
	return nil
}

func (se *shipElephant) PutElephantIntoRefrigerator(Elephant, Refrigerator) error {
	fmt.Println("用 ship 做 PutElephantIntoRefrigerator(Elephant, Refrigerator) error")
	return nil
}

func (se *shipElephant) CloseTheDoorOfRefrigerator(Refrigerator) error {
	fmt.Println("用 ship 做 CloseTheDoorOfRefrigerator(Refrigerator) error")
	return nil
}
