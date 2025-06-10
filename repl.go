package main

import (
	"strings"
)

func cleanInput(text string) []string {
	lowerText := strings.ToLower(text)
	sliceOfWords := strings.Fields(lowerText)
	return sliceOfWords
}
