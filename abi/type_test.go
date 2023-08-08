package abi

import (
	"testing"

	"github.com/stretchr/testify/assert"
	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/ast"
)

func TestTypeNormalizeStateMutability(t *testing.T) {
	tests := []struct {
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

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Builder{}
			result := b.normalizeStateMutability(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestTypeIsMappingType(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"MappingType", "mapping(uint => address)", true},
		{"NonMappingType", "uint256", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isMappingType(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestTypeIsContractType(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"ContractType", "contract MyContract", true},
		{"NonContractType", "uint256", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isContractType(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestResolveType(t *testing.T) {
	tests := []struct {
		name     string
		input    *ast.TypeDescription
		expected string
	}{
		{"MappingType", &ast.TypeDescription{TypeIdentifier: "t_mapping(uint => address)"}, "mapping"},
		{"StructType", &ast.TypeDescription{TypeIdentifier: "t_struct"}, "struct"},
		{"ContractType", &ast.TypeDescription{TypeIdentifier: "t_contract"}, "contract"},
		{"EnumType", &ast.TypeDescription{TypeIdentifier: "t_enum"}, "enum"},
		{"ErrorType", &ast.TypeDescription{TypeIdentifier: "t_error"}, "error"},
		{"OtherType", &ast.TypeDescription{TypeIdentifier: "t_other", TypeString: "uint256"}, "uint256"},
	}

	tr := &TypeResolver{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tr.ResolveType(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestResolveMappingType(t *testing.T) {
	tests := []struct {
		name            string
		input           *ast.TypeDescription
		expectedInputs  []MethodIO
		expectedOutputs []MethodIO
	}{
		{
			"MappingType",
			&ast.TypeDescription{TypeString: "mapping(uint => address)"},
			[]MethodIO{
				{
					Name:         "",
					Type:         "uint256",
					InternalType: "uint256",
				},
			},
			[]MethodIO{
				{
					Name:         "",
					Type:         "address",
					InternalType: "address",
				},
			},
		},
	}

	tr := &TypeResolver{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inputs, outputs := tr.ResolveMappingType(tt.input)
			assert.Equal(t, tt.expectedInputs, inputs)
			assert.Equal(t, tt.expectedOutputs, outputs)
		})
	}
}

func TestDiscoverType(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected Type
	}{
		{
			"UintType",
			"uint",
			Type{
				Type:         "uint256",
				InternalType: "uint256",
				Outputs:      make([]Type, 0),
			},
		},
		{
			"AddressType",
			"address",
			Type{
				Type:         "address",
				InternalType: "address",
				Outputs:      make([]Type, 0),
			},
		},
	}

	tr := &TypeResolver{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tr.discoverType(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}
