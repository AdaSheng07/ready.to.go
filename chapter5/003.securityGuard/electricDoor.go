package main

import "fmt"

type ElectricDoor struct {
}

func (*ElectricDoor) Open() {
	fmt.Println("ElectricDoor Open")
}

func (*ElectricDoor) Close() {
	fmt.Println("ElectricDoor Close")
}
