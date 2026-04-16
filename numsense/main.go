package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func checkForOddEvenNumbers(num int) string {
	if num%2 == 0 {
		return "even"

	}
	return "odd"
}

func CheckPrime(num int) string {
	if num < 2 {
		return "Not Prime"
	}

	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return "Not Prime"
		}
	}

	return "Prime"
}

func perfectSquare(num int) string {
	if num < 0 {
		return "not perfect square"
	}
	root := int(math.Sqrt(float64(num)))
	if root*root == num {
		return "perfect square"
	}
	return "not perfect square"
}

func CheckForPositive(num int) string {
	if num < 0 {
		return "Negative"
	}

	return "Positive"
}

func main() {
	var input string
	if len(os.Args) > 1 {
		input = os.Args[1]
	} else {
		fmt.Print("Enter a number >> ")
		reader := bufio.NewReader(os.Stdin)
		input, _ = reader.ReadString('\n')
		input = strings.TrimSpace(input)
	}

	if input == "" {
		fmt.Fprintln(os.Stderr, "error: no input provided")
		os.Exit(1)
	}

	line, err := strconv.Atoi(input)

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: '%s' is not a valid integer\n", input)
		os.Exit(1)
	}
		
	fmt.Printf("%d →  %s , %s , %s , %s \n", line, CheckForPositive(line), checkForOddEvenNumbers(line), CheckPrime(line), perfectSquare(line))
	
}
