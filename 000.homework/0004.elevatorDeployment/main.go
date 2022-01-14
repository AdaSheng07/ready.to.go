package main

import "fmt"

func main() {
	// 实例化一个elevatorDeploymentService
	elevatorDeploymentSvc := &elevatorDeploymentService{d: GetElevatorDeployment()}
	e := getFakeElevatorInfoFromInput3()
	// e := getElevatorInfoFromInput()
	fmt.Println(elevatorDeploymentSvc.DeployElevator(e))
}

func getElevatorInfoFromInput() *Elevator {
	// 录入总楼层数
	var totalFloors int
	fmt.Print("总楼层数：")
	fmt.Scanln(&totalFloors)

	// 录入电梯当前所在的楼层数
	var currentFloor int
	fmt.Print("电梯当前所在楼层：")
	fmt.Scanln(&currentFloor)

	// 录入电梯的目标楼层
	targetFloors := []int{}
	var ans string
	var targetFloor int
	for {
		fmt.Print("目标楼层：")
		fmt.Scanf("%d\n", &targetFloor)
		targetFloors = append(targetFloors, targetFloor)
		fmt.Print("继续录入？(输入n停止录入，其它任意键继续)")
		fmt.Scanln(&ans)
		if ans == "n" {
			break
		}
	}
	return &Elevator{
		totalFloors:  totalFloors,
		currentFloor: currentFloor,
		targetFloors: targetFloors,
	}
}

func getFakeElevatorInfoFromInput1() *Elevator {
	return &Elevator{
		totalFloors:  7,
		currentFloor: 3,
		targetFloors: []int{2, 1, 5, 4, 6},
	}
}

func getFakeElevatorInfoFromInput2() *Elevator {
	return &Elevator{
		totalFloors:  7,
		currentFloor: 5,
		targetFloors: []int{7, 2, 3, 4, 6},
	}
}

func getFakeElevatorInfoFromInput3() *Elevator {
	return &Elevator{
		totalFloors:  5,
		currentFloor: 4,
		targetFloors: []int{2},
	}
}
