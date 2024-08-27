package utils

import (
	"testing"
)

// Import your utils package

func TestSmallestSideArea(t *testing.T) {
	tests := []struct {
		l, w, h int
		want    int
	}{
		{2, 3, 4, 6},  // 2*3 is the smallest side area
		{1, 1, 10, 1}, // 1*1 is the smallest side area
		{5, 1, 3, 3},  // 1*3 is the smallest side area
		{7, 4, 5, 20}, // 4*5 is the smallest side area
	}

	for _, tt := range tests {
		got := SmallestSideArea(tt.l, tt.w, tt.h)
		if got != tt.want {
			t.Errorf("SmallestSideArea(%d, %d, %d) got %d, want %d", tt.l, tt.w, tt.h, got, tt.want)
		}
	}
}
