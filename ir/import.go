package ir

import (
	ast_pb "github.com/unpackdev/protos/dist/go/ast"
	ir_pb "github.com/unpackdev/protos/dist/go/ir"
	"github.com/unpackdev/solgo/ast"
)

// Import represents an import statement in the IR.
type Import struct {
	Unit         *ast.Import     `json:"ast"`
	Id           int64           `json:"id"`
	NodeType     ast_pb.NodeType `json:"node_type"`
	AbsolutePath string          `json:"absolute_path"`
	File         string          `json:"file"`
	UnitAlias    string          `json:"unit_alias"`
	SourceUnitId int64           `json:"source_unit_id"`
	ContractId   int64           `json:"contract_id"`
}

// GetId returns the unique identifier of the import statement.
func (i *Import) GetId() int64 {
	return i.Id
}

// GetNodeType returns the type of the node in the AST.
func (i *Import) GetNodeType() ast_pb.NodeType {
	return i.NodeType
}

// GetAST returns the AST (Abstract Syntax Tree) for the import statement.
func (i *Import) GetAST() *ast.Import {
	return i.Unit
}

// GetAbsolutePath returns the absolute path of the imported file.
func (i *Import) GetAbsolutePath() string {
	return i.AbsolutePath
}

// GetFile returns the file name of the imported file.
func (i *Import) GetFile() string {
	return i.File
}

// GetUnitAlias returns the alias used for the imported unit.
func (i *Import) GetUnitAlias() string {
	return i.UnitAlias
}

// GetSourceUnitId returns the ID of the source unit where the import statement is used.
func (i *Import) GetSourceUnitId() int64 {
	return i.SourceUnitId
}

// GetContractId returns the ID of the contract associated with the source unit.
func (i *Import) GetContractId() int64 {
	return i.ContractId
}

// ToProto returns the protocol buffer version of the import statement.
func (i *Import) ToProto() *ir_pb.Import {
	proto := &ir_pb.Import{
		Id:           i.GetId(),
		NodeType:     i.GetNodeType(),
		SourceUnitId: i.GetSourceUnitId(),
		ContractId:   i.GetContractId(),
		AbsolutePath: i.GetAbsolutePath(),
		File:         i.GetFile(),
		UnitAlias:    i.GetUnitAlias(),
	}

	return proto
}

// processImport processes the import statement and returns the Import.
func (b *Builder) processImport(unit *ast.Import) *Import {
	toReturn := &Import{
		Unit:         unit,
		Id:           unit.GetId(),
		NodeType:     unit.GetType(),
		AbsolutePath: unit.GetAbsolutePath(),
		File:         unit.GetFile(),
		UnitAlias:    unit.GetUnitAlias(),
		SourceUnitId: unit.GetSourceUnit(),
	}

	sourceUnit := b.astBuilder.GetTree().GetById(unit.GetSourceUnit())
	if sourceUnit != nil {
		if su, ok := sourceUnit.(*ast.SourceUnit[ast.Node[ast_pb.SourceUnit]]); ok {
			toReturn.ContractId = su.GetContract().GetId()
		}
	}

	return toReturn
}
