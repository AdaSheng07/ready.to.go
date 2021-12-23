package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	deferGuess()
	openFile()
}

// defer pitfall
// given what we already know about defer function, the code below should have duration about 5 seconds
// however, the result is in nanosecond level, why?
func deferGuess() {
	startTime := time.Now()
	fmt.Println("start time:", startTime)
	defer fmt.Println("duration in the middle: ", time.Now().Sub(startTime)) // in nanosecond level
	defer func() {
		fmt.Println("use closure to calculate duration: ", time.Now().Sub(startTime)) // about 5 seconds
	}()
	time.Sleep(time.Second * 5) // 5 seconds
	fmt.Println("finish time:", time.Now())
	defer fmt.Println("duration at the end: ", time.Now().Sub(startTime)) // about 5 seconds
}

// a classic usage of defer function: file operation, open and close
func openFile() {
	filename := "/test.txt"
	fileObj, err := os.Open(filename)
	if err != nil {
		fmt.Println("error:", err) // error: open /test.txt: no such file or directory
		os.Exit(1)                 // exit the procedure
	}
	// defer + function, this function is bounded with fileObj
	defer fileObj.Close()
}
