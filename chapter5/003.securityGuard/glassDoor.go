package main

import "fmt"

type GlassDoor struct {
}

func (*GlassDoor) Open() {
	fmt.Println("ElectricDoor Open")
}

func (*GlassDoor) Close() {
	fmt.Println("ElectricDoor Close")
}
