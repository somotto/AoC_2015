package main

import (
	"bufio"
	"fmt"
	"os"

	"day_6/utils" // Replace 'yourmodule' with the actual module name
)

// Define the grid dimensions
const gridSize = 1000

// Create a type alias for the grid
type Grid [gridSize][gridSize]int

// ApplyCommand updates the grid based on the command
func ApplyCommand(grid *Grid, command string) {
	// Split command into parts using the utility function
	parts := utils.Fields(command)

	// Determine the action and corresponding state
	action := parts[0] // 'turn' or 'toggle'
	var state int
	if action == "toggle" {
		state = -1 // Toggle
		// For toggle commands, coordinates are in parts[1] and parts[3]
		x1, y1 := ParseCoordinates(parts[1])
		x2, y2 := ParseCoordinates(parts[3])
		// Apply the toggle command to the grid
		for i := x1; i <= x2; i++ {
			for j := y1; j <= y2; j++ {
				grid[i][j] ^= 1 // Toggle the light (0 <-> 1)
			}
		}
	} else if action == "turn" {
		// For turn commands, determine if it's 'on' or 'off'
		if parts[1] == "on" {
			state = 1 // Turn on
		} else if parts[1] == "off" {
			state = 0 // Turn off
		} else {
			fmt.Printf("Unknown turn action: %s\n", parts[1])
			return
		}
		// For turn commands, coordinates are in parts[2] and parts[4]
		x1, y1 := ParseCoordinates(parts[2])
		x2, y2 := ParseCoordinates(parts[4])
		// Apply the turn command to the grid
		for i := x1; i <= x2; i++ {
			for j := y1; j <= y2; j++ {
				grid[i][j] = state // Turn on or off
			}
		}
	} else {
		fmt.Printf("Unknown action: %s\n", action)
	}
}


// parseCoordinates converts coordinate strings to integers using the utility functions
func ParseCoordinates(coord string) (int, int) {
	parts := utils.Split(coord, ",")
	x := utils.Atoi(parts[0])
	y := utils.Atoi(parts[1])
	return x, y
}

// CountLightsOn counts the number of lights that are on
func CountLightsOn(grid *Grid) int {
	count := 0
	for i := 0; i < gridSize; i++ {
		for j := 0; j < gridSize; j++ {
			if grid[i][j] == 1 {
				count++
			}
		}
	}
	return count
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()
	// Initialize the grid
	var grid Grid

	// Read input commands
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		command := scanner.Text()
		ApplyCommand(&grid, command)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	// Count and print the number of lights that are on
	lightsOn := CountLightsOn(&grid)
	fmt.Printf("Total lights on: %d\n", lightsOn)
}
