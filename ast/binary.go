package ast

import (
	"encoding/json"

	v3 "github.com/cncf/xds/go/xds/type/v3"
	ast_pb "github.com/unpackdev/protos/dist/go/ast"
	"github.com/unpackdev/solgo/parser"
)

// BinaryOperation represents a binary operation in a Solidity source file.
// A binary operation is an operation that operates on two operands like +, -, *, / etc.
type BinaryOperation struct {
	*ASTBuilder // Embedding the ASTBuilder to provide common functionality across all AST nodes.

	// Id is the unique identifier of the binary operation.
	Id int64 `json:"id"`
	// IsConstant indicates whether the binary operation is a constant.
	Constant bool `json:"is_constant"`
	// IsPure indicates whether the binary operation is pure (i.e., it does not read or modify state).
	Pure bool `json:"is_pure"`
	// NodeType is the type of the node.
	// For a BinaryOperation, this is always NodeType_BINARY_OPERATION.
	NodeType ast_pb.NodeType `json:"node_type"`
	// Src contains source information about the node, such as its line and column numbers in the source file.
	Src SrcNode `json:"src"`
	// Operator is the operator of the binary operation.
	Operator ast_pb.Operator `json:"operator"`
	// LeftExpression is the left operand of the binary operation.
	LeftExpression Node[NodeType] `json:"left_expression"`
	// RightExpression is the right operand of the binary operation.
	RightExpression Node[NodeType] `json:"right_expression"`
	// TypeDescription is the type description of the binary operation.
	TypeDescription *TypeDescription `json:"type_description"`
}

// NewBinaryOperationExpression is a constructor function that initializes a new BinaryOperation with a unique ID and the NodeType set to NodeType_BINARY_OPERATION.
func NewBinaryOperationExpression(b *ASTBuilder) *BinaryOperation {
	return &BinaryOperation{
		ASTBuilder: b,
		Id:         b.GetNextID(),
		NodeType:   ast_pb.NodeType_BINARY_OPERATION,
	}
}

// SetReferenceDescriptor sets the reference descriptions of the BinaryOperation node.
func (a *BinaryOperation) SetReferenceDescriptor(refId int64, refDesc *TypeDescription) bool {
	return false
}

// GetId is a getter method that returns the unique identifier of the binary operation.
func (a *BinaryOperation) GetId() int64 {
	return a.Id
}

// GetType is a getter method that returns the node type of the binary operation.
func (a *BinaryOperation) GetType() ast_pb.NodeType {
	return a.NodeType
}

// GetSrc is a getter method that returns the source information of the binary operation.
func (a *BinaryOperation) GetSrc() SrcNode {
	return a.Src
}

// GetOperator is a getter method that returns the operator of the binary operation.
func (a *BinaryOperation) GetOperator() ast_pb.Operator {
	return a.Operator
}

// GetLeftExpression is a getter method that returns the left operand of the binary operation.
func (a *BinaryOperation) GetLeftExpression() Node[NodeType] {
	return a.LeftExpression
}

// GetRightExpression is a getter method that returns the right operand of the binary operation.
func (a *BinaryOperation) GetRightExpression() Node[NodeType] {
	return a.RightExpression
}

// GetTypeDescription is a getter method that returns the type description of the left operand of the binary operation.
func (a *BinaryOperation) GetTypeDescription() *TypeDescription {
	if a.TypeDescription == nil {
		a.TypeDescription = a.LeftExpression.GetTypeDescription()
	}

	return a.TypeDescription
}

// GetNodes is a getter method that returns a slice of the operands of the binary operation.
func (a *BinaryOperation) GetNodes() []Node[NodeType] {
	return []Node[NodeType]{a.LeftExpression, a.RightExpression}
}

// IsConstant is a getter method that returns whether the binary operation is a constant.
func (a *BinaryOperation) IsConstant() bool {
	return a.Constant
}

// IsPure is a getter method that returns whether the binary operation is pure.
func (a *BinaryOperation) IsPure() bool {
	return a.Pure
}

// UnmarshalJSON sets the BinaryOperation node data from the JSON byte array.
func (a *BinaryOperation) UnmarshalJSON(data []byte) error {
	var tempMap map[string]json.RawMessage
	if err := json.Unmarshal(data, &tempMap); err != nil {
		return err
	}

	if id, ok := tempMap["id"]; ok {
		if err := json.Unmarshal(id, &a.Id); err != nil {
			return err
		}
	}

	if isConstant, ok := tempMap["is_constant"]; ok {
		if err := json.Unmarshal(isConstant, &a.Constant); err != nil {
			return err
		}
	}

	if isPure, ok := tempMap["is_pure"]; ok {
		if err := json.Unmarshal(isPure, &a.Pure); err != nil {
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

	return nil
}

// ToProto is a method that returns the protobuf representation of the binary operation.
func (a *BinaryOperation) ToProto() NodeType {
	proto := ast_pb.BinaryOperation{
		Id:              a.GetId(),
		IsConstant:      a.IsConstant(),
		IsPure:          a.IsPure(),
		NodeType:        a.GetType(),
		Src:             a.GetSrc().ToProto(),
		Operator:        a.GetOperator(),
		LeftExpression:  a.GetLeftExpression().ToProto().(*v3.TypedStruct),
		RightExpression: a.GetRightExpression().ToProto().(*v3.TypedStruct),
		TypeDescription: a.GetTypeDescription().ToProto(),
	}

	return NewTypedStruct(&proto, "BinaryOperation")
}

// ParseAddSub is a method that parses addition and subtraction operations.
func (a *BinaryOperation) ParseAddSub(
	unit *SourceUnit[Node[ast_pb.SourceUnit]],
	contractNode Node[NodeType],
	fnNode Node[NodeType],
	bodyNode *BodyNode,
	vDeclar *VariableDeclaration,
	expNode Node[NodeType],
	ctx *parser.AddSubOperationContext,
) Node[NodeType] {
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

	a.Operator = ast_pb.Operator_ADDITION
	if ctx.Sub() != nil {
		a.Operator = ast_pb.Operator_SUBTRACTION
	}

	expression := NewExpression(a.ASTBuilder)
	a.LeftExpression = expression.Parse(
		unit, contractNode, fnNode, bodyNode, vDeclar, a, ctx.Expression(0),
	)

	a.RightExpression = expression.Parse(
		unit, contractNode, fnNode, bodyNode, vDeclar, a, ctx.Expression(1),
	)

	a.TypeDescription = a.LeftExpression.GetTypeDescription()

	return a
}

// ParseOrderComparison is a method that parses order comparison operations.
func (a *BinaryOperation) ParseOrderComparison(
	unit *SourceUnit[Node[ast_pb.SourceUnit]],
	contractNode Node[NodeType],
	fnNode Node[NodeType],
	bodyNode *BodyNode,
	vDeclar *VariableDeclaration,
	expNode Node[NodeType],
	ctx *parser.OrderComparisonContext,
) Node[NodeType] {
	a.Src = SrcNode{
		Id:     a.GetNextID(),
		Line:   int64(ctx.GetStart().GetLine()),
		Column: int64(ctx.GetStart().GetColumn()),
		Start:  int64(ctx.GetStart().GetStart()),
		End:    int64(ctx.GetStop().GetStop()),
		Length: int64(ctx.GetStop().GetStop() - ctx.GetStart().GetStart() + 1),
		ParentIndex: func() int64 {
			if vDeclar != nil {
				return vDeclar.GetId()
			}

			if expNode != nil {
				return expNode.GetId()
			}

			return bodyNode.GetId()
		}(),
	}

	if ctx.GreaterThanOrEqual() != nil {
		a.Operator = ast_pb.Operator_GREATER_THAN_OR_EQUAL
	} else if ctx.LessThanOrEqual() != nil {
		a.Operator = ast_pb.Operator_LESS_THAN_OR_EQUAL
	} else if ctx.GreaterThan() != nil {
		a.Operator = ast_pb.Operator_GREATER_THAN
	} else if ctx.LessThan() != nil {
		a.Operator = ast_pb.Operator_LESS_THAN
	}

	expression := NewExpression(a.ASTBuilder)
	a.LeftExpression = expression.Parse(
		unit, contractNode, fnNode, bodyNode, vDeclar, a, ctx.Expression(0),
	)

	a.RightExpression = expression.Parse(
		unit, contractNode, fnNode, bodyNode, vDeclar, a, ctx.Expression(1),
	)

	a.TypeDescription = &TypeDescription{
		TypeIdentifier: "t_bool",
		TypeString:     "bool",
	}

	return a
}

// ParseMulDivMod is a method that parses multiplication, division, and modulo operations.
func (a *BinaryOperation) ParseMulDivMod(
	unit *SourceUnit[Node[ast_pb.SourceUnit]],
	contractNode Node[NodeType],
	fnNode Node[NodeType],
	bodyNode *BodyNode,
	vDeclar *VariableDeclaration,
	expNode Node[NodeType],
	ctx *parser.MulDivModOperationContext,
) Node[NodeType] {
	a.Src = SrcNode{
		Id:     a.GetNextID(),
		Line:   int64(ctx.GetStart().GetLine()),
		Column: int64(ctx.GetStart().GetColumn()),
		Start:  int64(ctx.GetStart().GetStart()),
		End:    int64(ctx.GetStop().GetStop()),
		Length: int64(ctx.GetStop().GetStop() - ctx.GetStart().GetStart() + 1),
		ParentIndex: func() int64 {
			if vDeclar != nil {
				return vDeclar.GetId()
			}

			if expNode != nil {
				return expNode.GetId()
			}

			return bodyNode.GetId()
		}(),
	}

	if ctx.Mul() != nil {
		a.Operator = ast_pb.Operator_MULTIPLICATION
	} else if ctx.Div() != nil {
		a.Operator = ast_pb.Operator_DIVISION
	} else if ctx.Mod() != nil {
		a.Operator = ast_pb.Operator_MODULO
	}

	expression := NewExpression(a.ASTBuilder)
	a.LeftExpression = expression.Parse(
		unit, contractNode, fnNode, bodyNode, vDeclar, a, ctx.Expression(0),
	)

	a.RightExpression = expression.Parse(
		unit, contractNode, fnNode, bodyNode, vDeclar, a, ctx.Expression(1),
	)

	if a.RightExpression.GetTypeDescription() == nil {
		a.RightExpression.SetReferenceDescriptor(
			a.LeftExpression.GetId(),
			a.LeftExpression.GetTypeDescription(),
		)
	}

	a.TypeDescription = a.LeftExpression.GetTypeDescription()

	return a
}

// ParseEqualityComparison is a method that parses equality comparison operations.
func (a *BinaryOperation) ParseEqualityComparison(
	unit *SourceUnit[Node[ast_pb.SourceUnit]],
	contractNode Node[NodeType],
	fnNode Node[NodeType],
	bodyNode *BodyNode,
	vDeclar *VariableDeclaration,
	expNode Node[NodeType],
	ctx *parser.EqualityComparisonContext,
) Node[NodeType] {
	a.Src = SrcNode{
		Id:     a.GetNextID(),
		Line:   int64(ctx.GetStart().GetLine()),
		Column: int64(ctx.GetStart().GetColumn()),
		Start:  int64(ctx.GetStart().GetStart()),
		End:    int64(ctx.GetStop().GetStop()),
		Length: int64(ctx.GetStop().GetStop() - ctx.GetStart().GetStart() + 1),
		ParentIndex: func() int64 {
			if vDeclar != nil {
				return vDeclar.GetId()
			}

			if expNode != nil {
				return expNode.GetId()
			}

			return bodyNode.GetId()
		}(),
	}

	if ctx.Equal() != nil {
		a.Operator = ast_pb.Operator_EQUAL
	} else if ctx.NotEqual() != nil {
		a.Operator = ast_pb.Operator_NOT_EQUAL
	}

	expression := NewExpression(a.ASTBuilder)
	a.LeftExpression = expression.Parse(
		unit, contractNode, fnNode, bodyNode, vDeclar, a, ctx.Expression(0),
	)

	a.RightExpression = expression.Parse(
		unit, contractNode, fnNode, bodyNode, vDeclar, a, ctx.Expression(1),
	)

	a.TypeDescription = &TypeDescription{
		TypeIdentifier: "t_bool",
		TypeString:     "bool",
	}

	return a
}

// ParseOr is a method that parses or comparison operations.
func (a *BinaryOperation) ParseOr(
	unit *SourceUnit[Node[ast_pb.SourceUnit]],
	contractNode Node[NodeType],
	fnNode Node[NodeType],
	bodyNode *BodyNode,
	vDeclar *VariableDeclaration,
	expNode Node[NodeType],
	ctx *parser.OrOperationContext,
) Node[NodeType] {
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

	a.Operator = ast_pb.Operator_OR

	expression := NewExpression(a.ASTBuilder)
	a.LeftExpression = expression.Parse(
		unit, contractNode, fnNode, bodyNode, vDeclar, a, ctx.Expression(0),
	)

	a.RightExpression = expression.Parse(
		unit, contractNode, fnNode, bodyNode, vDeclar, a, ctx.Expression(1),
	)

	a.TypeDescription = a.LeftExpression.GetTypeDescription()

	return a
}
