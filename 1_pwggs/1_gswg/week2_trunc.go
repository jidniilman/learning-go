package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input float64
	fmt.Printf("Input your floating point : ")
	fmt.Scan(&input)
	fmt.Printf(strconv.Itoa(int(input)))
}
