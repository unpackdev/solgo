package ast

import (
	"github.com/goccy/go-json"

	ast_pb "github.com/unpackdev/protos/dist/go/ast"
	"github.com/unpackdev/solgo/parser"
)

// YulSwitchStatement represents a switch statement in Yul assembly.
type YulSwitchStatement struct {
	*ASTBuilder // Embedded ASTBuilder for utility functions.

	Id       int64            `json:"id"`        // Id is the unique identifier for the switch statement.
	NodeType ast_pb.NodeType  `json:"node_type"` // NodeType specifies the type of the node.
	Src      SrcNode          `json:"src"`       // Src provides source location details of the switch statement.
	Cases    []Node[NodeType] `json:"cases"`     // Cases holds the different cases of the switch statement.
}

// NewYulSwitchStatement creates and initializes a new YulSwitchStatement.
func NewYulSwitchStatement(b *ASTBuilder) *YulSwitchStatement {
	return &YulSwitchStatement{
		ASTBuilder: b,
		Id:         b.GetNextID(),
		NodeType:   ast_pb.NodeType_YUL_SWITCH,
		Cases:      make([]Node[NodeType], 0),
	}
}

// SetReferenceDescriptor is a placeholder method for setting reference descriptors.
// Currently always returns false.
func (y *YulSwitchStatement) SetReferenceDescriptor(refId int64, refDesc *TypeDescription) bool {
	return false
}

// GetId retrieves the unique identifier of the YulSwitchStatement.
func (y *YulSwitchStatement) GetId() int64 {
	return y.Id
}

// GetType returns the NodeType of the YulSwitchStatement.
func (y *YulSwitchStatement) GetType() ast_pb.NodeType {
	return y.NodeType
}

// GetSrc provides the source location details of the YulSwitchStatement.
func (y *YulSwitchStatement) GetSrc() SrcNode {
	return y.Src
}

// GetNodes returns a list of nodes associated with the YulSwitchStatement.
func (y *YulSwitchStatement) GetNodes() []Node[NodeType] {
	toReturn := make([]Node[NodeType], 0)
	toReturn = append(toReturn, y.Cases...)
	return toReturn
}

// GetTypeDescription provides a description of the YulSwitchStatement's type.
// Always returns an empty TypeDescription.
func (y *YulSwitchStatement) GetTypeDescription() *TypeDescription {
	return &TypeDescription{}
}

func (y *YulSwitchStatement) GetCases() []Node[NodeType] {
	return y.Cases
}

// ToProto converts the YulSwitchStatement into its protobuf representation.
// Note: This method currently returns an empty Statement.
func (y *YulSwitchStatement) ToProto() NodeType {
	toReturn := ast_pb.YulSwitchStatement{
		Id:       y.GetId(),
		NodeType: y.GetType(),
		Src:      y.GetSrc().ToProto(),
		Cases:    make([]*ast_pb.YulSwitchCaseStatement, 0),
	}

	for _, ycase := range y.GetCases() {
		toReturn.Cases = append(toReturn.Cases, ycase.ToProto().(*ast_pb.YulSwitchCaseStatement))
	}

	return NewTypedStruct(&toReturn, "YulSwitchStatement")
}

// UnmarshalJSON unmarshals a given JSON byte array into a YulSwitchStatement.
// Currently, this is a placeholder and does nothing.
func (f *YulSwitchStatement) UnmarshalJSON(data []byte) error {
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

	if cases, ok := tempMap["cases"]; ok {
		var nodes []json.RawMessage
		if err := json.Unmarshal(cases, &nodes); err != nil {
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
			f.Cases = append(f.Cases, node)
		}
	}

	return nil
}

// Parse populates the YulSwitchStatement fields based on the provided YulSwitchStatementContext.
func (y *YulSwitchStatement) Parse(
	unit *SourceUnit[Node[ast_pb.SourceUnit]],
	contractNode Node[NodeType],
	fnNode Node[NodeType],
	bodyNode *BodyNode,
	assemblyNode *Yul,
	statementNode *YulStatement,
	ctx *parser.YulSwitchStatementContext,
) Node[NodeType] {
	// Set the source location details from context.
	y.Src = SrcNode{
		Line:        int64(ctx.GetStart().GetLine()),
		Column:      int64(ctx.GetStart().GetColumn()),
		Start:       int64(ctx.GetStart().GetStart()),
		End:         int64(ctx.GetStop().GetStop()),
		Length:      int64(ctx.GetStop().GetStop() - ctx.GetStart().GetStart() + 1),
		ParentIndex: statementNode.GetId(),
	}

	// Parse all switch cases if present.
	if ctx.AllYulSwitchCase() != nil {
		for _, switchCase := range ctx.AllYulSwitchCase() {
			caseStatement := NewYulSwitchCaseStatement(y.ASTBuilder)
			y.Cases = append(y.Cases, caseStatement.Parse(
				unit, contractNode, fnNode, bodyNode, assemblyNode, statementNode, y,
				switchCase.(*parser.YulSwitchCaseContext),
			))
		}
	}

	return y
}
