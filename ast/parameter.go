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
	Scope           int64                  `json:"scope,omitempty"`
	Name            string                 `json:"name"`
	TypeName        *TypeName              `json:"type_name,omitempty"`
	StorageLocation ast_pb.StorageLocation `json:"storage_location,omitempty"`
	Visibility      ast_pb.Visibility      `json:"visibility,omitempty"`
	StateMutability ast_pb.Mutability      `json:"state_mutability,omitempty"`
	Constant        bool                   `json:"constant,omitempty"`
	StateVariable   bool                   `json:"state_variable,omitempty"`
}

func NewParameter(b *ASTBuilder) *Parameter {
	return &Parameter{
		ASTBuilder:      b,
		NodeType:        ast_pb.NodeType_VARIABLE_DECLARATION,
		Visibility:      ast_pb.Visibility_INTERNAL,
		StateMutability: ast_pb.Mutability_MUTABLE,
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

func (p *Parameter) GetName() string {
	return p.Name
}

func (p *Parameter) Parse(unit *SourceUnit[Node[ast_pb.SourceUnit]], fnNode Node[NodeType], plNode Node[ast_pb.ParametersList], ctx *parser.ParameterDeclarationContext) {
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
	typeName.Parse(unit, fnNode, p.GetId(), ctx.TypeName())
	p.TypeName = typeName
}

func (p *Parameter) ParseEventParameter(unit *SourceUnit[Node[ast_pb.SourceUnit]], fnNode Node[NodeType], plNode Node[ast_pb.ParametersList], ctx parser.IEventParameterContext) {
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

	p.StorageLocation = ast_pb.StorageLocation_MEMORY

	typeName := NewTypeName(p.ASTBuilder)
	typeName.Parse(unit, fnNode, p.GetId(), ctx.TypeName())

	if typeName.TypeDescription != nil {
		switch typeName.TypeDescription.TypeIdentifier {
		case "t_address":
			p.StateMutability = ast_pb.Mutability_NONPAYABLE
		case "t_address_payable":
			p.StateMutability = ast_pb.Mutability_PAYABLE
		}
	}

	p.TypeName = typeName
}

func (p *Parameter) ParseStructParameter(unit *SourceUnit[Node[ast_pb.SourceUnit]], contractNode Node[NodeType], structNode *StructDefinition, ctx parser.IStructMemberContext) {
	p.Id = p.GetNextID()
	p.Src = SrcNode{
		Id:          p.GetNextID(),
		Line:        int64(ctx.GetStart().GetLine()),
		Column:      int64(ctx.GetStart().GetColumn()),
		Start:       int64(ctx.GetStart().GetStart()),
		End:         int64(ctx.GetStop().GetStop()),
		Length:      int64(ctx.GetStop().GetStop() - ctx.GetStart().GetStart() + 1),
		ParentIndex: structNode.GetId(),
	}
	p.Scope = contractNode.GetId()

	if ctx.Identifier() != nil {
		p.Name = ctx.Identifier().GetText()
	}

	//p.StorageLocation = p.getStorageLocationFromCtx(ctx)

	typeName := NewTypeName(p.ASTBuilder)
	typeName.Parse(unit, contractNode, p.GetId(), ctx.TypeName())
	p.TypeName = typeName
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
