---

# Text Analyzer (Go)

This is a simple Go program that reads text input from the user and reports the number of **words**, **characters**, and **lines**.

## 📖 Overview
The program demonstrates:
- Reading input from standard input (`os.Stdin`) using `bufio.Reader`.
- Counting characters, words, and lines with helper functions.
- Basic string manipulation using the `strings` package.

## ⚙️ Features
- **Character count**: Counts all characters, including spaces.
- **Word count**: Counts words separated by whitespace.
- **Line count**: Counts lines separated by newline characters (`\n`).

## 🧩 Code Structure
- `countChars(word string) int` → Counts characters.
- `countLines(word string) int` → Counts lines.
- `countWords(word string) int` → Counts words.
- `main()` → Reads user input and prints results.

## ▶️ Usage
1. Save the code in a file, e.g. `main.go`.
2. Run the program:
   ```bash
   go run main.go
   ```
3. Enter some text when prompted:
   ```
   Enter a text: Hello World
   ```
4. Output:
   ```
   ★ Words: 2 | Chars: 11 | Lines: 1
   ```

## 📝 Example
Input:
```
Go is fun
Go is powerful
```

Output:
```
★ Words: 6 | Chars: 27 | Lines: 2
```

## 🚀 Future Improvements
- Support multi-line input until EOF instead of stopping at the first newline.
- Add options to ignore spaces or punctuation in character count.
- Provide detailed statistics (average word length, longest line, etc.).

---
