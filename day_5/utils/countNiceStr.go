package utils

import (
	"bufio"
	"os"
)

func CountNiceStrings(filename string) (int, error) {
    file, err := os.Open(filename)
    if err != nil {
        return 0, err
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    niceCount := 0

    for scanner.Scan() {
        line := scanner.Text()
        if IsNice(line) {
            niceCount++
        }
    }

    if err := scanner.Err(); err != nil {
        return 0, err
    }

    return niceCount, nil
}
