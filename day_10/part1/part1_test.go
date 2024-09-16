package part1

import (
	"testing"
)

func TestLookAndSay(t *testing.T) {
	tests := []struct {
		sequence string
		want     string
	}{
		{"1", "11"},
		{"11", "21"},
		{"21", "1211"},
		{"1211", "111221"},
		{"111221", "312211"},
	}

	for _, tt := range tests {
		got := LookAndSay(tt.sequence)
		if got != tt.want {
			t.Errorf("LookAndSay(%q) = %q, want %q", tt.sequence, got, tt.want)
		}
	}
}

func TestLengthOf40thLookAndSay(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{"1", 82350},
		{"3113322113", 329356}, 
	}

	for _, tt := range tests {
		got := LengthOf40thLookAndSay(tt.input)
		if got != tt.want {
			t.Errorf("LengthOf40thLookAndSay(%q) = %d, want %d", tt.input, got, tt.want)
		}
	}
}
