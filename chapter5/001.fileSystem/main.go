package main

import "fmt"

func main() {
	var data string
	{
		var equipment IOInterface = &Soft{}
		data = equipment.Read()
		fmt.Println(data)
	}
	{
		var equipment IOInterface = &Mag{}
		data = equipment.Read()
		fmt.Println(data)
	}
	{
		var equipment IOInterface = &Paper{}
		data = equipment.Read()
		fmt.Println(data)
	}
	{
		var equipment IOInterface = &Sata{}
		data = equipment.Read()
		fmt.Println(data)
	}
	// 这套标准就是接口
}

type IOInterface interface {
	// ⬇️ 很像方法，是什么呢？
	Read() string
}

type Soft struct {
}

func (Soft) Read() string {
	return "软盘....."
}

type Mag struct {
}

func (Mag) Read() string {
	return "磁盘......"
}

type Paper struct {
}

func (Paper) Read() string {
	return "- - - 纸带 - - - "
}

type Sata struct {
}

func (Sata) Read() string {
	return "Sata."
}
