package main

import "fmt"

func main() {
	// 实例化一个elevatorDeploymentService
	elevatorDeploymentSvc := &elevatorDeploymentService{d: GetElevatorDeployment()}

	e2 := getFakeElevatorInfoFromInput2()
	fmt.Println(elevatorDeploymentSvc.DeployElevator(e2))
	e3 := getFakeElevatorInfoFromInput3()
	fmt.Println(elevatorDeploymentSvc.DeployElevator(e3))
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
	var ans1 string
	var targetFloor int
	for {
		fmt.Print("目标楼层：")
		fmt.Scanf("%d\n", &targetFloor)
		targetFloors = append(targetFloors, targetFloor)
		fmt.Print("继续录入？(输入n停止录入，其它任意键继续)")
		fmt.Scanln(&ans1)
		if ans1 == "n" {
			break
		}
	}

	// 录入电梯的运行规则选择
	var upgrade bool
	var ans2 string
	fmt.Printf("是否按照改进后的规则运行？(输入n按照原有规则运行，其它任意键默认按照改进后的规则运行)")
	fmt.Scanln(&ans2)
	if ans2 == "n" {
		upgrade = false
	} else {
		upgrade = true
	}

	return &Elevator{
		totalFloors:  totalFloors,
		currentFloor: currentFloor,
		targetFloors: targetFloors,
		upgrade:      upgrade,
	}
}

func getFakeElevatorInfoFromInput0() *Elevator {
	return &Elevator{
		totalFloors:  7,
		currentFloor: 3,
		targetFloors: []int{2, 1, 5, 4, 6},
		upgrade:      false,
	}
}

func getFakeElevatorInfoFromInput1() *Elevator {
	return &Elevator{
		totalFloors:  7,
		currentFloor: 3,
		targetFloors: []int{2, 1, 5, 4, 6},
		upgrade:      true,
	}
}

func getFakeElevatorInfoFromInput2() *Elevator {
	return &Elevator{
		totalFloors:  7,
		currentFloor: 5,
		targetFloors: []int{7, 2, 3, 4, 6},
		upgrade:      false,
	}
}

func getFakeElevatorInfoFromInput3() *Elevator {
	return &Elevator{
		totalFloors:  7,
		currentFloor: 5,
		targetFloors: []int{7, 2, 3, 4, 6},
		upgrade:      true,
	}
}

func getFakeElevatorInfoFromInput4() *Elevator {
	return &Elevator{
		totalFloors:  5,
		currentFloor: 4,
		targetFloors: []int{2},
	}
}
