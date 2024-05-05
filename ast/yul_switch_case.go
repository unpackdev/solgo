package ast

import (
	"github.com/goccy/go-json"

	v3 "github.com/cncf/xds/go/xds/type/v3"
	ast_pb "github.com/unpackdev/protos/dist/go/ast"
	"github.com/unpackdev/solgo/parser"
)

// YulSwitchCaseStatement represents an individual case statement within a Yul switch structure.
type YulSwitchCaseStatement struct {
	*ASTBuilder // Embedded ASTBuilder for utility functions.

	Id       int64           `json:"id"`       // Id is the unique identifier for the switch case statement.
	NodeType ast_pb.NodeType `json:"nodeType"` // NodeType specifies the type of the node.
	Src      SrcNode         `json:"src"`      // Src provides source location details of the switch case statement.
	Case     Node[NodeType]  `json:"case"`     // Case holds the condition for the switch case.
	Body     Node[NodeType]  `json:"body"`     // Body represents the block of code to execute for this case.
}

// NewYulSwitchCaseStatement creates and initializes a new YulSwitchCaseStatement.
func NewYulSwitchCaseStatement(b *ASTBuilder) *YulSwitchCaseStatement {
	return &YulSwitchCaseStatement{
		ASTBuilder: b,
		Id:         b.GetNextID(),
		NodeType:   ast_pb.NodeType_YUL_SWITCH_CASE,
	}
}

// SetReferenceDescriptor is a placeholder method for setting reference descriptors.
// Currently always returns false.
func (y *YulSwitchCaseStatement) SetReferenceDescriptor(refId int64, refDesc *TypeDescription) bool {
	return false
}

// GetId retrieves the unique identifier of the YulSwitchCaseStatement.
func (y *YulSwitchCaseStatement) GetId() int64 {
	return y.Id
}

// GetType returns the NodeType of the YulSwitchCaseStatement.
func (y *YulSwitchCaseStatement) GetType() ast_pb.NodeType {
	return y.NodeType
}

// GetSrc provides the source location details of the YulSwitchCaseStatement.
func (y *YulSwitchCaseStatement) GetSrc() SrcNode {
	return y.Src
}

// GetNodes returns a list of nodes associated with the YulSwitchCaseStatement.
func (y *YulSwitchCaseStatement) GetNodes() []Node[NodeType] {
	toReturn := make([]Node[NodeType], 0)
	toReturn = append(toReturn, y.Case)
	toReturn = append(toReturn, y.Body)
	return toReturn
}

// GetTypeDescription provides a description of the YulSwitchCaseStatement's type.
// Always returns an empty TypeDescription.
func (y *YulSwitchCaseStatement) GetTypeDescription() *TypeDescription {
	return &TypeDescription{}
}

// GetCase returns case identifier
func (y *YulSwitchCaseStatement) GetCase() Node[NodeType] {
	return y.Case
}

// GetBody returns body of the associated nodes
func (y *YulSwitchCaseStatement) GetBody() Node[NodeType] {
	return y.Body
}

// ToProto converts the YulSwitchCaseStatement into its protobuf representation.
// Note: This method currently returns an empty Statement.
func (y *YulSwitchCaseStatement) ToProto() NodeType {
	toReturn := &ast_pb.YulSwitchCaseStatement{
		Id:       y.GetId(),
		NodeType: y.GetType(),
		Src:      y.GetSrc().ToProto(),
	}

	if y.GetCase() != nil {
		toReturn.Case = y.GetCase().ToProto().(*v3.TypedStruct)
	}

	if y.GetBody() != nil {
		toReturn.Body = y.GetBody().ToProto().(*v3.TypedStruct)
	}

	return toReturn
}

// UnmarshalJSON unmarshals a given JSON byte array into a YulSwitchCaseStatement.
// Currently, this is a placeholder and does nothing.
func (f *YulSwitchCaseStatement) UnmarshalJSON(data []byte) error {
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

	if fcase, ok := tempMap["case"]; ok {
		if err := json.Unmarshal(fcase, &f.Case); err != nil {
			var tempNodeMap map[string]json.RawMessage
			if err := json.Unmarshal(fcase, &tempNodeMap); err != nil {
				return err
			}

			var tempNodeType ast_pb.NodeType
			if err := json.Unmarshal(tempNodeMap["nodeType"], &tempNodeType); err != nil {
				return err
			}

			node, err := unmarshalNode(fcase, tempNodeType)
			if err != nil {
				return err
			}
			f.Case = node
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

// Parse populates the YulSwitchCaseStatement fields based on the provided YulSwitchCaseContext.
func (y *YulSwitchCaseStatement) Parse(
	unit *SourceUnit[Node[ast_pb.SourceUnit]],
	contractNode Node[NodeType],
	fnNode Node[NodeType],
	bodyNode *BodyNode,
	assemblyNode *Yul,
	statementNode *YulStatement,
	parentNode Node[NodeType],
	ctx *parser.YulSwitchCaseContext,
) Node[NodeType] {
	// Set the source location details from context.
	y.Src = SrcNode{
		Line:        int64(ctx.GetStart().GetLine()),
		Column:      int64(ctx.GetStart().GetColumn()),
		Start:       int64(ctx.GetStart().GetStart()),
		End:         int64(ctx.GetStop().GetStop()),
		Length:      int64(ctx.GetStop().GetStop() - ctx.GetStart().GetStart() + 1),
		ParentIndex: parentNode.GetId(),
	}

	// Parse the Yul literal if present.
	if ctx.YulLiteral() != nil {
		literalStatement := NewYulLiteralStatement(y.ASTBuilder)
		y.Case = literalStatement.Parse(
			unit, contractNode, fnNode, bodyNode, assemblyNode, statementNode, y,
			ctx.YulLiteral().(*parser.YulLiteralContext),
		)
	}

	// Parse the Yul block if present.
	if ctx.YulBlock() != nil {
		block := NewYulBlockStatement(y.ASTBuilder)
		y.Body = block.Parse(
			unit, contractNode, fnNode, bodyNode, assemblyNode, statementNode, nil, y,
			ctx.YulBlock().(*parser.YulBlockContext),
		)
	}

	return y
}
