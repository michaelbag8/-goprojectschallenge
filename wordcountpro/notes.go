package main

import(
	"bufio"
	"strings"
)


func CountLinesWithBufio(word string)int{
	count := 0
	scanner := bufio.NewScanner(strings.NewReader(word))
	for scanner.Scan(){
		count++
	}
	return count
}

func countLinesWithCount(word string)int{
	return strings.Count(word, "\n") + 1
}
