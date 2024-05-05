package ast

import (
	v3 "github.com/cncf/xds/go/xds/type/v3"
	"github.com/goccy/go-json"
	ast_pb "github.com/unpackdev/protos/dist/go/ast"
	"github.com/unpackdev/solgo/parser"
)

// YulForStatement represents a YUL for statement in the abstract syntax tree.
type YulForStatement struct {
	*ASTBuilder

	// Id is the unique identifier of the YUL for statement.
	Id int64 `json:"id"`

	// NodeType is the type of the YUL for statement node.
	NodeType ast_pb.NodeType `json:"nodeType"`

	// Src is the source location information of the YUL for statement.
	Src SrcNode `json:"src"`

	// Pre is the initialization node before the loop.
	Pre Node[NodeType] `json:"pre"`

	// Post is the post-execution node after each iteration of the loop.
	Post Node[NodeType] `json:"post"`

	// Condition is the condition node of the loop.
	Condition Node[NodeType] `json:"condition"`

	// Body is the body of the loop.
	Body Node[NodeType] `json:"body"`
}

// NewYulForStatement creates a new YulForStatement instance.
func NewYulForStatement(b *ASTBuilder) *YulForStatement {
	return &YulForStatement{
		ASTBuilder: b,
		Id:         b.GetNextID(),
		NodeType:   ast_pb.NodeType_YUL_FOR,
	}
}

// SetReferenceDescriptor sets the reference descriptions of the YulForStatement node.
func (y *YulForStatement) SetReferenceDescriptor(refId int64, refDesc *TypeDescription) bool {
	return false
}

// GetId returns the ID of the YulForStatement.
func (y *YulForStatement) GetId() int64 {
	return y.Id
}

// GetType returns the NodeType of the YulForStatement.
func (y *YulForStatement) GetType() ast_pb.NodeType {
	return y.NodeType
}

// GetSrc returns the source information of the YulForStatement.
func (y *YulForStatement) GetSrc() SrcNode {
	return y.Src
}

// GetNodes returns a list containing the condition and body nodes.
func (y *YulForStatement) GetNodes() []Node[NodeType] {
	toReturn := make([]Node[NodeType], 0)
	toReturn = append(toReturn, y.Condition)
	toReturn = append(toReturn, y.Body)
	toReturn = append(toReturn, y.Pre)
	toReturn = append(toReturn, y.Post)
	return toReturn
}

// GetTypeDescription returns the type description of the YulForStatement.
func (y *YulForStatement) GetTypeDescription() *TypeDescription {
	return &TypeDescription{}
}

// GetPre returns the initialization node before the loop.
func (y *YulForStatement) GetPre() Node[NodeType] {
	return y.Pre
}

// GetPost returns the post-execution node after each iteration of the loop.
func (y *YulForStatement) GetPost() Node[NodeType] {
	return y.Post
}

// GetCondition returns the condition node of the loop.
func (y *YulForStatement) GetCondition() Node[NodeType] {
	return y.Condition
}

// GetBody returns the body of the loop.
func (y *YulForStatement) GetBody() Node[NodeType] {
	return y.Body
}

// ToProto converts the YulForStatement to its protocol buffer representation.
func (y *YulForStatement) ToProto() NodeType {
	toReturn := ast_pb.YulForStatement{
		Id:       y.GetId(),
		NodeType: y.GetType(),
		Src:      y.GetSrc().ToProto(),
	}

	if y.GetPre() != nil {
		toReturn.Pre = y.GetPre().ToProto().(*v3.TypedStruct)
	}

	if y.GetPost() != nil {
		toReturn.Post = y.GetPost().ToProto().(*v3.TypedStruct)
	}

	if y.GetCondition() != nil {
		toReturn.Condition = y.GetCondition().ToProto().(*v3.TypedStruct)
	}

	if y.GetBody() != nil {
		toReturn.Body = y.GetBody().ToProto().(*v3.TypedStruct)
	}

	return NewTypedStruct(&toReturn, "YulForStatement")
}

// UnmarshalJSON unmarshals a given JSON byte array into a YulForStatement node.
func (f *YulForStatement) UnmarshalJSON(data []byte) error {
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

	if pre, ok := tempMap["pre"]; ok {
		if err := json.Unmarshal(pre, &f.Pre); err != nil {
			var tempNodeMap map[string]json.RawMessage
			if err := json.Unmarshal(pre, &tempNodeMap); err != nil {
				return err
			}

			var tempNodeType ast_pb.NodeType
			if err := json.Unmarshal(tempNodeMap["nodeType"], &tempNodeType); err != nil {
				return err
			}

			node, err := unmarshalNode(pre, tempNodeType)
			if err != nil {
				return err
			}
			f.Pre = node
		}
	}

	if post, ok := tempMap["post"]; ok {
		if err := json.Unmarshal(post, &f.Post); err != nil {
			var tempNodeMap map[string]json.RawMessage
			if err := json.Unmarshal(post, &tempNodeMap); err != nil {
				return err
			}

			var tempNodeType ast_pb.NodeType
			if err := json.Unmarshal(tempNodeMap["nodeType"], &tempNodeType); err != nil {
				return err
			}

			node, err := unmarshalNode(post, tempNodeType)
			if err != nil {
				return err
			}
			f.Post = node
		}
	}

	if condition, ok := tempMap["condition"]; ok {
		if err := json.Unmarshal(condition, &f.Condition); err != nil {
			var tempNodeMap map[string]json.RawMessage
			if err := json.Unmarshal(condition, &tempNodeMap); err != nil {
				return err
			}

			var tempNodeType ast_pb.NodeType
			if err := json.Unmarshal(tempNodeMap["nodeType"], &tempNodeType); err != nil {
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
			if err := json.Unmarshal(tempNodeMap["nodeType"], &tempNodeType); err != nil {
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

// Parse parses a YUL for statement.
func (y *YulForStatement) Parse(
	unit *SourceUnit[Node[ast_pb.SourceUnit]],
	contractNode Node[NodeType],
	fnNode Node[NodeType],
	bodyNode *BodyNode,
	assemblyNode *Yul,
	statementNode *YulStatement,
	ctx *parser.YulForStatementContext,
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

	if ctx.GetInit() != nil {
		blockStatement := NewYulBlockStatement(y.ASTBuilder)
		y.Pre = blockStatement.Parse(
			unit, contractNode, fnNode, bodyNode, assemblyNode, statementNode, nil, y,
			ctx.GetInit().(*parser.YulBlockContext),
		)
	}

	if ctx.GetPost() != nil {
		blockStatement := NewYulBlockStatement(y.ASTBuilder)
		y.Post = blockStatement.Parse(
			unit, contractNode, fnNode, bodyNode, assemblyNode, statementNode, nil, y,
			ctx.GetPost().(*parser.YulBlockContext),
		)
	}

	if ctx.GetBody() != nil {
		blockStatement := NewYulBlockStatement(y.ASTBuilder)
		y.Body = blockStatement.Parse(
			unit, contractNode, fnNode, bodyNode, assemblyNode, statementNode, nil, y,
			ctx.GetBody().(*parser.YulBlockContext),
		)
	}

	return y
}
