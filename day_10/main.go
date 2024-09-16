package main

import (
	"day_10/part1"
	"day_10/part2"
	"fmt"
)
func main() {
	inputSequence := "3113322113"
	length := part1.LengthOf40thLookAndSay(inputSequence)
	length2 := part2.LengthOf50thLookAndSay(inputSequence)


	fmt.Println(length)
	fmt.Println(length2)
}