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

func (r Runner) Run(startRacing *sync.WaitGroup, wg *sync.WaitGroup) {
	defer wg.Done()
	startRacing.Wait()
	start := time.Now()
	fmt.Println(r.name, "start running", start)
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

	startRacingWG := sync.WaitGroup{}
	startRacingWG.Add(1)

	for i := 1; i <= runnerCount; i++ {
		runners = append(runners, Runner{
			name: fmt.Sprintf("%d", i),
		})
	}

	for _, runner := range runners {
		go runner.Run(&startRacingWG, &wg)
	}

	fmt.Println("Ready...")
	time.Sleep(1 * time.Second)
	fmt.Println("GO!!!")

	startRacingWG.Done()
	wg.Wait()
	fmt.Println("END")
}
