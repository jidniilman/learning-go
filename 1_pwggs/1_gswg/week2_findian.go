package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var input string

func main() {
	fmt.Printf("Input your string : ")

	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		input = strings.ToLower(scanner.Text())
	}

	if string(input[0]) == "i" && string(input[len(input) - 1]) == "n" && strings.Contains(input, "a") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}

	// If you want the order of "i" and "n" is ignored
	// algorithmicWay()
	// fmt.Println("------------------------------------")
	// shortWay()

}

// Ignored order
// func algorithmicWay() {
// 	var iFound bool
// 	var aFound bool
// 	var nFound bool
// 	for i := 0; i < len(input); i++ {
// 		if string(input[i]) == "i" {
// 			iFound = true
// 		}
// 		if iFound {
// 			if string(input[i]) == "a" {
// 				aFound = true
// 			}
// 		}
// 		if aFound {
// 			if string(input[i]) == "n" {
// 				nFound = true
// 			}
// 		}
// 	}
//
// 	if iFound && aFound && nFound {
// 		fmt.Println("Found!")
// 	} else {
// 		fmt.Println("Not Found!")
// 	}
// }

// // Ignored order
// func shortWay() {
// 	substr := strings.TrimLeft(strings.TrimRight(input,"n"),"i")
// 	if strings.Contains(substr, "a") {
// 		fmt.Println("Found!")
// 	} else {
// 		fmt.Println("Not Found!")
// 	}
// }