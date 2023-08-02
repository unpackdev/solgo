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

func (p *Parameter) GetScope() int64 {
	return p.Id
}

func (p *Parameter) GetTypeDescription() *TypeDescription {
	// Enum value type can have type name as nil
	if p.TypeName != nil {
		return p.TypeName.TypeDescription
	}
	return &TypeDescription{}
}

func (p *Parameter) GetVisibility() ast_pb.Visibility {
	return p.Visibility
}

func (p *Parameter) GetStateMutability() ast_pb.Mutability {
	return p.StateMutability
}

func (p *Parameter) GetStorageLocation() ast_pb.StorageLocation {
	return p.StorageLocation
}

func (p *Parameter) IsConstant() bool {
	return p.Constant
}

func (p *Parameter) IsStateVariable() bool {
	return p.StateVariable
}

func (p *Parameter) GetTypeName() *TypeName {
	return p.TypeName
}

func (p *Parameter) GetNodes() []Node[NodeType] {
	return nil
}

func (p *Parameter) ToProto() *ast_pb.Parameter {
	toReturn := &ast_pb.Parameter{
		Id:              p.GetId(),
		Name:            p.GetName(),
		NodeType:        p.GetType(),
		Src:             p.GetSrc().ToProto(),
		Scope:           p.GetScope(),
		Constant:        p.IsConstant(),
		StateVariable:   p.IsStateVariable(),
		StateMutability: p.GetStateMutability(),
		Visibility:      p.GetVisibility(),
		StorageLocation: p.GetStorageLocation(),
	}

	if p.GetTypeName() != nil {
		toReturn.TypeName = p.GetTypeName().ToProto().(*ast_pb.TypeName)
	}

	if p.GetTypeDescription() != nil {
		toReturn.TypeDescription = p.GetTypeDescription().ToProto()
	}

	return toReturn
}

func (p *Parameter) Parse(unit *SourceUnit[Node[ast_pb.SourceUnit]], fnNode Node[NodeType], plNode Node[*ast_pb.ParameterList], ctx *parser.ParameterDeclarationContext) {
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

func (p *Parameter) ParseEventParameter(unit *SourceUnit[Node[ast_pb.SourceUnit]], fnNode Node[NodeType], plNode Node[*ast_pb.ParameterList], ctx parser.IEventParameterContext) {
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

func (p *Parameter) ParseErrorParameter(unit *SourceUnit[Node[ast_pb.SourceUnit]], fnNode Node[NodeType], plNode Node[*ast_pb.ParameterList], ctx parser.IErrorParameterContext) {
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
