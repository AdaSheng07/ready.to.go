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
}

func main() {

	// 初始化数据
	personCount := 1000 // 一共有1000人
	persons := []Person{}
	fatRates := []float64{}
	fatRatesLock := sync.RWMutex{}

	//fatRateList := &FatRateList{
	//	FatRates: []float64{},}
	// 1000人的体脂率

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
				fatRates = append(fatRates, persons[i].FatRate)
				break
			}
		}
	}

	//fmt.Println("查询排名")
	count := 1000
	// 查询排名
	go func() {
		for c := 0; c < count; c++ {
			fmt.Println(persons[c].Name, "号开始查询排名：")
			time.Sleep(1 * time.Millisecond)
			persons[c].lock.Lock()
			GetRank(c, fatRates, &fatRatesLock, persons)
			persons[c].lock.Unlock()
		}
	}()

	//fmt.Println("更新并打印排名")
	// 更新并打印排名
	go func() {
		for c := 0; c < count; c++ {
			fmt.Println(persons[c].Name, "号开始更新排名：")
			time.Sleep(1 * time.Millisecond)
			persons[c].lock.Lock()
			for {
				rand.Seed(time.Now().UnixNano())
				delta := rand.Float64()*0.4 - 0.2 // 体脂率波动范围[-0.2, 0.2]
				// 保证体脂率是非负数，防止非法输入
				if base+delta >= 0 {
					persons[c].FatRate = base + delta
					fatRates[persons[c].Rank] = base + delta
					break
				}
			}
			GetRank(c, fatRates, &fatRatesLock, persons)
			persons[c].lock.Unlock()
		}
	}()

	time.Sleep(10 * time.Second)
}

func GetRank(c int, fatRates []float64, fatRatesLock *sync.RWMutex, persons []Person) {
	fatRatesLock.Lock()
	sort.Float64s(fatRates)
	fatRatesLock.Unlock()
	fatRatesLock.RLock()
	index, _ := strconv.Atoi(persons[c].Name)
	persons[c].Rank = sort.SearchFloat64s(fatRates, persons[index].FatRate)
	fatRatesLock.RUnlock()
	fmt.Println(persons[c].Name, "号的体脂率排名为：", persons[c].Rank)
}
