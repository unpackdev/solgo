package ast

import (
	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/parser"
)

type Receive struct {
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
	Payable          bool                  `json:"payable"`
}

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
func (f *Receive) SetReferenceDescriptor(refId int64, refDesc *TypeDescription) bool {
	return false
}

func (f *Receive) GetId() int64 {
	return f.Id
}

func (f *Receive) GetSrc() SrcNode {
	return f.Src
}

func (f *Receive) GetType() ast_pb.NodeType {
	return f.NodeType
}

func (f *Receive) GetNodes() []Node[NodeType] {
	return f.Body.Statements
}

func (f *Receive) GetTypeDescription() *TypeDescription {
	return nil
}

func (f *Receive) GetModifiers() []*ModifierInvocation {
	return f.Modifiers
}

func (f *Receive) GetOverrides() []*OverrideSpecifier {
	return f.Overrides
}

func (f *Receive) GetParameters() *ParameterList {
	return f.Parameters
}

func (f *Receive) GetReturnParameters() *ParameterList {
	return f.ReturnParameters
}

func (f *Receive) GetBody() *BodyNode {
	return f.Body
}

func (f *Receive) GetKind() ast_pb.NodeType {
	return f.Kind
}

func (f *Receive) IsImplemented() bool {
	return f.Implemented
}

func (f *Receive) IsVirtual() bool {
	return f.Virtual
}

func (f *Receive) GetVisibility() ast_pb.Visibility {
	return f.Visibility
}

func (f *Receive) GetStateMutability() ast_pb.Mutability {
	return f.StateMutability
}

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

func (f *Receive) getVisibilityFromCtx(ctx *parser.ReceiveFunctionDefinitionContext) ast_pb.Visibility {
	for _, visibility := range ctx.AllExternal() {
		if visibility.GetText() == "external" {
			f.Visibility = ast_pb.Visibility_EXTERNAL
		}
	}

	return ast_pb.Visibility_INTERNAL
}

func (f *Receive) getStateMutabilityFromCtx(ctx *parser.ReceiveFunctionDefinitionContext) ast_pb.Mutability {
	for _, stateMutability := range ctx.AllPayable() {
		if stateMutability.GetText() == "payable" {
			f.StateMutability = ast_pb.Mutability_PAYABLE
		}
	}

	return ast_pb.Mutability_NONPAYABLE
}
