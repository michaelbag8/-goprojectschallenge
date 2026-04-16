# Go Number Analyzer

## Overview

This is a simple Go command-line application that analyzes a given integer and reports several of its mathematical properties. It determines whether the number is:

* Positive or Negative
* Odd or Even
* Prime or Not Prime
* A Perfect Square or Not

The program accepts input either as a command-line argument or via standard input (interactive prompt).

---

## Features

* Handles user input from both CLI arguments and terminal prompt
* Validates input to ensure it is an integer
* Efficient prime number checking
* Uses Go’s standard library (`math`, `strconv`, `bufio`, etc.)
* Clean and readable output format

---

## Installation

1. Ensure you have Go installed (version 1.18 or later recommended).
2. Clone this repository or copy the source code into a file (e.g., `main.go`).

```bash
git clone <your-repo-url>
cd <your-project-folder>
```

---

## Usage

### Run with Command-Line Argument

```bash
go run main.go 25
```

### Run in Interactive Mode

```bash
go run main.go
```

Then enter a number when prompted:

```
Enter a number >> 16
```

---

## Example Output

```bash
25 →  Positive , odd , Not Prime , not perfect square
```

```bash
16 →  Positive , even , Not Prime , perfect square
```

---

## Function Breakdown

### `CheckForPositive(num int) string`

Determines if a number is **Positive** or **Negative**.

### `checkForOddEvenNumbers(num int) string`

Returns whether the number is **odd** or **even**.

### `CheckPrime(num int) string`

Checks if the number is a **Prime** number using an optimized loop up to √n.

### `perfectSquare(num int) string`

Determines whether the number is a **perfect square**.

---

## Error Handling

* If no input is provided → program exits with an error
* If input is not a valid integer → displays a clear error message

Example:

```bash
error: 'abc' is not a valid integer
```

---

## Dependencies

This project uses only Go’s standard library:

* `fmt`
* `os`
* `bufio`
* `strconv`
* `strings`
* `math`

---

## Future Improvements

* Support for floating-point numbers
* Add unit tests
* Return structured output (e.g., JSON)
* Extend with more mathematical properties (factorial, divisors, etc.)

---

