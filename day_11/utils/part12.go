package utils

import (
	"strings"
)

// Increment the password by treating it as a base-26 number.
func IncrementPassword(password string) string {
	runes := []rune(password)
	for i := len(runes) - 1; i >= 0; i-- {
		if runes[i] == 'z' {
			runes[i] = 'a'
		} else {
			runes[i]++
			break
		}
	}
	return string(runes)
}

// Check for increasing straight of at least 3 letters (e.g., abc, bcd, cde).
func HasIncreasingStraight(password string) bool {
	for i := 0; i < len(password)-2; i++ {
		if password[i+1] == password[i]+1 && password[i+2] == password[i]+2 {
			return true
		}
	}
	return false
}

// Check if the password contains any forbidden letters (i, o, l).
func HasNoForbiddenLetters(password string) bool {
	return !strings.ContainsAny(password, "iol")
}

// Check for at least two different non-overlapping pairs of letters.
func HasTwoPairs(password string) bool {
	pairs := 0
	for i := 0; i < len(password)-1; i++ {
		if password[i] == password[i+1] {
			pairs++
			i++ // Skip the next character to ensure non-overlapping pairs.
		}
		if pairs >= 2 {
			return true
		}
	}
	return false
}

// Check if the password satisfies all the requirements.
func IsValidPassword(password string) bool {
	return HasIncreasingStraight(password) && HasNoForbiddenLetters(password) && HasTwoPairs(password)
}

// Find the next valid password.
func FindNextPassword(currentPassword string) string {
	for {
		currentPassword = IncrementPassword(currentPassword)
		if IsValidPassword(currentPassword) {
			return currentPassword
		}
	}
}
