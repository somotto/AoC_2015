package main

import (
	"bufio"
	"fmt"
	"os"
)

func processString(s string) (int, int) {
	codeCount := len(s)
	memoryCount := 0
	i := 1 // Start after the opening quote

	for i < len(s)-1 { // Stop before the closing quote
		if s[i] == '\\' {
			if i+1 < len(s) {
				if s[i+1] == '\\' || s[i+1] == '"' {
					memoryCount++
					i += 2
				} else if s[i+1] == 'x' && i+3 < len(s) {
					memoryCount++
					i += 4
				} else {
					memoryCount++
					i++
				}
			} else {
				memoryCount++
				i++
			}
		} else {
			memoryCount++
			i++
		}
	}

	return codeCount, memoryCount
}

func processFile(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	totalCodeCount := 0
	totalMemoryCount := 0
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		codeCount, memoryCount := processString(scanner.Text())
		totalCodeCount += codeCount
		totalMemoryCount += memoryCount
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return totalCodeCount - totalMemoryCount, nil
}

func main() {
	filename := "input.txt" // Replace with the actual filename
	result, err := processFile(filename)
	if err != nil {
		fmt.Printf("Error processing file: %v\n", err)
		return
	}
	fmt.Printf("The difference is: %d\n", result)
}