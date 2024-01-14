package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	var countBytes, countLines, countWords, countCharacters bool

	flag.BoolVar(&countBytes, "c", false, "Count the number of bytes in the input.")
	flag.BoolVar(&countLines, "l", false, "Count the number of lines in the input.")
	flag.BoolVar(&countWords, "w", false, "Count the number of words in the input.")
	flag.BoolVar(&countCharacters, "m", false, "Count the number of characters in the input.")
	flag.Parse()

	filePath := flag.Arg(0)

	if filePath == "" {
		fileContent, err := io.ReadAll(os.Stdin)

		if err != nil {
			fmt.Printf("Error reading standard input: %s\n", err)
			os.Exit(1)
		}
		
		processInput(fileContent, countBytes, countLines, countWords, countCharacters, "")
	} else {
		fileContent, err := os.ReadFile(filePath)

		if err != nil {
			fmt.Printf("Error reading file: %s\n", err)
			os.Exit(1)
		}

		processInput(fileContent, countBytes, countLines, countWords, countCharacters, filePath)
	}
}

func countFileContentLines(fileContent []byte) int {
	numberOfLines := 0
	for _, char := range fileContent {
		if char == '\n' {
			numberOfLines++
		}
	}
	return numberOfLines
}

func processInput(content []byte, numberOfBytes, numberOfLines, numberOfWords, numberOfCharacters bool, filePath string) {
	byteCount := len(content)
	lineCount := countFileContentLines(content)
	wordCount := len(strings.Fields(string(content)))
	characterCount := utf8.RuneCountInString(string(content))

	if numberOfBytes {
		fmt.Printf("%8d %s\n", byteCount, filePath)
	} else if numberOfLines {
		fmt.Printf("%8d %s\n", lineCount, filePath)
	} else if numberOfWords {
		fmt.Printf("%8d %s\n", wordCount, filePath)
	} else if numberOfCharacters {
		fmt.Printf("%8d %s\n", characterCount, filePath)
	} else {
		fmt.Printf("%8d %8d %8d %s\n", byteCount, lineCount, wordCount, filePath)
	}
}