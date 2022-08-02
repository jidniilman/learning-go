package main

import (
	"fmt"
)

func main() {
	fmt.Println(kWeakestRows([][]int{
		{1, 1, 0, 0, 0},
		{1, 1, 1, 1, 0},
		{1, 0, 0, 0, 0},
		{1, 1, 0, 0, 0},
		{1, 1, 1, 1, 1}}, 3))

	fmt.Println(kWeakestRows([][]int{
		{1, 0, 0, 0},
		{1, 1, 1, 1},
		{1, 0, 0, 0},
		{1, 0, 0, 0}}, 2))
}

type Row struct {
	index int
	power int
}

func kWeakestRows(mat [][]int, k int) []int {
	power := make([]Row, len(mat))
	for i := 0; i < len(mat); i++ {
		power[i].index, power[i].power = i, calcPower(mat[i])
	}
	sort(power)
	order := make([]int, k)
	for i := range order {
		order[i] = power[i].index
	}
	return order
}

func calcPower(rows []int) int {
	power := 0
	for _, row := range rows {
		if row == 1 {
			power += row
			continue
		}
		return power
	}
	return power
}

func sort(slice []Row) {
	for i := 0; i < len(slice)-1; i++ {
		for j := 0; j < len(slice)-i-1; j++ {
			if slice[j].power > slice[j+1].power {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}
}
