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

	niceCount2 := utils.CountNiceStrPart2(filename)
	
	fmt.Printf("Number of nice strings: %d\n", niceCount)
	fmt.Printf("Number of nice strings: %d\n", niceCount2)
}
