package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Runner struct {
	name string
}

func (r Runner) Run(wg *sync.WaitGroup) {
	defer wg.Done()
	start := time.Now()
	fmt.Println(r.name, "start running")
	rand.Seed(time.Now().UnixNano())
	time.Sleep(time.Duration(rand.Uint64()%10) * time.Second)
	finish := time.Now()
	fmt.Println(r.name, "finish line! Time:", finish.Sub(start))
}

func main() {
	runnerCount := 10
	runners := []Runner{}
	wg := sync.WaitGroup{}
	wg.Add(runnerCount)

	for i := 0; i < runnerCount; i++ {
		runners = append(runners, Runner{
			name: fmt.Sprintf("%d", i),
		})
	}

	for _, runner := range runners {
		go runner.Run(&wg)
	}
	wg.Wait()
	fmt.Println("END")
}
