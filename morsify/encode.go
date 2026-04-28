package main

import (
	"strings"
)

var letterToMorse = map[string]string{
	"A": ".-",
	"B": "-...",
	"C": "-.-.",
	"D": "-..",
	"E": ".",
	"F": "..-.",
	"G": "--.",
	"H": "....",
	"I": "..",
	"J": ".---",
	"K": "-.-",
	"L": ".-..",
	"M": "--",
	"N": "-.",
	"O": "---",
	"P": ".--.",
	"Q": "--.-",
	"R": ".-.",
	"S": "...",
	"T": "-",
	"U": "..-",
	"V": "...-",
	"W": ".--",
	"X": "-..-",
	"Y": "-.--",
	"Z": "--..",
}

func encode(input string, textToMorse map[string]string) string {
	input = Helper(input)

	input = strings.ToUpper(input)

	var result strings.Builder

	for _, char := range input {
		letter := string(char)
		if code, exist := textToMorse[letter]; exist {
			result.WriteString(code)
			result.WriteString(" ")
		}

	}
	return result.String()

}

