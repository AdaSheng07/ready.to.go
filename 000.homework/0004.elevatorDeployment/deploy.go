package main

import (
	"fmt"
	"time"
)

func GetElevatorDeployment() *Deploy {
	return &Deploy{}
}

type Deploy struct {
	finalFloor   int           // 电梯最终停靠的楼层
	timeStamp    string        // 电梯运行时间戳：随电梯运行Operation()不断更新
	timeDuration time.Duration // 电梯整体运行时间统计：随电梯运行Operation()不断计算更新
}

// Operation 根据电梯目标楼层切片，部署电梯的运行顺序，返回电梯最终停留的楼层
func (d *Deploy) Operation(elevator *Elevator) (int, time.Duration) {
	// 无人按电梯：如果没有目标楼层，targetFloors切片为空，电梯所在楼层不变
	if len(elevator.targetFloors) == 0 {
		d.finalFloor, d.timeDuration = d.noRequest(elevator)
	}
	// 一人按电梯：如果目标楼层只有一个，直接运行，不用考虑电梯运行方向
	if len(elevator.targetFloors) == 1 {
		d.finalFloor, d.timeDuration = d.oneRequest(elevator)
	}
	// 多人按电梯：如果目标楼层超过两个，需要规划电梯运行的优先方向和到达目标楼层的顺序：orderOfDockedFloors
	if len(elevator.targetFloors) >= 2 {
		d.finalFloor, d.timeDuration = d.multipleRequests(elevator)
	}
	return elevator.currentFloor, d.timeDuration
}

// multipleRequests 多人按电梯时候，遵循特定的规则运行
func (d *Deploy) multipleRequests(elevator *Elevator) (int, time.Duration) {
	// 根据电梯的初始化信息，规划电梯应该以 GetOrderOfDockedFloors(elevator) 更新得到的顺序依次到达各目标楼层
	elevator.GetOrderOfDockedFloors(elevator)
	// 电梯开始运行前，设定时间戳，初始化层数和sumOfFloors
	d.timeStamp = time.Now().Format("2006-01-02 15:04:05")
	sumOfFloors := 0
	fmt.Printf("%s: 电梯在 %d 层，开始运行\n", d.timeStamp, elevator.currentFloor)
	// 电梯按照规定的顺序依次到达目标楼层
	for _, floor := range elevator.orderOfDockedFloors {
		// 电梯每移动一层，需要一秒，计算当前楼层与目标楼层之间的层数，层数 * (1 秒) = 当前楼层到目标楼层的耗时
		time.Sleep(time.Second * time.Duration(abs(floor-elevator.currentFloor)))
		// 计算电梯走过的层数之和
		sumOfFloors += abs(floor - elevator.currentFloor)
		// 更新到达目标楼层之后的时间戳
		d.timeStamp = time.Now().Format("2006-01-02 15:04:05")
		// 打印电梯的工作日志
		fmt.Printf("%s: 电梯停靠在 %d 层，开门，关门\n", d.timeStamp, floor)
		// 电梯目前所在楼层elevator.currentFloor更新为刚刚到达的目标楼层floor
		elevator.currentFloor = floor
	}
	// 将走完所有目标楼层所需时间存入d.timeDuration
	d.timeDuration += time.Second * time.Duration(sumOfFloors)
	// 打印电梯的工作日志
	fmt.Printf("%s: 最终电梯停靠在 %d 层，共用时 %v\n", d.timeStamp, elevator.currentFloor, d.timeDuration)
	return elevator.currentFloor, d.timeDuration
}

// oneRequest 只有一层需要使用电梯，电梯直接到达目标楼层
func (d *Deploy) oneRequest(elevator *Elevator) (int, time.Duration) {
	// 电梯开始运行前，设定时间戳
	d.timeStamp = time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("%s: 电梯在 %d 层，开始运行\n", d.timeStamp, elevator.currentFloor)
	// 电梯每移动一层，需要一秒，计算当前楼层与目标楼层之间的层数，层数 * (1 秒) = 当前楼层到目标楼层的耗时
	time.Sleep(time.Second * time.Duration(abs(elevator.targetFloors[0]-elevator.currentFloor)))
	// 将到达目标楼层所需时间存入d.timeDuration
	d.timeDuration = time.Second * time.Duration(abs(elevator.targetFloors[0]-elevator.currentFloor))
	// 更新到达目标楼层之后的时间戳
	d.timeStamp = time.Now().Format("2006-01-02 15:04:05")
	// 电梯目前所在楼层elevator.currentFloor更新为到达的目标楼层elevator.targetFloors[0]
	elevator.currentFloor = elevator.targetFloors[0]
	// 打印电梯的工作日志
	fmt.Printf("%s: 电梯停靠在 %d 层，开门，关门\n", d.timeStamp, elevator.targetFloors[0])
	fmt.Printf("%s: 最终电梯停靠在 %d 层，共用时 %v\n", d.timeStamp, elevator.currentFloor, d.timeDuration)
	// 返回电梯当前停靠楼层和运行时间
	return elevator.currentFloor, d.timeDuration
}

// noRequest 无人使用电梯，电梯停靠层不变
func (d *Deploy) noRequest(elevator *Elevator) (int, time.Duration) {
	// 电梯开始运行前，设定时间戳
	d.timeStamp = time.Now().Format("2006-01-02 15:04:05")
	// 无人使用电梯，电梯运行时间为0
	d.timeDuration = time.Second * time.Duration(0)
	// 打印电梯的工作日志
	fmt.Printf("%s: 无电梯请求，电梯不动，共用时 %v\n", d.timeStamp, d.timeDuration)
	// 返回电梯当前停靠楼层和运行时间
	return elevator.currentFloor, d.timeDuration
}

// abs 用于计算绝对值，e.g. 当前所在楼层与目标楼层的差距
func abs(i int) int {
	result := i
	if i < 0 {
		result *= -1
		return result
	}
	return result
}
