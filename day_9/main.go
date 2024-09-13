package main

import (
	"day_9/part1"
	"day_9/part2"
	"fmt"
)

func main() {
    distances1, locations1, err := part1.ReadInput("input.txt")
    if err != nil {
		fmt.Println("Error reading input file:", err)
        return
    }
	distances2, locations2, err2 := part2.ReadInput2("input.txt")
	if err2 != nil {
		fmt.Println("Error reading input file:", err)
        return
    }

    fmt.Println("Shortest route distance:", part1.ShortestRoute(locations1, distances1))
	fmt.Println("Longest route distance:", part2.LongestRoute(locations2, distances2))
}
