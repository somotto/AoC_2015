package utils

import (
	"bufio"
	"fmt"
	"os"
)

// Count nice strings
func CountNiceStrPart2(filename string) int {
    niceCount := 0

    // Open the file
    file, err := os.Open(filename)
    if err != nil {
        fmt.Println("Error opening file:", err)
        return 0
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        if HasNonOverlappingPair(line) && HasRepeatWithOneBetween(line) {
            niceCount++
        }
    }

    if err := scanner.Err(); err != nil {
        fmt.Println("Error reading file:", err)
    }

    return niceCount
}