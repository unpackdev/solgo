package ast

import (
	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/parser"
)

type Constructor[T ast_pb.Function] struct {
	*ASTBuilder

	Id               int64             `json:"id"`
	NodeType         ast_pb.NodeType   `json:"node_type"`
	Src              SrcNode           `json:"src"`
	Kind             ast_pb.NodeType   `json:"kind"`
	StateMutability  ast_pb.Mutability `json:"state_mutability"`
	Visibility       ast_pb.Visibility `json:"visibility"`
	Implemented      bool              `json:"implemented"`
	Parameters       *ParameterList[T] `json:"parameters"`
	ReturnParameters *ParameterList[T] `json:"return_parameters"`
	Scope            int64             `json:"scope"`
	Body             *BodyNode         `json:"body"`
}

func NewConstructor[T ast_pb.Function](b *ASTBuilder) *Constructor[T] {
	return &Constructor[T]{
		ASTBuilder:      b,
		Id:              b.GetNextID(),
		NodeType:        ast_pb.NodeType_FUNCTION_DEFINITION,
		Kind:            ast_pb.NodeType_CONSTRUCTOR,
		StateMutability: ast_pb.Mutability_NONPAYABLE,
	}
}

func (c *Constructor[T]) GetId() int64 {
	return c.Id
}

func (c *Constructor[T]) GetSrc() SrcNode {
	return c.Src
}

func (c *Constructor[T]) GetType() ast_pb.NodeType {
	return c.NodeType
}

func (c *Constructor[T]) GetNodes() []Node[NodeType] {
	return c.Body.Statements
}

func (c *Constructor[T]) GetTypeDescription() *TypeDescription {
	return nil
}

func (c *Constructor[T]) GetParameters() *ParameterList[T] {
	return c.Parameters
}

func (c *Constructor[T]) GetReturnParameters() *ParameterList[T] {
	return c.ReturnParameters
}

func (c *Constructor[T]) GetBody() *BodyNode {
	return c.Body
}

func (c *Constructor[T]) GetKind() ast_pb.NodeType {
	return c.Kind
}

func (c *Constructor[T]) IsImplemented() bool {
	return c.Implemented
}

func (c *Constructor[T]) GetVisibility() ast_pb.Visibility {
	return c.Visibility
}

func (c *Constructor[T]) GetStateMutability() ast_pb.Mutability {
	return c.StateMutability
}

func (c *Constructor[T]) GetScope() int64 {
	return c.Scope
}

func (c *Constructor[T]) ToProto() NodeType {
	return ast_pb.Function{}
}

func (c *Constructor[T]) Parse(
	unit *SourceUnit[Node[ast_pb.SourceUnit]],
	contractNode Node[NodeType],
	ctx *parser.ConstructorDefinitionContext,
) Node[NodeType] {
	c.Scope = contractNode.GetId()
	c.Implemented = ctx.Block() != nil && !ctx.Block().IsEmpty()

	c.Src = SrcNode{
		Id:          c.GetNextID(),
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

	params := NewParameterList[T](c.ASTBuilder)
	params.Parse(unit, c, ctx.ParameterList())
	c.Parameters = params

	c.ReturnParameters = &ParameterList[T]{
		Id: c.GetNextID(),
		Src: SrcNode{
			Id:          c.GetNextID(),
			Line:        int64(ctx.GetStart().GetLine()),
			Column:      int64(ctx.GetStart().GetColumn()),
			Start:       int64(ctx.GetStart().GetStart()),
			End:         int64(ctx.GetStop().GetStop()),
			Length:      int64(ctx.GetStop().GetStop() - ctx.GetStart().GetStart() + 1),
			ParentIndex: c.Id,
		},
		NodeType:   ast_pb.NodeType_PARAMETER_LIST,
		Parameters: []*Parameter{},
	}

	if ctx.Block() != nil && !ctx.Block().IsEmpty() {
		bodyNode := NewBodyNode(c.ASTBuilder)
		bodyNode.ParseBlock(unit, contractNode, c, ctx.Block())
		c.Body = bodyNode

		if ctx.Block().AllUncheckedBlock() != nil {
			for _, uncheckedCtx := range ctx.Block().AllUncheckedBlock() {
				bodyNode := NewBodyNode(c.ASTBuilder)
				bodyNode.ParseUncheckedBlock(unit, contractNode, c, uncheckedCtx)
				c.Body.Statements = append(c.Body.Statements, bodyNode)
			}
		}
	}

	return c
}

func (c *Constructor[T]) getVisibilityFromCtx(ctx *parser.ConstructorDefinitionContext) ast_pb.Visibility {
	visibilityMap := map[string]ast_pb.Visibility{
		"public":   ast_pb.Visibility_PUBLIC,
		"internal": ast_pb.Visibility_INTERNAL,
	}

	// Check each visibility context in the map
	if len(ctx.AllPublic()) > 0 {
		if v, ok := visibilityMap["public"]; ok {
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
