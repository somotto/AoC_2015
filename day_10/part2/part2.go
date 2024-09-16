package part2

import "day_10/part1"

// LengthOf50thLookAndSay calculates the length of the 50th term
func LengthOf50thLookAndSay(input string) int {
	sequence := input
	for i := 0; i < 50; i++ {
		sequence = part1.LookAndSay(sequence)
	}
	return len(sequence)
}
