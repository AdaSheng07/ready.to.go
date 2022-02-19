package main

import "fmt"

type SteelDoor struct {
}

func (*SteelDoor) Open() {
	fmt.Println("SteelDoor Open")
}

func (*SteelDoor) Close() {
	fmt.Println("SteelDoor Close")
}
