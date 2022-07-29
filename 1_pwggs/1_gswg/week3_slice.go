package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	// Declare Integer Slice with capacity 3
	sliceInteger := make([]int, 0, 3)

	// Loop until 'X'
	for {
		var input string
		fmt.Printf("Input your Integer, or X to quit : ")
		fmt.Scan(&input)

		// If X then stop
		if input == "X" {
			fmt.Println("Final Sorted Slice : ", sliceInteger)
			break
		}

		integer, err := strconv.Atoi(input)

		// If not integer, then skip loop
		if err != nil {
			fmt.Println("Not an Integer or X, Skipping.")
			continue
		}

		sliceInteger = append(sliceInteger, integer)
		// Sort the slice Ascending
		sort.Sort(sort.IntSlice(sliceInteger))
		fmt.Println("Sorted Slice : ", sliceInteger)
	}
}
