package utils

import (
	"testing"
)

// Import your utils package

func TestFindRibbonFeet(t *testing.T) {
	tests := []struct {
		l, w, h int
		want    int
	}{
		{2, 3, 4, 34},  // 2*(2+3) is the smallest perimeter (10 feet) + 2*3*4 = 24 feet for the bow, total 34 feet
		{1, 1, 10, 14}, // 2*(1+1) is the smallest perimeter (4 feet) + 1*1*10 = 10 feet for the bow, total 14 feet
		{5, 1, 3, 23},  // 2*(1+3) is the smallest perimeter (8 feet) + 5*1*3 = 15 feet for the bow, total 23 feet
		{7, 4, 5, 158}, // 2*(4+5) is the smallest perimeter (18 feet) + 7*4*5 = 56 feet for the bow, total 74 feet
	}

	for _, tt := range tests {
		got := FindRibbonFeet(tt.l, tt.w, tt.h)
		if got != tt.want {
			t.Errorf("FindRibbonFeet(%d, %d, %d) got %d, want %d", tt.l, tt.w, tt.h, got, tt.want)
		}
	}
}
