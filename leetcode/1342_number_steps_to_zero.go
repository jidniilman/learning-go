package main

import "fmt"

func main() {
	fmt.Println(numberOfSteps(14))
	fmt.Println(numberOfSteps(8))
	fmt.Println(numberOfSteps(123))
}

func numberOfSteps(num int) int {
	counter := 0
	for num != 0 {
		if num%2 == 0 {
			counter++
			num /= 2
		}
		if num%2 == 1 {
			counter++
			num -= 1
		}
	}
	return counter
}
