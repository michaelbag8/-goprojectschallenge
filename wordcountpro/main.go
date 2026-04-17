package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// countChars count characters, space and words
func countChars(word string) int{
	count := 0
	for i:=0; i <len(word); i++{
		count++
	}
	return count
}

//countLines count lines
func countLines(word string)int{
	words := strings.Split(word, "\n")
	return  len(words)
	
}

//countWords count only words
func countWords(word string)int{
	words := strings.Fields(word)
	return len(words)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a text: ")

	input, err := reader.ReadString('\n')
	if err != nil{
		fmt.Println("Error reading the input", input)
		return
	}


	fmt.Printf("★ Words: %d | Chars: %d | Lines: %d\n",
        countWords(input), countChars(input), countLines(input))
}
