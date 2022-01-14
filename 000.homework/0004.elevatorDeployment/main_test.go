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
		totalFloors:  5,
		currentFloor: 3,
	}
	deploy1 := Deploy{}
	deploy1.Operation(&elevator1)
	if deploy1.finalFloor != 3 {
		t.Fatalf("预期电梯不动，实际电梯停靠在 %d 层", deploy1.finalFloor)
	}
	if deploy1.timeDuration != time.Second*time.Duration(0) {
		t.Fatalf("预期电梯不动，花费 0 秒，实际电梯运行时间为：%v\n", deploy1.timeDuration)
	}
	if deploy1.deployStrategy != "无电梯请求，电梯不动，共用时 0s\n" {
		t.Fatalf("电梯部署有误，无请求时电梯应不动，实际电梯运行至：%+v 楼\n", elevator1.orderOfDockedFloors)
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
	deploy2.Operation(&elevator2)
	if deploy2.finalFloor != 3 {
		t.Fatalf("预期电梯停在 3 楼，实际电梯停靠在 %d 层", elevator2.currentFloor)
	}
	if deploy2.timeDuration != time.Second*time.Duration(2) {
		t.Fatalf("预期电梯运行花费 2 秒，实际电梯运行时间为：%v\n", deploy2.timeDuration)
	}
	if deploy2.deployStrategy != "电梯在 1 层，开始运行\n电梯停靠在 3 层，开门，关门\n最终电梯停靠在 3 层，共用时 2s\n" {
		t.Fatalf("电梯部署有误，正确运行顺序为：3，实际电梯运行顺序为：%v\n", elevator2.orderOfDockedFloors)
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
	deploy3.Operation(&elevator3)
	if deploy3.finalFloor != 2 {
		t.Fatalf("预期电梯停在 2 楼，实际电梯停靠在 %d 层", deploy3.finalFloor)
	}
	if deploy3.timeDuration != time.Second*time.Duration(3) {
		t.Fatalf("预期电梯运行花费 3 秒，实际电梯运行时间为：%v\n", deploy3.timeDuration)
	}
	if deploy3.deployStrategy != "电梯在 3 层，开始运行\n电梯停靠在 4 层，开门，关门\n电梯停靠在 2 层，开门，关门\n最终电梯停靠在 2 层，共用时 3s\n" {
		t.Fatalf("电梯部署有误，正确运行顺序为：4, 2，实际电梯运行顺序为：%v\n", elevator3.orderOfDockedFloors)
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
	deploy4.Operation(&elevator4)
	if deploy4.finalFloor != 2 {
		t.Fatalf("预期电梯停在 2 楼，实际电梯停靠在 %d 层", deploy4.finalFloor)
	}
	if deploy4.timeDuration != time.Second*time.Duration(5) {
		t.Fatalf("预期电梯运行花费 5 秒，实际电梯运行时间为：%v\n", deploy4.timeDuration)
	}
	if deploy4.deployStrategy != "电梯在 3 层，开始运行\n电梯停靠在 4 层，开门，关门\n电梯停靠在 5 层，开门，关门\n电梯停靠在 2 层，开门，关门\n最终电梯停靠在 2 层，共用时 5s\n" {
		t.Fatalf("电梯部署有误，正确运行顺序为：4, 5, 2，实际电梯运行顺序为：%v\n", elevator4.orderOfDockedFloors)
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
	deploy5.Operation(&elevator5)
	if deploy5.finalFloor != 2 {
		t.Fatalf("预期电梯停在 2 楼，实际电梯停靠在 %d 层", deploy5.finalFloor)
	}
	if deploy5.timeDuration != time.Second*time.Duration(2) {
		t.Fatalf("预期电梯运行花费 2 秒，实际电梯运行时间为：%v\n", deploy5.timeDuration)
	}
	if deploy5.deployStrategy != "电梯在 4 层，开始运行\n电梯停靠在 2 层，开门，关门\n最终电梯停靠在 2 层，共用时 2s\n" {
		t.Fatalf("电梯部署有误，正确运行顺序为：2，实际电梯运行顺序为：%v\n", elevator5.orderOfDockedFloors)
	}
}

/*
补充案例6：
楼层有5层，电梯在3层。上来一些人后，目标楼层： 4楼、2楼、5楼。
电梯先向上到4楼，然后到5楼，之后转头到2楼，最后停在2楼。
*/
func TestCase6(t *testing.T) {
	elevator6 := Elevator{
		totalFloors:  5,
		currentFloor: 3,
		targetFloors: []int{4, 2, 5},
	}
	deploy6 := Deploy{}
	deploy6.Operation(&elevator6)
	if deploy6.finalFloor != 2 {
		t.Fatalf("预期电梯停在 2 楼，实际电梯停靠在 %d 层", deploy6.finalFloor)
	}
	if deploy6.timeDuration != time.Second*time.Duration(5) {
		t.Fatalf("预期电梯运行花费 5 秒，实际电梯运行时间为：%v\n", deploy6.timeDuration)
	}
	if deploy6.deployStrategy != "电梯在 3 层，开始运行\n电梯停靠在 4 层，开门，关门\n电梯停靠在 5 层，开门，关门\n电梯停靠在 2 层，开门，关门\n最终电梯停靠在 2 层，共用时 5s\n" {
		t.Fatalf("电梯部署有误，正确运行顺序为：4, 5, 2，实际电梯运行顺序为：%v\n", elevator6.orderOfDockedFloors)
	}
}

/*
补充案例7：
楼层有7层，电梯在3层。上来一些人后，目标楼层： 2楼、1楼、5楼、4楼、6楼。
电梯先向下到2楼，然后到1楼，之后转头到4楼，5楼，6楼，最后停在6楼。
*/
func TestCase7(t *testing.T) {
	elevatorDeploymentSvc := &elevatorDeploymentService{d: GetElevatorDeployment()}
	elevator7 := getFakeElevatorInfoFromInput1()
	elevatorDeploymentSvc.DeployElevator(elevator7)
	if elevatorDeploymentSvc.d.finalFloor != 6 {
		t.Fatalf("预期电梯停在 6 楼，实际电梯停靠在 %d 层", elevatorDeploymentSvc.d.finalFloor)
	}
	if elevatorDeploymentSvc.d.timeDuration != time.Second*time.Duration(7) {
		t.Fatalf("预期电梯运行花费 7 秒，实际电梯运行时间为：%v\n", elevatorDeploymentSvc.d.timeDuration)
	}
	if elevatorDeploymentSvc.d.deployStrategy != "电梯在 3 层，开始运行\n电梯停靠在 2 层，开门，关门\n电梯停靠在 1 层，开门，关门\n电梯停靠在 4 层，开门，关门\n电梯停靠在 5 层，开门，关门\n电梯停靠在 6 层，开门，关门\n最终电梯停靠在 6 层，共用时 7s\n" {
		t.Fatalf("电梯部署有误，正确运行顺序为2, 1, 4, 5, 6，实际电梯运行顺序为：%v\n", elevator7.orderOfDockedFloors)
	}
}

/*
补充案例8：
楼层有7层，电梯在5层。上来一些人后，目标楼层： 7楼、2楼、3楼、4楼、6楼。
电梯先向上到6楼，然后到7楼，之后转头到4楼、3楼、2楼，最后停在2楼。
*/
func TestCase8(t *testing.T) {
	elevatorDeploymentSvc := &elevatorDeploymentService{d: GetElevatorDeployment()}
	elevator8 := getFakeElevatorInfoFromInput2()
	elevatorDeploymentSvc.DeployElevator(elevator8)
	if elevatorDeploymentSvc.d.finalFloor != 2 {
		t.Fatalf("预期电梯停在 2 楼，实际电梯停靠在 %d 层", elevatorDeploymentSvc.d.finalFloor)
	}
	if elevatorDeploymentSvc.d.timeDuration != time.Second*time.Duration(7) {
		t.Fatalf("预期电梯运行花费 7 秒，实际电梯运行时间为：%v\n", elevatorDeploymentSvc.d.timeDuration)
	}
	if elevatorDeploymentSvc.d.deployStrategy != "电梯在 5 层，开始运行\n电梯停靠在 6 层，开门，关门\n电梯停靠在 7 层，开门，关门\n电梯停靠在 4 层，开门，关门\n电梯停靠在 3 层，开门，关门\n电梯停靠在 2 层，开门，关门\n最终电梯停靠在 2 层，共用时 7s\n" {
		t.Fatalf("电梯部署有误，正确运行顺序为6, 7, 4, 3, 2，实际电梯运行顺序为：%v\n", elevator8.orderOfDockedFloors)
	}
}
