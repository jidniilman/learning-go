package main

import (
	"fmt"
	"math"
)

func main() {
	var inputAcceleration, inputVelocity, inputInitialDisplacement, time float64

	fmt.Printf("Input your Acceleration (a) (float64) : ")
	fmt.Scan(&inputAcceleration)

	fmt.Printf("Input your Initial Velocity (Vo) (float64) : ")
	fmt.Scan(&inputVelocity)

	fmt.Printf("Input your Initial Displacement (So) (float64) : ")
	fmt.Scan(&inputInitialDisplacement)

	fmt.Printf("Input your Time (t) (float64) : ")
	fmt.Scan(&time)

	fn := GenDisplaceFn(inputAcceleration, inputVelocity, inputInitialDisplacement)

	fmt.Println(fn(time))
}

func GenDisplaceFn(acceleration, velocity, initialDisplacement float64) func(float64) float64 {
	return func(time float64) (displacement float64) {
		displacement = (0.5 * acceleration * math.Pow(time, 2)) + (velocity * time) + initialDisplacement
		return
	}
}
