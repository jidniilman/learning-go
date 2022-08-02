package main

import "fmt"

func main() {
	fmt.Println("Output:", romanToInt("III"))
	fmt.Println("Output:", romanToInt("LVIII"))
	fmt.Println("Output:", romanToInt("MCMXCIV"))
}

func romanToInt(s string) int {
	number := 0
	for i := 0; i < len(s); i++ {
		if i < len(s)-1 && string(s[i]) == "I" && string(s[i+1]) == "V" {
			number += 4
			i++
			continue
		}
		if i < len(s)-1 && string(s[i]) == "I" && string(s[i+1]) == "X" {
			number += 9
			i++
			continue
		}
		if i < len(s)-1 && string(s[i]) == "X" && string(s[i+1]) == "L" {
			number += 40
			i++
			continue
		}
		if i < len(s)-1 && string(s[i]) == "X" && string(s[i+1]) == "C" {
			number += 90
			i++
			continue
		}
		if i < len(s)-1 && string(s[i]) == "C" && string(s[i+1]) == "D" {
			number += 400
			i++
			continue
		}
		if i < len(s)-1 && string(s[i]) == "C" && string(s[i+1]) == "M" {
			number += 900
			i++
			continue
		}
		number += romanCharToInt(string(s[i]))
	}

	return number
}

func romanCharToInt(r string) int {
	switch r {
	case "I":
		return 1
	case "V":
		return 5
	case "X":
		return 10
	case "L":
		return 50
	case "C":
		return 100
	case "D":
		return 500
	case "M":
		return 1000
	default:
		return 0
	}
}
