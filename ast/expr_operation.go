package ast

import (
	"encoding/json"

	v3 "github.com/cncf/xds/go/xds/type/v3"
	ast_pb "github.com/unpackdev/protos/dist/go/ast"
	"github.com/unpackdev/solgo/parser"
)

// ExprOperation represents an expression operation in the Solidity abstract syntax tree (AST).
type ExprOperation struct {
	*ASTBuilder

	Id               int64              `json:"id"`                // Unique identifier for the expression operation
	NodeType         ast_pb.NodeType    `json:"node_type"`         // Type of the node (EXPRESSION_OPERATION for expression operation)
	Src              SrcNode            `json:"src"`               // Source information about the expression operation
	LeftExpression   Node[NodeType]     `json:"left_expression"`   // Left expression in the operation
	RightExpression  Node[NodeType]     `json:"right_expression"`  // Right expression in the operation
	TypeDescriptions []*TypeDescription `json:"type_descriptions"` // Type descriptions of the expressions
}

// NewExprOperationExpression creates a new ExprOperation instance.
func NewExprOperationExpression(b *ASTBuilder) *ExprOperation {
	return &ExprOperation{
		ASTBuilder:       b,
		Id:               b.GetNextID(),
		NodeType:         ast_pb.NodeType_EXPRESSION_OPERATION,
		TypeDescriptions: make([]*TypeDescription, 0),
	}
}

// SetReferenceDescriptor sets the reference descriptions of the ExprOperation node.
func (b *ExprOperation) SetReferenceDescriptor(refId int64, refDesc *TypeDescription) bool {
	return false
}

// GetId returns the unique identifier of the expression operation.
func (f *ExprOperation) GetId() int64 {
	return f.Id
}

// GetType returns the type of the node, which is 'EXPRESSION_OPERATION' for an expression operation.
func (f *ExprOperation) GetType() ast_pb.NodeType {
	return f.NodeType
}

// GetSrc returns the source information about the expression operation.
func (f *ExprOperation) GetSrc() SrcNode {
	return f.Src
}

// GetTypeDescription returns the type description of the expression operation.
func (f *ExprOperation) GetTypeDescription() *TypeDescription {
	return f.TypeDescriptions[0]
}

// GetNodes returns the nodes representing the left and right expressions of the operation.
func (f *ExprOperation) GetNodes() []Node[NodeType] {
	return []Node[NodeType]{f.LeftExpression, f.RightExpression}
}

// GetLeftExpression returns the left expression in the operation.
func (f *ExprOperation) GetLeftExpression() Node[NodeType] {
	return f.LeftExpression
}

// GetRightExpression returns the right expression in the operation.
func (f *ExprOperation) GetRightExpression() Node[NodeType] {
	return f.RightExpression
}

func (f *ExprOperation) UnmarshalJSON(data []byte) error {
	var tempMap map[string]json.RawMessage
	if err := json.Unmarshal(data, &tempMap); err != nil {
		return err
	}

	if id, ok := tempMap["id"]; ok {
		if err := json.Unmarshal(id, &f.Id); err != nil {
			return err
		}
	}

	if nodeType, ok := tempMap["node_type"]; ok {
		if err := json.Unmarshal(nodeType, &f.NodeType); err != nil {
			return err
		}
	}

	if src, ok := tempMap["src"]; ok {
		if err := json.Unmarshal(src, &f.Src); err != nil {
			return err
		}
	}

	if leftExpression, ok := tempMap["left_expression"]; ok {
		if err := json.Unmarshal(leftExpression, &f.LeftExpression); err != nil {
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
			f.LeftExpression = node
		}
	}

	if rightExpression, ok := tempMap["right_expression"]; ok {
		if err := json.Unmarshal(rightExpression, &f.RightExpression); err != nil {
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
			f.RightExpression = node
		}
	}

	if typeDescriptions, ok := tempMap["type_descriptions"]; ok {
		if err := json.Unmarshal(typeDescriptions, &f.TypeDescriptions); err != nil {
			return err
		}
	}

	return nil

}

// ToProto returns the protobuf representation of the expression operation.
func (f *ExprOperation) ToProto() NodeType {
	proto := ast_pb.ExprOperation{
		Id:              f.GetId(),
		NodeType:        f.GetType(),
		Src:             f.GetSrc().ToProto(),
		LeftExpression:  f.GetLeftExpression().ToProto().(*v3.TypedStruct),
		RightExpression: f.GetRightExpression().ToProto().(*v3.TypedStruct),
		TypeDescription: f.GetTypeDescription().ToProto(),
	}

	return NewTypedStruct(&proto, "ExprOperation")
}

// Parse parses an expression operation from the provided parser.ExpOperationContext and updates the current instance.
func (f *ExprOperation) Parse(
	unit *SourceUnit[Node[ast_pb.SourceUnit]],
	contractNode Node[NodeType],
	fnNode Node[NodeType],
	bodyNode *BodyNode,
	vDeclar *VariableDeclaration,
	expNode Node[NodeType],
	ctx *parser.ExpOperationContext,
) Node[NodeType] {
	f.Id = f.GetNextID()
	f.Src = SrcNode{
		Id:     f.GetNextID(),
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

			if bodyNode != nil {
				return bodyNode.GetId()
			}

			if fnNode != nil {
				return fnNode.GetId()
			}

			return contractNode.GetId()
		}(),
	}

	expression := NewExpression(f.ASTBuilder)

	f.LeftExpression = expression.Parse(unit, contractNode, fnNode, bodyNode, vDeclar, f, ctx.Expression(0))
	f.RightExpression = expression.Parse(unit, contractNode, fnNode, bodyNode, vDeclar, f, ctx.Expression(1))

	f.TypeDescriptions = append(f.TypeDescriptions, f.LeftExpression.GetTypeDescription())
	f.TypeDescriptions = append(f.TypeDescriptions, f.RightExpression.GetTypeDescription())

	return f
}
