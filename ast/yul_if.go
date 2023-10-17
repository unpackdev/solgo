package ast

import (
	"encoding/json"

	v3 "github.com/cncf/xds/go/xds/type/v3"
	ast_pb "github.com/unpackdev/protos/dist/go/ast"
	"github.com/unpackdev/solgo/parser"
)

type YulIfStatement struct {
	*ASTBuilder

	Id        int64           `json:"id"`
	NodeType  ast_pb.NodeType `json:"node_type"`
	Src       SrcNode         `json:"src"`
	Condition Node[NodeType]  `json:"condition"`
	Body      Node[NodeType]  `json:"body"`
}

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

func (y *YulIfStatement) GetId() int64 {
	return y.Id
}

func (y *YulIfStatement) GetType() ast_pb.NodeType {
	return y.NodeType
}

func (y *YulIfStatement) GetSrc() SrcNode {
	return y.Src
}

func (y *YulIfStatement) GetNodes() []Node[NodeType] {
	toReturn := make([]Node[NodeType], 0)
	toReturn = append(toReturn, y.Condition)
	toReturn = append(toReturn, y.Body)
	return toReturn
}

func (y *YulIfStatement) GetTypeDescription() *TypeDescription {
	return &TypeDescription{}
}

func (y *YulIfStatement) GetCondition() Node[NodeType] {
	return y.Condition
}

func (y *YulIfStatement) GetBody() Node[NodeType] {
	return y.Body
}

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
		Id:          y.GetNextID(),
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
