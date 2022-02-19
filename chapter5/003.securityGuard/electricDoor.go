package main

import "fmt"

var _ Door = &ElectricDoor{}

type ElectricDoor struct {
}

func (*ElectricDoor) Open() {
	fmt.Println("ElectricDoor Open")
}

func (*ElectricDoor) Close() {
	fmt.Println("ElectricDoor Close")
}

func (*ElectricDoor) Unlock() {
	fmt.Println("ElectricDoor Unlock")
}

func (*ElectricDoor) Lock() {
	fmt.Println("ElectricDoor Lock")
}
