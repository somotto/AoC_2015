package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var wires = map[string]uint16{}
var instructions = map[string]string{}

// Helper functions for bitwise operations
func AND(x, y uint16) uint16 {
	return x & y
}

func OR(x, y uint16) uint16 {
	return x | y
}

func LSHIFT(x uint16, n int) uint16 {
	return x << n
}

func RSHIFT(x uint16, n int) uint16 {
	return x >> n
}

func NOT(x uint16) uint16 {
	return ^x & 0xFFFF
}

// Function to compute the signal for a given wire
func getSignal(wire string) uint16 {
	// If the wire is a digit, convert and return it
	if val, err := strconv.Atoi(wire); err == nil {
		return uint16(val)
	}

	// If the signal is already computed, return it
	if val, ok := wires[wire]; ok {
		return val
	}

	// Get the operation associated with the wire
	operation := instructions[wire]
	parts := strings.Fields(operation)

	var signal uint16

	switch len(parts) {
	case 1:
		// Direct assignment, e.g., '123 -> x'
		signal = getSignal(parts[0])
	case 2:
		// NOT operation, e.g., 'NOT x -> y'
		signal = NOT(getSignal(parts[1]))
	case 3:
		// Other operations, e.g., 'x AND y -> z', 'x LSHIFT 2 -> f'
		switch parts[1] {
		case "AND":
			signal = AND(getSignal(parts[0]), getSignal(parts[2]))
		case "OR":
			signal = OR(getSignal(parts[0]), getSignal(parts[2]))
		case "LSHIFT":
			shift, _ := strconv.Atoi(parts[2])
			signal = LSHIFT(getSignal(parts[0]), shift)
		case "RSHIFT":
			shift, _ := strconv.Atoi(parts[2])
			signal = RSHIFT(getSignal(parts[0]), shift)
		}
	}

	// Store the computed signal in the wires map
	wires[wire] = signal
	return signal
}

func main() {
	// Read the input from the file
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	// Parse the input and fill the instructions map
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		parts := strings.Split(line, " -> ")
		instructions[parts[1]] = parts[0]
	}

	// Compute and print the signal for wire "a"
	signalA := getSignal("a")
	fmt.Printf("The signal for wire 'a' is: %d\n", signalA)

	// Override wire 'b' with the signal from wire 'a'
	instructions["b"] = strconv.Itoa(int(signalA))

	// Reset all wires
	wires = map[string]uint16{}
	instructionsCopy := make(map[string]string)
	for k, v := range instructions {
		instructionsCopy[k] = v
	}
	instructions = instructionsCopy

	// Recompute the signal for wire 'a'
	newSignalA := getSignal("a")
	fmt.Printf("The new signal for wire 'a' after overriding 'b' and resetting is: %d\n", newSignalA)
}
