package abis

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNormalizeTypeName(t *testing.T) {
	tests := []struct {
		name     string
		typeName string
		expected string
	}{
		{
			name:     "Test uint",
			typeName: "uint",
			expected: "uint256",
		},
		{
			name:     "Test int",
			typeName: "int",
			expected: "int256",
		},
		{
			name:     "Test addresspayable",
			typeName: "addresspayable",
			expected: "address",
		},
		{
			name:     "Test other",
			typeName: "other",
			expected: "other",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := normalizeTypeName(tt.typeName)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestIsMappingType(t *testing.T) {
	tests := []struct {
		name     string
		typeName string
		expected bool
	}{
		{
			name:     "Test mapping",
			typeName: "mapping(address=>uint256)",
			expected: true,
		},
		{
			name:     "Test non-mapping",
			typeName: "uint256",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isMappingType(tt.typeName)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestParseMappingType(t *testing.T) {
	tests := []struct {
		name        string
		abi         string
		wantMatched bool
		wantIn      []string
		wantOut     []string
	}{
		{
			name:        "Test Mapping",
			abi:         "mapping(address=>uint256)",
			wantMatched: true,
			wantIn:      []string{"address"},
			wantOut:     []string{"uint256"},
		},
		{
			name:        "Test Nested Mapping",
			abi:         "mapping(address=>mapping(address=>bool))",
			wantMatched: true,
			wantIn:      []string{"address", "address"},
			wantOut:     []string{"bool"},
		},
		{
			name:        "Test Struct",
			abi:         "struct(Person{name=>string, age=>uint256})",
			wantMatched: false,
			wantIn:      []string{},
			wantOut:     []string{},
		},
		{
			name:        "Test Tuple",
			abi:         "tuple(address, uint256, bool)",
			wantMatched: false,
			wantIn:      []string{},
			wantOut:     []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMatched, gotIn, gotOut := parseMappingType(tt.abi)
			assert.Equal(t, tt.wantMatched, gotMatched)
			assert.Equal(t, tt.wantIn, gotIn)
			assert.Equal(t, tt.wantOut, gotOut)
		})
	}
}
