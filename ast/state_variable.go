package ast

import (
	"strings"

	v3 "github.com/cncf/xds/go/xds/type/v3"
	ast_pb "github.com/unpackdev/protos/dist/go/ast"
	"github.com/unpackdev/solgo/parser"
)

// StateVariableDeclaration represents a state variable declaration in the Solidity abstract syntax tree (AST).
type StateVariableDeclaration struct {
	*ASTBuilder                            // Embedding the ASTBuilder for common functionality
	Id              int64                  `json:"id"`                // Unique identifier for the state variable declaration
	Name            string                 `json:"name"`              // Name of the state variable
	Constant        bool                   `json:"is_constant"`       // Indicates if the state variable is constant
	StateVariable   bool                   `json:"is_state_variable"` // Indicates if the declaration is a state variable
	NodeType        ast_pb.NodeType        `json:"node_type"`         // Type of the node (VARIABLE_DECLARATION for state variable declaration)
	Src             SrcNode                `json:"src"`               // Source information about the state variable declaration
	Scope           int64                  `json:"scope"`             // Scope of the state variable declaration
	TypeDescription *TypeDescription       `json:"type_description"`  // Type description of the state variable declaration
	Visibility      ast_pb.Visibility      `json:"visibility"`        // Visibility of the state variable declaration
	StorageLocation ast_pb.StorageLocation `json:"storage_location"`  // Storage location of the state variable declaration
	StateMutability ast_pb.Mutability      `json:"mutability"`        // State mutability of the state variable declaration
	TypeName        *TypeName              `json:"type_name"`         // Type name of the state variable
	InitialValue    Node[NodeType]         `json:"initial_value"`     // Initial value of the state variable
}

// NewStateVariableDeclaration creates a new StateVariableDeclaration instance.
func NewStateVariableDeclaration(b *ASTBuilder) *StateVariableDeclaration {
	return &StateVariableDeclaration{
		ASTBuilder:      b,
		Id:              b.GetNextID(),
		NodeType:        ast_pb.NodeType_VARIABLE_DECLARATION,
		StateMutability: ast_pb.Mutability_MUTABLE,
		StorageLocation: ast_pb.StorageLocation_DEFAULT,
		StateVariable:   true,
	}
}

// SetReferenceDescriptor sets the reference descriptions of the StateVariableDeclaration node.
func (v *StateVariableDeclaration) SetReferenceDescriptor(refId int64, refDesc *TypeDescription) bool {
	v.TypeDescription = refDesc

	if v.TypeName != nil {
		v.TypeName.SetReferenceDescriptor(refId, refDesc)
	}

	// Lets update the parent node as well in case that type description is not set...
	parentNodeId := v.GetSrc().GetParentIndex()

	if parentNodeId > 0 {
		if parentNode := v.GetTree().GetById(parentNodeId); parentNode != nil {
			if parentNode.GetTypeDescription() == nil {
				parentNode.SetReferenceDescriptor(refId, refDesc)
			}
		}
	}

	return false
}

// GetId returns the unique identifier of the state variable declaration.
func (v *StateVariableDeclaration) GetId() int64 {
	return v.Id
}

// GetType returns the type of the node, which is 'VARIABLE_DECLARATION' for a state variable declaration.
func (v *StateVariableDeclaration) GetType() ast_pb.NodeType {
	return v.NodeType
}

// GetSrc returns the source information about the state variable declaration.
func (v *StateVariableDeclaration) GetSrc() SrcNode {
	return v.Src
}

// GetTypeDescription returns the type description of the state variable declaration.
func (v *StateVariableDeclaration) GetTypeDescription() *TypeDescription {
	return v.TypeDescription
}

// GetName returns the name of the state variable.
func (v *StateVariableDeclaration) GetName() string {
	return v.Name
}

// GetVisibility returns the visibility of the state variable declaration.
func (v *StateVariableDeclaration) GetVisibility() ast_pb.Visibility {
	return v.Visibility
}

// GetStorageLocation returns the storage location of the state variable declaration.
func (v *StateVariableDeclaration) GetStorageLocation() ast_pb.StorageLocation {
	return v.StorageLocation
}

// GetStateMutability returns the state mutability of the state variable declaration.
func (v *StateVariableDeclaration) GetStateMutability() ast_pb.Mutability {
	return v.StateMutability
}

// GetTypeName returns the type name of the state variable declaration.
func (v *StateVariableDeclaration) GetTypeName() *TypeName {
	return v.TypeName
}

// GetReferencedDeclaration returns the referenced declaration of the type name in the state variable declaration.
func (v *StateVariableDeclaration) GetReferencedDeclaration() int64 {
	return v.TypeName.ReferencedDeclaration
}

// GetNodes returns the type name node in the state variable declaration.
func (v *StateVariableDeclaration) GetNodes() []Node[NodeType] {
	toReturn := []Node[NodeType]{}

	if v.TypeName != nil {
		toReturn = append(toReturn, v.TypeName)
	}

	if v.InitialValue != nil {
		toReturn = append(toReturn, v.InitialValue)
	}

	return toReturn
}

// GetScope returns the scope of the state variable declaration.
func (v *StateVariableDeclaration) GetScope() int64 {
	return v.Scope
}

// IsConstant returns whether the state variable declaration is constant.
func (v *StateVariableDeclaration) IsConstant() bool {
	return v.Constant
}

// IsStateVariable returns whether the declaration is a state variable.
func (v *StateVariableDeclaration) IsStateVariable() bool {
	return v.StateVariable
}

// GetInitialValue returns the initial value of the state variable declaration.
func (v *StateVariableDeclaration) GetInitialValue() Node[NodeType] {
	return v.InitialValue
}

// ToProto returns the protobuf representation of the state variable declaration.
func (v *StateVariableDeclaration) ToProto() NodeType {
	proto := ast_pb.StateVariable{
		Id:              v.GetId(),
		NodeType:        v.GetType(),
		Src:             v.GetSrc().ToProto(),
		Name:            v.GetName(),
		Visibility:      v.GetVisibility(),
		StorageLocation: v.GetStorageLocation(),
		StateMutability: v.GetStateMutability(),
		Scope:           v.GetScope(),
		IsConstant:      v.IsConstant(),
		IsStateVariable: v.IsStateVariable(),
	}

	if v.GetTypeName() != nil {
		proto.TypeName = v.GetTypeName().ToProto().(*ast_pb.TypeName)
	}

	if v.GetTypeDescription() != nil {
		proto.TypeDescription = v.GetTypeDescription().ToProto()
	}

	if v.GetInitialValue() != nil {
		proto.InitialValue = v.GetInitialValue().ToProto().(*v3.TypedStruct)
	}

	return NewTypedStruct(&proto, "Variable")
}

// Parse parses a state variable declaration from the provided parser.StateVariableDeclarationContext and updates the current instance.
func (v *StateVariableDeclaration) Parse(
	unit *SourceUnit[Node[ast_pb.SourceUnit]],
	contractNode Node[NodeType],
	bodyCtx parser.IContractBodyElementContext,
	ctx *parser.StateVariableDeclarationContext,
) {
	v.Name = ctx.Identifier().GetText()
	v.Src = SrcNode{
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
		v.Constant = constantCtx != nil
	}

	typeName := NewTypeName(v.ASTBuilder)
	typeName.Parse(unit, nil, v.Id, ctx.GetType_())
	v.TypeName = typeName
	v.TypeDescription = typeName.TypeDescription

	if ctx.GetInitialValue() != nil {
		expression := NewExpression(v.ASTBuilder)
		v.InitialValue = expression.Parse(unit, contractNode, nil, nil, nil, v, v.GetId(), ctx.GetInitialValue())
		v.TypeDescription = v.InitialValue.GetTypeDescription()
	}

	// This is going to be a fallback...
	// Now this is a severe hack designed to provide fallback functionality to a level, however, it is not a proper
	// fallback implementation. This is because the parser itself does not handle properly function () payable {} that are supported
	// by older versions of the solidity. So instead of doing proper solution for which I'll need time, I'm "fixing it" by doing this.
	// NOTICE: This is a temporary fix and should be removed as soon as possible.
	if strings.Contains(ctx.GetText(), "function") {
		v.NodeType = ast_pb.NodeType_FALLBACK
		v.StateVariable = false
		v.TypeName.TypeDescription = &TypeDescription{
			TypeString:     "fallback",
			TypeIdentifier: "$_t_fallback",
		}
		v.TypeDescription = typeName.TypeDescription
		return
	}

	v.currentStateVariables = append(v.currentStateVariables, v)
}

// ParseGlobal parses a state variable declaration from the provided parser.StateVariableDeclarationContext and updates the current instance.
func (v *StateVariableDeclaration) ParseGlobal(
	ctx *parser.StateVariableDeclarationContext,
) {
	v.Name = ctx.Identifier().GetText()
	v.Src = SrcNode{
		Line:        int64(ctx.GetStart().GetLine()),
		Column:      int64(ctx.GetStart().GetColumn()),
		Start:       int64(ctx.GetStart().GetStart()),
		End:         int64(ctx.GetStop().GetStop()),
		Length:      int64(ctx.GetStop().GetStop() - ctx.GetStart().GetStart() + 1),
		ParentIndex: 0,
	}
	v.Visibility = v.getVisibilityFromCtx(ctx)

	for _, immutableCtx := range ctx.AllImmutable() {
		if immutableCtx != nil {
			v.StateMutability = ast_pb.Mutability_IMMUTABLE
		}
	}

	for _, constantCtx := range ctx.AllConstant() {
		v.Constant = constantCtx != nil
	}

	typeName := NewTypeName(v.ASTBuilder)
	typeName.Parse(nil, nil, v.Id, ctx.GetType_())
	v.TypeName = typeName
	v.TypeDescription = typeName.TypeDescription

	if ctx.GetInitialValue() != nil {
		expression := NewExpression(v.ASTBuilder)
		v.InitialValue = expression.Parse(nil, nil, nil, nil, nil, v, v.GetId(), ctx.GetInitialValue())
		v.TypeDescription = v.InitialValue.GetTypeDescription()
	}

	v.globalDefinitions = append(v.globalDefinitions, v)
}

// ParseGlobalConstant parses a global constant state variable declaration from the provided parser.ConstantVariableDeclarationContext and updates the current instance.
func (v *StateVariableDeclaration) ParseGlobalConstant(
	ctx *parser.ConstantVariableDeclarationContext,
) {
	if ctx.Identifier() != nil {
		v.Name = ctx.Identifier().GetText()
	}

	v.Src = SrcNode{
		Line:        int64(ctx.GetStart().GetLine()),
		Column:      int64(ctx.GetStart().GetColumn()),
		Start:       int64(ctx.GetStart().GetStart()),
		End:         int64(ctx.GetStop().GetStop()),
		Length:      int64(ctx.GetStop().GetStop() - ctx.GetStart().GetStart() + 1),
		ParentIndex: 0,
	}
	v.Visibility = ast_pb.Visibility_PUBLIC
	v.Constant = true

	typeName := NewTypeName(v.ASTBuilder)
	typeName.Parse(nil, nil, v.Id, ctx.GetType_())
	v.TypeName = typeName
	v.TypeDescription = typeName.TypeDescription

	if ctx.GetInitialValue() != nil {
		expression := NewExpression(v.ASTBuilder)
		v.InitialValue = expression.Parse(nil, nil, nil, nil, nil, v, v.GetId(), ctx.GetInitialValue())
		v.TypeDescription = v.InitialValue.GetTypeDescription()
	}

	v.globalDefinitions = append(v.globalDefinitions, v)
}

// ParseGlobalVariable parses a global variable declaration from the provided parser.VariableDeclarationContext and updates the current instance.
func (v *StateVariableDeclaration) ParseGlobalVariable(
	ctx *parser.VariableDeclarationContext,
) {
	v.Name = ctx.Identifier().GetText()
	v.Src = SrcNode{
		Line:        int64(ctx.GetStart().GetLine()),
		Column:      int64(ctx.GetStart().GetColumn()),
		Start:       int64(ctx.GetStart().GetStart()),
		End:         int64(ctx.GetStop().GetStop()),
		Length:      int64(ctx.GetStop().GetStop() - ctx.GetStart().GetStart() + 1),
		ParentIndex: 0,
	}
	v.Visibility = ast_pb.Visibility_PUBLIC
	v.Constant = true

	typeName := NewTypeName(v.ASTBuilder)
	typeName.Parse(nil, nil, v.Id, ctx.GetType_())
	v.TypeName = typeName
	v.TypeDescription = typeName.TypeDescription

	v.globalDefinitions = append(v.globalDefinitions, v)
}

// getVisibilityFromCtx extracts visibility information from the parser context.
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

func (b *ASTBuilder) EnterStateVariableDeclaration(ctx *parser.StateVariableDeclarationContext) {
	stateVar := NewStateVariableDeclaration(b)
	stateVar.ParseGlobal(ctx)
}

func (b *ASTBuilder) EnterConstantVariableDeclaration(ctx *parser.ConstantVariableDeclarationContext) {
	stateVar := NewStateVariableDeclaration(b)
	stateVar.ParseGlobalConstant(ctx)
}

func (b *ASTBuilder) EnterVariableDeclaration(ctx *parser.VariableDeclarationContext) {
	child := NewStateVariableDeclaration(b)
	child.ParseGlobalVariable(ctx)
}
