package ast

import (
	"github.com/goccy/go-json"

	ast_pb "github.com/unpackdev/protos/dist/go/ast"
	"github.com/unpackdev/solgo/parser"
)

// The Constructor struct represents a constructor function in a Solidity contract.
type Constructor struct {
	// Embedding the ASTBuilder to provide common functionality
	*ASTBuilder

	// The unique identifier for the constructor
	Id int64 `json:"id"`

	// The type of the node, which is 'FUNCTION_DEFINITION' for a constructor
	NodeType ast_pb.NodeType `json:"node_type"`

	// The source information about the constructor, such as its line and column numbers in the source file
	Src SrcNode `json:"src"`

	// The kind of the node, which is 'CONSTRUCTOR' for a constructor
	Kind ast_pb.NodeType `json:"kind"`

	// The state mutability of the constructor, which is 'NONPAYABLE' by default
	StateMutability ast_pb.Mutability `json:"state_mutability"`

	// The visibility of the constructor
	Visibility ast_pb.Visibility `json:"visibility"`

	// Whether the constructor is implemented or not
	Implemented bool `json:"implemented"`

	// The modifiers of the constructor
	Modifiers []*ModifierInvocation `json:"modifiers"`

	// The parameters of the constructor
	Parameters *ParameterList `json:"parameters"`

	// The return parameters of the constructor, which are always empty for a constructor
	ReturnParameters *ParameterList `json:"return_parameters"`

	// The scope of the constructor, which is the id of the contract that the constructor belongs to
	Scope int64 `json:"scope"`

	// The body of the constructor, which is a block of statements
	Body *BodyNode `json:"body"`
}

// NewConstructor creates a new Constructor instance.
func NewConstructor(b *ASTBuilder) *Constructor {
	return &Constructor{
		ASTBuilder:      b,
		Id:              b.GetNextID(),
		NodeType:        ast_pb.NodeType_FUNCTION_DEFINITION,
		Kind:            ast_pb.NodeType_CONSTRUCTOR,
		StateMutability: ast_pb.Mutability_NONPAYABLE,
		Modifiers:       make([]*ModifierInvocation, 0),
	}
}

// SetReferenceDescriptor sets the reference descriptions of the Constructor node.
func (c *Constructor) SetReferenceDescriptor(refId int64, refDesc *TypeDescription) bool {
	return false
}

// GetId returns the unique identifier of the constructor.
func (c *Constructor) GetId() int64 {
	return c.Id
}

// GetSrc returns the source information about the constructor.
func (c *Constructor) GetSrc() SrcNode {
	return c.Src
}

// GetType returns the type of the node, which is 'FUNCTION_DEFINITION' for a constructor.
func (c *Constructor) GetType() ast_pb.NodeType {
	return c.NodeType
}

// GetNodes returns the statements in the body of the constructor.
func (c *Constructor) GetNodes() []Node[NodeType] {
	toReturn := []Node[NodeType]{}
	toReturn = append(toReturn, c.Body.GetNodes()...)
	toReturn = append(toReturn, c.GetParameters().GetNodes()...)
	toReturn = append(toReturn, c.GetReturnParameters().GetNodes()...)

	for _, m := range c.GetModifiers() {
		toReturn = append(toReturn, m.GetNodes()...)
	}

	return toReturn
}

// GetTypeDescription returns the type description of the constructor, which is nil as constructors do not have a type description.
func (c *Constructor) GetTypeDescription() *TypeDescription {
	return &TypeDescription{
		TypeString:     "constructor",
		TypeIdentifier: "$_t_constructor",
	}
}

func (c *Constructor) GetModifiers() []*ModifierInvocation {
	return c.Modifiers
}

// GetParameters returns the parameters of the constructor.
func (c *Constructor) GetParameters() *ParameterList {
	return c.Parameters
}

// GetReturnParameters returns the return parameters of the constructor, which are always empty for a constructor.
func (c *Constructor) GetReturnParameters() *ParameterList {
	return c.ReturnParameters
}

// GetBody returns the body of the constructor.
func (c *Constructor) GetBody() *BodyNode {
	return c.Body
}

// GetKind returns the kind of the node, which is 'CONSTRUCTOR' for a constructor.
func (c *Constructor) GetKind() ast_pb.NodeType {
	return c.Kind
}

// IsImplemented returns whether the constructor is implemented or not.
func (c *Constructor) IsImplemented() bool {
	return c.Implemented
}

// GetVisibility returns the visibility of the constructor.
func (c *Constructor) GetVisibility() ast_pb.Visibility {
	return c.Visibility
}

// GetStateMutability returns the state mutability of the constructor.
func (c *Constructor) GetStateMutability() ast_pb.Mutability {
	return c.StateMutability
}

// GetScope returns the scope of the constructor, which is the id of the contract that the constructor belongs to.
func (c *Constructor) GetScope() int64 {
	return c.Scope
}

func (c *Constructor) UnmarshalJSON(data []byte) error {
	var tempMap map[string]json.RawMessage
	if err := json.Unmarshal(data, &tempMap); err != nil {
		return err
	}

	if id, ok := tempMap["id"]; ok {
		if err := json.Unmarshal(id, &c.Id); err != nil {
			return err
		}
	}

	if nodeType, ok := tempMap["node_type"]; ok {
		if err := json.Unmarshal(nodeType, &c.NodeType); err != nil {
			return err
		}
	}

	if kind, ok := tempMap["kind"]; ok {
		if err := json.Unmarshal(kind, &c.Kind); err != nil {
			return err
		}
	}

	if src, ok := tempMap["src"]; ok {
		if err := json.Unmarshal(src, &c.Src); err != nil {
			return err
		}
	}

	if implemented, ok := tempMap["implemented"]; ok {
		if err := json.Unmarshal(implemented, &c.Implemented); err != nil {
			return err
		}
	}

	if visibility, ok := tempMap["visibility"]; ok {
		if err := json.Unmarshal(visibility, &c.Visibility); err != nil {
			return err
		}
	}

	if sm, ok := tempMap["state_mutability"]; ok {
		if err := json.Unmarshal(sm, &c.StateMutability); err != nil {
			return err
		}
	}

	if modifiers, ok := tempMap["modifiers"]; ok {
		if err := json.Unmarshal(modifiers, &c.Modifiers); err != nil {
			return err
		}
	}

	if params, ok := tempMap["parameters"]; ok {
		if err := json.Unmarshal(params, &c.Parameters); err != nil {
			return err
		}
	}

	if retParams, ok := tempMap["return_parameters"]; ok {
		if err := json.Unmarshal(retParams, &c.ReturnParameters); err != nil {
			return err
		}
	}

	if scope, ok := tempMap["scope"]; ok {
		if err := json.Unmarshal(scope, &c.Scope); err != nil {
			return err
		}
	}

	if body, ok := tempMap["body"]; ok {
		if err := json.Unmarshal(body, &c.Body); err != nil {
			return err
		}
	}

	return nil
}

// ToProto returns the protobuf representation of the constructor.
func (c *Constructor) ToProto() NodeType {
	proto := ast_pb.Function{
		Id:               c.GetId(),
		NodeType:         c.GetType(),
		Kind:             c.GetKind(),
		Src:              c.GetSrc().ToProto(),
		Implemented:      c.IsImplemented(),
		Scope:            c.GetScope(),
		Visibility:       c.GetVisibility(),
		StateMutability:  c.GetStateMutability(),
		Parameters:       c.GetParameters().ToProto(),
		ReturnParameters: c.GetReturnParameters().ToProto(),
	}

	if c.GetBody() != nil {
		proto.Body = c.GetBody().ToProto().(*ast_pb.Body)
	}

	if c.GetTypeDescription() != nil {
		proto.TypeDescription = c.GetTypeDescription().ToProto()
	}

	for _, m := range c.GetModifiers() {
		proto.Modifiers = append(proto.Modifiers, m.ToProto().(*ast_pb.ModifierInvocation))
	}

	return NewTypedStruct(&proto, "Function")
}

// Parse parses a constructor from the provided parser.ConstructorDefinitionContext and returns the corresponding Constructor.
func (c *Constructor) Parse(
	unit *SourceUnit[Node[ast_pb.SourceUnit]],
	contractNode Node[NodeType],
	ctx *parser.ConstructorDefinitionContext,
) Node[NodeType] {
	c.Scope = contractNode.GetId()
	c.Implemented = ctx.Block() != nil && !ctx.Block().IsEmpty()

	c.Src = SrcNode{
		Line:        int64(ctx.GetStart().GetLine()),
		Column:      int64(ctx.GetStart().GetColumn()),
		Start:       int64(ctx.GetStart().GetStart()),
		End:         int64(ctx.GetStop().GetStop()),
		Length:      int64(ctx.GetStop().GetStop() - ctx.GetStart().GetStart() + 1),
		ParentIndex: contractNode.GetId(),
	}

	for _, payableCtx := range ctx.AllPayable() {
		if payableCtx.GetText() == "payable" {
			c.StateMutability = ast_pb.Mutability_PAYABLE
		}
	}

	c.Visibility = c.getVisibilityFromCtx(ctx)

	params := NewParameterList(c.ASTBuilder)
	if ctx.ParameterList() != nil {
		params.Parse(unit, c, ctx.ParameterList())
	} else {
		params.Src = c.Src
		params.Src.ParentIndex = c.Id
	}
	c.Parameters = params

	returnParams := NewParameterList(c.ASTBuilder)
	returnParams.Src = c.Src
	returnParams.Src.ParentIndex = c.Id
	c.ReturnParameters = returnParams

	for _, modifierCtx := range ctx.AllModifierInvocation() {
		modifier := NewModifierInvocation(c.ASTBuilder)
		modifier.Parse(unit, contractNode, c, nil, modifierCtx)
		c.Modifiers = append(c.Modifiers, modifier)
	}

	if ctx.Block() != nil && !ctx.Block().IsEmpty() {
		bodyNode := NewBodyNode(c.ASTBuilder, false)
		bodyNode.ParseBlock(unit, contractNode, c, ctx.Block())
		c.Body = bodyNode

		if ctx.Block().AllUncheckedBlock() != nil {
			for _, uncheckedCtx := range ctx.Block().AllUncheckedBlock() {
				bodyNode := NewBodyNode(c.ASTBuilder, false)
				bodyNode.ParseUncheckedBlock(unit, contractNode, c, uncheckedCtx)
				c.Body.Statements = append(c.Body.Statements, bodyNode)
			}
		}
	}

	c.currentFunctions = append(c.currentFunctions, c)
	return c
}

// getVisibilityFromCtx returns the visibility of the constructor based on the provided parser.ConstructorDefinitionContext.
func (c *Constructor) getVisibilityFromCtx(ctx *parser.ConstructorDefinitionContext) ast_pb.Visibility {
	visibilityMap := map[string]ast_pb.Visibility{
		"public":   ast_pb.Visibility_PUBLIC,
		"internal": ast_pb.Visibility_INTERNAL,
	}

	if len(ctx.AllPublic()) > 0 {
		if v, ok := visibilityMap["public"]; ok {
			return v
		}
	} else if len(ctx.AllInternal()) > 0 {
		if v, ok := visibilityMap["internal"]; ok {
			return v
		}
	}

	return ast_pb.Visibility_INTERNAL
}
