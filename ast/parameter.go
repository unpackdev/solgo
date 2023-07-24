package ast

import (
	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/parser"
)

type Parameter struct {
	*ASTBuilder

	Id              int64                  `json:"id"`
	NodeType        ast_pb.NodeType        `json:"node_type"`
	Src             SrcNode                `json:"src"`
	Scope           int64                  `json:"scope"`
	Name            string                 `json:"name"`
	TypeName        *TypeName              `json:"type_name,omitempty"`
	StorageLocation ast_pb.StorageLocation `json:"storage_location"`
	Visibility      ast_pb.Visibility      `json:"visibility"`
	Mutability      ast_pb.Mutability      `json:"mutability"`
}

func NewParameter(b *ASTBuilder) *Parameter {
	return &Parameter{
		ASTBuilder: b,
		NodeType:   ast_pb.NodeType_VARIABLE_DECLARATION,
		Visibility: ast_pb.Visibility_INTERNAL,
		Mutability: ast_pb.Mutability_MUTABLE,
	}
}

func (p *Parameter) GetId() int64 {
	return p.Id
}

func (p *Parameter) GetType() ast_pb.NodeType {
	return p.NodeType
}

func (p *Parameter) GetSrc() SrcNode {
	return p.Src
}

func (p *Parameter) Parse(unit *SourceUnit[Node], fnNode Node, plNode Node, ctx *parser.ParameterDeclarationContext) {
	p.Id = p.GetNextID()
	p.Src = SrcNode{
		Id:          p.GetNextID(),
		Line:        int64(ctx.GetStart().GetLine()),
		Column:      int64(ctx.GetStart().GetColumn()),
		Start:       int64(ctx.GetStart().GetStart()),
		End:         int64(ctx.GetStop().GetStop()),
		Length:      int64(ctx.GetStop().GetStop() - ctx.GetStart().GetStart() + 1),
		ParentIndex: plNode.GetId(),
	}
	p.Scope = fnNode.GetId()

	if ctx.Identifier() != nil {
		p.Name = ctx.Identifier().GetText()
	}

	p.StorageLocation = p.getStorageLocationFromCtx(ctx)

	typeName := NewTypeName(p.ASTBuilder)
	typeName.Parse(unit, fnNode, p, ctx.TypeName())

}

func (p *Parameter) getStorageLocationFromCtx(ctx *parser.ParameterDeclarationContext) ast_pb.StorageLocation {
	storageLocationMap := map[string]ast_pb.StorageLocation{
		"memory":   ast_pb.StorageLocation_MEMORY,
		"storage":  ast_pb.StorageLocation_STORAGE,
		"calldata": ast_pb.StorageLocation_CALLDATA,
	}

	if ctx.DataLocation() != nil {
		if s, ok := storageLocationMap[ctx.DataLocation().GetText()]; ok {
			return s
		}
	}

	return ast_pb.StorageLocation_MEMORY
}
