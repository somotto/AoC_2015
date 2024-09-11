package main

import (
	"testing"
)

func TestEncodeString(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{`abc"def\ghi`, `"abc\"def\\ghi"`},
		{"", `""`},
		{"hello", `"hello"`},
	}

	for _, tt := range tests {
		got := encodeString(tt.input)
		if got != tt.want {
			t.Errorf("encodeString(%q) got %q, want %q", tt.input, got, tt.want)
		}
	}
}

func TestMainFunction(t *testing.T) {
	tests := []struct {
		input              []string
		wantEncodedDiff int
	}{
		{
			input: []string{
				`abc"def\ghi`,
				`simple`,
				`\backslash\`,
			},
			wantEncodedDiff: 10, // Adjust based on manual calculations
		},
	}

	for _, tt := range tests {
		totalOriginalLength := 0
		totalEncodedLength := 0

		for _, line := range tt.input {
			originalLength := len(line)
			encoded := encodeString(line)
			encodedLength := len(encoded)

			totalOriginalLength += originalLength
			totalEncodedLength += encodedLength
		}

		got := totalEncodedLength - totalOriginalLength
		if got != tt.wantEncodedDiff {
			t.Errorf("MainFunction() got %d, want %d", got, tt.wantEncodedDiff)
		}
	}
}
