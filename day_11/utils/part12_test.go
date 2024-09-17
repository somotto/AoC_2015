package utils

import (
	"testing"
)

func TestIncrementPassword(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"abcd", "abce"},
		{"xyzz", "xzaa"},
		{"zzzz", "aaaa"},
		{"cqjxjnds", "cqjxjndt"},
	}

	for _, tt := range tests {
		got := IncrementPassword(tt.input)
		if got != tt.want {
			t.Errorf("incrementPassword(%q) = %q; want %q", tt.input, got, tt.want)
		}
	}
}

func TestHasIncreasingStraight(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"abc", true},
		{"abd", false},
		{"hijklmmn", true},
		{"xyz", true},
		{"aabbcc", false},
	}

	for _, tt := range tests {
		got := HasIncreasingStraight(tt.input)
		if got != tt.want {
			t.Errorf("hasIncreasingStraight(%q) = %v; want %v", tt.input, got, tt.want)
		}
	}
}

func TestHasNoForbiddenLetters(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"abcdefg", true},
		{"hijklmmn", false},
		{"oopqrst", false},
		{"lmnopqr", false},
		{"abcdefgh", true},
	}

	for _, tt := range tests {
		got := HasNoForbiddenLetters(tt.input)
		if got != tt.want {
			t.Errorf("hasNoForbiddenLetters(%q) = %v; want %v", tt.input, got, tt.want)
		}
	}
}

func TestHasTwoPairs(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"aabb", true},
		{"abcdd", false},
		{"aabcdd", true},
		{"aaa", false},  
		{"abbc", false}, 
	}

	for _, tt := range tests {
		got := HasTwoPairs(tt.input)
		if got != tt.want {
			t.Errorf("hasTwoPairs(%q) = %v; want %v", tt.input, got, tt.want)
		}
	}
}

func TestIsValidPassword(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"abcdefgh", false},  
		{"abcdffaa", true},   
		{"ghjaabcc", true},   
		{"hijklmmn", false},  
		{"abbceffg", false},  
		{"abbcegjk", false},  
	}

	for _, tt := range tests {
		got := IsValidPassword(tt.input)
		if got != tt.want {
			t.Errorf("isValidPassword(%q) = %v; want %v", tt.input, got, tt.want)
		}
	}
}

func TestFindNextPassword(t *testing.T) {
	tests := []struct {
		currentPassword string
		want            string
	}{
		{"abcdefgh", "abcdffaa"},  
		{"ghijklmn", "ghjaabcc"},  
		{"cqjxjnds", "cqjxxyzz"},  
		{"zzzzzzzz", "aaaaaabc"},  
	}

	for _, tt := range tests {
		got := FindNextPassword(tt.currentPassword)
		if got != tt.want {
			t.Errorf("FindNextPassword(%q) = %q; want %q", tt.currentPassword, got, tt.want)
		}
	}
}
