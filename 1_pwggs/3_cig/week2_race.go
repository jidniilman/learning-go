package main

import (
	"fmt"
	"time"
)

var x int = 0

func main() {
	// Race condition is a situation when the two or more concurrency (or goroutines in golang) do read/write/other operation
	// to the same shared data/variable at the same time. This condition made some interleavings that expected or unexpected.
	fmt.Println("This simple program is showing you how race condition in concurrency can occur.")
	fmt.Println("The program try to increment x from 0 to 6 by using 2 goroutines.")
	fmt.Println("See the code for further explanation.")
	fmt.Println("Initial x = ", x)

	// How it's different? try to remove the 'go' keyword below, and the result must be in sequential order 1 2 3 4 5 6
	// But if you execute Add() with goroutines (with go keyword), sometime the order is messed up. In my case it's often
	// 1 3 4 2 5 6
	// Why this is happen? because these Add() functions below executed asynchronously and not waiting the other one finish first
	go Add()
	go Add()
	fmt.Println("Wait 1 seconds for all goroutines to finish")
	time.Sleep(1 * time.Second)
	fmt.Println("Final x is : ", x)
}

func Add() {
	for i := 0; i < 3; i++ {
		x++
		fmt.Println("x now is : ", x)
	}
}
