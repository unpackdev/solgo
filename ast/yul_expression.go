package ast

import (
	"github.com/goccy/go-json"

	v3 "github.com/cncf/xds/go/xds/type/v3"
	ast_pb "github.com/unpackdev/protos/dist/go/ast"
	"github.com/unpackdev/solgo/parser"
)

// YulExpressionStatement represents a YUL expression statement in the abstract syntax tree.
type YulExpressionStatement struct {
	*ASTBuilder

	// Id is the unique identifier of the YUL expression statement.
	Id int64 `json:"id"`

	// NodeType is the type of the YUL expression statement node.
	NodeType ast_pb.NodeType `json:"nodeType"`

	// Src is the source location information of the YUL expression statement.
	Src SrcNode `json:"src"`

	// Expression is the expression node within the statement.
	Expression Node[NodeType] `json:"expression"`
}

// NewYulExpressionStatement creates a new YulExpressionStatement instance.
func NewYulExpressionStatement(b *ASTBuilder) *YulExpressionStatement {
	return &YulExpressionStatement{
		ASTBuilder: b,
		Id:         b.GetNextID(),
		NodeType:   ast_pb.NodeType_YUL_EXPRESSION,
	}
}

// SetReferenceDescriptor sets the reference descriptions of the YulExpressionStatement node.
func (y *YulExpressionStatement) SetReferenceDescriptor(refId int64, refDesc *TypeDescription) bool {
	return false
}

// GetId returns the ID of the YulExpressionStatement.
func (y *YulExpressionStatement) GetId() int64 {
	return y.Id
}

// GetType returns the NodeType of the YulExpressionStatement.
func (y *YulExpressionStatement) GetType() ast_pb.NodeType {
	return y.NodeType
}

// GetSrc returns the source information of the YulExpressionStatement.
func (y *YulExpressionStatement) GetSrc() SrcNode {
	return y.Src
}

// GetNodes returns a list containing the expression node.
func (y *YulExpressionStatement) GetNodes() []Node[NodeType] {
	toReturn := make([]Node[NodeType], 0)
	toReturn = append(toReturn, y.Expression)
	return toReturn
}

// GetTypeDescription returns the type description of the YulExpressionStatement.
func (y *YulExpressionStatement) GetTypeDescription() *TypeDescription {
	return &TypeDescription{}
}

// GetExpression returns the expression node within the statement.
func (y *YulExpressionStatement) GetExpression() Node[NodeType] {
	return y.Expression
}

// ToProto converts the YulExpressionStatement to its protocol buffer representation.
func (y *YulExpressionStatement) ToProto() NodeType {
	toReturn := ast_pb.YulExpressionStatement{
		Id:       y.GetId(),
		NodeType: y.GetType(),
		Src:      y.GetSrc().ToProto(),
	}

	if y.GetExpression() != nil {
		toReturn.Expression = y.GetExpression().ToProto().(*v3.TypedStruct)
	}

	return NewTypedStruct(&toReturn, "YulExpressionStatement")
}

// UnmarshalJSON unmarshals a given JSON byte array into a YulExpressionStatement node.
func (f *YulExpressionStatement) UnmarshalJSON(data []byte) error {
	var tempMap map[string]json.RawMessage
	if err := json.Unmarshal(data, &tempMap); err != nil {
		return err
	}

	if id, ok := tempMap["id"]; ok {
		if err := json.Unmarshal(id, &f.Id); err != nil {
			return err
		}
	}

	if nodeType, ok := tempMap["nodeType"]; ok {
		if err := json.Unmarshal(nodeType, &f.NodeType); err != nil {
			return err
		}
	}

	if src, ok := tempMap["src"]; ok {
		if err := json.Unmarshal(src, &f.Src); err != nil {
			return err
		}
	}

	if expression, ok := tempMap["expression"]; ok {
		if err := json.Unmarshal(expression, &f.Expression); err != nil {
			var tempNodeMap map[string]json.RawMessage
			if err := json.Unmarshal(expression, &tempNodeMap); err != nil {
				return err
			}

			var tempNodeType ast_pb.NodeType
			if err := json.Unmarshal(tempNodeMap["nodeType"], &tempNodeType); err != nil {
				return err
			}

			node, err := unmarshalNode(expression, tempNodeType)
			if err != nil {
				return err
			}
			f.Expression = node
		}
	}

	return nil
}

// Parse parses a YulExpressionStatement node.
func (y *YulExpressionStatement) Parse(
	unit *SourceUnit[Node[ast_pb.SourceUnit]],
	contractNode Node[NodeType],
	fnNode Node[NodeType],
	bodyNode *BodyNode,
	assemblyNode *Yul,
	statementNode *YulStatement,
	parentNode Node[NodeType],
	ctx *parser.YulExpressionContext,
) Node[NodeType] {
	y.Src = SrcNode{
		Line:        int64(ctx.GetStart().GetLine()),
		Column:      int64(ctx.GetStart().GetColumn()),
		Start:       int64(ctx.GetStart().GetStart()),
		End:         int64(ctx.GetStart().GetStop()),
		Length:      int64(ctx.GetStart().GetStop() - ctx.GetStart().GetStart() + 1),
		ParentIndex: parentNode.GetId(),
	}

	if ctx.YulLiteral() != nil {
		literalStatement := NewYulLiteralStatement(y.ASTBuilder)
		y.Expression = literalStatement.Parse(
			unit, contractNode, fnNode, bodyNode, assemblyNode, statementNode, parentNode,
			ctx.YulLiteral().(*parser.YulLiteralContext),
		)
	}

	if ctx.YulFunctionCall() != nil {
		fcStatement := NewYulFunctionCallStatement(y.ASTBuilder)
		y.Expression = fcStatement.Parse(
			unit, contractNode, fnNode, bodyNode, assemblyNode, statementNode, parentNode,
			ctx.YulFunctionCall().(*parser.YulFunctionCallContext),
		)
	}

	return y
}

// ParseYulExpression parses a YUL expression statement.
func ParseYulExpression(
	b *ASTBuilder,
	unit *SourceUnit[Node[ast_pb.SourceUnit]],
	contractNode Node[NodeType],
	fnNode Node[NodeType],
	bodyNode *BodyNode,
	assemblyNode *Yul,
	statementNode *YulStatement,
	assignmentNode *parser.YulAssignmentContext,
	variableNode *parser.YulVariableDeclarationContext,
	parentNode Node[NodeType],
	ctx parser.IYulExpressionContext,
) Node[NodeType] {
	if ctx.YulLiteral() != nil {
		literalStatement := NewYulLiteralStatement(b)
		return literalStatement.Parse(
			unit, contractNode, fnNode, bodyNode, assemblyNode, statementNode, parentNode,
			ctx.YulLiteral().(*parser.YulLiteralContext),
		)
	}

	if ctx.YulFunctionCall() != nil {
		fcStatement := NewYulFunctionCallStatement(b)
		return fcStatement.Parse(
			unit, contractNode, fnNode, bodyNode, assemblyNode, statementNode, parentNode,
			ctx.YulFunctionCall().(*parser.YulFunctionCallContext),
		)
	}

	return nil
}
