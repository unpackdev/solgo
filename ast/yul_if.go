package ast

import (
	"encoding/json"

	v3 "github.com/cncf/xds/go/xds/type/v3"
	ast_pb "github.com/unpackdev/protos/dist/go/ast"
	"github.com/unpackdev/solgo/parser"
)

// YulIfStatement represents an if statement in the abstract syntax tree.
type YulIfStatement struct {
	*ASTBuilder

	// Id is the unique identifier of the YUL if statement.
	Id int64 `json:"id"`

	// NodeType is the type of the YUL if statement node.
	NodeType ast_pb.NodeType `json:"node_type"`

	// Src is the source location information of the YUL if statement.
	Src SrcNode `json:"src"`

	// Condition is the condition expression of the if statement.
	Condition Node[NodeType] `json:"condition"`

	// Body is the body of the if statement.
	Body Node[NodeType] `json:"body"`
}

// NewYulIfStatement creates a new YulIfStatement with the provided AST builder.
func NewYulIfStatement(b *ASTBuilder) *YulIfStatement {
	return &YulIfStatement{
		ASTBuilder: b,
		Id:         b.GetNextID(),
		NodeType:   ast_pb.NodeType_YUL_IF,
	}
}

// SetReferenceDescriptor sets the reference descriptions of the YulIfStatement node.
func (y *YulIfStatement) SetReferenceDescriptor(refId int64, refDesc *TypeDescription) bool {
	return false
}

// GetId returns the ID of the YulIfStatement.
func (y *YulIfStatement) GetId() int64 {
	return y.Id
}

// GetType returns the NodeType of the YulIfStatement.
func (y *YulIfStatement) GetType() ast_pb.NodeType {
	return y.NodeType
}

// GetSrc returns the source information of the YulIfStatement.
func (y *YulIfStatement) GetSrc() SrcNode {
	return y.Src
}

// GetNodes returns a list containing the Condition and Body nodes.
func (y *YulIfStatement) GetNodes() []Node[NodeType] {
	toReturn := make([]Node[NodeType], 0)
	toReturn = append(toReturn, y.Condition)
	toReturn = append(toReturn, y.Body)
	return toReturn
}

// GetTypeDescription returns the type description of the YulIfStatement.
func (y *YulIfStatement) GetTypeDescription() *TypeDescription {
	return &TypeDescription{}
}

// GetCondition returns the condition expression of the if statement.
func (y *YulIfStatement) GetCondition() Node[NodeType] {
	return y.Condition
}

// GetBody returns the body of the if statement.
func (y *YulIfStatement) GetBody() Node[NodeType] {
	return y.Body
}

// ToProto converts the YulIfStatement to its protocol buffer representation.
func (y *YulIfStatement) ToProto() NodeType {
	toReturn := ast_pb.YulIfStatement{
		Id:       y.GetId(),
		NodeType: y.GetType(),
		Src:      y.GetSrc().ToProto(),
	}

	if y.GetCondition() != nil {
		toReturn.Condition = y.GetCondition().ToProto().(*v3.TypedStruct)
	}

	if y.GetBody() != nil {
		toReturn.Condition = y.GetBody().ToProto().(*v3.TypedStruct)
	}

	return NewTypedStruct(&toReturn, "YulIfStatement")
}

// UnmarshalJSON unmarshals a given JSON byte array into a YulIfStatement node.
func (f *YulIfStatement) UnmarshalJSON(data []byte) error {
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

	if condition, ok := tempMap["condition"]; ok {
		if err := json.Unmarshal(condition, &f.Condition); err != nil {
			var tempNodeMap map[string]json.RawMessage
			if err := json.Unmarshal(condition, &tempNodeMap); err != nil {
				return err
			}

			var tempNodeType ast_pb.NodeType
			if err := json.Unmarshal(tempNodeMap["node_type"], &tempNodeType); err != nil {
				return err
			}

			node, err := unmarshalNode(condition, tempNodeType)
			if err != nil {
				return err
			}
			f.Condition = node
		}
	}

	if body, ok := tempMap["body"]; ok {
		if err := json.Unmarshal(body, &f.Body); err != nil {
			var tempNodeMap map[string]json.RawMessage
			if err := json.Unmarshal(body, &tempNodeMap); err != nil {
				return err
			}

			var tempNodeType ast_pb.NodeType
			if err := json.Unmarshal(tempNodeMap["node_type"], &tempNodeType); err != nil {
				return err
			}

			node, err := unmarshalNode(body, tempNodeType)
			if err != nil {
				return err
			}
			f.Body = node
		}
	}

	return nil
}

// Parse converts a YulIfStatementContext to a YulIfStatement node.
func (y *YulIfStatement) Parse(
	unit *SourceUnit[Node[ast_pb.SourceUnit]],
	contractNode Node[NodeType],
	fnNode Node[NodeType],
	bodyNode *BodyNode,
	assemblyNode *Yul,
	statementNode *YulStatement,
	ctx *parser.YulIfStatementContext,
) Node[NodeType] {
	y.Src = SrcNode{
		Line:        int64(ctx.GetStart().GetLine()),
		Column:      int64(ctx.GetStart().GetColumn()),
		Start:       int64(ctx.GetStart().GetStart()),
		End:         int64(ctx.GetStop().GetStop()),
		Length:      int64(ctx.GetStop().GetStop() - ctx.GetStart().GetStart() + 1),
		ParentIndex: assemblyNode.GetId(),
	}

	if ctx.GetCond() != nil {
		y.Condition = ParseYulExpression(
			y.ASTBuilder, unit, contractNode, fnNode, bodyNode, assemblyNode, statementNode, nil, nil,
			y, ctx.GetCond(),
		)
	}

	if ctx.GetBody() != nil {
		blockStatement := NewYulBlockStatement(y.ASTBuilder)
		y.Body = blockStatement.Parse(
			unit, contractNode, fnNode, bodyNode, assemblyNode, statementNode, ctx, y,
			ctx.GetBody().(*parser.YulBlockContext),
		)
	}

	return y
}
