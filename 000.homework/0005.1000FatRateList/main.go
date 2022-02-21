package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"sync"
	"time"
)

type Person struct {
	Name    string
	FatRate float64
	Rank    int
	lock    sync.Mutex
	pCond   sync.Cond
}

type FatRateList struct {
	FatRates []float64
	lock     sync.Mutex
}

//func (p *Person) GetRank(name string, persons []Person, fatRates []float64) (rank int) {
//	p.lock.Lock()
//	defer p.lock.Unlock()
//	sort.Float64s(fatRates)
//	i, _ := strconv.Atoi(name)
//	return sort.SearchFloat64s(fatRates, persons[i].FatRate)
//}

func main() {

	// 初始化数据
	personCount := 1000 // 一共有1000人
	persons := []Person{}
	fatRateList := &FatRateList{
		FatRates: []float64{},
	} // 1000人的体脂率

	// 生成随机数base，base范围在[0, 0.4]
	rand.Seed(time.Now().UnixNano())
	base := rand.Float64() * 0.4

	for i := 0; i < personCount; i++ {
		persons = append(persons, Person{
			Name: fmt.Sprintf("%d", i),
		})
		for {
			rand.Seed(time.Now().UnixNano())
			delta := rand.Float64()*0.4 - 0.2 // 体脂率波动范围[-0.2, 0.2]
			// 保证体脂率是非负数，防止非法输入
			if base+delta >= 0 {
				persons[i].FatRate = base + delta
				fatRateList.FatRates = append(fatRateList.FatRates, base+delta)
				break
			}
		}
	}

	fmt.Println("查询排名")
	count := 1000
	for c := 0; c < count; c++ {
		// 查询排名
		go func(c int) {
			fmt.Println(persons[c].Name, "号开始查询排名：")
			persons[c].lock.Lock()
			fatRateList.lock.Lock()
			sort.Float64s(fatRateList.FatRates)
			fatRateList.lock.Unlock()
			index, _ := strconv.Atoi(persons[c].Name)
			persons[c].Rank = sort.SearchFloat64s(fatRateList.FatRates, persons[index].FatRate)
			fmt.Println(persons[c].Name, "号的体脂率排名为：", persons[c].Rank)
			persons[c].lock.Unlock()
		}(c)
	}

	fmt.Println("更新并打印排名")
	for c := 0; c < count; c++ {
		// 更新并打印排名
		go func(c int) {
			fmt.Println(persons[c].Name, "号开始更新排名：")
			persons[c].lock.Lock()
			fatRateList.lock.Lock()
			for {
				rand.Seed(time.Now().UnixNano())
				delta := rand.Float64()*0.4 - 0.2 // 体脂率波动范围[-0.2, 0.2]
				// 保证体脂率是非负数，防止非法输入
				if base+delta >= 0 {
					persons[c].FatRate = base + delta
					fatRateList.FatRates[persons[c].Rank] = base + delta
					break
				}
			}
			sort.Float64s(fatRateList.FatRates)
			fatRateList.lock.Unlock()
			index, _ := strconv.Atoi(persons[c].Name)
			persons[c].Rank = sort.SearchFloat64s(fatRateList.FatRates, persons[index].FatRate)
			fmt.Println(persons[c].Name, "号的体脂率排名为：", persons[c].Rank)
			persons[c].lock.Unlock()
		}(c)
	}

	time.Sleep(5 * time.Second)
}
