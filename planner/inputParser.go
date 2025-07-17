package planner

import "strings"

func ParseInput(input string) ([]string, error) {
	input = strings.TrimSpace(input)
	words := strings.Fields(input)
	var parsedInput []string
	for _, word := range words {
		word = strings.ToLower(word)
		parsedInput = append(parsedInput, word)
	}
	return parsedInput, nil
}
