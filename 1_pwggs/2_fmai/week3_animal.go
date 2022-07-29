package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func main() {
	// Declare the animal hard-coded data
	animals := map[string]Animal{}
	animals["cow"] = Animal{
		food:       "grass",
		locomotion: "walk",
		noise:      "moo",
	}
	animals["bird"] = Animal{
		food:       "worms",
		locomotion: "fly",
		noise:      "peep",
	}
	animals["snake"] = Animal{
		food:       "mice",
		locomotion: "slither",
		noise:      "hsss",
	}

	// Read Requests
	fmt.Println("Input Animal name and Command (X to exit) - Ex. cow eat")
	fmt.Print(">")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		commands := strings.Split(scanner.Text(), " ")
		if commands[0] == "X" {
			break
		}
		switch commands[1] {
		case "eat":
			animals[commands[0]].Eat()
		case "move":
			animals[commands[0]].Move()
		case "speak":
			animals[commands[0]].Speak()
		default:
			fmt.Println("Unrecognized command. Skipping")
		}

		fmt.Print(">")
	}
}

func (a Animal) Eat() {
	fmt.Println(a.food)
}

func (a Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a Animal) Speak() {
	fmt.Println(a.noise)
}
