package main

import "fmt"

func main() {
	//e := getElevatorInfoFromInput()
	//deploy := Deploy{}
	//deploy.Operation(e)

	// 实例化一个elevatorDeploymentService
	elevatorDeploymentSvc := &elevatorDeploymentService{d: GetElevatorDeployment()}
	//e := getFakeElevatorInfoFromInput()
	//elevatorDeploymentSvc.DeployElevator(e)
	elevator := getElevatorInfoFromInput()
	elevatorDeploymentSvc.DeployElevator(elevator)
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
		fmt.Print("继续录入？(y/n)")
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

func getFakeElevatorInfoFromInput() *Elevator {
	return &Elevator{
		totalFloors:  7,
		currentFloor: 3,
		targetFloors: []int{2, 1, 5, 4, 6},
	}
}
