package main

import "fmt"

func main() {
	fmt.Println(maximumWealth([][]int{
		{1, 2, 3},
		{3, 2, 1}}))
	fmt.Println(maximumWealth([][]int{
		{1, 5},
		{7, 3},
		{3, 5}}))
	fmt.Println(maximumWealth([][]int{
		{2, 8, 7},
		{7, 1, 3},
		{1, 9, 5}}))
}

func maximumWealth(accounts [][]int) int {
	maxWealth := 0
	amount := 0
	for _, account := range accounts {
		amount = 0
		for _, bank := range account {
			amount += bank
		}
		if amount > maxWealth {
			maxWealth = amount
		}
	}
	return maxWealth
}
