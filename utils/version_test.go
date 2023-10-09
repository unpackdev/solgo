package utils

import (
	"testing"
)

func TestParseSemanticVersion(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  SemanticVersion
	}{
		{
			name:  "standard version",
			input: "1.2.3",
			want:  SemanticVersion{Major: 1, Minor: 2, Patch: 3},
		},
		{
			name:  "zero version",
			input: "0.0.0",
			want:  SemanticVersion{Major: 0, Minor: 0, Patch: 0},
		},
		// Add more test cases as needed.
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ParseSemanticVersion(tt.input)
			if got != tt.want {
				t.Errorf("ParseSemanticVersion(%q) = %v; want %v", tt.input, got, tt.want)
			}
		})
	}
}
