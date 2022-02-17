package main

import (
	"fmt"
	"reflect"
)

func main() {
	var refrigerator Refrigerator
	fmt.Println(refrigerator.Size)
	var elephant Elephant
	fmt.Println(elephant.Name)

	// 2. 接口是引用类型

	//var ip *int      // pointer
	//fmt.Println(*ip) // panic: runtime error: invalid memory address or nil pointer dereference

	var map1 map[string]string
	map1 = map[string]string{} // 实例化引用类型map
	fmt.Println(map1)

	// var putER PutElephantIntoRefrigerator
	// 3. 实例化时，接口不能直接使用变量
	//putER = &PutElephantIntoRefrigerator() // this is illegal
	// putER.OpenTheDoorOfRefrigerator(refrigerator)
	// panic: runtime error: invalid memory address or nil pointer dereference

	// 接口的实例化
	var putER2 PutElephantIntoRefrigerator
	putER2 = &PutElephantIntoRefrigeratorImplement{}
	// var putER2 PutElephantIntoRefrigerator = PutElephantIntoRefrigeratorImplement{} // this is also legal
	putER2.OpenTheDoorOfRefrigerator(refrigerator)
	putER2.PutElephantIntoRefrigerator(elephant, refrigerator)
	putER2.CloseTheDoorOfRefrigerator(refrigerator)

	// 同一个对象可以实现多个接口，e.g. Open 和 Close
	var open Open
	open = Refrigerator{
		Size: 4,
	}
	fmt.Println(open)
	var close Close = Refrigerator{
		Size: 4,
	}
	fmt.Println(close)

	// 接口的组合与嵌套
	var door Box = &Door{}
	door.Open()
	door.Close()

	open = door
	// door = open
	// Cannot use 'open' (type Open) as the type Box Type does not implement 'Box' as some methods are missing: Close() error

	// 9. 空接口 = `any` = 黑盒子
	var i interface{}
	i = 3
	fmt.Println(reflect.TypeOf(i), "value:", i) // int value: 3
	i = 3.4
	fmt.Println(reflect.TypeOf(i), "value:", i) // float64 value: 3.4
	i = Refrigerator{}
	fmt.Println(reflect.TypeOf(i), "value:", i) // main.Refrigerator value: {0}
}

type PutElephantIntoRefrigerator interface {
	// 1. 接口定义只有方法：接口中定义的各种方法签名，不能有值或者其它成员变量
	OpenTheDoorOfRefrigerator(Refrigerator) error
	PutElephantIntoRefrigerator(Elephant, Refrigerator) error
	CloseTheDoorOfRefrigerator(Refrigerator) error
	//name string // error: 接口中不能定义变量
}

type TestTypeImplementInterface func()

func (t TestTypeImplementInterface) OpenTheDoorOfRefrigerator(refrigerator Refrigerator) error {
	return nil
}

func (t TestTypeImplementInterface) PutElephantIntoRefrigerator(elephant Elephant, refrigerator Refrigerator) error {
	return nil
}

func (t TestTypeImplementInterface) CloseTheDoorOfRefrigerator(refrigerator Refrigerator) error {
	return nil
}

// 3. 实例化时，接口不能直接使用变量，必须指向实现
type PutElephantIntoRefrigeratorImplement struct {
}

func (operation PutElephantIntoRefrigeratorImplement) OpenTheDoorOfRefrigerator(refrigerator Refrigerator) error {
	fmt.Println("open the door of refrigerator")
	return nil
}

func (operation PutElephantIntoRefrigeratorImplement) PutElephantIntoRefrigerator(elephant Elephant, refrigerator Refrigerator) error {
	fmt.Println("put elephant into refrigerator")
	return nil
}

func (operation PutElephantIntoRefrigeratorImplement) CloseTheDoorOfRefrigerator(refrigerator Refrigerator) error {
	fmt.Println("close the door of refrigerator")
	return nil
}

type Open interface {
	Open() error
}

type Close interface {
	Close() error
}

type Box interface {
	Open
	Close
}

type Door struct {
}

func (d *Door) Open() error {
	fmt.Println("Open the door!")
	return nil
}

func (d *Door) Close() error {
	fmt.Println("Close the door!")
	return nil
}

type Refrigerator struct {
	Size int
}

func (Refrigerator) Open() error {
	fmt.Println("Open!")
	return nil
}

func (Refrigerator) Close() error {
	fmt.Println("Close!")
	return nil
}

type Elephant struct {
	Name string
}
