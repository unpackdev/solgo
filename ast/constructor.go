package ast

import (
	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/parser"
)

type Constructor struct {
	*ASTBuilder

	Id               int64             `json:"id"`
	NodeType         ast_pb.NodeType   `json:"node_type"`
	Src              SrcNode           `json:"src"`
	Kind             ast_pb.NodeType   `json:"kind"`
	StateMutability  ast_pb.Mutability `json:"state_mutability"`
	Visibility       ast_pb.Visibility `json:"visibility"`
	Implemented      bool              `json:"implemented"`
	Parameters       *ParameterList    `json:"parameters"`
	ReturnParameters *ParameterList    `json:"return_parameters"`
	Scope            int64             `json:"scope"`
	Body             *BodyNode         `json:"body"`
}

func NewConstructor(b *ASTBuilder) *Constructor {
	return &Constructor{
		ASTBuilder:      b,
		Id:              b.GetNextID(),
		NodeType:        ast_pb.NodeType_FUNCTION_DEFINITION,
		Kind:            ast_pb.NodeType_CONSTRUCTOR,
		StateMutability: ast_pb.Mutability_NONPAYABLE,
	}
}

func (c *Constructor) GetId() int64 {
	return c.Id
}

func (c *Constructor) GetSrc() SrcNode {
	return c.Src
}

func (c *Constructor) GetType() ast_pb.NodeType {
	return c.NodeType
}

func (c *Constructor) GetNodes() []Node[NodeType] {
	return c.Body.Statements
}

func (c *Constructor) GetTypeDescription() *TypeDescription {
	return nil
}

func (c *Constructor) GetParameters() *ParameterList {
	return c.Parameters
}

func (c *Constructor) GetReturnParameters() *ParameterList {
	return c.ReturnParameters
}

func (c *Constructor) GetBody() *BodyNode {
	return c.Body
}

func (c *Constructor) GetKind() ast_pb.NodeType {
	return c.Kind
}

func (c *Constructor) IsImplemented() bool {
	return c.Implemented
}

func (c *Constructor) GetVisibility() ast_pb.Visibility {
	return c.Visibility
}

func (c *Constructor) GetStateMutability() ast_pb.Mutability {
	return c.StateMutability
}

func (c *Constructor) GetScope() int64 {
	return c.Scope
}

func (c *Constructor) ToProto() NodeType {
	return ast_pb.Function{}
}

func (c *Constructor) Parse(
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

func (c *Constructor) getVisibilityFromCtx(ctx *parser.ConstructorDefinitionContext) ast_pb.Visibility {
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
