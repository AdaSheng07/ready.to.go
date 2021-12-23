package main

import (
	"fmt"
	"runtime/debug"
)

func recoverSample() {
	// adding defer with recover will catch panic in the program
	// with this fragment, program will not drop out automatically
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("fatal error discovered here:", r)
			debug.PrintStack()
		}
	}()

	defer catchPanicUpgraded()

	defer catchPanic()

	var nameScore map[string]int
	nameScore["lisa"] = 100 // panic: assignment to entry in nil map

}

// catchPanic: refactor func() to get it
// in this case, the panic will not be caught, why?
// this is because when we use catchPanic, the call process of func() has escaped the running of recoverSample
// they are not in the same stack anymore, therefore panic error cannot be caught
func catchPanic() {
	func() {
		if r := recover(); r != nil {
			fmt.Println("fatal error discovered:", r)
			debug.PrintStack()
		}
	}()
}

// upgrade it: this one will catch panic
func catchPanicUpgraded() {
	if r := recover(); r != nil {
		fmt.Println("fatal error discovered finally:", r)
		debug.PrintStack()
	}
}
