package ast

import (
	"encoding/json"

	v3 "github.com/cncf/xds/go/xds/type/v3"
	ast_pb "github.com/unpackdev/protos/dist/go/ast"
	"github.com/unpackdev/solgo/parser"
)

// YulVariable represents a variable declaration in Yul assembly.
type YulVariable struct {
	*ASTBuilder // Embedded ASTBuilder for utility functions.

	Id        int64            `json:"id"`        // Id is a unique identifier for the variable.
	NodeType  ast_pb.NodeType  `json:"node_type"` // NodeType specifies the type of the node.
	Src       SrcNode          `json:"src"`       // Src contains the source location details of the variable.
	Let       bool             `json:"let"`       // Let indicates if the variable is declared with a "let" keyword.
	Value     Node[NodeType]   `json:"value"`     // Value holds the initialized value of the variable.
	Variables []*YulIdentifier `json:"variables"` // Variables contains a list of Yul identifiers.
}

// NewYulVariable initializes a new YulVariable with default values.
func NewYulVariable(b *ASTBuilder) *YulVariable {
	return &YulVariable{
		ASTBuilder: b,
		Id:         b.GetNextID(),
		NodeType:   ast_pb.NodeType_YUL_VARIABLE_DECLARATION,
		Variables:  make([]*YulIdentifier, 0),
	}
}

// SetReferenceDescriptor is a placeholder method for setting reference descriptors.
// Currently always returns false.
func (y *YulVariable) SetReferenceDescriptor(refId int64, refDesc *TypeDescription) bool {
	return false
}

// GetId retrieves the unique identifier of the YulVariable.
func (y *YulVariable) GetId() int64 {
	return y.Id
}

// GetType returns the NodeType of the YulVariable.
func (y *YulVariable) GetType() ast_pb.NodeType {
	return y.NodeType
}

// GetSrc provides the source location details of the YulVariable.
func (y *YulVariable) GetSrc() SrcNode {
	return y.Src
}

// GetNodes returns a list of nodes associated with the YulVariable, including its initialized value.
func (y *YulVariable) GetNodes() []Node[NodeType] {
	toReturn := make([]Node[NodeType], 0)
	toReturn = append(toReturn, y.Value)
	return toReturn
}

// GetTypeDescription provides a description of the YulVariable's type.
// Always returns an empty TypeDescription.
func (y *YulVariable) GetTypeDescription() *TypeDescription {
	return &TypeDescription{}
}

// IsLet returns whenever variable is using let
func (y *YulVariable) IsLet() bool {
	return y.Let
}

// GetValue returns variable value declaration
func (y *YulVariable) GetValue() Node[NodeType] {
	return y.Value
}

// GetVariables returns variable discovered variables
func (y *YulVariable) GetVariables() []*YulIdentifier {
	return y.Variables
}

// ToProto converts the YulVariable into its protobuf representation.
// Note: This method currently returns an empty Statement.
func (y *YulVariable) ToProto() NodeType {
	toReturn := ast_pb.YulVariableStatement{
		Id:       y.GetId(),
		NodeType: y.GetType(),
		Src:      y.GetSrc().ToProto(),
		Let:      y.IsLet(),
		Value:    y.GetValue().ToProto().(*v3.TypedStruct),
	}

	for _, variable := range y.GetVariables() {
		toReturn.Variables = append(
			toReturn.Variables,
			variable.ToProto().(*v3.TypedStruct),
		)
	}

	return NewTypedStruct(&toReturn, "YulVariableStatement")
}

// UnmarshalJSON unmarshals a given JSON byte array into a YulVariable.
func (y *YulVariable) UnmarshalJSON(data []byte) error {
	var tempMap map[string]json.RawMessage
	if err := json.Unmarshal(data, &tempMap); err != nil {
		return err
	}

	if id, ok := tempMap["id"]; ok {
		if err := json.Unmarshal(id, &y.Id); err != nil {
			return err
		}
	}

	if nodeType, ok := tempMap["node_type"]; ok {
		if err := json.Unmarshal(nodeType, &y.NodeType); err != nil {
			return err
		}
	}

	if src, ok := tempMap["src"]; ok {
		if err := json.Unmarshal(src, &y.Src); err != nil {
			return err
		}
	}

	if let, ok := tempMap["let"]; ok {
		if err := json.Unmarshal(let, &y.Let); err != nil {
			return err
		}
	}

	if variables, ok := tempMap["variables"]; ok {
		var nodes []json.RawMessage
		if err := json.Unmarshal(variables, &nodes); err != nil {
			return err
		}

		for _, tempNode := range nodes {
			var tempIdentifier *YulIdentifier
			if err := json.Unmarshal(tempNode, &tempIdentifier); err != nil {
				return err
			}

			y.Variables = append(y.Variables, tempIdentifier)
		}
	}

	if value, ok := tempMap["value"]; ok {
		if err := json.Unmarshal(value, &y.Value); err != nil {
			var tempNodeMap map[string]json.RawMessage
			if err := json.Unmarshal(value, &tempNodeMap); err != nil {
				return err
			}

			var tempNodeType ast_pb.NodeType
			if err := json.Unmarshal(tempNodeMap["node_type"], &tempNodeType); err != nil {
				return err
			}

			node, err := unmarshalNode(value, tempNodeType)
			if err != nil {
				return err
			}
			y.Value = node
		}
	}

	return nil
}

// Parse populates the YulVariable fields by parsing the provided YulVariableDeclarationContext.
func (y *YulVariable) Parse(
	unit *SourceUnit[Node[ast_pb.SourceUnit]],
	contractNode Node[NodeType],
	fnNode Node[NodeType],
	bodyNode *BodyNode,
	assemblyNode *Yul,
	statementNode *YulStatement,
	ctx *parser.YulVariableDeclarationContext,
) Node[NodeType] {
	// Set source location details from context.
	y.Src = SrcNode{
		Id:          y.GetNextID(),
		Line:        int64(ctx.GetStart().GetLine()),
		Column:      int64(ctx.GetStart().GetColumn()),
		Start:       int64(ctx.GetStart().GetStart()),
		End:         int64(ctx.GetStop().GetStop()),
		Length:      int64(ctx.GetStop().GetStop() - ctx.GetStart().GetStart() + 1),
		ParentIndex: statementNode.GetId(),
	}

	// Determine if "let" keyword is present.
	y.Let = ctx.YulLet() != nil

	// Parse declared variables.
	for _, variable := range ctx.GetVariables() {
		y.Variables = append(y.Variables, &YulIdentifier{
			Id:       y.GetNextID(),
			Name:     variable.GetText(),
			NodeType: ast_pb.NodeType_YUL_VARIABLE_NAME,
			Src: SrcNode{
				Id:          y.GetNextID(),
				Line:        int64(variable.GetLine()),
				Column:      int64(variable.GetColumn()),
				Start:       int64(variable.GetStart()),
				End:         int64(variable.GetStop()),
				Length:      int64(variable.GetStop() - variable.GetStart() + 1),
				ParentIndex: y.GetId(),
			},
		})
	}

	// Parse expression if present.
	if ctx.YulExpression() != nil {
		expression := ctx.YulExpression()
		yExpression := NewYulExpressionStatement(y.ASTBuilder)
		y.Value = yExpression.Parse(
			unit, contractNode, fnNode, bodyNode, assemblyNode, statementNode, y,
			expression.(*parser.YulExpressionContext),
		)
	}

	return y
}
