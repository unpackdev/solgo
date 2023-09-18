package ast

import (
	ast_pb "github.com/unpackdev/protos/dist/go/ast" // Import for AST protocol buffer definitions.
	"github.com/unpackdev/solgo/parser"              // Import for the solgo parser.
)

// Receive represents a receive function definition node in the abstract syntax tree (AST).
// It encapsulates information about the characteristics and properties of a receive function within a contract.
type Receive struct {
	*ASTBuilder                            // Embedded ASTBuilder for building the AST.
	Id               int64                 `json:"id"`                // Unique identifier for the Receive node.
	NodeType         ast_pb.NodeType       `json:"node_type"`         // Type of the AST node.
	Kind             ast_pb.NodeType       `json:"kind"`              // Kind of the receive function.
	Src              SrcNode               `json:"src"`               // Source location information.
	Implemented      bool                  `json:"implemented"`       // Indicates whether the function is implemented.
	Visibility       ast_pb.Visibility     `json:"visibility"`        // Visibility of the receive function.
	StateMutability  ast_pb.Mutability     `json:"state_mutability"`  // State mutability of the receive function.
	Modifiers        []*ModifierInvocation `json:"modifiers"`         // List of modifier invocations applied to the receive function.
	Overrides        []*OverrideSpecifier  `json:"overrides"`         // List of override specifiers for the receive function.
	Parameters       *ParameterList        `json:"parameters"`        // List of parameters for the receive function.
	ReturnParameters *ParameterList        `json:"return_parameters"` // List of return parameters for the receive function.
	Body             *BodyNode             `json:"body"`              // Body of the receive function.
	Virtual          bool                  `json:"virtual"`           // Indicates whether the function is virtual.
	Payable          bool                  `json:"payable"`           // Indicates whether the function is payable.
}

// NewReceiveDefinition creates a new Receive node with default values and returns it.
func NewReceiveDefinition(b *ASTBuilder) *Receive {
	return &Receive{
		ASTBuilder:      b,
		Id:              b.GetNextID(),
		NodeType:        ast_pb.NodeType_FUNCTION_DEFINITION,
		Kind:            ast_pb.NodeType_RECEIVE,
		StateMutability: ast_pb.Mutability_NONPAYABLE,
		Modifiers:       make([]*ModifierInvocation, 0),
		Overrides:       make([]*OverrideSpecifier, 0),
	}
}

// SetReferenceDescriptor sets the reference descriptions of the Receive node.
// This function is not yet implemented and returns false.
func (f *Receive) SetReferenceDescriptor(refId int64, refDesc *TypeDescription) bool {
	return false
}

// GetId returns the unique identifier of the Receive node.
func (f *Receive) GetId() int64 {
	return f.Id
}

// GetSrc returns the source location information of the Receive node.
func (f *Receive) GetSrc() SrcNode {
	return f.Src
}

// GetType returns the type of the AST node, which is NodeType_FUNCTION_DEFINITION for a receive function.
func (f *Receive) GetType() ast_pb.NodeType {
	return f.NodeType
}

// GetNodes returns a slice of child nodes within the body of the receive function.
func (f *Receive) GetNodes() []Node[NodeType] {
	return f.Body.Statements
}

// GetTypeDescription returns the type description associated with the Receive node (currently returns nil).
func (f *Receive) GetTypeDescription() *TypeDescription {
	return &TypeDescription{
		TypeString:     "receive",
		TypeIdentifier: "$_t_receive",
	}
}

// GetModifiers returns a list of modifier invocations applied to the Receive node.
func (f *Receive) GetModifiers() []*ModifierInvocation {
	return f.Modifiers
}

// GetOverrides returns a list of override specifiers for the Receive node.
func (f *Receive) GetOverrides() []*OverrideSpecifier {
	return f.Overrides
}

// GetParameters returns the list of parameters for the Receive node.
func (f *Receive) GetParameters() *ParameterList {
	return f.Parameters
}

// GetReturnParameters returns the list of return parameters for the Receive node.
func (f *Receive) GetReturnParameters() *ParameterList {
	return f.ReturnParameters
}

// GetBody returns the body of the Receive node.
func (f *Receive) GetBody() *BodyNode {
	return f.Body
}

// GetKind returns the kind of the Receive node, which is NodeType_RECEIVE.
func (f *Receive) GetKind() ast_pb.NodeType {
	return f.Kind
}

// IsImplemented returns true if the Receive function is implemented, false otherwise.
func (f *Receive) IsImplemented() bool {
	return f.Implemented
}

// IsVirtual returns true if the Receive function is virtual, false otherwise.
func (f *Receive) IsVirtual() bool {
	return f.Virtual
}

// GetVisibility returns the visibility of the Receive function.
func (f *Receive) GetVisibility() ast_pb.Visibility {
	return f.Visibility
}

// GetStateMutability returns the state mutability of the Receive function.
func (f *Receive) GetStateMutability() ast_pb.Mutability {
	return f.StateMutability
}

// ToProto converts the Receive node to its corresponding protocol buffer representation.
func (f *Receive) ToProto() NodeType {
	proto := ast_pb.Receive{
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

	return NewTypedStruct(&proto, "Receive")
}

// Parse populates the properties of the Receive node by parsing the corresponding context and information.
// It returns the populated Receive node.
func (f *Receive) Parse(
	unit *SourceUnit[Node[ast_pb.SourceUnit]],
	contractNode Node[NodeType],
	bodyCtx parser.IContractBodyElementContext,
	ctx *parser.ReceiveFunctionDefinitionContext,
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
	params.Src = f.Src
	params.Src.ParentIndex = f.Id
	f.Parameters = params

	returnParams := NewParameterList(f.ASTBuilder)
	returnParams.Src = f.Src
	returnParams.Src.ParentIndex = f.Id
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

// getVisibilityFromCtx determines the visibility of the Receive function based on the context.
func (f *Receive) getVisibilityFromCtx(ctx *parser.ReceiveFunctionDefinitionContext) ast_pb.Visibility {
	for _, visibility := range ctx.AllExternal() {
		if visibility.GetText() == "external" {
			f.Visibility = ast_pb.Visibility_EXTERNAL
		}
	}

	return ast_pb.Visibility_INTERNAL
}

// getStateMutabilityFromCtx determines the state mutability of the Receive function based on the context.
func (f *Receive) getStateMutabilityFromCtx(ctx *parser.ReceiveFunctionDefinitionContext) ast_pb.Mutability {
	for _, stateMutability := range ctx.AllPayable() {
		if stateMutability.GetText() == "payable" {
			f.StateMutability = ast_pb.Mutability_PAYABLE
		}
	}

	return ast_pb.Mutability_NONPAYABLE
}
