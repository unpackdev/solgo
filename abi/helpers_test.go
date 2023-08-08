package abi

import (
	"testing"

	"github.com/stretchr/testify/assert"
	ast_pb "github.com/txpull/protos/dist/go/ast"
)

func TestNormalizeStateMutability(t *testing.T) {
	builder := &Builder{}

	testCases := []struct {
		name     string
		input    ast_pb.Mutability
		expected string
	}{
		{"Pure", ast_pb.Mutability_PURE, "pure"},
		{"View", ast_pb.Mutability_VIEW, "view"},
		{"NonPayable", ast_pb.Mutability_NONPAYABLE, "nonpayable"},
		{"Payable", ast_pb.Mutability_PAYABLE, "payable"},
		{"Default", 0, "view"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := builder.normalizeStateMutability(tc.input)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestIsMappingType(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected bool
	}{
		{"MappingType", "mapping(uint => address)", true},
		{"NonMappingType", "uint256", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := isMappingType(tc.input)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestIsContractType(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected bool
	}{
		{"ContractType", "contract MyContract", true},
		{"NonContractType", "uint256", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := isContractType(tc.input)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestParseMappingType(t *testing.T) {
	testCases := []struct {
		name            string
		input           string
		expectedSuccess bool
		expectedInput   []string
		expectedOutput  []string
	}{
		{"ValidMapping", "mapping(uint => address)", true, []string{"uint"}, []string{"address"}},
		{"NestedMapping", "mapping(uint => mapping(address => bool))", true, []string{"uint", "address"}, []string{"bool"}},
		{"InvalidMapping", "uint256", false, []string{}, []string{}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			success, input, output := parseMappingType(tc.input)
			assert.Equal(t, tc.expectedSuccess, success)
			assert.Equal(t, tc.expectedInput, input)
			assert.Equal(t, tc.expectedOutput, output)
		})
	}
}

func TestNormalizeTypeName(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{"Uint", "uint", "uint256"},
		{"Int", "int", "int256"},
		{"Bool", "bool", "bool"},
		{"Bytes", "bytes", "bytes"},
		{"String", "string", "string"},
		{"Address", "address", "address"},
		{"AddressPayable", "addresspayable", "address"},
		{"Tuple", "tuple", "tuple"},
		{"Default", "customType", "customType"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := normalizeTypeName(tc.input)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestNormalizeTypeNameWithStatus(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
		status   bool
	}{
		{"Uint", "uint", "uint256", true},
		{"Int", "int", "int256", true},
		{"Bool", "bool", "bool", true},
		{"Bytes", "bytes", "bytes", true},
		{"String", "string", "string", true},
		{"Address", "address", "address", true},
		{"AddressPayable", "addresspayable", "address", true},
		{"Tuple", "tuple", "tuple", true},
		{"Default", "customType", "customType", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, status := normalizeTypeNameWithStatus(tc.input)
			assert.Equal(t, tc.expected, result)
			assert.Equal(t, tc.status, status)
		})
	}
}
