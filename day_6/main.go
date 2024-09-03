package main

import (
	"bufio"
	"day_6/utils"
	"fmt"
	"os"
)

// Define the grid dimensions
const gridSize = 1000

// Create type aliases for the grids
type GridPart1 [gridSize][gridSize]bool
type GridPart2 [gridSize][gridSize]int

// ApplyCommand updates both grids based on the command
func ApplyCommand(gridPart1 *GridPart1, gridPart2 *GridPart2, command string) {
	parts := utils.Fields(command)
	var action string
	var x1, y1, x2, y2 int

	if parts[0] == "toggle" {
		action = "toggle"
		x1, y1 = ParseCoordinates(parts[1])
		x2, y2 = ParseCoordinates(parts[3])
	} else if parts[0] == "turn" {
		action = parts[1] // "on" or "off"
		x1, y1 = ParseCoordinates(parts[2])
		x2, y2 = ParseCoordinates(parts[4])
	} else {
		fmt.Printf("Unknown action: %s\n", parts[0])
		return
	}

	for i := x1; i <= x2; i++ {
		for j := y1; j <= y2; j++ {
			// Part 1: On/Off state
			switch action {
			case "on":
				gridPart1[i][j] = true
			case "off":
				gridPart1[i][j] = false
			case "toggle":
				gridPart1[i][j] = !gridPart1[i][j]
			}

			// Part 2: Brightness levels
			switch action {
			case "on":
				gridPart2[i][j]++
			case "off":
				if gridPart2[i][j] > 0 {
					gridPart2[i][j]--
				}
			case "toggle":
				gridPart2[i][j] += 2
			}
		}
	}
}

// ParseCoordinates converts coordinate strings to integers
func ParseCoordinates(coord string) (int, int) {
	parts := utils.Split(coord, ",")
	x := utils.Atoi(parts[0])
	y := utils.Atoi(parts[1])
	return x, y
}

// CountLightsOn counts the number of lights that are on (Part 1)
func CountLightsOn(grid *GridPart1) int {
	count := 0
	for i := 0; i < gridSize; i++ {
		for j := 0; j < gridSize; j++ {
			if grid[i][j] {
				count++
			}
		}
	}
	return count
}

// CalculateTotalBrightness calculates the total brightness of all lights (Part 2)
func CalculateTotalBrightness(grid *GridPart2) int {
	total := 0
	for i := 0; i < gridSize; i++ {
		for j := 0; j < gridSize; j++ {
			total += grid[i][j]
		}
	}
	return total
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	// Initialize the grids
	var gridPart1 GridPart1
	var gridPart2 GridPart2

	// Read input commands
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		command := scanner.Text()
		ApplyCommand(&gridPart1, &gridPart2, command)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	// Calculate and print results for both parts
	lightsOn := CountLightsOn(&gridPart1)
	totalBrightness := CalculateTotalBrightness(&gridPart2)
	fmt.Printf("Part 1 - Total lights on: %d\n", lightsOn)
	fmt.Printf("Part 2 - Total brightness: %d\n", totalBrightness)
}