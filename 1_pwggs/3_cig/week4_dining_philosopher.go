package main

import (
	"fmt"
	"sync"
)

type ChopStick struct {
	sync.Mutex
	number int
}

type Philosopher struct {
	number          int
	numberEat       int
	leftCS, rightCS *ChopStick
}

var wg sync.WaitGroup

func main() {
	// Make 5 Chopstick
	chopSticks := make([]*ChopStick, 5)
	for i := 0; i < 5; i++ {
		chopSticks[i] = &ChopStick{
			Mutex:  sync.Mutex{},
			number: i,
		}
	}
	// Make 5 Philosopher (numbered) and shared adjacent pair of chopstick
	philosophers := make([]*Philosopher, 5)
	for i := 0; i < 5; i++ {
		philosophers[i] = &Philosopher{
			number:  i,
			leftCS:  chopSticks[i],
			rightCS: chopSticks[(i+1)%5],
		}
	}
	// Start Eating 15 time (each philosopher eat 3 times), with only 2 concurrent philosopher can eat at same time
	// Ask permission from host (main goroutine)
	limit := make(chan int, 2)
	wg.Add(15)
	for i := 0; i < 5; i++ {
		go philosophers[i].eat(limit)
	}
	wg.Wait()
	fmt.Println("Eating Finished")
}

func (p Philosopher) eat(limit chan int) {
	for {
		// Only can eat 3 times, philosopher pick chopstick in any order
		if p.numberEat < 3 {
			// Add limit guard (only 2 philosopher can eat at same time)
			limit <- 1
			p.leftCS.Lock()
			p.rightCS.Lock()
			fmt.Printf("Philosopher %d Start Eating Plate %d \n", p.number+1, p.numberEat+1)
			fmt.Printf("Philosopher %d Finished Eating Plate %d \n", p.number+1, p.numberEat+1)
			p.rightCS.Unlock()
			p.leftCS.Unlock()
			p.numberEat++
			// Free limit
			<-limit
			wg.Done()
		}
	}
}
