package ir

import (
	"testing"

	"github.com/stretchr/testify/assert"
	ast_pb "github.com/unpackdev/protos/dist/go/ast"
)

func TestRootSourceUnitMethods(t *testing.T) {
	// Create a new RootSourceUnit instance
	rootSourceUnitInstance := &RootSourceUnit{
		NodeType:          ast_pb.NodeType(1),
		EntryContractId:   1,
		EntryContractName: "TestContract",
		ContractsCount:    1,
		Contracts:         []*Contract{{Id: 1, Name: "TestContract"}},
		Links:             []*Link{{Location: "https://unpack.dev"}},
	}

	// Test GetNodeType method
	assert.Equal(t, ast_pb.NodeType(1), rootSourceUnitInstance.GetNodeType())

	// Test GetContracts method
	assert.Equal(t, []*Contract{{Id: 1, Name: "TestContract"}}, rootSourceUnitInstance.GetContracts())

	// Test GetContractByName method
	assert.Equal(t, &Contract{Id: 1, Name: "TestContract"}, rootSourceUnitInstance.GetContractByName("TestContract"))

	// Test GetContractById method
	assert.Equal(t, &Contract{Id: 1, Name: "TestContract"}, rootSourceUnitInstance.GetContractById(1))

	// Test GetEntryContract method
	assert.Equal(t, &Contract{Id: 1, Name: "TestContract"}, rootSourceUnitInstance.GetEntryContract())

	// Test HasContracts method
	assert.Equal(t, true, rootSourceUnitInstance.HasContracts())

	// Test GetContractsCount method
	assert.Equal(t, int32(1), rootSourceUnitInstance.GetContractsCount())

	// Test GetLinks method
	assert.Equal(t, []*Link{{Location: "https://unpack.dev"}}, rootSourceUnitInstance.GetLinks())
}
