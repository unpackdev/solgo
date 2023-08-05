package ast

import (
	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/parser"
)

type Fallback struct {
	*ASTBuilder

	Id               int64                 `json:"id"`
	NodeType         ast_pb.NodeType       `json:"node_type"`
	Kind             ast_pb.NodeType       `json:"kind"`
	Src              SrcNode               `json:"src"`
	Implemented      bool                  `json:"implemented"`
	Visibility       ast_pb.Visibility     `json:"visibility"`
	StateMutability  ast_pb.Mutability     `json:"state_mutability"`
	Modifiers        []*ModifierInvocation `json:"modifiers"`
	Overrides        []*OverrideSpecifier  `json:"overrides"`
	Parameters       *ParameterList        `json:"parameters"`
	ReturnParameters *ParameterList        `json:"return_parameters"`
	Body             *BodyNode             `json:"body"`
	Virtual          bool                  `json:"virtual"`
}

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
// We don't need to do any reference description updates here, at least for now...
func (f *Fallback) SetReferenceDescriptor(refId int64, refDesc *TypeDescription) bool {
	return false
}

func (f *Fallback) GetId() int64 {
	return f.Id
}

func (f *Fallback) GetSrc() SrcNode {
	return f.Src
}

func (f *Fallback) GetType() ast_pb.NodeType {
	return f.NodeType
}

func (f *Fallback) GetNodes() []Node[NodeType] {
	return f.Body.Statements
}

func (f *Fallback) GetTypeDescription() *TypeDescription {
	return nil
}

func (f *Fallback) GetModifiers() []*ModifierInvocation {
	return f.Modifiers
}

func (f *Fallback) GetOverrides() []*OverrideSpecifier {
	return f.Overrides
}

func (f *Fallback) GetParameters() *ParameterList {
	return f.Parameters
}

func (f *Fallback) GetReturnParameters() *ParameterList {
	return f.ReturnParameters
}

func (f *Fallback) GetBody() *BodyNode {
	return f.Body
}

func (f *Fallback) GetKind() ast_pb.NodeType {
	return f.Kind
}

func (f *Fallback) GetVisibility() ast_pb.Visibility {
	return f.Visibility
}

func (f *Fallback) GetStateMutability() ast_pb.Mutability {
	return f.StateMutability
}

func (f *Fallback) IsVirtual() bool {
	return f.Virtual
}

func (f *Fallback) IsImplemented() bool {
	return f.Implemented
}

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

	return NewTypedStruct(&proto, "Fallback")
}

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

	// Set function parameters if they exist.

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

	// Set function modifiers.
	for _, modifierCtx := range ctx.AllModifierInvocation() {
		modifier := NewModifierInvocation(f.ASTBuilder)
		modifier.Parse(unit, contractNode, f, nil, modifierCtx)
		f.Modifiers = append(f.Modifiers, modifier)
	}

	// Set function override specifier.
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

func (f *Fallback) getVisibilityFromCtx(ctx *parser.FallbackFunctionDefinitionContext) ast_pb.Visibility {
	for _, visibility := range ctx.AllExternal() {
		if visibility.GetText() == "external" {
			f.Visibility = ast_pb.Visibility_EXTERNAL
		}
	}

	return ast_pb.Visibility_INTERNAL
}

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
