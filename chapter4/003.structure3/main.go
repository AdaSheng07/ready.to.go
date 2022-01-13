package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

func main() {
	var num1, num2 int = 2, 3

	p := calculator{
		a: 0,
		b: 1,
	}
	fmt.Printf("%d + %d = %d\n", p.a, p.b, p.Add())

	old := calculator{
		a: num1,
		b: num2,
	}
	fmt.Printf("%d + %d = %d\n", old.a, old.b, old.Add())

	var c1 NewCalculator1
	c1.old.a = 3
	c1.old.b = 4
	fmt.Println("Nest structure calculator into structure NewCalculator1 w name 'old':")
	fmt.Printf("%d + %d = %d\n", c1.old.a, c1.old.b, c1.old.Add())
	fmt.Printf("%d + %d = %d\n", c1.old.a, c1.old.b, c1.Add())
	// c1.Add(): unsolved reference error if Add() is not redefined on *NewCalculator1

	c2 := NewCalculator2{}
	c2.a = 5
	c2.b = 6
	fmt.Println("Nest structure calculator into structure NewCalculator2 w/o name:")
	fmt.Printf("%d + %d = %d\n", c2.a, c2.b, c2.Add())
	c2.calculator.a = 7
	c2.calculator.b = 8
	fmt.Printf("%d + %d = %d\n", c2.calculator.a, c2.calculator.b, c2.calculator.Add())

	c3 := NewCalculator2{}
	c3.a = 9
	c3.b = 10
	fmt.Printf("%d + %d = %d\n", c3.a, c3.b, c3.Add())
	fmt.Printf("%d^2 + %d^2 = %d\n", c3.a, c3.b, c3.QuadraticSum())

	// instantiate StructB otherwise it will panic, because *StructB is a <nil> pointer
	c4 := StructA{&StructB{}}
	c4.c = 14
	fmt.Println(c4.c)

	c5 := StructB{}
	c5.options["key"] = "value" // panic: assignment to entry in nil map
	fmt.Println(c5.c)
	fmt.Println(c5.options)
	c6 := StructB{
		options: map[string]string{},
	}
	c6.options["key"] = "value"
	fmt.Println(c6.c)
	fmt.Println(c6.options)
}

type calculator struct {
	a, b int
}

func (p *calculator) Add() int {
	return p.a + p.b
}

func (p *calculator) QuadraticSum() int {
	return p.a*p.a + p.b*p.b
}

type NewCalculator1 struct {
	old calculator // a combination of two structures
}

func (c1 *NewCalculator1) Add() int {
	return c1.old.Add() // redefine Add() on *NewCalculator, otherwise there is no access to c1.Add()
}

type NewCalculator2 struct {
	calculator // nesting a anonymous structure calculator
}

type MyCommand struct {
	cobra.Command
}

type StructA struct {
	// instantiate StructB otherwise it will panic, because *StructB is a <nil> pointer
	*StructB
}
type StructB struct {
	c       int
	options map[string]string
}
