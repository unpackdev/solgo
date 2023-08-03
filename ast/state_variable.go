package ast

import (
	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/parser"
)

type StateVariableDeclaration struct {
	*ASTBuilder

	Id              int64                  `json:"id"`
	Name            string                 `json:"name"`
	IsConstant      bool                   `json:"is_constant"`
	IsStateVariable bool                   `json:"is_state_variable"`
	NodeType        ast_pb.NodeType        `json:"node_type"`
	Src             SrcNode                `json:"src"`
	Scope           int64                  `json:"scope"`
	TypeDescription *TypeDescription       `json:"type_description"`
	Visibility      ast_pb.Visibility      `json:"visibility"`
	StorageLocation ast_pb.StorageLocation `json:"storage_location"`
	StateMutability ast_pb.Mutability      `json:"mutability"`
	TypeName        *TypeName              `json:"type_name"`
}

func NewStateVariableDeclaration(b *ASTBuilder) *StateVariableDeclaration {
	return &StateVariableDeclaration{
		ASTBuilder:      b,
		Id:              b.GetNextID(),
		NodeType:        ast_pb.NodeType_VARIABLE_DECLARATION,
		StateMutability: ast_pb.Mutability_MUTABLE,
		StorageLocation: ast_pb.StorageLocation_DEFAULT,
		IsStateVariable: true,
	}
}

// SetReferenceDescriptor sets the reference descriptions of the StateVariableDeclaration node.
func (v *StateVariableDeclaration) SetReferenceDescriptor(refId int64, refDesc *TypeDescription) bool {
	return false
}

func (v *StateVariableDeclaration) GetId() int64 {
	return v.Id
}

func (v *StateVariableDeclaration) GetType() ast_pb.NodeType {
	return v.NodeType
}

func (v *StateVariableDeclaration) GetSrc() SrcNode {
	return v.Src
}

func (v *StateVariableDeclaration) GetTypeDescription() *TypeDescription {
	return v.TypeDescription
}

func (v *StateVariableDeclaration) GetName() string {
	return v.Name
}

func (v *StateVariableDeclaration) GetVisibility() ast_pb.Visibility {
	return v.Visibility
}

func (v *StateVariableDeclaration) GetStorageLocation() ast_pb.StorageLocation {
	return v.StorageLocation
}

func (v *StateVariableDeclaration) GetStateMutability() ast_pb.Mutability {
	return v.StateMutability
}

func (v *StateVariableDeclaration) GetTypeName() *TypeName {
	return v.TypeName
}

func (v *StateVariableDeclaration) GetReferencedDeclaration() int64 {
	return v.TypeName.ReferencedDeclaration
}

func (v *StateVariableDeclaration) GetNodes() []Node[NodeType] {
	return nil
}

func (v *StateVariableDeclaration) GetScope() int64 {
	return v.Scope
}

func (v *StateVariableDeclaration) ToProto() NodeType {
	proto := ast_pb.StateVariable{
		Id:              v.Id,
		NodeType:        v.NodeType,
		Src:             v.Src.ToProto(),
		Name:            v.Name,
		Visibility:      v.Visibility,
		StorageLocation: v.StorageLocation,
		StateMutability: v.StateMutability,
		Scope:           v.Scope,
		IsConstant:      v.IsConstant,
		IsStateVariable: v.IsStateVariable,
	}

	if v.GetTypeName() != nil {
		proto.TypeName = v.GetTypeName().ToProto().(*ast_pb.TypeName)
	}

	if v.GetTypeDescription() != nil {
		proto.TypeDescription = v.GetTypeDescription().ToProto()
	}

	return NewTypedStruct(&proto, "Variable")
}

func (v *StateVariableDeclaration) Parse(
	unit *SourceUnit[Node[ast_pb.SourceUnit]],
	contractNode Node[NodeType],
	bodyCtx parser.IContractBodyElementContext,
	ctx *parser.StateVariableDeclarationContext,
) {
	v.Name = ctx.Identifier().GetText()
	v.Src = SrcNode{
		Id:          v.GetNextID(),
		Line:        int64(ctx.GetStart().GetLine()),
		Column:      int64(ctx.GetStart().GetColumn()),
		Start:       int64(ctx.GetStart().GetStart()),
		End:         int64(ctx.GetStop().GetStop()),
		Length:      int64(ctx.GetStop().GetStop() - ctx.GetStart().GetStart() + 1),
		ParentIndex: contractNode.GetId(),
	}
	v.Scope = contractNode.GetId()
	v.Visibility = v.getVisibilityFromCtx(ctx)

	for _, immutableCtx := range ctx.AllImmutable() {
		if immutableCtx != nil {
			v.StateMutability = ast_pb.Mutability_IMMUTABLE
		}
	}

	for _, constantCtx := range ctx.AllConstant() {
		v.IsConstant = constantCtx != nil
	}

	typeName := NewTypeName(v.ASTBuilder)
	typeName.Parse(unit, nil, v.Id, ctx.GetType_())
	v.TypeName = typeName
	v.TypeDescription = typeName.TypeDescription

	v.currentStateVariables = append(v.currentStateVariables, v)
}

func (v *StateVariableDeclaration) getVisibilityFromCtx(ctx *parser.StateVariableDeclarationContext) ast_pb.Visibility {
	visibilityMap := map[string]ast_pb.Visibility{
		"public":   ast_pb.Visibility_PUBLIC,
		"private":  ast_pb.Visibility_PRIVATE,
		"internal": ast_pb.Visibility_INTERNAL,
		"external": ast_pb.Visibility_EXTERNAL,
	}

	// Check each visibility context in the map
	if len(ctx.AllPublic()) > 0 {
		if v, ok := visibilityMap["public"]; ok {
			return v
		}
	} else if len(ctx.AllPrivate()) > 0 {
		if v, ok := visibilityMap["private"]; ok {
			return v
		}
	} else if len(ctx.AllInternal()) > 0 {
		if v, ok := visibilityMap["internal"]; ok {
			return v
		}
	}

	// If no visibility context matches, return the default value
	return ast_pb.Visibility_INTERNAL
}
