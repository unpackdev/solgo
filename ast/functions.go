package ast

import (
	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/parser"
)

type Function struct {
	*ASTBuilder

	Id               int64                `json:"id"`
	Name             string               `json:"name"`
	NodeType         ast_pb.NodeType      `json:"node_type"`
	Kind             ast_pb.NodeType      `json:"kind"`
	Src              SrcNode              `json:"src"`
	Body             *BodyNode            `json:"body"`
	Implemented      bool                 `json:"implemented"`
	Visibility       ast_pb.Visibility    `json:"visibility"`
	StateMutability  ast_pb.Mutability    `json:"state_mutability"`
	Virtual          bool                 `json:"virtual"`
	Modifiers        []ModifierDefinition `json:"modifiers"`
	Overrides        []OverrideSpecifier  `json:"overrides"`
	Parameters       *ParameterList       `json:"parameters"`
	ReturnParameters *ParameterList       `json:"return_parameters"`
	Scope            int64                `json:"scope"`
}

func NewFunction(b *ASTBuilder) *Function {
	return &Function{
		ASTBuilder:  b,
		NodeType:    ast_pb.NodeType_FUNCTION_DEFINITION,
		Kind:        ast_pb.NodeType_KIND_FUNCTION,
		Modifiers:   make([]ModifierDefinition, 0),
		Overrides:   make([]OverrideSpecifier, 0),
		Implemented: true,
	}
}

func (f Function) GetId() int64 {
	return f.Id
}

func (f Function) GetType() ast_pb.NodeType {
	return f.NodeType
}

func (f Function) GetSrc() SrcNode {
	return f.Src
}

func (f Function) GetParameters() *ParameterList {
	return f.Parameters
}

func (f Function) GetReturnParameters() *ParameterList {
	return f.ReturnParameters
}

func (f Function) GetBody() *BodyNode {
	return f.Body
}

func (f Function) GetKind() ast_pb.NodeType {
	return f.Kind
}

func (f Function) IsImplemented() bool {
	return f.Implemented
}

func (f Function) GetModifiers() []ModifierDefinition {
	return f.Modifiers
}

func (f Function) GetOverrides() []OverrideSpecifier {
	return f.Overrides
}

func (f Function) GetVisibility() ast_pb.Visibility {
	return f.Visibility
}

func (f Function) GetStateMutability() ast_pb.Mutability {
	return f.StateMutability
}

func (f Function) IsVirtual() bool {
	return f.Virtual
}

func (f Function) GetScope() int64 {
	return f.Scope
}

func (f Function) GetName() string {
	return f.Name
}

func (f Function) GetTypeDescription() *TypeDescription {
	return nil
}

func (f Function) GetNodes() []Node[NodeType] {
	return f.Body.GetNodes()
}

func (f Function) ToProto() NodeType {
	return ast_pb.Function{}
}

func (f Function) Parse(
	unit *SourceUnit[Node[ast_pb.SourceUnit]],
	contractNode Node[NodeType],
	bodyCtx parser.IContractBodyElementContext,
	ctx *parser.FunctionDefinitionContext,
) Node[NodeType] {
	f.Id = f.GetNextID()
	f.Scope = contractNode.GetId()
	if ctx.Identifier() != nil {
		f.Name = ctx.Identifier().GetText()
	}
	f.Implemented = ctx.Block() != nil && !ctx.Block().IsEmpty()
	f.Src = SrcNode{
		Id:          f.GetNextID(),
		Line:        int64(ctx.GetStart().GetLine()),
		Column:      int64(ctx.GetStart().GetColumn()),
		Start:       int64(ctx.GetStart().GetStart()),
		End:         int64(ctx.GetStop().GetStop()),
		Length:      int64(ctx.GetStop().GetStop() - ctx.GetStart().GetStart() + 1),
		ParentIndex: contractNode.GetId(),
	}

	// Set function visibility state.
	f.Visibility = f.getVisibilityFromCtx(ctx)

	// Set function state mutability.
	f.StateMutability = f.getStateMutabilityFromCtx(ctx)

	// Set if function is virtual.
	f.Virtual = f.getVirtualState(ctx)

	// Set function modifiers.
	for _, modifierCtx := range ctx.AllModifierInvocation() {
		modifier := NewModifierDefinition(f.ASTBuilder)
		modifier.Parse(unit, f, modifierCtx)
		f.Modifiers = append(f.Modifiers, *modifier)
	}

	// Set function override specifier.
	for _, overrideCtx := range ctx.AllOverrideSpecifier() {
		overrideSpecifier := NewOverrideSpecifier(f.ASTBuilder)
		overrideSpecifier.Parse(unit, f, overrideCtx)
		f.Overrides = append(f.Overrides, *overrideSpecifier)
	}

	// Set function parameters if they exist.
	if len(ctx.AllParameterList()) > 0 {
		params := NewParameterList(f.ASTBuilder)
		params.Parse(unit, f, ctx.AllParameterList()[0])
		f.Parameters = params
	}

	// Set function return parameters if they exist.
	// @TODO: Consider traversing through body to discover name of the return parameters even
	// if they are not defined in (name uint) format.
	if ctx.GetReturnParameters() != nil {
		returnParams := NewParameterList(f.ASTBuilder)
		returnParams.Parse(unit, f, ctx.GetReturnParameters())
		f.ReturnParameters = returnParams
	}

	if ctx.Block() != nil && !ctx.Block().IsEmpty() {
		bodyNode := NewBodyNode(f.ASTBuilder)
		bodyNode.ParseBlock(unit, contractNode, f, ctx.Block())
		f.Body = bodyNode

		// In case at any point we discover that the function is not implemented, we set the
		// implemented flag to false.
		if !bodyNode.Implemented {
			f.Implemented = false
		}

		if ctx.Block().AllUncheckedBlock() != nil {
			for _, uncheckedCtx := range ctx.Block().AllUncheckedBlock() {
				bodyNode := NewBodyNode(f.ASTBuilder)
				bodyNode.ParseUncheckedBlock(unit, contractNode, f, uncheckedCtx)
				f.Body.Statements = append(f.Body.Statements, bodyNode)

				// In case at any point we discover that the function is not implemented, we set the
				// implemented flag to false.
				if !bodyNode.Implemented {
					f.Implemented = false
				}
			}
		}
	}

	return f
}

func (f Function) getVisibilityFromCtx(ctx *parser.FunctionDefinitionContext) ast_pb.Visibility {
	visibilityMap := map[string]ast_pb.Visibility{
		"public":   ast_pb.Visibility_PUBLIC,
		"private":  ast_pb.Visibility_PRIVATE,
		"internal": ast_pb.Visibility_INTERNAL,
		"external": ast_pb.Visibility_EXTERNAL,
	}

	for _, visibility := range ctx.AllVisibility() {
		if v, ok := visibilityMap[visibility.GetText()]; ok {
			return v
		}
	}

	return ast_pb.Visibility_INTERNAL
}

func (f Function) getStateMutabilityFromCtx(ctx *parser.FunctionDefinitionContext) ast_pb.Mutability {
	mutabilityMap := map[string]ast_pb.Mutability{
		"payable": ast_pb.Mutability_PAYABLE,
		"pure":    ast_pb.Mutability_PURE,
		"view":    ast_pb.Mutability_VIEW,
	}

	for _, stateMutability := range ctx.AllStateMutability() {
		if m, ok := mutabilityMap[stateMutability.GetText()]; ok {
			return m
		}
	}

	return ast_pb.Mutability_NONPAYABLE
}

func (f *Function) getVirtualState(ctx *parser.FunctionDefinitionContext) bool {
	for _, virtual := range ctx.AllVirtual() {
		if virtual.GetText() == "virtual" {
			return true
		}
	}

	return false
}
