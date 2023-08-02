package ast

import (
	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/parser"
)

type Declaration struct {
	*ASTBuilder

	IsConstant      bool                   `json:"is_constant"`
	Id              int64                  `json:"id"`
	StateMutability ast_pb.Mutability      `json:"state_mutability"`
	Name            string                 `json:"name"`
	NodeType        ast_pb.NodeType        `json:"node_type"`
	Scope           int64                  `json:"scope"`
	Src             SrcNode                `json:"src"`
	IsStateVariable bool                   `json:"is_state_variable"`
	StorageLocation ast_pb.StorageLocation `json:"storage_location"`
	TypeName        *TypeName              `json:"type_name"`
	Visibility      ast_pb.Visibility      `json:"visibility"`
}

func NewDeclaration(b *ASTBuilder) *Declaration {
	return &Declaration{
		ASTBuilder:      b,
		Id:              b.GetNextID(),
		IsStateVariable: false,
		IsConstant:      false,
	}
}

func (d *Declaration) GetId() int64 {
	return d.Id
}

func (d *Declaration) GetType() ast_pb.NodeType {
	return d.NodeType
}

func (d *Declaration) GetSrc() SrcNode {
	return d.Src
}

func (d *Declaration) GetName() string {
	return d.Name
}

func (d *Declaration) GetTypeName() *TypeName {
	return d.TypeName
}

func (d *Declaration) GetScope() int64 {
	return d.Scope
}

func (d *Declaration) GetStateMutability() ast_pb.Mutability {
	return d.StateMutability
}

func (d *Declaration) GetVisibility() ast_pb.Visibility {
	return d.Visibility
}

func (d *Declaration) GetStorageLocation() ast_pb.StorageLocation {
	return d.StorageLocation
}

func (d *Declaration) GetIsConstant() bool {
	return d.IsConstant
}

func (d *Declaration) GetIsStateVariable() bool {
	return d.IsStateVariable
}

func (d *Declaration) GetPathNode() *PathNode {
	return nil
}

func (d *Declaration) GetReferencedDeclaration() int64 {
	return 0
}

func (d *Declaration) GetKeyType() *TypeName {
	return nil
}

func (d *Declaration) GetValueType() *TypeName {
	return nil
}

func (d *Declaration) GetTypeDescription() *TypeDescription {
	if d.TypeName != nil {
		return d.TypeName.GetTypeDescription()
	}
	return nil
}

func (d *Declaration) GetNodes() []Node[NodeType] {
	return nil
}

func (d *Declaration) ToProto() *ast_pb.Declaration {
	toReturn := &ast_pb.Declaration{
		Id:              d.Id,
		Name:            d.Name,
		NodeType:        d.NodeType,
		Scope:           d.Scope,
		Src:             d.Src.ToProto(),
		Mutability:      d.StateMutability,
		StorageLocation: d.StorageLocation,
		Visibility:      d.Visibility,
		IsConstant:      d.IsConstant,
		IsStateVariable: d.IsStateVariable,
		TypeDescription: d.GetTypeDescription().ToProto(),
	}

	if d.GetTypeName() != nil {
		toReturn.TypeName = d.GetTypeName().ToProto().(*ast_pb.TypeName)
	}

	return toReturn
}

func (d *Declaration) ParseVariableDeclaration(
	unit *SourceUnit[Node[ast_pb.SourceUnit]],
	contractNode Node[NodeType],
	fnNode Node[NodeType],
	bodyNode *BodyNode,
	vDeclar *VariableDeclaration,
	ctx parser.IVariableDeclarationContext,
) {
	d.NodeType = ast_pb.NodeType_VARIABLE_DECLARATION
	d.Src = SrcNode{
		Id:          d.GetNextID(),
		Line:        int64(ctx.GetStart().GetLine()),
		Column:      int64(ctx.GetStart().GetColumn()),
		Start:       int64(ctx.GetStart().GetStart()),
		End:         int64(ctx.GetStop().GetStop()),
		Length:      int64(ctx.GetStop().GetStop() - ctx.GetStart().GetStart() + 1),
		ParentIndex: vDeclar.GetId(),
	}

	d.StorageLocation = getStorageLocationFromDataLocationCtx(ctx.DataLocation())
	d.Visibility = ast_pb.Visibility_INTERNAL
	d.StateMutability = ast_pb.Mutability_MUTABLE

	if ctx.Identifier() != nil {
		d.Name = ctx.Identifier().GetText()
	}

	d.Scope = bodyNode.GetId()

	typeName := NewTypeName(d.ASTBuilder)
	typeName.Parse(unit, fnNode, d.GetId(), ctx.TypeName())
	d.TypeName = typeName
}
