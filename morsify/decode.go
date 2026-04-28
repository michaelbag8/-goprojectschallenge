package main

import (
	"fmt"
	"os"
	"strings"
)

func Helper(input string)string{
	input = strings.TrimSpace(input)

	if input == " "{
		fmt.Println("input can not be empty")
		os.Exit(1)
	}

	return input
}
func decode(input string, morseToText map[string]string) string {
	input = Helper(input)

	reversedMap := buildReverseMap(morseToText)
	codes := strings.Split(input," ")
	var result strings.Builder

	for _, char := range codes {
		if letter, exist := reversedMap[char]; exist {
			result.WriteString(letter)
			
		}

	}
	return result.String()

}


func buildReverseMap(textToMorse map[string]string) map[string]string {
	reversedMap := make(map[string]string)
	for letter, code := range textToMorse {
		reversedMap[code] = letter
	}
	return reversedMap
}