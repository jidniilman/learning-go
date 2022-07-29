package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	identity := map[string]string{}
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("Input your name : ")
	if scanner.Scan() {
		identity["name"] = scanner.Text()
	}

	fmt.Printf("Input your address : ")
	if scanner.Scan() {
		identity["address"] = scanner.Text()
	}

	jsonMarshal, _ := json.Marshal(identity)
	fmt.Println("JSON Marshal Result : ", string(jsonMarshal))
}
