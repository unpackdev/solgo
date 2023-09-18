package ast

import (
	ast_pb "github.com/unpackdev/protos/dist/go/ast"
	"github.com/unpackdev/solgo/parser"
)

// Declaration is a struct that contains information about a variable declaration in the AST.
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

// NewDeclaration creates a new Declaration instance.
func NewDeclaration(b *ASTBuilder) *Declaration {
	return &Declaration{
		ASTBuilder:      b,
		Id:              b.GetNextID(),
		IsStateVariable: false,
		IsConstant:      false,
	}
}

// SetReferenceDescriptor sets the reference descriptions of the VariableDeclaration node.
func (v *Declaration) SetReferenceDescriptor(refId int64, refDesc *TypeDescription) bool {
	if v.TypeName != nil {
		return v.TypeName.SetReferenceDescriptor(refId, refDesc)
	}

	return false
}

// GetId returns the ID of the Declaration.
func (d *Declaration) GetId() int64 {
	return d.Id
}

// GetType returns the NodeType of the Declaration.
func (d *Declaration) GetType() ast_pb.NodeType {
	return d.NodeType
}

// GetSrc returns the SrcNode of the Declaration.
func (d *Declaration) GetSrc() SrcNode {
	return d.Src
}

// GetName returns the name of the Declaration.
func (d *Declaration) GetName() string {
	return d.Name
}

// GetTypeName returns the TypeName of the Declaration.
func (d *Declaration) GetTypeName() *TypeName {
	return d.TypeName
}

// GetScope returns the scope of the Declaration.
func (d *Declaration) GetScope() int64 {
	return d.Scope
}

// GetStateMutability returns the state mutability of the Declaration.
func (d *Declaration) GetStateMutability() ast_pb.Mutability {
	return d.StateMutability
}

// GetVisibility returns the visibility of the Declaration.
func (d *Declaration) GetVisibility() ast_pb.Visibility {
	return d.Visibility
}

// GetStorageLocation returns the storage location of the Declaration.
func (d *Declaration) GetStorageLocation() ast_pb.StorageLocation {
	return d.StorageLocation
}

// GetIsConstant returns whether or not the Declaration is constant.
func (d *Declaration) GetIsConstant() bool {
	return d.IsConstant
}

// GetIsStateVariable returns whether or not the Declaration is a state variable.
func (d *Declaration) GetIsStateVariable() bool {
	return d.IsStateVariable
}

// GetTypeDescription returns the TypeDescription of the Declaration.
func (d *Declaration) GetTypeDescription() *TypeDescription {
	if d.TypeName != nil {
		return d.TypeName.GetTypeDescription()
	}
	return nil
}

// GetNodes returns the nodes associated with the Declaration.
func (d *Declaration) GetNodes() []Node[NodeType] {
	if d.TypeName != nil {
		return []Node[NodeType]{d.TypeName}
	}

	return nil
}

// ToProto converts the Declaration to its corresponding protocol buffer representation.
func (d *Declaration) ToProto() NodeType {
	toReturn := &ast_pb.Declaration{
		Id:              d.GetId(),
		Name:            d.GetName(),
		NodeType:        d.GetType(),
		Scope:           d.GetScope(),
		Src:             d.GetSrc().ToProto(),
		Mutability:      d.GetStateMutability(),
		StorageLocation: d.GetStorageLocation(),
		Visibility:      d.GetVisibility(),
		IsConstant:      d.GetIsConstant(),
		IsStateVariable: d.GetIsStateVariable(),
		TypeDescription: d.GetTypeDescription().ToProto(),
	}

	if d.GetTypeName() != nil {
		toReturn.TypeName = d.GetTypeName().ToProto().(*ast_pb.TypeName)
	}

	return toReturn
}

// ParseVariableDeclaration parses a VariableDeclaration and stores the relevant information in the Declaration.
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

	d.currentVariables = append(d.currentVariables, d)
}
