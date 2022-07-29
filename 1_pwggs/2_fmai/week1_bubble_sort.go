package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var inputString string

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("Input your 10 Integer Sequence - Ex. 1 10 2 9 -5 2 -4 3 8 7 : ")
	if scanner.Scan() {
		inputString = scanner.Text()
	}

	intSlice := GetIntSlice(inputString)
	fmt.Println("Initial Integer Slice : ", intSlice)
	BubbleSort(intSlice)
	fmt.Println("Sorted Integer Slice : ", intSlice)
}

func GetIntSlice(str string) []int {
	ints := make([]int, 0)
	for _, char := range strings.Fields(str) {
		num, _ := strconv.Atoi(char)
		ints = append(ints, num)
	}
	return ints
}

func BubbleSort(slice []int) {
	for i := 0; i < len(slice)-1; i++ {
		// fmt.Println("Iteration ", i+1)
		for j := 0; j < len(slice)-i-1; j++ {
			// fmt.Printf("Compare %d with %d\n", slice[j], slice[j+1])
			if slice[j] > slice[j+1] {
				Swap(slice, j)
			}
		}
	}
}

func Swap(slice []int, index int) {
	slice[index], slice[index+1] = slice[index+1], slice[index]
}
