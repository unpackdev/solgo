package ir

import (
	"testing"

	"github.com/stretchr/testify/assert"
	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/ast"
)

func TestContractMethods(t *testing.T) {
	// Initialize a Contract instance for testing
	contract := &Contract{
		Id:             1,
		NodeType:       ast_pb.NodeType_CONTRACT_DEFINITION,
		Kind:           ast_pb.NodeType_CONTRACT,
		Name:           "TestContract",
		License:        "MIT",
		Language:       LanguageSolidity,
		AbsolutePath:   "/path/to/contract",
		Symbols:        []ast.Symbol{},
		Imports:        []*Import{{}},
		Pragmas:        []*Pragma{{}},
		StateVariables: []*StateVariable{{}},
		Structs:        []*Struct{{}},
		Enums:          []*Enum{{}},
		Events:         []*Event{{}},
		Errors:         []*Error{{}},
		Constructor:    &Constructor{},
		Functions:      []*Function{{}},
		Fallback:       &Fallback{},
		Receive:        &Receive{},
	}

	assert.Equal(t, int64(1), contract.GetId())
	assert.Equal(t, ast_pb.NodeType_CONTRACT_DEFINITION, contract.GetNodeType())
	assert.Equal(t, "TestContract", contract.GetName())
	assert.Equal(t, "MIT", contract.GetLicense())
	assert.Equal(t, "/path/to/contract", contract.GetAbsolutePath())
	assert.Equal(t, LanguageSolidity, contract.GetLanguage())
	assert.Equal(t, ast_pb.NodeType_CONTRACT, contract.GetKind())
	assert.Equal(t, []*Import{{}}, contract.GetImports())
	assert.Equal(t, []*Pragma{{}}, contract.GetPragmas())
	assert.Equal(t, []*StateVariable{{}}, contract.GetStateVariables())
	assert.Equal(t, []*Struct{{}}, contract.GetStructs())
	assert.Equal(t, []*Enum{{}}, contract.GetEnums())
	assert.Equal(t, []*Event{{}}, contract.GetEvents())
	assert.Equal(t, []*Error{{}}, contract.GetErrors())
	assert.Equal(t, &Constructor{}, contract.GetConstructor())
	assert.Equal(t, []*Function{{}}, contract.GetFunctions())
	assert.Equal(t, &Fallback{}, contract.GetFallback())
	assert.Equal(t, &Receive{}, contract.GetReceive())
	assert.Equal(t, []ast.Symbol{}, contract.GetSymbols())
}
