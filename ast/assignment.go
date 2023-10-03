package ast

import (
	"encoding/json"
	"reflect"

	v3 "github.com/cncf/xds/go/xds/type/v3"
	ast_pb "github.com/unpackdev/protos/dist/go/ast"
	"github.com/unpackdev/solgo/parser"
	"go.uber.org/zap"
)

// Assignment represents an assignment statement in the AST.
type Assignment struct {
	*ASTBuilder

	Id                    int64            `json:"id"`                               // Unique identifier for the Assignment node.
	NodeType              ast_pb.NodeType  `json:"node_type"`                        // Type of the AST node.
	Src                   SrcNode          `json:"src"`                              // Source location information.
	Expression            Node[NodeType]   `json:"expression,omitempty"`             // Expression for the assignment (if used).
	Operator              ast_pb.Operator  `json:"operator,omitempty"`               // Operator used in the assignment.
	LeftExpression        Node[NodeType]   `json:"left_expression,omitempty"`        // Left-hand side expression.
	RightExpression       Node[NodeType]   `json:"right_expression,omitempty"`       // Right-hand side expression.
	ReferencedDeclaration int64            `json:"referenced_declaration,omitempty"` // Referenced declaration identifier (if used).
	TypeDescription       *TypeDescription `json:"type_description,omitempty"`       // Type description associated with the Assignment node.
}

// NewAssignment creates a new Assignment node with a given ASTBuilder.
func NewAssignment(b *ASTBuilder) *Assignment {
	return &Assignment{
		ASTBuilder: b,
		Id:         b.GetNextID(),
		NodeType:   ast_pb.NodeType_ASSIGNMENT,
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
	toReturn := []Node[NodeType]{}
	if a.Expression != nil {
		toReturn = append(toReturn, a.Expression)
	}

	if a.LeftExpression != nil {
		toReturn = append(toReturn, a.LeftExpression)
	}

	if a.RightExpression != nil {
		toReturn = append(toReturn, a.RightExpression)
		toReturn = append(toReturn, a.RightExpression.GetNodes()...)
	}

	return toReturn
}

// GetExpression returns the expression of the Assignment node.
func (a *Assignment) GetExpression() Node[NodeType] {
	return a.Expression
}

// GetOperator returns the operator of the Assignment node.
func (a *Assignment) GetOperator() ast_pb.Operator {
	return a.Operator
}

// GetLeftExpression returns the left expression of the Assignment node.
func (a *Assignment) GetLeftExpression() Node[NodeType] {
	return a.LeftExpression
}

// GetRightExpression returns the right expression of the Assignment node.
func (a *Assignment) GetRightExpression() Node[NodeType] {
	return a.RightExpression
}

// GetReferencedDeclaration returns the referenced declaration of the Assignment node.
func (a *Assignment) GetReferencedDeclaration() int64 {
	return a.ReferencedDeclaration
}

// UnmarshalJSON sets the Assignment node data from the JSON byte array.
func (a *Assignment) UnmarshalJSON(data []byte) error {
	var tempMap map[string]json.RawMessage
	if err := json.Unmarshal(data, &tempMap); err != nil {
		return err
	}

	if id, ok := tempMap["id"]; ok {
		if err := json.Unmarshal(id, &a.Id); err != nil {
			return err
		}
	}

	if nodeType, ok := tempMap["node_type"]; ok {
		if err := json.Unmarshal(nodeType, &a.NodeType); err != nil {
			return err
		}
	}

	if src, ok := tempMap["src"]; ok {
		if err := json.Unmarshal(src, &a.Src); err != nil {
			return err
		}
	}

	if referencedDeclaration, ok := tempMap["referenced_declaration"]; ok {
		if err := json.Unmarshal(referencedDeclaration, &a.ReferencedDeclaration); err != nil {
			return err
		}
	}

	if typeDescription, ok := tempMap["type_description"]; ok {
		if err := json.Unmarshal(typeDescription, &a.TypeDescription); err != nil {
			return err
		}
	}

	if operator, ok := tempMap["operator"]; ok {
		if err := json.Unmarshal(operator, &a.Operator); err != nil {
			return err
		}
	}

	if leftExpression, ok := tempMap["left_expression"]; ok {
		if err := json.Unmarshal(leftExpression, &a.LeftExpression); err != nil {
			var tempNodeMap map[string]json.RawMessage
			if err := json.Unmarshal(leftExpression, &tempNodeMap); err != nil {
				return err
			}

			var tempNodeType ast_pb.NodeType
			if err := json.Unmarshal(tempNodeMap["node_type"], &tempNodeType); err != nil {
				return err
			}

			node, err := unmarshalNode(leftExpression, tempNodeType)
			if err != nil {
				return err
			}
			a.LeftExpression = node
		}
	}

	if rightExpression, ok := tempMap["right_expression"]; ok {
		if err := json.Unmarshal(rightExpression, &a.RightExpression); err != nil {
			var tempNodeMap map[string]json.RawMessage
			if err := json.Unmarshal(rightExpression, &tempNodeMap); err != nil {
				return err
			}

			var tempNodeType ast_pb.NodeType
			if err := json.Unmarshal(tempNodeMap["node_type"], &tempNodeType); err != nil {
				return err
			}

			node, err := unmarshalNode(rightExpression, tempNodeType)
			if err != nil {
				return err
			}
			a.RightExpression = node
		}
	}

	if expression, ok := tempMap["expression"]; ok {
		if err := json.Unmarshal(expression, &a.Expression); err != nil {
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
			a.Expression = node
		}
	}

	return nil
}

// ToProto returns a protobuf representation of the Assignment node.
func (a *Assignment) ToProto() NodeType {
	proto := ast_pb.Assignment{
		Id:                    a.GetId(),
		NodeType:              a.GetType(),
		Src:                   a.GetSrc().ToProto(),
		ReferencedDeclaration: a.GetReferencedDeclaration(),
		Operator:              a.GetOperator(),
	}

	if a.GetExpression() != nil {
		proto.Expression = a.GetExpression().ToProto().(*v3.TypedStruct)
	}

	if a.GetLeftExpression() != nil {
		proto.LeftExpression = a.GetLeftExpression().ToProto().(*v3.TypedStruct)
	}

	if a.GetRightExpression() != nil {
		proto.RightExpression = a.GetRightExpression().ToProto().(*v3.TypedStruct)
	}

	if a.GetTypeDescription() != nil {
		proto.TypeDescription = a.GetTypeDescription().ToProto()
	}

	return NewTypedStruct(&proto, "Assignment")
}

// SetReferenceDescriptor sets the reference descriptions of the Assignment node.
func (a *Assignment) SetReferenceDescriptor(refId int64, refDesc *TypeDescription) bool {
	a.ReferencedDeclaration = refId
	a.TypeDescription = refDesc
	return true
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
	// Setting the source location information.
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

	// Parsing the expression and setting the type description.
	expression := NewExpression(a.ASTBuilder)
	a.Expression = expression.Parse(unit, contractNode, fnNode, bodyNode, nil, nil, ctx)
	a.TypeDescription = a.Expression.GetTypeDescription()
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
	// Setting the type and source location information.
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

	// Parsing the operator.
	a.Operator = parseOperator(ctx.AssignOp())

	// Parsing left and right expressions.
	expression := NewExpression(a.ASTBuilder)
	a.LeftExpression = expression.Parse(unit, contractNode, fnNode, bodyNode, vDeclar, a, ctx.Expression(0))
	a.RightExpression = expression.Parse(unit, contractNode, fnNode, bodyNode, vDeclar, a, ctx.Expression(1))

	// Setting the type description based on the left expression.
	a.TypeDescription = a.LeftExpression.GetTypeDescription()

	// If the left expression is nil, set the reference descriptor for the right expression.
	if a.TypeDescription == nil {
		a.LeftExpression.SetReferenceDescriptor(a.RightExpression.GetId(), a.RightExpression.GetTypeDescription())
		a.TypeDescription = a.RightExpression.GetTypeDescription()
	}

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
