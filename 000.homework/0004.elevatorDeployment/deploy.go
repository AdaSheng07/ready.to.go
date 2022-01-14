package main

import (
	"fmt"
	"log"
	"time"
)

func GetElevatorDeployment() *Deploy {
	return &Deploy{}
}

type Deploy struct {
	finalFloor     int           // 电梯最终停靠的楼层
	timeDuration   time.Duration // 电梯整体运行时间统计：随电梯运行Operation()不断计算更新
	deployStrategy string        // 电梯部署的整体方案打印
}

// Operation 根据电梯目标楼层切片，部署电梯的运行顺序，打印电梯部署的整体方案
func (d *Deploy) Operation(elevator *Elevator) {
	d.deployStrategy = "" // 初始化 deployStrategy 为空字符串
	switch {
	case len(elevator.targetFloors) == 0:
		// 无人按电梯：没有目标楼层，targetFloors切片为空，电梯所在楼层不变
		_ = d.noRequest(elevator)
	case len(elevator.targetFloors) >= 1:
		// 有人按电梯：需要规划电梯运行的优先方向和到达目标楼层的顺序：orderOfDockedFloors
		_ = d.request(elevator)
	}
}

// request 有人使用电梯，电梯遵循特定的规则运行，需要先排序目标楼层
func (d *Deploy) request(elevator *Elevator) (err error) {
	err = elevator.GetOrderOfDockedFloors(elevator) // 根据电梯的初始化信息，规划电梯应该以 GetOrderOfDockedFloors(elevator) 更新得到的顺序依次到达各目标楼层，如有错误，返回
	if err != nil {
		log.Println("规划电梯停靠的楼层顺序时出错：", err)
		return err
	}
	sumOfFloors := 0 // 初始化电梯走过的层数之和 sumOfFloors
	// 电梯开始运行
	d.deployStrategy = fmt.Sprintf("电梯在 %d 层，开始运行\n", elevator.currentFloor) // 将电梯工作日志存入d.deployStrategy
	// 电梯按照规定的顺序依次到达目标楼层
	for _, floor := range elevator.orderOfDockedFloors {
		arrivingAtTargetFloor(elevator.currentFloor, floor)          // 从当前楼层到目标楼层所需要等待的时间
		sumOfFloors += abs(floor - elevator.currentFloor)            // 计算电梯走过的层数之和
		d.deployStrategy += fmt.Sprintf("电梯停靠在 %d 层，开门，关门\n", floor) // 将电梯工作日志存入d.deployStrategy
		elevator.currentFloor = floor                                // 将电梯目前所在楼层 elevator.currentFloor 更新为刚刚到达的目标楼层 floor
	}
	d.finalFloor = elevator.currentFloor                                                            // 将电梯最终所在楼层存入d.finalFloor
	d.timeDuration = time.Second * time.Duration(sumOfFloors)                                       // 将走完所有目标楼层所需时间存入d.timeDuration
	d.deployStrategy += fmt.Sprintf("最终电梯停靠在 %d 层，共用时 %v\n", elevator.currentFloor, d.timeDuration) // 将电梯工作日志存入d.deployStrategy
	return nil
}

// noRequest 无人使用电梯，电梯停靠层不变
func (d *Deploy) noRequest(elevator *Elevator) (err error) {
	err = elevator.GetOrderOfDockedFloors(elevator) // 根据电梯的初始化信息，规划电梯应该以 GetOrderOfDockedFloors(elevator) 更新得到的顺序依次到达各目标楼层，如有错误，返回
	if err != nil {
		log.Println("规划电梯停靠的楼层顺序时出错：", err)
		return err
	}
	d.timeDuration = time.Second * time.Duration(0) // 无人使用电梯，电梯运行时间为0
	d.finalFloor = elevator.currentFloor            // 将电梯最终所在楼层存入d.finalFloor
	d.deployStrategy = fmt.Sprintf("无电梯请求，电梯不动，共用时 %v\n", d.timeDuration)
	return nil
}

// arrivingAtTargetFloor 从当前楼层到目标楼层所需要等待的时间（秒）
func arrivingAtTargetFloor(currentFloor, floor int) {
	// 电梯每移动一层，需要一秒，计算当前楼层与目标楼层之间的层数，层数 * (1 秒) = 当前楼层到目标楼层的耗时
	time.Sleep(time.Second * time.Duration(abs(floor-currentFloor)))
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
