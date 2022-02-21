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
	lock    sync.RWMutex
}

func main() {

	// 初始化数据
	personCount := 1000 // 一共有1000人
	persons := []Person{}
	fatRates := []float64{}
	fatRatesLock := sync.RWMutex{}

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
	wgLookup := sync.WaitGroup{}
	wgLookup.Add(count)
	wgUpdate := sync.WaitGroup{}
	wgUpdate.Add(count)
	// 查询排名
	for c := 0; c < count; c++ {
		go func(p []Person, fr []float64, lock *sync.RWMutex, c int) {
			defer wgLookup.Done()
			fmt.Println(p[c].Name, "号开始查询排名：")
			GetRank(c, fr, lock, p)
		}(persons, fatRates, &fatRatesLock, c)
	}

	//fmt.Println("更新并打印排名")
	// 更新并打印排名
	for c := 0; c < count; c++ {
		go func(p []Person, fr []float64, lock *sync.RWMutex, c int) {
			defer wgUpdate.Done()
			fmt.Println(persons[c].Name, "号开始更新排名：")
			for {
				rand.Seed(time.Now().UnixNano())
				delta := rand.Float64()*0.4 - 0.2 // 体脂率波动范围[-0.2, 0.2]
				// 保证体脂率是非负数，防止非法输入
				if base+delta >= 0 {
					persons[c].lock.Lock()
					persons[c].FatRate = base + delta
					persons[c].lock.Unlock()
					fatRatesLock.Lock()
					fatRates[persons[c].Rank] = base + delta
					fatRatesLock.Unlock()
					break
				}
			}
			GetRank(c, fatRates, &fatRatesLock, persons)
		}(persons, fatRates, &fatRatesLock, c)
	}
	wgLookup.Wait()
	wgUpdate.Wait()
	fmt.Println("end")

}

func GetRank(c int, fatRates []float64, fatRatesLock *sync.RWMutex, persons []Person) {
	fatRatesLock.Lock()
	sort.Float64s(fatRates)
	fatRatesLock.Unlock()
	fatRatesLock.RLock()
	index, _ := strconv.Atoi(persons[c].Name)
	persons[c].lock.Lock()
	persons[c].Rank = sort.SearchFloat64s(fatRates, persons[index].FatRate) + 1
	persons[c].lock.Unlock()
	fatRatesLock.RUnlock()
	fmt.Println(persons[c].Name, "号的体脂率排名为：", persons[c].Rank)
}
