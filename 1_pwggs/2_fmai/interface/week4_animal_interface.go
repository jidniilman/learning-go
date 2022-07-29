package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Cow struct {
}

type Bird struct {
}

type Snake struct {
}

type Animal interface {
	Eat()
	Move()
	Speak()
}

func main() {
	animals := map[string]Animal{}

	// Read Requests
	fmt.Println("Input Commands (X to exit) - Ex. newanimal joe cow")
	fmt.Print(">")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		commands := strings.Split(scanner.Text(), " ")
		if commands[0] == "X" {
			break
		}
		if commands[0] != "newanimal" && commands[0] != "query" {
			fmt.Println("Unrecognized Command. Skipping")
			fmt.Print(">")
			continue
		}

		if commands[0] == "newanimal" {
			switch commands[2] {
			case "cow":
				animals[commands[1]] = Cow{}
				fmt.Println("Created it!")
			case "bird":
				animals[commands[1]] = Bird{}
				fmt.Println("Created it!")
			case "snake":
				animals[commands[1]] = Snake{}
				fmt.Println("Created it!")
			default:
				fmt.Println("Unrecognized Animal. Skipping")
			}
		}

		if commands[0] == "query" {
			switch commands[2] {
			case "eat":
				animals[commands[1]].Eat()
			case "move":
				animals[commands[1]].Move()
			case "speak":
				animals[commands[1]].Speak()
			default:
				fmt.Println("Unrecognized command. Skipping")
			}
		}
		fmt.Print(">")
	}
}

func (c Cow) Eat() {
	fmt.Println("grass")
}

func (c Cow) Move() {
	fmt.Println("walk")
}

func (c Cow) Speak() {
	fmt.Println("moo")
}

func (b Bird) Eat() {
	fmt.Println("worms")
}

func (b Bird) Move() {
	fmt.Println("fly")
}

func (b Bird) Speak() {
	fmt.Println("peep")
}

func (s Snake) Eat() {
	fmt.Println("mice")
}

func (s Snake) Move() {
	fmt.Println("slither")
}

func (s Snake) Speak() {
	fmt.Println("hsss")
}
