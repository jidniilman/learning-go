package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Person struct {
	fname [20]rune
	lname [20]rune
}

func main() {
	// Declare the persons slice
	persons := make([]Person, 0)

	// Read File Input (Local Path)
	var inputFile string
	fmt.Printf("Input your TXT File Name (Local Path) - Ex. text.txt : ")
	fmt.Scan(&inputFile)

	// Read data from file
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	// Read Line by Line
	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		fullName := strings.Split(fileScanner.Text(), " ")
		persons = append(persons, Person{
			fname: stringTo20Rune(fullName[0]),
			lname: stringTo20Rune(fullName[1]),
		})
	}

	// Print output
	for i, person := range persons {
		fmt.Printf("Person %d First Name : %s with Last Name : %s\n", i+1, runesToString(person.fname), runesToString(person.lname))
	}
}

func stringTo20Rune(str string) (res [20]rune) {
	for i, rune := range str {
		if i >= 20 {
			break
		}
		res[i] = rune
	}
	return
}

func runesToString(runes [20]rune) (res string) {
	res = ""
	for _, rune := range runes {
		res += string(rune)
	}
	return
}
