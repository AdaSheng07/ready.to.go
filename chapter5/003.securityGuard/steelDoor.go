package main

import "fmt"

var _ Door = &SteelDoor{}

// Cannot use '&SteelDoor{}' (type *SteelDoor) as the type Door Type does not implement 'Door' as some methods are missing: Unlock() Lock()

type SteelDoor struct {
}

func (*SteelDoor) Unlock() {
	fmt.Println("SteelDoor Unlock")
}

func (*SteelDoor) Lock() {
	fmt.Println("SteelDoor Lock")

}

func (*SteelDoor) Open() {
	fmt.Println("SteelDoor Open")
}

func (*SteelDoor) Close() {
	fmt.Println("SteelDoor Close")
}
