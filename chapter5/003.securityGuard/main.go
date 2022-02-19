package main

import "fmt"

func main() {
	security := Assets{
		assets: []Asset{
			&GlassDoor{},
			&WoodDoor{},
			&SteelDoor{},
			&ElectricDoor{},
		},
	}
	fmt.Println("上班")
	security.OnDuty()
	fmt.Println("下班")
	security.OffDuty()
	fmt.Println("DONE")
}
