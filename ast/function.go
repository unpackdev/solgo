package ast

import (
	"strings"

	v3 "github.com/cncf/xds/go/xds/type/v3"
	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/parser"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/structpb"
)

type Function struct {
	*ASTBuilder

	Id                    int64                 `json:"id"`
	Name                  string                `json:"name"`
	NodeType              ast_pb.NodeType       `json:"node_type"`
	Kind                  ast_pb.NodeType       `json:"kind"`
	Src                   SrcNode               `json:"src"`
	Body                  *BodyNode             `json:"body"`
	Implemented           bool                  `json:"implemented"`
	Visibility            ast_pb.Visibility     `json:"visibility"`
	StateMutability       ast_pb.Mutability     `json:"state_mutability"`
	Virtual               bool                  `json:"virtual"`
	Modifiers             []*ModifierInvocation `json:"modifiers"`
	Overrides             []*OverrideSpecifier  `json:"overrides"`
	Parameters            *ParameterList        `json:"parameters"`
	ReturnParameters      *ParameterList        `json:"return_parameters"`
	Scope                 int64                 `json:"scope"`
	ReferencedDeclaration int64                 `json:"referenced_declaration,omitempty"`
	TypeDescription       *TypeDescription      `json:"type_description"`
}

func NewFunction(b *ASTBuilder) *Function {
	return &Function{
		ASTBuilder:  b,
		NodeType:    ast_pb.NodeType_FUNCTION_DEFINITION,
		Kind:        ast_pb.NodeType_KIND_FUNCTION,
		Modifiers:   make([]*ModifierInvocation, 0),
		Overrides:   make([]*OverrideSpecifier, 0),
		Implemented: true,
	}
}

// SetReferenceDescriptor sets the reference descriptions of the Function node.
func (f *Function) SetReferenceDescriptor(refId int64, refDesc *TypeDescription) bool {
	f.ReferencedDeclaration = refId
	f.TypeDescription = refDesc
	return false
}

func (f *Function) GetId() int64 {
	return f.Id
}

func (f *Function) GetType() ast_pb.NodeType {
	return f.NodeType
}

func (f *Function) GetSrc() SrcNode {
	return f.Src
}

func (f *Function) GetParameters() *ParameterList {
	return f.Parameters
}

func (f *Function) GetReturnParameters() *ParameterList {
	return f.ReturnParameters
}

func (f *Function) GetBody() *BodyNode {
	return f.Body
}

func (f *Function) GetKind() ast_pb.NodeType {
	return f.Kind
}

func (f *Function) IsImplemented() bool {
	return f.Implemented
}

func (f *Function) GetModifiers() []*ModifierInvocation {
	return f.Modifiers
}

func (f *Function) GetOverrides() []*OverrideSpecifier {
	return f.Overrides
}

func (f *Function) GetVisibility() ast_pb.Visibility {
	return f.Visibility
}

func (f *Function) GetStateMutability() ast_pb.Mutability {
	return f.StateMutability
}

func (f *Function) IsVirtual() bool {
	return f.Virtual
}

func (f *Function) GetScope() int64 {
	return f.Scope
}

func (f *Function) GetName() string {
	return f.Name
}

func (f *Function) GetTypeDescription() *TypeDescription {
	return f.TypeDescription
}

func (f *Function) GetNodes() []Node[NodeType] {
	return f.Body.GetNodes()
}

func (f *Function) GetReferencedDeclaration() int64 {
	return f.ReferencedDeclaration
}

func (f *Function) ToProto() NodeType {
	proto := ast_pb.Function{
		Id:                    f.GetId(),
		Name:                  f.GetName(),
		NodeType:              f.GetType(),
		Kind:                  f.GetKind(),
		Src:                   f.GetSrc().ToProto(),
		ReferencedDeclaration: f.GetReferencedDeclaration(),
		Implemented:           f.IsImplemented(),
		Virtual:               f.IsVirtual(),
		Scope:                 f.GetScope(),
		Visibility:            f.GetVisibility(),
		StateMutability:       f.GetStateMutability(),
		Modifiers:             make([]*ast_pb.ModifierInvocation, 0),
		Overrides:             make([]*ast_pb.OverrideSpecifier, 0),
		Parameters:            f.GetParameters().ToProto(),
		ReturnParameters:      f.GetReturnParameters().ToProto(),
		Body:                  f.GetBody().ToProto().(*ast_pb.Body),
	}

	if f.GetTypeDescription() != nil {
		proto.TypeDescription = f.GetTypeDescription().ToProto()
	}

	for _, modifier := range f.GetModifiers() {
		proto.Modifiers = append(proto.Modifiers, modifier.ToProto().(*ast_pb.ModifierInvocation))
	}

	for _, override := range f.GetOverrides() {
		proto.Overrides = append(proto.Overrides, override.ToProto())
	}

	// Marshal the Pragma into JSON
	jsonBytes, err := protojson.Marshal(&proto)
	if err != nil {
		panic(err)
	}

	s := &structpb.Struct{}
	if err := protojson.Unmarshal(jsonBytes, s); err != nil {
		panic(err)
	}

	return &v3.TypedStruct{
		TypeUrl: "github.com/txpull/protos/txpull.v1.ast.Function",
		Value:   s,
	}
}

/**

Modifiers             []*ModifierInvocation `json:"modifiers"`
Overrides             []*OverrideSpecifier  `json:"overrides"`
Parameters            *ParameterList        `json:"parameters"`
ReturnParameters      *ParameterList        `json:"return_parameters"`

TypeDescription       *TypeDescription      `json:"type_description"`
**/

func (f *Function) Parse(
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

	// Set function parameters if they exist.
	params := NewParameterList(f.ASTBuilder)
	if len(ctx.AllParameterList()) > 0 {
		params.Parse(unit, f, ctx.AllParameterList()[0])
	} else {
		params.Src = f.Src
		params.Src.ParentIndex = f.Id
	}
	f.Parameters = params

	// Set function return parameters if they exist.
	// @TODO: Consider traversing through body to discover name of the return parameters even
	// if they are not defined in (name uint) format.
	returnParams := NewParameterList(f.ASTBuilder)
	if ctx.GetReturnParameters() != nil {
		returnParams.Parse(unit, f, ctx.GetReturnParameters())
	} else {
		returnParams.Src = f.Src
		returnParams.Src.ParentIndex = f.Id
	}
	f.ReturnParameters = returnParams

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
	} else {
		bodyNode := NewBodyNode(f.ASTBuilder)
		bodyNode.Src = f.Src
		bodyNode.Src.ParentIndex = f.Id
		f.Body = bodyNode
	}

	f.TypeDescription = f.buildTypeDescription()
	f.currentFunctions = append(f.currentFunctions, f)
	return f
}

func (f *Function) buildTypeDescription() *TypeDescription {
	typeString := "function("
	typeIdentifier := "t_function_"
	typeStrings := make([]string, 0)
	typeIdentifiers := make([]string, 0)

	for _, paramType := range f.GetParameters().GetParameterTypes() {
		typeStrings = append(typeStrings, paramType.TypeString)
		typeIdentifiers = append(typeIdentifiers, "$_"+paramType.TypeIdentifier)
	}
	typeString += strings.Join(typeStrings, ",") + ")"
	typeIdentifier += strings.Join(typeIdentifiers, "$")

	return &TypeDescription{
		TypeString:     typeString,
		TypeIdentifier: typeIdentifier,
	}
}

func (f *Function) getVisibilityFromCtx(ctx *parser.FunctionDefinitionContext) ast_pb.Visibility {
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

func (f *Function) getStateMutabilityFromCtx(ctx *parser.FunctionDefinitionContext) ast_pb.Mutability {
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
