package ast

import (
	"github.com/goccy/go-json"

	v3 "github.com/cncf/xds/go/xds/type/v3"
	ast_pb "github.com/unpackdev/protos/dist/go/ast"
	"github.com/unpackdev/solgo/parser"
)

// YulAssignment represents a Yul assignment structure in the AST.
type YulAssignment struct {
	*ASTBuilder

	Id            int64            `json:"id"`
	NodeType      ast_pb.NodeType  `json:"nodeType"`
	Src           SrcNode          `json:"src"`
	VariableNames []*YulIdentifier `json:"variableNames"`
	Value         Node[NodeType]   `json:"value"`
}

// NewYulAssignment initializes a new instance of the YulAssignment structure.
func NewYulAssignment(b *ASTBuilder) *YulAssignment {
	return &YulAssignment{
		ASTBuilder:    b,
		Id:            b.GetNextID(),
		NodeType:      ast_pb.NodeType_YUL_ASSIGNMENT,
		VariableNames: make([]*YulIdentifier, 0),
	}
}

// SetReferenceDescriptor sets the reference descriptions of the YulAssignment node.
// It currently always returns false, and is a placeholder for future extensions.
func (y *YulAssignment) SetReferenceDescriptor(refId int64, refDesc *TypeDescription) bool {
	return false
}

// GetId retrieves the ID of the YulAssignment node.
func (y *YulAssignment) GetId() int64 {
	return y.Id
}

// GetType retrieves the NodeType of the YulAssignment node.
func (y *YulAssignment) GetType() ast_pb.NodeType {
	return y.NodeType
}

// GetSrc retrieves the source of the YulAssignment node.
func (y *YulAssignment) GetSrc() SrcNode {
	return y.Src
}

// GetNodes retrieves the child nodes of the YulAssignment node.
func (y *YulAssignment) GetNodes() []Node[NodeType] {
	toReturn := make([]Node[NodeType], 0)
	toReturn = append(toReturn, y.Value)

	for _, variable := range y.VariableNames {
		toReturn = append(toReturn, variable)
	}

	return toReturn
}

// GetTypeDescription retrieves the type description of the YulAssignment node.
// It currently returns an empty TypeDescription, and is a placeholder for future extensions.
func (y *YulAssignment) GetTypeDescription() *TypeDescription {
	return &TypeDescription{}
}

// GetVariableNames retrieves the variable names associated with the YulAssignment node.
func (y *YulAssignment) GetVariableNames() []*YulIdentifier {
	return y.VariableNames
}

// GetValue retrieves the value assigned in the YulAssignment node.
func (y *YulAssignment) GetValue() Node[NodeType] {
	return y.Value
}

// ToProto converts the YulAssignment node to its corresponding protocol buffer representation.
func (y *YulAssignment) ToProto() NodeType {
	toReturn := ast_pb.YulAssignmentStatement{
		Id:            y.GetId(),
		NodeType:      y.GetType(),
		Src:           y.GetSrc().ToProto(),
		VariableNames: make([]*v3.TypedStruct, 0),
		Value:         y.GetValue().ToProto().(*v3.TypedStruct),
	}

	for _, ycase := range y.GetVariableNames() {
		toReturn.VariableNames = append(toReturn.VariableNames, ycase.ToProto().(*v3.TypedStruct))
	}

	return NewTypedStruct(&toReturn, "YulAssignmentStatement")
}

// UnmarshalJSON unmarshals a given JSON byte array into a YulAssignment node.
func (f *YulAssignment) UnmarshalJSON(data []byte) error {
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

	if arguments, ok := tempMap["variableNames"]; ok {
		var nodes []json.RawMessage
		if err := json.Unmarshal(arguments, &nodes); err != nil {
			return err
		}

		for _, tempNode := range nodes {
			var node *YulIdentifier
			if err := json.Unmarshal(tempNode, &node); err != nil {
				return err
			}

			f.VariableNames = append(f.VariableNames, node)
		}
	}

	if value, ok := tempMap["value"]; ok {
		if err := json.Unmarshal(value, &f.Value); err != nil {
			var tempNodeMap map[string]json.RawMessage
			if err := json.Unmarshal(value, &tempNodeMap); err != nil {
				return err
			}

			var tempNodeType ast_pb.NodeType
			if err := json.Unmarshal(tempNodeMap["nodeType"], &tempNodeType); err != nil {
				return err
			}

			node, err := unmarshalNode(value, tempNodeType)
			if err != nil {
				return err
			}
			f.Value = node
		}
	}

	return nil
}

// Parse parses a given YulAssignmentContext into the YulAssignment node.
func (y *YulAssignment) Parse(
	unit *SourceUnit[Node[ast_pb.SourceUnit]],
	contractNode Node[NodeType],
	fnNode Node[NodeType],
	bodyNode *BodyNode,
	assemblyNode *Yul,
	statementNode *YulStatement,
	ctx *parser.YulAssignmentContext,
) Node[NodeType] {
	y.Src = SrcNode{
		Line:        int64(ctx.GetStart().GetLine()),
		Column:      int64(ctx.GetStart().GetColumn()),
		Start:       int64(ctx.GetStart().GetStart()),
		End:         int64(ctx.GetStop().GetStop()),
		Length:      int64(ctx.GetStop().GetStop() - ctx.GetStart().GetStart() + 1),
		ParentIndex: assemblyNode.GetId(),
	}

	if ctx.AllYulPath() != nil {
		for _, path := range ctx.AllYulPath() {
			for _, identifier := range path.AllYulIdentifier() {
				y.VariableNames = append(y.VariableNames, &YulIdentifier{
					Id:       y.GetNextID(),
					NodeType: ast_pb.NodeType_YUL_IDENTIFIER,
					Src: SrcNode{
						Line:        int64(identifier.GetSymbol().GetLine()),
						Column:      int64(identifier.GetSymbol().GetColumn()),
						Start:       int64(identifier.GetSymbol().GetStart()),
						End:         int64(identifier.GetSymbol().GetStop()),
						Length:      int64(identifier.GetSymbol().GetStop() - identifier.GetSymbol().GetStart() + 1),
						ParentIndex: y.GetId(),
					},
					Name: identifier.GetText(),
				})
			}
		}
	}

	if ctx.YulExpression() != nil {
		yExpression := NewYulExpressionStatement(y.ASTBuilder)
		y.Value = yExpression.Parse(
			unit, contractNode, fnNode, bodyNode, assemblyNode, statementNode,
			y, ctx.YulExpression().(*parser.YulExpressionContext),
		)
	}

	return y
}
