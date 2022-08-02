package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(fizzBuzz(3))
	fmt.Println(fizzBuzz(5))
	fmt.Println(fizzBuzz(15))
}

func fizzBuzz(n int) []string {
	// strings := make([]string, 0, 10000)
	var strings [10000]string
	for i := 1; i < n+1; i++ {
		if i%3 == 0 && i%5 == 0 {
			// strings = append(strings, "FizzBuzz")
			strings[i] = "FizzBuzz"
			continue
		}
		if i%3 == 0 {
			// strings = append(strings, "Fizz")
			strings[i] = "Fizz"
			continue
		}
		if i%5 == 0 {
			// strings = append(strings, "Buzz")
			strings[i] = "Buzz"
			continue
		}
		// strings = append(strings, strconv.Itoa(i))
		strings[i] = strconv.Itoa(i)
	}
	return strings[1 : n+1]
}
