package utils

import (
	"testing"
)

func TestWrappingPaper(t *testing.T) {
	tests := []struct {
		l, w, h int
		want    int
	}{
		{2, 3, 4, 58}, 
		{1, 1, 10, 43}, 
		{5, 1, 3, 49},  
		{7, 4, 5, 186}, 
	}

	for _, tt := range tests {
		got := WrappingPaper(tt.l, tt.w, tt.h)
		if got != tt.want {
			t.Errorf("WrappingPaper(%d, %d, %d) got %d, want %d", tt.l, tt.w, tt.h, got, tt.want)
		}
	}
}
