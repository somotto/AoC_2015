package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Encode a string by adding surrounding quotes and escaping special characters
func encodeString(s string) string {
    var encoded strings.Builder
    encoded.WriteString("\"") 

    for _, r := range s {
        switch r {
        case '\\':
            encoded.WriteString("\\\\") 
        case '"':
            encoded.WriteString("\\\"") 
        default:
            encoded.WriteRune(r) 
        }
    }

    encoded.WriteString("\"") 
    return encoded.String()
}

func main() {
    file, err := os.Open("../input.txt")
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer file.Close()

    var totalOriginalLength, totalEncodedLength int

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()

        originalLength := len(line)

        encoded := encodeString(line)
        encodedLength := len(encoded)

        totalOriginalLength += originalLength
        totalEncodedLength += encodedLength
    }

    if err := scanner.Err(); err != nil {
        fmt.Println("Error reading file:", err)
    }

    fmt.Printf("Difference: %d\n", totalEncodedLength-totalOriginalLength)
}
