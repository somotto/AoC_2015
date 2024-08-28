package utils

import (
	"testing"
)

func TestFindRibbonFeet(t *testing.T) {
	tests := []struct {
		l, w, h int
		want    int
	}{
		{2, 3, 4, 34}, 
		{1, 1, 10, 14}, 
		{5, 1, 3, 23},  
		{7, 4, 5, 158},
	}

	for _, tt := range tests {
		got := FindRibbonFeet(tt.l, tt.w, tt.h)
		if got != tt.want {
			t.Errorf("FindRibbonFeet(%d, %d, %d) got %d, want %d", tt.l, tt.w, tt.h, got, tt.want)
		}
	}
}
