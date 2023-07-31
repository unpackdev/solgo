package ast

import (
	"reflect"

	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/parser"
	"go.uber.org/zap"
)

// Assignment represents an assignment statement in the AST.
type Assignment struct {
	*ASTBuilder

	Id                    int64            `json:"id"`
	NodeType              ast_pb.NodeType  `json:"node_type"`
	Src                   SrcNode          `json:"src"`
	Expression            Node[NodeType]   `json:"expression,omitempty"`
	Operator              ast_pb.Operator  `json:"operator,omitempty"`
	LeftExpression        Node[NodeType]   `json:"left_expression,omitempty"`
	RightExpression       Node[NodeType]   `json:"right_expression,omitempty"`
	ReferencedDeclaration int64            `json:"referenced_declaration,omitempty"`
	TypeDescription       *TypeDescription `json:"type_description,omitempty"`
}

// NewAssignment creates a new Assignment node with a given ASTBuilder.
func NewAssignment(b *ASTBuilder) *Assignment {
	return &Assignment{
		ASTBuilder: b,
		Id:         b.GetNextID(),
		NodeType:   ast_pb.NodeType_EXPRESSION_STATEMENT,
	}
}

// GetId returns the ID of the Assignment node.
func (a *Assignment) GetId() int64 {
	return a.Id
}

// GetType returns the NodeType of the Assignment node.
func (a *Assignment) GetType() ast_pb.NodeType {
	return a.NodeType
}

// GetSrc returns the SrcNode of the Assignment node.
func (a *Assignment) GetSrc() SrcNode {
	return a.Src
}

// GetTypeDescription returns the TypeDescription of the Assignment node.
func (a *Assignment) GetTypeDescription() *TypeDescription {
	return a.TypeDescription
}

// GetNodes returns the child nodes of the Assignment node.
func (a *Assignment) GetNodes() []Node[NodeType] {
	return nil
}

// ToProto returns a protobuf representation of the Assignment node.
func (a *Assignment) ToProto() NodeType {
	return ast_pb.Statement{}
}

// SetReferenceDescriptor sets the reference descriptions of the Assignment node.
func (a *Assignment) SetReferenceDescriptor(refId int64, refDesc *TypeDescription) bool {
	a.ReferencedDeclaration = refId
	a.TypeDescription = refDesc
	return false
}

// ParseStatement parses an expression statement context into the Assignment node.
func (a *Assignment) ParseStatement(
	unit *SourceUnit[Node[ast_pb.SourceUnit]],
	contractNode Node[NodeType],
	fnNode Node[NodeType],
	bodyNode *BodyNode,
	parentNode Node[NodeType],
	eCtx *parser.ExpressionStatementContext,
	ctx *parser.AssignmentContext,
) {
	a.Src = SrcNode{
		Id:     a.GetNextID(),
		Line:   int64(eCtx.GetStart().GetLine()),
		Column: int64(eCtx.GetStart().GetColumn()),
		Start:  int64(eCtx.GetStart().GetStart()),
		End:    int64(eCtx.GetStop().GetStop()),
		Length: int64(eCtx.GetStop().GetStop() - eCtx.GetStart().GetStart() + 1),
		ParentIndex: func() int64 {
			if parentNode != nil {
				return parentNode.GetId()
			}

			if bodyNode != nil {
				return bodyNode.GetId()
			}

			if fnNode != nil {
				return fnNode.GetId()
			}

			return contractNode.GetId()
		}(),
	}

	expression := NewExpression(a.ASTBuilder)
	a.Expression = expression.Parse(unit, contractNode, fnNode, bodyNode, nil, nil, ctx)
}

// Parse parses an assignment context into the Assignment node.
func (a *Assignment) Parse(
	unit *SourceUnit[Node[ast_pb.SourceUnit]],
	contractNode Node[NodeType],
	fnNode Node[NodeType],
	bodyNode *BodyNode,
	vDeclar *VariableDeclaration,
	expNode Node[NodeType],
	ctx *parser.AssignmentContext,
) Node[NodeType] {
	a.NodeType = ast_pb.NodeType_ASSIGNMENT
	a.Src = SrcNode{
		Id:     a.GetNextID(),
		Line:   int64(ctx.GetStart().GetLine()),
		Column: int64(ctx.GetStart().GetColumn()),
		Start:  int64(ctx.GetStart().GetStart()),
		End:    int64(ctx.GetStop().GetStop()),
		Length: int64(ctx.GetStop().GetStop() - ctx.GetStart().GetStart() + 1),
		ParentIndex: func() int64 {
			if expNode != nil {
				return expNode.GetId()
			}

			if vDeclar != nil {
				return vDeclar.GetId()
			}

			return bodyNode.GetId()
		}(),
	}

	a.Operator = parseOperator(ctx.AssignOp())
	expression := NewExpression(a.ASTBuilder)
	a.LeftExpression = expression.Parse(unit, contractNode, fnNode, bodyNode, vDeclar, a, ctx.Expression(0))
	a.RightExpression = expression.Parse(unit, contractNode, fnNode, bodyNode, vDeclar, a, ctx.Expression(1))

	// What we are going to do here is take left expression type description and assign it to the
	// asignment type description. This is because the left expression is the one that holds the
	// type description of the assignment.
	a.TypeDescription = a.LeftExpression.GetTypeDescription()

	return a
}

// parseOperator parses the assignment operator from the context.
func parseOperator(operator parser.IAssignOpContext) ast_pb.Operator {
	if operator == nil {
		return ast_pb.Operator_O_DEFAULT
	}

	switch {
	case operator.Assign() != nil:
		return ast_pb.Operator_EQUAL
	case operator.AssignAdd() != nil:
		return ast_pb.Operator_PLUS_EQUAL
	case operator.AssignSub() != nil:
		return ast_pb.Operator_MINUS_EQUAL
	case operator.AssignMul() != nil:
		return ast_pb.Operator_MUL_EQUAL
	case operator.AssignDiv() != nil:
		return ast_pb.Operator_DIVISION
	case operator.AssignMod() != nil:
		return ast_pb.Operator_MOD_EQUAL
	case operator.AssignBitAnd() != nil:
		return ast_pb.Operator_AND_EQUAL
	case operator.AssignBitOr() != nil:
		return ast_pb.Operator_OR_EQUAL
	case operator.AssignBitXor() != nil:
		return ast_pb.Operator_XOR_EQUAL
	case operator.AssignShl() != nil:
		return ast_pb.Operator_SHIFT_LEFT_EQUAL
	case operator.AssignShr() != nil:
		return ast_pb.Operator_SHIFT_RIGHT_EQUAL
	case operator.AssignBitAnd() != nil:
		return ast_pb.Operator_BIT_AND_EQUAL
	case operator.AssignBitOr() != nil:
		return ast_pb.Operator_BIT_OR_EQUAL
	case operator.AssignBitXor() != nil:
		return ast_pb.Operator_BIT_XOR_EQUAL
	case operator.AssignSar() != nil:
		return ast_pb.Operator_POW_EQUAL
	default:
		zap.L().Warn(
			"Assignment operator not recognized",
			zap.String("type", reflect.TypeOf(operator).String()),
		)
		return ast_pb.Operator_O_DEFAULT
	}
}
