package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	question2()
	fmt.Println("------------------------------")
	question3()
	fmt.Println("------------------------------")
	question4()
	fmt.Println()
	fmt.Println("------------------------------")
	question5()
	fmt.Println("------------------------------")
}

func question2() {
	i, _ := strconv.Atoi("10")
	y := i * 2
	fmt.Println(y)
}

func question3() {
	s := strings.Replace("ianianian", "ni", "in", 2)
	fmt.Println(s)
}

func question4() {
	x := 7
	switch {
	case x > 3:
		fmt.Printf("1")
	case x > 5:
		fmt.Printf("2")
	case x == 7:
		fmt.Printf("3")
	default:
		fmt.Printf("4")
	}
}

func question5() {
	var xtemp int
	x1 := 0
	x2 := 1
	for x := 0; x < 5; x++ {
		xtemp = x2
		x2 = x2 + x1
		x1 = xtemp
	}
	fmt.Println(x2)
}

func question6() {
	// var x int
	// var y *int
	// z := 3
	// y = &z
	// x = &y
	// result : compiler error, x cannot have &&z
}
