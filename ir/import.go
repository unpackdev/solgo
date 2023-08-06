package ir

import (
	ast_pb "github.com/txpull/protos/dist/go/ast"
	ir_pb "github.com/txpull/protos/dist/go/ir"
	"github.com/txpull/solgo/ast"
)

type Import struct {
	Unit         *ast.Import     `json:"-"`
	Id           int64           `json:"id"`
	NodeType     ast_pb.NodeType `json:"node_type"`
	AbsolutePath string          `json:"absolute_path"`
	File         string          `json:"file"`
	UnitAlias    string          `json:"unit_alias"`
	SourceUnitId int64           `json:"source_unit_id"`
	ContractId   int64           `json:"contract_id"`
}

func (i *Import) GetId() int64 {
	return i.Id
}

func (i *Import) GetNodeType() ast_pb.NodeType {
	return i.NodeType
}

func (i *Import) GetAST() *ast.Import {
	return i.Unit
}

func (i *Import) GetAbsolutePath() string {
	return i.AbsolutePath
}

func (i *Import) GetFile() string {
	return i.File
}

func (i *Import) GetUnitAlias() string {
	return i.UnitAlias
}

func (i *Import) GetSourceUnitId() int64 {
	return i.SourceUnitId
}

func (i *Import) GetContractId() int64 {
	return i.ContractId
}

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
		su := sourceUnit.(*ast.SourceUnit[ast.Node[ast_pb.SourceUnit]])
		toReturn.ContractId = su.GetContract().GetId()
	}

	return toReturn
}
