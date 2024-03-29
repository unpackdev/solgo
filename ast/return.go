package ast

import (
	"github.com/goccy/go-json"

	v3 "github.com/cncf/xds/go/xds/type/v3"
	ast_pb "github.com/unpackdev/protos/dist/go/ast"
	"github.com/unpackdev/solgo/parser"
)

// ReturnStatement represents a return statement in the AST.
type ReturnStatement struct {
	*ASTBuilder

	Id                       int64           `json:"id"`
	NodeType                 ast_pb.NodeType `json:"node_type"`
	Src                      SrcNode         `json:"src"`
	FunctionReturnParameters int64           `json:"function_return_parameters"`
	Expression               Node[NodeType]  `json:"expression"`
}

// NewReturnStatement creates a new instance of ReturnStatement using the provided ASTBuilder.
func NewReturnStatement(b *ASTBuilder) *ReturnStatement {
	return &ReturnStatement{
		ASTBuilder: b,
		Id:         b.GetNextID(),
		NodeType:   ast_pb.NodeType_RETURN_STATEMENT,
	}
}

// SetReferenceDescriptor sets the reference descriptions of the ReturnStatement node.
func (r *ReturnStatement) SetReferenceDescriptor(refId int64, refDesc *TypeDescription) bool {
	return false
}

// GetId returns the ID of the ReturnStatement node.
func (r *ReturnStatement) GetId() int64 {
	return r.Id
}

// GetType returns the NodeType of the ReturnStatement node.
func (r *ReturnStatement) GetType() ast_pb.NodeType {
	return r.NodeType
}

// GetSrc returns the source information of the ReturnStatement node.
func (r *ReturnStatement) GetSrc() SrcNode {
	return r.Src
}

// GetExpression returns the expression associated with the ReturnStatement node.
func (r *ReturnStatement) GetExpression() Node[NodeType] {
	return r.Expression
}

// GetFunctionReturnParameters returns the ID of the function's return parameters.
func (r *ReturnStatement) GetFunctionReturnParameters() int64 {
	return r.FunctionReturnParameters
}

// GetTypeDescription returns the type description of the ReturnStatement's expression.
func (r *ReturnStatement) GetTypeDescription() *TypeDescription {
	if r.Expression == nil {
		return &TypeDescription{
			TypeString:     "void",
			TypeIdentifier: "$_t_return_void",
		}
	}

	return r.Expression.GetTypeDescription()
}

// GetNodes returns a list of child nodes contained in the ReturnStatement.
func (r *ReturnStatement) GetNodes() []Node[NodeType] {
	return []Node[NodeType]{r.Expression}
}

func (r *ReturnStatement) UnmarshalJSON(data []byte) error {
	var tempMap map[string]json.RawMessage
	if err := json.Unmarshal(data, &tempMap); err != nil {
		return err
	}

	if id, ok := tempMap["id"]; ok {
		if err := json.Unmarshal(id, &r.Id); err != nil {
			return err
		}
	}

	if nodeType, ok := tempMap["node_type"]; ok {
		if err := json.Unmarshal(nodeType, &r.NodeType); err != nil {
			return err
		}
	}

	if src, ok := tempMap["src"]; ok {
		if err := json.Unmarshal(src, &r.Src); err != nil {
			return err
		}
	}

	if functionReturnParameters, ok := tempMap["function_return_parameters"]; ok {
		if err := json.Unmarshal(functionReturnParameters, &r.FunctionReturnParameters); err != nil {
			return err
		}
	}

	if expression, ok := tempMap["expression"]; ok {
		if err := json.Unmarshal(expression, &r.Expression); err != nil {
			var tempNodeMap map[string]json.RawMessage
			if err := json.Unmarshal(expression, &tempNodeMap); err != nil {
				return err
			}

			var tempNodeType ast_pb.NodeType
			if err := json.Unmarshal(tempNodeMap["node_type"], &tempNodeType); err != nil {
				return err
			}

			node, err := unmarshalNode(expression, tempNodeType)
			if err != nil {
				return err
			}
			r.Expression = node
		}
	}

	return nil

}

// ToProto converts the ReturnStatement into its corresponding Protocol Buffers representation.
func (r *ReturnStatement) ToProto() NodeType {
	proto := ast_pb.Return{
		Id:                       r.GetId(),
		NodeType:                 r.GetType(),
		Src:                      r.Src.ToProto(),
		FunctionReturnParameters: r.GetFunctionReturnParameters(),
	}

	if r.GetExpression() != nil {
		proto.Expression = r.GetExpression().ToProto().(*v3.TypedStruct)
	}

	if r.GetTypeDescription() != nil {
		proto.TypeDescription = r.GetTypeDescription().ToProto()
	}

	return NewTypedStruct(&proto, "Return")
}

// Parse parses the ReturnStatement node from the provided context.
func (r *ReturnStatement) Parse(
	unit *SourceUnit[Node[ast_pb.SourceUnit]],
	contractNode Node[NodeType],
	fnNode Node[NodeType],
	bodyNode *BodyNode,
	ctx *parser.ReturnStatementContext,
) Node[NodeType] {
	r.Src = SrcNode{
		Line:        int64(ctx.GetStart().GetLine()),
		Column:      int64(ctx.GetStart().GetColumn()),
		Start:       int64(ctx.GetStart().GetStart()),
		End:         int64(ctx.GetStop().GetStop()),
		Length:      int64(ctx.GetStop().GetStop() - ctx.GetStart().GetStart() + 1),
		ParentIndex: fnNode.GetId(),
	}

	if fnCtx, ok := fnNode.(*Function); ok {
		if fnCtx.GetReturnParameters() != nil {
			r.FunctionReturnParameters = fnCtx.GetId()
		}
	} else if fnCtx, ok := fnNode.(*TryStatement); ok {
		r.FunctionReturnParameters = fnCtx.GetId()
	}

	if ctx.Expression() != nil {
		expression := NewExpression(r.ASTBuilder)
		r.Expression = expression.Parse(unit, contractNode, fnNode, bodyNode, nil, r, r.GetId(), ctx.Expression())
	}

	return r
}
