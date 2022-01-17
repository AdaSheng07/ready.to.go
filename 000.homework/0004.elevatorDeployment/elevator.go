package main

import (
	"fmt"
	"sort"
)

type Elevator struct {
	totalFloors  int   // 电梯总层数
	currentFloor int   // 电梯当前所在楼层
	targetFloors []int // 电梯目标楼层
	upgrade      bool  // 电梯按照默认的规则运行（false），或者按照改进后的规则运行（true）

	orderOfDockedFloors []int // 电梯应该按照规则停靠目标楼层的顺序：需要 GetOrderOfDockedFloors() 来计算
}

/*
电梯当前所在楼层和电梯目标楼层切片的第一个元素决定了电梯的运行方向：上/下
第一个目标楼层 > 电梯当前所在楼层：电梯向上运行
第一个目标楼层 < 电梯当前所在楼层：电梯向下运行

在电梯向上或者向下过程中途其它楼层如果也是目标楼层，到对应楼层时暂停、开门、关门，重新继续原来方向行进
如果已经到达目的地，但还有同方向的目的楼层时，电梯维持原有方向继续行进，直到该方向最后一个楼层
到最后一个目的后，如果另一个方向还有未到达楼层，则电梯调转方向，逐个送到最后目的地，直到所有目标都送达
*/

// GetOrderOfDockedFloors 检查录入数据，利用升序/降序排列后的电梯目标楼层切片，规划电梯停靠的楼层顺序
func (e *Elevator) GetOrderOfDockedFloors(elevator *Elevator) (err error) {
	// 对电梯初始化数据进行检查
	if e.totalFloors < e.currentFloor {
		err = fmt.Errorf("电梯总楼层数量小于电梯当前所在楼层，输入有误")
		return
	}
	if len(e.targetFloors) > 0 {
		for _, floor := range e.targetFloors {
			if e.totalFloors < floor {
				err = fmt.Errorf("电梯总楼层数量小于目标楼层，输入有误")
				return
			}
			continue
		}
	}
	// 如果目标楼层切片为空，不需进行排序，直接跳出
	if len(e.targetFloors) == 0 {
		return nil
	}
	switch e.upgrade {
	case false:
		fmt.Println("按照原有规则运行电梯：")
		// 升序/降序排列电梯目标楼层切片，并确定电梯刚开始运行时的方向
		targetFloorsSorted, direction := e.setStartingDirection()
		// 根据排序后的目标楼层切片和电梯一开始的运行方向，规划电梯到达所有目标楼层的应有顺序
		e.deployOrderOfFloors(direction, targetFloorsSorted)
	case true:
		fmt.Println("按照改进后的原则运行电梯：")
		// 根据改进版电梯运行规则，升序/降序排列电梯目标楼层切片，并确定电梯刚开始运行时的方向
		targetFloorsSorted, direction := e.sortTargetFloorsUpgrade()
		// 根据排序后的目标楼层切片和电梯一开始的运行方向，规划电梯到达所有目标楼层的应有顺序
		e.deployOrderOfFloors(direction, targetFloorsSorted)
	}
	return nil
}

// deployOrderOfFloors 规划所有目标楼层的到达顺序
func (e *Elevator) deployOrderOfFloors(direction bool, targetFloorsSorted []int) {
	switch direction {
	// direction 为true时电梯向上运行，为false时电梯向下运行
	case true:
		for i, floor := range targetFloorsSorted {
			if floor <= e.currentFloor {
				continue
			} else {
				// 电梯起始时向上运行，所有楼层的到达顺序规划：
				e.orderOfDockedFloors = append(targetFloorsSorted[i:], sortDescending(targetFloorsSorted[:i])...)
				break
			}
		}
	case false:
		for i, floor := range targetFloorsSorted {
			if floor >= e.currentFloor {
				continue
			} else {
				// 电梯起始时向下运行，所有楼层的到达顺序规划：
				e.orderOfDockedFloors = append(targetFloorsSorted[i:], sortAscending(targetFloorsSorted[:i])...)
				break
			}
		}
	}
}

// setStartingDirection 判断电梯一开始的运行方向，并根据方向上/下，将电梯目标楼层切片的拷贝升序/降序排列
func (e *Elevator) setStartingDirection() (result []int, dir bool) {
	result = []int{}
	result = append(result, e.targetFloors...) // 做电梯目标楼层切片的拷贝
	// 判断电梯刚开始运行时的方向：
	// 如果第一个目标楼层 > 电梯当前所在楼层：电梯向上运行，将电梯目标楼层升序排列
	if e.targetFloors[0] > e.currentFloor {
		result = sortAscending(result) // 升序排列电梯目标楼层切片
		dir = true                     // 电梯向上运行
	} else {
		// 第一个目标楼层 < 电梯当前所在楼层：电梯向下运行
		result = sortDescending(result) // 降序排列电梯目标楼层切片
		dir = false                     // 电梯向下运行
	}
	// 返回升序/降序排列后的楼层切片拷贝以及电梯的初始运行方向
	return result, dir
}

// sortTargetFloorUpgrade 重新判断电梯的运行方向，改进电梯运行方向判断规则：
// 电梯根据当前的位置，优先向目标多的方向行进，直到该方向全部送达后重新根据当前的电梯位置做相同规则的选择。
// 电梯的所有目标楼层和电梯的当前所在楼层组成切片，一起从小到大排序：
// - 如果电梯当前所在楼层处于切片的前半段，即向上方的目标更多：电梯先向上运行
// - 如果电梯当前所在楼层处于切片的后半段，即向下方的目标更多：电梯先向下运行
func (e *Elevator) sortTargetFloorsUpgrade() (result []int, dir bool) {
	result = []int{}
	result = append(result, e.targetFloors...) // 做电梯目标楼层切片的拷贝
	// 将电梯当前所在楼层与目标楼层切片的中间值比较
	// 电梯当前所在楼层 <= 目标楼层切片中间值：电梯向上运行
	if e.currentFloor <= result[(len(result)/2)] {
		result = sortAscending(result) // 升序排列电梯目标楼层切片
		dir = true                     // 电梯向上运行
	} else {
		// 电梯当前所在楼层 > 目标楼层切片中间值：电梯向下运行
		result = sortDescending(result) // 降序排列电梯目标楼层切片
		dir = false                     // 电梯向下运行
	}
	// 返回升序/降序排列后的楼层切片拷贝以及电梯的初始运行方向
	return result, dir
}

// sortAscending 带有返回值的升序排列：从小到大
func sortAscending(slice []int) []int {
	sort.Ints(slice)
	return slice
}

// sortDescending 带有返回值的降序排列：从大到小
func sortDescending(slice []int) []int {
	sort.Sort(sort.Reverse(sort.IntSlice(slice)))
	return slice
}
