package main

import (
	"fmt"
)

func checkForOddEvenNumbers(num int) string {
	if num%2 == 0 {
		return "even"

	}
	return "odd"
}

func CheckPrime(num int) string {
	if num <= 1 {
		return "negative number"
	}

	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return "Not a prime"
		}
	}

	return "Prime"
}

// func primeNumbers(num int) string {
// 	if isPrime(num) {
// 		return "Prime"
// 	}
// 	return "Not a prime"

// }

func CheckForPostive(num int) string {
	if num < 0 {
		return "Negative"
	}

	return "Positive"
}
func numberCount(num int) {
	for i := 1; i <= num; i++ {
		fmt.Print(i, " ")
	}
}

func numbersCount(num []int) int {
	count := 0
	for i := 1; i <= len(num); i++ {
		count++
	}
	return count
}

func main() {
	numbers := []int{1, 2, 4, 8, 8, 7, 9, 30, 60, 5, 8, 5, 9, 10, 11}
	fmt.Println(numbersCount(numbers))
	numberCount(4)
	fmt.Println(CheckPrime(13))
	fmt.Println(checkForOddEvenNumbers(3))
	fmt.Println(CheckForPostive(10))

}

/*

TITLE: NumSense

CONCEPT: fmt, strconv, os.Stdin, bufio

RULES:
→ Write everything in Go using only the standard library.
→ The program must compile and run without errors.
→ The project must be structured inside a folder: thecodecrafterthon-day1/
→ Include a README.md explaining how to build and run the program.
→ The program must run from the terminal.

PROJECT DESCRIPTION:
A CLI number classifier that reads a number and tells you whether it's positive, negative, even, odd, prime, or a perfect square.

CLI USAGE EXAMPLES:
> ./numsense 28
  ✦ 28 → positive, even, not prime, not perfect square
> ./numsense -7
  ✦ -7 → negative, odd

REQUIREMENTS:
→ Accept a number as CLI arg or prompt for it.
→ Classify: sign, parity, primality, perfect square.
→ Print all labels on one line.

VALIDATION RULES:
→ Reject non-numeric input with a clear error.
→ Handle 0 and 1 as edge cases.
→ Do NOT crash on empty input.

PROGRAM BEHAVIOR:
→ Single-shot. One number in, one line out.
→ Handle whitespace in input.
→ Case-insensitive where applicable.

THINKING GUIDANCE:
→ How does strconv.Atoi differ from strconv.ParseFloat?
→ What's the most efficient primality check for small numbers?
→ Which standard library packages do you need for this task?

BONUS (OPTIONAL):
→ Accept multiple numbers separated by spaces.
→ Support a --json flag for structured output.
*/
