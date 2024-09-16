package main

import (
	"day_10/part1"
	"fmt"
)
func main() {
	inputSequence := "3113322113"
	length := part1.LengthOf40thLookAndSay(inputSequence)
	fmt.Println(length)
}