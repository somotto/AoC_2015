package utils

import (
	"testing"
)

// Import your utils package

func TestWrappingPaper(t *testing.T) {
	tests := []struct {
		l, w, h int
		want    int
	}{
		{2, 3, 4, 58},  // Surface area: 2*(2*3 + 3*4 + 4*2) = 52, Smallest side area: 2*3 = 6, Total = 58
		{1, 1, 10, 43}, // Surface area: 2*(1*1 + 1*10 + 10*1) = 42, Smallest side area: 1*1 = 1, Total = 43
		{5, 1, 3, 49},  // Surface area: 2*(5*1 + 1*3 + 3*5) = 46, Smallest side area: 1*3 = 4, Total = 50
		{7, 4, 5, 186}, // Surface area: 2*(7*4 + 4*5 + 5*7) = 150, Smallest side area: 4*5 = 8, Total = 158
	}

	for _, tt := range tests {
		got := WrappingPaper(tt.l, tt.w, tt.h)
		if got != tt.want {
			t.Errorf("WrappingPaper(%d, %d, %d) got %d, want %d", tt.l, tt.w, tt.h, got, tt.want)
		}
	}
}
