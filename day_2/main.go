package main

import (
	"bufio"
	"daytwo/utils"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("dimension.txt")

	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	TotalWrappingPaper := 0
	TotalRibbonFeet := 0

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		dimensions := strings.Split(line, "x")
		var l, w, h int
		fmt.Sscanf(dimensions[0], "%d", &l)
		fmt.Sscanf(dimensions[1], "%d", &w)
		fmt.Sscanf(dimensions[2], "%d", &h)

		TotalWrappingPaper += utils.WrappingPaper(l, w, h)
		TotalRibbonFeet += utils.FindRibbonFeet(l, w, h)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println(TotalWrappingPaper)
	fmt.Println(TotalRibbonFeet)
}
