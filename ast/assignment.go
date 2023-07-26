package ast

import (
	"reflect"

	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/parser"
	"go.uber.org/zap"
)

type Assignment struct {
	*ASTBuilder

	Id              int64           `json:"id"`
	NodeType        ast_pb.NodeType `json:"node_type"`
	Src             SrcNode         `json:"src"`
	Expression      Node[NodeType]  `json:"expression,omitempty"`
	Operator        ast_pb.Operator `json:"operator,omitempty"`
	LeftExpression  Node[NodeType]  `json:"left_expression,omitempty"`
	RightExpression Node[NodeType]  `json:"right_expression,omitempty"`
}

func NewAssignment(b *ASTBuilder) *Assignment {
	return &Assignment{
		ASTBuilder: b,
		Id:         b.GetNextID(),
		NodeType:   ast_pb.NodeType_EXPRESSION_STATEMENT,
	}
}

func (a *Assignment) GetId() int64 {
	return a.Id
}

func (a *Assignment) GetType() ast_pb.NodeType {
	return a.NodeType
}

func (a *Assignment) GetSrc() SrcNode {
	return a.Src
}

func (a *Assignment) GetTypeDescription() *TypeDescription {
	return nil
}

func (a *Assignment) GetNodes() []Node[NodeType] {
	return nil
}

func (a *Assignment) ToProto() NodeType {
	return ast_pb.Statement{}
}

func (a *Assignment) ParseStatement(
	unit *SourceUnit[Node[ast_pb.SourceUnit]],
	contractNode Node[NodeType],
	fnNode Node[NodeType],
	bodyNode *BodyNode,
	eCtx *parser.ExpressionStatementContext,
	ctx *parser.AssignmentContext,
) {
	a.Src = SrcNode{
		Id:          a.GetNextID(),
		Line:        int64(eCtx.GetStart().GetLine()),
		Column:      int64(eCtx.GetStart().GetColumn()),
		Start:       int64(eCtx.GetStart().GetStart()),
		End:         int64(eCtx.GetStop().GetStop()),
		Length:      int64(eCtx.GetStop().GetStop() - eCtx.GetStart().GetStart() + 1),
		ParentIndex: bodyNode.GetId(),
	}

	expression := NewExpression(a.ASTBuilder)
	a.Expression = expression.Parse(unit, contractNode, fnNode, bodyNode, nil, nil, ctx)
}

func (a *Assignment) Parse(
	unit *SourceUnit[Node[ast_pb.SourceUnit]],
	contractNode Node[NodeType],
	fnNode Node[NodeType],
	bodyNode *BodyNode,
	vDeclar *VariableDeclaration,
	expNode Node[NodeType],
	ctx *parser.AssignmentContext,
) {
	a.NodeType = ast_pb.NodeType_ASSIGNMENT
	a.Src = SrcNode{
		Id:          a.GetNextID(),
		Line:        int64(ctx.GetStart().GetLine()),
		Column:      int64(ctx.GetStart().GetColumn()),
		Start:       int64(ctx.GetStart().GetStart()),
		End:         int64(ctx.GetStop().GetStop()),
		Length:      int64(ctx.GetStop().GetStop() - ctx.GetStart().GetStart() + 1),
		ParentIndex: bodyNode.GetId(),
	}

	operator := ctx.AssignOp()
	if operator != nil {
		if operator.Assign() != nil {
			a.Operator = ast_pb.Operator_EQUAL
		} else if operator.AssignAdd() != nil {
			a.Operator = ast_pb.Operator_PLUS_EQUAL
		} else if operator.AssignSub() != nil {
			a.Operator = ast_pb.Operator_MINUS_EQUAL
		} else if operator.AssignMul() != nil {
			a.Operator = ast_pb.Operator_MUL_EQUAL
		} else if operator.AssignDiv() != nil {
			a.Operator = ast_pb.Operator_DIVISION
		} else if operator.AssignMod() != nil {
			a.Operator = ast_pb.Operator_MOD_EQUAL
		} else if operator.AssignBitAnd() != nil {
			a.Operator = ast_pb.Operator_AND_EQUAL
		} else if operator.AssignBitOr() != nil {
			a.Operator = ast_pb.Operator_OR_EQUAL
		} else if operator.AssignBitXor() != nil {
			a.Operator = ast_pb.Operator_XOR_EQUAL
		} else if operator.AssignShl() != nil {
			a.Operator = ast_pb.Operator_SHIFT_LEFT_EQUAL
		} else if operator.AssignShr() != nil {
			a.Operator = ast_pb.Operator_SHIFT_RIGHT_EQUAL
		} else if operator.AssignBitAnd() != nil {
			a.Operator = ast_pb.Operator_BIT_AND_EQUAL
		} else if operator.AssignBitOr() != nil {
			a.Operator = ast_pb.Operator_BIT_OR_EQUAL
		} else if operator.AssignBitXor() != nil {
			a.Operator = ast_pb.Operator_BIT_XOR_EQUAL
		} else if operator.AssignSar() != nil {
			a.Operator = ast_pb.Operator_POW_EQUAL
		} else {
			zap.L().Warn(
				"Assignment operator not recognized",
				zap.String("type", reflect.TypeOf(operator).String()),
			)
		}
	}

	expression := NewExpression(a.ASTBuilder)
	a.LeftExpression = expression.Parse(
		unit, contractNode, fnNode, bodyNode, vDeclar, expNode, ctx.Expression(0),
	)

	a.RightExpression = expression.Parse(
		unit, contractNode, fnNode, bodyNode, vDeclar, expNode, ctx.Expression(1),
	)
}
