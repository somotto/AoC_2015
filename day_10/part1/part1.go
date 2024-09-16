package part1

import (
	"strconv"
)

// lookAndSay generates the next term in the sequence
func LookAndSay(sequence string) string {
	var result string
	i := 0
	for i < len(sequence) {
		count := 1
		for i+1 < len(sequence) && sequence[i] == sequence[i+1] {
			i++
			count++
		}
		result += strconv.Itoa(count) + string(sequence[i])
		i++
	}
	return result
}

// lengthOf40thLookAndSay calculates the length of the 40th term
func LengthOf40thLookAndSay(input string) int {
	sequence := input
	for i := 0; i < 40; i++ {
		sequence = LookAndSay(sequence)
	}
	return len(sequence)
}


