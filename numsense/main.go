package main

import(
	"fmt"
	"code/functions"

)
func checkForOddEvenNumbers(num int) int {
	if num%2 == 0 {
		fmt.Println("even")
	} else {
		fmt.Println("odd")
	}
	return num
}

func primeNumber(num int) {
	if num <= 1 {
		fmt.Println("Not a prime")
		return
	}

	for i := 2; i < num; i++ {
		if num%i == 0 {
			fmt.Println("Not a prime")
			return
		}
	}

	fmt.Println("Prime")
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


/*
func applyRules(inputFile string) error {
	result, err := transformation(inputFile)
	if err != nil {
		return err
	}

	return os.WriteFile(outputFile, []byte(result), 0644)
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . <input.txt> <output.txt>")
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	if err := applyRules(inputFile, outputFile); err != nil {
		fmt.Println("Error:", err)
	}
}*/



func main() {
	even, odd := functions.CountEvenOdd([]int{1, 2, 4, 8, 8, 7, 9, 5, 8, 5, 9, 10, 11})
	numbers := []int{1, 2, 4, 8, 8, 7, 9, 30, 60,5, 8, 5, 9, 10, 11}

	fmt.Println("Even: ", even)

	fmt.Println("Odd: ", odd)

	functions.PrintList(10)

	fmt.Println()

	fmt.Println("Largest Number: ", functions.Largest(numbers))

	fmt.Println("Smallest Number: ",functions.Smallest(numbers))
}


/*
Using "slices package"

	num := []int{1, 2, 4, 8, 8, 7, 9, 5, 8, 5, 9, 10, 11}

	fmt.Println("Max: ", slices.Max(num))
	fmt.Println("Min: ", slices.Min(num))
*/
