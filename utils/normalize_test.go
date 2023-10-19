package utils

import (
	"testing"
)

func TestNormalize(t *testing.T) {
	normalizer := NewNormalizeType()

	tests := []struct {
		name  string
		input string
		want  NormalizationInfo
	}{
		{
			name:  "standard uint",
			input: "uint",
			want:  NormalizationInfo{TypeName: "uint256", Normalized: true},
		},
		{
			name:  "standard int",
			input: "int",
			want:  NormalizationInfo{TypeName: "int256", Normalized: true},
		},
		{
			name:  "enum type",
			input: "enum",
			want:  NormalizationInfo{TypeName: "uint8", Normalized: true},
		},
		{
			name:  "custom struct type",
			input: "CustomType",
			want:  NormalizationInfo{TypeName: "CustomType", Normalized: false},
		},
		// Add more test cases as needed.
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := normalizer.Normalize(tt.input)
			if got != tt.want {
				t.Errorf("Normalize(%q) = %v; want %v", tt.input, got, tt.want)
			}
		})
	}
}
