// Package ast defines data structures and methods for abstract syntax tree nodes used in a specific programming language.
// The package contains definitions for various AST nodes that represent different elements of the programming language's syntax.
package ast

import (
	ast_pb "github.com/txpull/protos/dist/go/ast" // Import for AST protocol buffer definitions.
	"github.com/txpull/solgo/parser"              // Import for the solgo parser.
)

// Fallback represents a fallback function definition node in the abstract syntax tree (AST).
// It encapsulates information about the characteristics and properties of a fallback function within a contract.
type Fallback struct {
	*ASTBuilder                            // Embedded ASTBuilder for building the AST.
	Id               int64                 `json:"id"`                // Unique identifier for the Fallback node.
	NodeType         ast_pb.NodeType       `json:"node_type"`         // Type of the AST node.
	Kind             ast_pb.NodeType       `json:"kind"`              // Kind of the fallback function.
	Src              SrcNode               `json:"src"`               // Source location information.
	Implemented      bool                  `json:"implemented"`       // Indicates whether the function is implemented.
	Visibility       ast_pb.Visibility     `json:"visibility"`        // Visibility of the fallback function.
	StateMutability  ast_pb.Mutability     `json:"state_mutability"`  // State mutability of the fallback function.
	Modifiers        []*ModifierInvocation `json:"modifiers"`         // List of modifier invocations applied to the fallback function.
	Overrides        []*OverrideSpecifier  `json:"overrides"`         // List of override specifiers for the fallback function.
	Parameters       *ParameterList        `json:"parameters"`        // List of parameters for the fallback function.
	ReturnParameters *ParameterList        `json:"return_parameters"` // List of return parameters for the fallback function.
	Body             *BodyNode             `json:"body"`              // Body of the fallback function.
	Virtual          bool                  `json:"virtual"`           // Indicates whether the function is virtual.
}

// NewFallbackDefinition creates a new Fallback node with default values and returns it.
func NewFallbackDefinition(b *ASTBuilder) *Fallback {
	return &Fallback{
		ASTBuilder:      b,
		Id:              b.GetNextID(),
		NodeType:        ast_pb.NodeType_FUNCTION_DEFINITION,
		Kind:            ast_pb.NodeType_FALLBACK,
		StateMutability: ast_pb.Mutability_NONPAYABLE,
		Modifiers:       make([]*ModifierInvocation, 0),
		Overrides:       make([]*OverrideSpecifier, 0),
	}
}

// SetReferenceDescriptor sets the reference descriptions of the Fallback node.
// This function currently returns false, as no reference description updates are performed.
func (f *Fallback) SetReferenceDescriptor(refId int64, refDesc *TypeDescription) bool {
	return false
}

// GetId returns the unique identifier of the Fallback node.
func (f *Fallback) GetId() int64 {
	return f.Id
}

// GetSrc returns the source location information of the Fallback node.
func (f *Fallback) GetSrc() SrcNode {
	return f.Src
}

// GetType returns the type of the AST node, which is NodeType_FUNCTION_DEFINITION for a fallback function.
func (f *Fallback) GetType() ast_pb.NodeType {
	return f.NodeType
}

// GetNodes returns a slice of child nodes within the body of the fallback function.
func (f *Fallback) GetNodes() []Node[NodeType] {
	return f.Body.Statements
}

// GetTypeDescription returns the type description associated with the Fallback node.
func (f *Fallback) GetTypeDescription() *TypeDescription {
	return &TypeDescription{
		TypeString:     "fallback",
		TypeIdentifier: "$_t_fallback",
	}
}

// GetModifiers returns a list of modifier invocations applied to the Fallback node.
func (f *Fallback) GetModifiers() []*ModifierInvocation {
	return f.Modifiers
}

// GetOverrides returns a list of override specifiers for the Fallback node.
func (f *Fallback) GetOverrides() []*OverrideSpecifier {
	return f.Overrides
}

// GetParameters returns the list of parameters for the Fallback node.
func (f *Fallback) GetParameters() *ParameterList {
	return f.Parameters
}

// GetReturnParameters returns the list of return parameters for the Fallback node.
func (f *Fallback) GetReturnParameters() *ParameterList {
	return f.ReturnParameters
}

// GetBody returns the body of the Fallback node.
func (f *Fallback) GetBody() *BodyNode {
	return f.Body
}

// GetKind returns the kind of the Fallback node, which is NodeType_FALLBACK.
func (f *Fallback) GetKind() ast_pb.NodeType {
	return f.Kind
}

// GetVisibility returns the visibility of the Fallback function.
func (f *Fallback) GetVisibility() ast_pb.Visibility {
	return f.Visibility
}

// GetStateMutability returns the state mutability of the Fallback function.
func (f *Fallback) GetStateMutability() ast_pb.Mutability {
	return f.StateMutability
}

// IsVirtual returns true if the Fallback function is virtual, false otherwise.
func (f *Fallback) IsVirtual() bool {
	return f.Virtual
}

// IsImplemented returns true if the Fallback function is implemented, false otherwise.
func (f *Fallback) IsImplemented() bool {
	return f.Implemented
}

// ToProto converts the Fallback node to its corresponding protocol buffer representation.
func (f *Fallback) ToProto() NodeType {
	proto := ast_pb.Fallback{
		Id:               f.GetId(),
		NodeType:         f.GetType(),
		Kind:             f.GetKind(),
		Src:              f.GetSrc().ToProto(),
		Virtual:          f.IsVirtual(),
		Implemented:      f.IsImplemented(),
		Visibility:       f.GetVisibility(),
		StateMutability:  f.GetStateMutability(),
		Parameters:       f.GetParameters().ToProto(),
		ReturnParameters: f.GetReturnParameters().ToProto(),
		Body:             f.GetBody().ToProto().(*ast_pb.Body),
	}

	for _, modifier := range f.GetModifiers() {
		proto.Modifiers = append(proto.Modifiers, modifier.ToProto().(*ast_pb.ModifierInvocation))
	}

	for _, override := range f.GetOverrides() {
		proto.Overrides = append(proto.Overrides, override.ToProto().(*ast_pb.OverrideSpecifier))
	}

	return NewTypedStruct(&proto, "Fallback")
}

// Parse populates the properties of the Fallback node by parsing the corresponding context and information.
// It returns the populated Fallback node.
func (f *Fallback) Parse(
	unit *SourceUnit[Node[ast_pb.SourceUnit]],
	contractNode Node[NodeType],
	bodyCtx parser.IContractBodyElementContext,
	ctx *parser.FallbackFunctionDefinitionContext,
) Node[NodeType] {
	f.Src = SrcNode{
		Id:          f.GetNextID(),
		Line:        int64(ctx.GetStart().GetLine()),
		Column:      int64(ctx.GetStart().GetColumn()),
		Start:       int64(ctx.GetStart().GetStart()),
		End:         int64(ctx.GetStop().GetStop()),
		Length:      int64(ctx.GetStop().GetStop() - ctx.GetStart().GetStart() + 1),
		ParentIndex: contractNode.GetId(),
	}
	f.Implemented = ctx.Block() != nil && !ctx.Block().IsEmpty()

	for _, virtual := range ctx.AllVirtual() {
		if virtual.GetText() == "virtual" {
			f.Virtual = true
		}
	}

	f.Visibility = f.getVisibilityFromCtx(ctx)
	f.StateMutability = f.getStateMutabilityFromCtx(ctx)

	params := NewParameterList(f.ASTBuilder)
	if len(ctx.AllParameterList()) > 0 {
		params.Parse(unit, f, ctx.AllParameterList()[0])
	} else {
		params.Src = f.Src
		params.Src.ParentIndex = f.Id
	}
	f.Parameters = params

	returnParams := NewParameterList(f.ASTBuilder)
	if ctx.GetReturnParameters() != nil {
		returnParams.Parse(unit, f, ctx.GetReturnParameters())
	} else {
		returnParams.Src = f.Src
		returnParams.Src.ParentIndex = f.Id
	}
	f.ReturnParameters = returnParams

	for _, modifierCtx := range ctx.AllModifierInvocation() {
		modifier := NewModifierInvocation(f.ASTBuilder)
		modifier.Parse(unit, contractNode, f, nil, modifierCtx)
		f.Modifiers = append(f.Modifiers, modifier)
	}

	for _, overrideCtx := range ctx.AllOverrideSpecifier() {
		overrideSpecifier := NewOverrideSpecifier(f.ASTBuilder)
		overrideSpecifier.Parse(unit, f, overrideCtx)
		f.Overrides = append(f.Overrides, overrideSpecifier)
	}

	if ctx.Block() != nil && !ctx.Block().IsEmpty() {
		bodyNode := NewBodyNode(f.ASTBuilder)
		bodyNode.ParseBlock(unit, contractNode, f, ctx.Block())
		f.Body = bodyNode

		if ctx.Block().AllUncheckedBlock() != nil {
			for _, uncheckedCtx := range ctx.Block().AllUncheckedBlock() {
				bodyNode := NewBodyNode(f.ASTBuilder)
				bodyNode.ParseUncheckedBlock(unit, contractNode, f, uncheckedCtx)
				f.Body.Statements = append(f.Body.Statements, bodyNode)
			}
		}
	}

	return f
}

// getVisibilityFromCtx determines the visibility of the Fallback function based on the context.
func (f *Fallback) getVisibilityFromCtx(ctx *parser.FallbackFunctionDefinitionContext) ast_pb.Visibility {
	for _, visibility := range ctx.AllExternal() {
		if visibility.GetText() == "external" {
			f.Visibility = ast_pb.Visibility_EXTERNAL
		}
	}

	return ast_pb.Visibility_INTERNAL
}

// getStateMutabilityFromCtx determines the state mutability of the Fallback function based on the context.
func (f *Fallback) getStateMutabilityFromCtx(ctx *parser.FallbackFunctionDefinitionContext) ast_pb.Mutability {
	mutabilityMap := map[string]ast_pb.Mutability{
		"payable":    ast_pb.Mutability_PAYABLE,
		"pure":       ast_pb.Mutability_PURE,
		"view":       ast_pb.Mutability_VIEW,
		"immutable":  ast_pb.Mutability_IMMUTABLE,
		"mutable":    ast_pb.Mutability_MUTABLE,
		"nonpayable": ast_pb.Mutability_NONPAYABLE,
	}

	for _, stateMutability := range ctx.AllStateMutability() {
		if m, ok := mutabilityMap[stateMutability.GetText()]; ok {
			return m
		}
	}

	return ast_pb.Mutability_NONPAYABLE
}
