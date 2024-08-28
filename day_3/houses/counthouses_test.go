package houses

import (
	"testing"
)

func TestCountHouses(t *testing.T) {
	tests := []struct {
		directions string
		want       int
	}{
		{">", 2},
		{"^>v<", 4},
		{"^v^v^v^v^v", 2},
		{"", 1},
		{">>^^<<vv", 8},
	}

	for _, tt := range tests {
		got := CountHouses(tt.directions)
		if got != tt.want {
			t.Errorf("CountHouses(%q) got %d, want %d", tt.directions, got, tt.want)
		}
	}
}
