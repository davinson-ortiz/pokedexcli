package main

import (
	"fmt"
)

func callbackHelp(cfg *config) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!\nUsage:\n")
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}
