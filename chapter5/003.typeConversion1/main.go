package main

import "fmt"

func main() {
	var refrigerator Refrigerator // 具体的实现
	var door Door                 // 具体的实现
	var op Open                   // 抽象的接口，实现是未知的
	var cl Close                  // 抽象的接口，实现是未知的
	var box Box
	// Refrigerator实现了Open、Close和Box接口，那么可以将Refrigerator转换回Open和Close吗？
	op = refrigerator  // 成功转换，特殊到一般
	cl = refrigerator  // 成功转换，特殊到一般
	box = refrigerator // 成功转换，特殊到一般
	op.Open()          // Open!
	cl.Close()         // Close!
	box.Open()         // Open!
	box.Close()        // Close!
	fmt.Println("open:", op, "close:", cl)
	op = &door  // 必须要有指针，因为Open和Close方法是在对象Door的指针上实现的，特殊到一般
	cl = &door  // 必须要有指针，因为Open和Close方法是在对象Door的指针上实现的，特殊到一般
	box = &door // 必须要有指针，因为Open和Close方法是在对象Door的指针上实现的，特殊到一般
	op.Open()   // Open the door!
	cl.Close()  // Close the door!
	box.Open()  // Open the door!
	box.Close() // Close the door!
	fmt.Println("open:", op, "close:", cl)
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

type Refrigerator struct {
	Size int
}

// Refrigerator 实现了Open接口、Close接口，同时也实现了Box接口
func (Refrigerator) Open() error {
	fmt.Println("Open!")
	return nil
}

func (Refrigerator) Close() error {
	fmt.Println("Close!")
	return nil
}

// Door 在指针上实现了Open接口、Close接口，同时也实现了Box接口
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
