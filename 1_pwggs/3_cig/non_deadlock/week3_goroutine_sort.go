package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

var wg sync.WaitGroup

func main() {
	var inputString string

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("Input your Integer Sequence - Ex. 1 10 2 9 -5 2 -4 3 8 7 0 6 5 11 13 12 : ")
	if scanner.Scan() {
		inputString = scanner.Text()
	}

	intSlice := GetIntSlice(inputString)
	fmt.Println("Initial Integers : ", intSlice)

	// Make 4 Partitions to split the job
	length := int(math.Ceil(float64(len(intSlice)) / float64(4)))
	slice1 := intSlice[0:length]
	slice2 := intSlice[length : 2*length]
	slice3 := intSlice[2*length : 3*length]
	slice4 := intSlice[3*length:]

	fmt.Println("Initial Partition 1 : ", slice1)
	fmt.Println("Initial Partition 2 : ", slice2)
	fmt.Println("Initial Partition 3 : ", slice3)
	fmt.Println("Initial Partition 4 : ", slice4)

	wg.Add(4)
	go Sort(1, slice1)
	go Sort(2, slice2)
	go Sort(3, slice3)
	go Sort(4, slice4)

	wg.Wait()

	fmt.Println("Merged Integers from 4 Partitions : ", intSlice)
	sort.Sort(sort.IntSlice(intSlice))
	fmt.Println("Sorted Integers from merged 4 Partitions : ", intSlice)
}

func GetIntSlice(str string) []int {
	ints := make([]int, 0)
	for _, char := range strings.Fields(str) {
		num, _ := strconv.Atoi(char)
		ints = append(ints, num)
	}
	return ints
}

func Sort(partition int, slice []int) {
	sort.Sort(sort.IntSlice(slice))
	fmt.Printf("Sorted Partition %d Finished : %v \n", partition, slice)
	wg.Done()
}
