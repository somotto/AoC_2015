package part2

import (
	"testing"
)

func TestLengthOf50thLookAndSay(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{"1", 1166642},
		{"3113322113", 4666278}, 
	}

	for _, tt := range tests {
		got := LengthOf50thLookAndSay(tt.input)
		if got != tt.want {
			t.Errorf("LengthOf50thLookAndSay(%q) = %d, want %d", tt.input, got, tt.want)
		}
	}
}
