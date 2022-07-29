package main

import (
	"fmt"
	"sync"
	"time"
)

type ChopStick struct {
	sync.Mutex
	number int
}

type Philosopher struct {
	number          int
	leftCS, rightCS *ChopStick
}

func main() {
	// Make 5 Chopstick
	chopSticks := make([]*ChopStick, 5)
	for i := 0; i < 5; i++ {
		chopSticks[i] = &ChopStick{
			Mutex:  sync.Mutex{},
			number: i,
		}
	}
	// Make 5 Philosopher
	philosophers := make([]*Philosopher, 5)
	for i := 0; i < 5; i++ {
		philosophers[i] = &Philosopher{
			number:  i,
			leftCS:  chopSticks[i],
			rightCS: chopSticks[(i+1)%5],
		}
	}
	// Start Eating
	for i := 0; i < 5; i++ {
		go philosophers[i].eat()
	}
	time.Sleep(2000 * time.Millisecond)
}

func (p Philosopher) eat() {
	for {
		if p.leftCS.number < p.rightCS.number {
			p.leftCS.Lock()
			p.rightCS.Lock()
			fmt.Printf("1Philosopher %d Eating\n", p.number)
			p.rightCS.Unlock()
			p.leftCS.Unlock()
		}
		if p.leftCS.number > p.rightCS.number {
			p.rightCS.Lock()
			p.leftCS.Lock()
			fmt.Printf("2Philosopher %d Eating\n", p.number)
			p.leftCS.Lock()
			p.rightCS.Lock()
		}
	}
}
