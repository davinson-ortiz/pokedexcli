package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			userInput := scanner.Text()
			cleanedInput := cleanInput(userInput)
			fmt.Printf("Your command was: %s\n", cleanedInput[0])
		}
	}
}

func cleanInput(text string) []string {
	lowerText := strings.ToLower(text)
	sliceOfWords := strings.Fields(lowerText)
	return sliceOfWords
}
