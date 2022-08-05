package main

import (
	"fmt"
)

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
	fmt.Println(twoSum([]int{3, 3}, 6))
	fmt.Println(twoSum([]int{0, 4, 3, 0}, 0))
	fmt.Println(twoSum([]int{-3, 4, 3, 90}, 0))
}

func twoSum(nums []int, target int) []int {
	for i := range nums {
		for j := range nums {
			if nums[i]+nums[j] == target && i != j {
				return []int{i, j}
			}
		}
	}
	return []int{}
}
