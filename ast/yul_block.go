package ast

import (
	"encoding/json"

	v3 "github.com/cncf/xds/go/xds/type/v3"
	ast_pb "github.com/unpackdev/protos/dist/go/ast"
	"github.com/unpackdev/solgo/parser"
)

// YulBlockStatement represents a YUL block statement in the abstract syntax tree.
type YulBlockStatement struct {
	*ASTBuilder

	// Id is the unique identifier of the YUL block statement.
	Id int64 `json:"id"`

	// NodeType is the type of the YUL block statement node.
	NodeType ast_pb.NodeType `json:"node_type"`

	// Src is the source location information of the YUL block statement.
	Src SrcNode `json:"src"`

	// Statements is a list of child nodes representing statements within the YUL block.
	Statements []Node[NodeType] `json:"statements"`
}

// NewYulBlockStatement creates a new YulBlockStatement instance.
func NewYulBlockStatement(b *ASTBuilder) *YulBlockStatement {
	return &YulBlockStatement{
		ASTBuilder: b,
		Id:         b.GetNextID(),
		NodeType:   ast_pb.NodeType_YUL_BLOCK,
		Statements: make([]Node[NodeType], 0),
	}
}

// SetReferenceDescriptor sets the reference descriptions of the YulBlockStatement node.
func (y *YulBlockStatement) SetReferenceDescriptor(refId int64, refDesc *TypeDescription) bool {
	return false
}

// GetId returns the ID of the YulBlockStatement.
func (y *YulBlockStatement) GetId() int64 {
	return y.Id
}

// GetType returns the NodeType of the YulBlockStatement.
func (y *YulBlockStatement) GetType() ast_pb.NodeType {
	return y.NodeType
}

// GetSrc returns the source information of the YulBlockStatement.
func (y *YulBlockStatement) GetSrc() SrcNode {
	return y.Src
}

// GetNodes returns the child nodes of the YulBlockStatement.
func (y *YulBlockStatement) GetNodes() []Node[NodeType] {
	toReturn := make([]Node[NodeType], 0)
	toReturn = append(toReturn, y.Statements...)
	return toReturn
}

// GetTypeDescription returns the type description of the YulBlockStatement.
func (y *YulBlockStatement) GetTypeDescription() *TypeDescription {
	return &TypeDescription{}
}

// GetStatements returns the list of statements within the YulBlockStatement.
func (y *YulBlockStatement) GetStatements() []Node[NodeType] {
	return y.Statements
}

// ToProto converts the YulBlockStatement to its protocol buffer representation.
func (y *YulBlockStatement) ToProto() NodeType {
	toReturn := ast_pb.YulBlockStatement{
		Id:       y.GetId(),
		NodeType: y.GetType(),
		Src:      y.GetSrc().ToProto(),
	}

	for _, statement := range y.GetStatements() {
		toReturn.Statements = append(
			toReturn.Statements,
			statement.ToProto().(*v3.TypedStruct),
		)
	}

	return NewTypedStruct(&toReturn, "YulBlockStatement")
}

// UnmarshalJSON unmarshals a given JSON byte array into a YulBlockStatement node.
func (f *YulBlockStatement) UnmarshalJSON(data []byte) error {
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

	if statements, ok := tempMap["statements"]; ok {
		var nodes []json.RawMessage
		if err := json.Unmarshal(statements, &nodes); err != nil {
			return err
		}

		for _, tempNode := range nodes {
			var tempNodeMap map[string]json.RawMessage
			if err := json.Unmarshal(tempNode, &tempNodeMap); err != nil {
				return err
			}

			var tempNodeType ast_pb.NodeType
			if err := json.Unmarshal(tempNodeMap["node_type"], &tempNodeType); err != nil {
				return err
			}

			node, err := unmarshalNode(tempNode, tempNodeType)
			if err != nil {
				return err
			}
			f.Statements = append(f.Statements, node)
		}
	}

	return nil
}

// Parse parses a YulBlockStatement node.
func (y *YulBlockStatement) Parse(
	unit *SourceUnit[Node[ast_pb.SourceUnit]],
	contractNode Node[NodeType],
	fnNode Node[NodeType],
	bodyNode *BodyNode,
	assemblyNode *Yul,
	statementNode *YulStatement,
	ifStatement *parser.YulIfStatementContext,
	parentNode Node[NodeType],
	ctx *parser.YulBlockContext,
) Node[NodeType] {
	y.Src = SrcNode{
		Id:          y.GetNextID(),
		Line:        int64(ctx.GetStart().GetLine()),
		Column:      int64(ctx.GetStart().GetColumn()),
		Start:       int64(ctx.GetStart().GetStart()),
		End:         int64(ctx.GetStop().GetStop()),
		Length:      int64(ctx.GetStop().GetStop() - ctx.GetStart().GetStart() + 1),
		ParentIndex: parentNode.GetId(),
	}

	if ctx.AllYulStatement() != nil {
		for _, statement := range ctx.AllYulStatement() {
			yStatement := NewYulStatement(y.ASTBuilder)

			y.Statements = append(y.Statements, yStatement.Parse(
				unit, contractNode, fnNode, bodyNode, assemblyNode, y, statement.(*parser.YulStatementContext),
			))
		}
	}

	return y
}
