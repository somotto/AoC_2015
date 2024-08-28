package houses

import (
	"testing"
)

func TestCountHousesWithSantaRobo(t *testing.T) {
	tests := []struct {
		directions string
		want       int
	}{
		{"", 1},
		{"^v", 3},
		{"^>v<", 3},
		{"^v^v^v^v^v", 11},
		{"<^", 3},
		{">>", 2},
	}

	for _, tt := range tests {
		got := CountHousesWithSantaRobo(tt.directions)
		if got != tt.want {
			t.Errorf("CountHousesWithRoboSanta(%q) got %d, want %d", tt.directions, got, tt.want)
		}
	}
}
