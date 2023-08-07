package ir

import (
	"testing"

	"github.com/stretchr/testify/assert"
	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/ast"
)

func TestImportMethods(t *testing.T) {
	// Create a new Import instance
	importInstance := &Import{
		Unit:         &ast.Import{},
		Id:           1,
		NodeType:     ast_pb.NodeType(1),
		AbsolutePath: "/path/to/import",
		File:         "import.sol",
		UnitAlias:    "alias",
		SourceUnitId: 1,
		ContractId:   1,
	}

	// Test GetId method
	assert.Equal(t, int64(1), importInstance.GetId())

	// Test GetNodeType method
	assert.Equal(t, ast_pb.NodeType(1), importInstance.GetNodeType())

	// Test GetAST method
	assert.IsType(t, &ast.Import{}, importInstance.GetAST())

	// Test GetAbsolutePath method
	assert.Equal(t, "/path/to/import", importInstance.GetAbsolutePath())

	// Test GetFile method
	assert.Equal(t, "import.sol", importInstance.GetFile())

	// Test GetUnitAlias method
	assert.Equal(t, "alias", importInstance.GetUnitAlias())

	// Test GetSourceUnitId method
	assert.Equal(t, int64(1), importInstance.GetSourceUnitId())

	// Test GetContractId method
	assert.Equal(t, int64(1), importInstance.GetContractId())
}
