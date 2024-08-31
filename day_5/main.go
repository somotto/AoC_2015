package main

import (
	"day_5/utils"
	"fmt"
)

func main() {
	filename := "input.txt" // replace with your input file name
	niceCount, err := utils.CountNiceStrings(filename)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Number of nice strings: %d\n", niceCount)
}
