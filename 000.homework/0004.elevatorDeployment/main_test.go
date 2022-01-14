package main

import (
	"testing"
	"time"
)

/*
案例1：
楼层有5层，所有电梯楼层没有人请求电梯，电梯不动
*/
func TestCase1(t *testing.T) {
	elevator1 := Elevator{
		totalFloors: 5,
	}
	deploy1 := Deploy{}
	positionOfElevator1, timeDurationOfElevator1 := deploy1.Operation(&elevator1)
	if positionOfElevator1 != elevator1.currentFloor {
		t.Fatalf("预期电梯不动，实际电梯停靠在 %d 层", positionOfElevator1)
	}
	if timeDurationOfElevator1 != time.Second*time.Duration(0) {
		t.Fatalf("预期电梯不动，花费 0 秒，实际电梯运行时间为：%v\n", timeDurationOfElevator1)
	}
}

/*
案例2：
楼层有5层，电梯在1层。三楼按电梯。电梯向三楼行进，并停在三楼。
*/
func TestCase2(t *testing.T) {
	elevator2 := Elevator{
		totalFloors:  5,
		currentFloor: 1,
		targetFloors: []int{3},
	}
	deploy2 := Deploy{}
	positionOfElevator2, timeDurationOfElevator2 := deploy2.Operation(&elevator2)
	if positionOfElevator2 != 3 {
		t.Fatalf("预期电梯停在 3 楼，实际电梯停靠在 %d 层", positionOfElevator2)
	}
	if timeDurationOfElevator2 != time.Second*time.Duration(2) {
		t.Fatalf("预期电梯运行花费 2 秒，实际电梯运行时间为：%v\n", timeDurationOfElevator2)
	}
}

/*
案例3：
楼层有5层，电梯在3层。上来一些人后，目标楼层： 4楼、2楼。
电梯先向上到4楼，然后转头到2楼，最后停在2楼。
*/
func TestCase3(t *testing.T) {
	elevator3 := Elevator{
		totalFloors:  5,
		currentFloor: 3,
		targetFloors: []int{4, 2},
	}
	deploy3 := Deploy{}
	positionOfElevator3, timeDurationOfElevator3 := deploy3.Operation(&elevator3)
	if positionOfElevator3 != 2 {
		t.Fatalf("预期电梯停在 2 楼，实际电梯停靠在 %d 层", positionOfElevator3)
	}
	if timeDurationOfElevator3 != time.Second*time.Duration(3) {
		t.Fatalf("预期电梯运行花费 3 秒，实际电梯运行时间为：%v\n", timeDurationOfElevator3)
	}
}

/*
案例4：
楼层有5层，电梯在3层。上来一些人后，目标楼层： 4楼、5楼、2楼。
电梯先向上到4楼，然后到5楼，之后转头到2楼，最后停在2楼。
*/
func TestCase4(t *testing.T) {
	elevator4 := Elevator{
		totalFloors:  5,
		currentFloor: 3,
		targetFloors: []int{4, 5, 2},
	}
	deploy4 := Deploy{}
	positionOfElevator4, timeDurationOfElevator4 := deploy4.Operation(&elevator4)
	if positionOfElevator4 != 2 {
		t.Fatalf("预期电梯停在 2 楼，实际电梯停靠在 %d 层", positionOfElevator4)
	}
	if timeDurationOfElevator4 != time.Second*time.Duration(5) {
		t.Fatalf("预期电梯运行花费 5 秒，实际电梯运行时间为：%v\n", timeDurationOfElevator4)
	}
}

/*
补充案例5：
楼层有5层，电梯在4层。2楼按电梯。电梯向2楼行进，并停在2楼。
*/
func TestCase5(t *testing.T) {
	elevator5 := Elevator{
		totalFloors:  5,
		currentFloor: 4,
		targetFloors: []int{2},
	}
	deploy5 := Deploy{}
	positionOfElevator5, timeDurationOfElevator5 := deploy5.Operation(&elevator5)
	if positionOfElevator5 != 2 {
		t.Fatalf("预期电梯停在 2 楼，实际电梯停靠在 %d 层", positionOfElevator5)
	}
	if timeDurationOfElevator5 != time.Second*time.Duration(2) {
		t.Fatalf("预期电梯运行花费 2 秒，实际电梯运行时间为：%v\n", timeDurationOfElevator5)
	}
}

/*
补充案例6：
楼层有5层，电梯在3层。上来一些人后，目标楼层： 4楼、5楼、2楼。
电梯先向上到4楼，然后到5楼，之后转头到2楼，最后停在2楼。
*/
func TestCase6(t *testing.T) {
	elevator6 := Elevator{
		totalFloors:  5,
		currentFloor: 3,
		targetFloors: []int{4, 2, 5},
	}
	deploy6 := Deploy{}
	positionOfElevator6, timeDurationOfElevator6 := deploy6.Operation(&elevator6)
	if positionOfElevator6 != 2 {
		t.Fatalf("预期电梯停在 2 楼，实际电梯停靠在 %d 层", positionOfElevator6)
	}
	if timeDurationOfElevator6 != time.Second*time.Duration(5) {
		t.Fatalf("预期电梯运行花费 5 秒，实际电梯运行时间为：%v\n", timeDurationOfElevator6)
	}
}

/*
补充案例7：
楼层有5层，电梯在3层。上来一些人后，目标楼层： 4楼、5楼、2楼。
电梯先向上到4楼，然后到5楼，之后转头到2楼，最后停在2楼。
*/
func TestCase7(t *testing.T) {
	elevator6 := Elevator{
		totalFloors:  7,
		currentFloor: 3,
		targetFloors: []int{2, 1, 5, 4, 6},
	}
	deploy7 := Deploy{}
	positionOfElevator7, timeDurationOfElevator7 := deploy7.Operation(&elevator6)
	if positionOfElevator7 != 6 {
		t.Fatalf("预期电梯停在 6 楼，实际电梯停靠在 %d 层", positionOfElevator7)
	}
	if timeDurationOfElevator7 != time.Second*time.Duration(7) {
		t.Fatalf("预期电梯运行花费 7 秒，实际电梯运行时间为：%v\n", timeDurationOfElevator7)
	}
}
