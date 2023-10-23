package ast

import (
	"encoding/json"

	v3 "github.com/cncf/xds/go/xds/type/v3"
	ast_pb "github.com/unpackdev/protos/dist/go/ast"
	"github.com/unpackdev/solgo/parser"
)

// YulFunctionDefinition represents a YUL function definition in the abstract syntax tree.
type YulFunctionDefinition struct {
	*ASTBuilder

	// Id is the unique identifier of the YUL function definition.
	Id int64 `json:"id"`

	// NodeType is the type of the YUL function definition node.
	NodeType ast_pb.NodeType `json:"node_type"`

	// Src is the source location information of the YUL function definition.
	Src SrcNode `json:"src"`

	// Arguments is a list of YUL identifiers representing function arguments.
	Arguments []*YulIdentifier `json:"arguments"`

	// Body is the body of the YUL function definition.
	Body Node[NodeType] `json:"body"`

	// ReturnParameters is a list of YUL identifiers representing return parameters.
	ReturnParameters []*YulIdentifier `json:"return_parameters"`
}

// NewYulFunctionDefinition creates a new YulFunctionDefinition instance.
func NewYulFunctionDefinition(b *ASTBuilder) *YulFunctionDefinition {
	return &YulFunctionDefinition{
		ASTBuilder: b,
		Id:         b.GetNextID(),
		NodeType:   ast_pb.NodeType_YUL_FUNCTION_DEFINITION,
	}
}

// SetReferenceDescriptor sets the reference descriptions of the YulFunctionDefinition node.
func (y *YulFunctionDefinition) SetReferenceDescriptor(refId int64, refDesc *TypeDescription) bool {
	return false
}

// GetId returns the ID of the YulFunctionDefinition.
func (y *YulFunctionDefinition) GetId() int64 {
	return y.Id
}

// GetType returns the NodeType of the YulFunctionDefinition.
func (y *YulFunctionDefinition) GetType() ast_pb.NodeType {
	return y.NodeType
}

// GetSrc returns the source information of the YulFunctionDefinition.
func (y *YulFunctionDefinition) GetSrc() SrcNode {
	return y.Src
}

// GetNodes returns a list containing the body node.
func (y *YulFunctionDefinition) GetNodes() []Node[NodeType] {
	toReturn := make([]Node[NodeType], 0)
	toReturn = append(toReturn, y.Body)

	for _, argument := range y.Arguments {
		toReturn = append(toReturn, argument)
	}

	for _, retParam := range y.ReturnParameters {
		toReturn = append(toReturn, retParam)
	}

	return toReturn
}

// GetTypeDescription returns the type description of the YulFunctionDefinition.
func (y *YulFunctionDefinition) GetTypeDescription() *TypeDescription {
	return &TypeDescription{}
}

// GetArguments returns the list of YUL identifiers representing function arguments.
func (y *YulFunctionDefinition) GetArguments() []*YulIdentifier {
	return y.Arguments
}

// GetBody returns the body of the YUL function definition.
func (y *YulFunctionDefinition) GetBody() Node[NodeType] {
	return y.Body
}

// GetReturnParameters returns the list of YUL identifiers representing return parameters.
func (y *YulFunctionDefinition) GetReturnParameters() []*YulIdentifier {
	return y.ReturnParameters
}

// ToProto converts the YulFunctionDefinition to its protocol buffer representation.
func (y *YulFunctionDefinition) ToProto() NodeType {
	toReturn := ast_pb.YulFunctionDefinition{
		Id:               y.GetId(),
		NodeType:         y.GetType(),
		Src:              y.GetSrc().ToProto(),
		Arguments:        make([]*v3.TypedStruct, 0),
		Body:             y.GetBody().ToProto().(*v3.TypedStruct),
		ReturnParameters: make([]*v3.TypedStruct, 0),
	}

	for _, ycase := range y.GetArguments() {
		toReturn.Arguments = append(toReturn.Arguments, ycase.ToProto().(*v3.TypedStruct))
	}

	for _, ycase := range y.GetReturnParameters() {
		toReturn.ReturnParameters = append(toReturn.ReturnParameters, ycase.ToProto().(*v3.TypedStruct))
	}

	return NewTypedStruct(&toReturn, "YulFunctionDefinition")
}

// UnmarshalJSON unmarshals a given JSON byte array into a YulFunctionDefinition node.
func (f *YulFunctionDefinition) UnmarshalJSON(data []byte) error {
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

	if arguments, ok := tempMap["arguments"]; ok {
		var nodes []json.RawMessage
		if err := json.Unmarshal(arguments, &nodes); err != nil {
			return err
		}

		for _, tempNode := range nodes {
			var tempIdentifier *YulIdentifier
			if err := json.Unmarshal(tempNode, &tempIdentifier); err != nil {
				return err
			}

			f.Arguments = append(f.Arguments, tempIdentifier)
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

	if retParams, ok := tempMap["return_parameters"]; ok {
		var nodes []json.RawMessage
		if err := json.Unmarshal(retParams, &nodes); err != nil {
			return err
		}

		for _, tempNode := range nodes {
			var tempIdentifier *YulIdentifier
			if err := json.Unmarshal(tempNode, &tempIdentifier); err != nil {
				return err
			}

			f.ReturnParameters = append(f.ReturnParameters, tempIdentifier)
		}
	}

	return nil
}

// Parse parses a YUL function definition.
func (y *YulFunctionDefinition) Parse(
	unit *SourceUnit[Node[ast_pb.SourceUnit]],
	contractNode Node[NodeType],
	fnNode Node[NodeType],
	bodyNode *BodyNode,
	assemblyNode *Yul,
	statementNode *YulStatement,
	ctx *parser.YulFunctionDefinitionContext,
) Node[NodeType] {
	y.Src = SrcNode{
		Line:        int64(ctx.GetStart().GetLine()),
		Column:      int64(ctx.GetStart().GetColumn()),
		Start:       int64(ctx.GetStart().GetStart()),
		End:         int64(ctx.GetStart().GetStop()),
		Length:      int64(ctx.GetStart().GetStop() - ctx.GetStart().GetStart() + 1),
		ParentIndex: statementNode.GetId(),
	}

	for _, argument := range ctx.GetArguments() {
		y.Arguments = append(y.Arguments, &YulIdentifier{
			Id:       y.GetNextID(),
			Name:     argument.GetText(),
			NodeType: ast_pb.NodeType_YUL_IDENTIFIER,
			Src: SrcNode{
				Line:        int64(argument.GetLine()),
				Column:      int64(argument.GetColumn()),
				Start:       int64(argument.GetStart()),
				End:         int64(argument.GetStop()),
				Length:      int64(argument.GetStop() - argument.GetStart() + 1),
				ParentIndex: y.GetId(),
			},
		})
	}

	if ctx.GetBody() != nil {
		blockStatement := NewYulBlockStatement(y.ASTBuilder)
		y.Body = blockStatement.Parse(
			unit, contractNode, fnNode, bodyNode, assemblyNode, statementNode, nil, y,
			ctx.GetBody().(*parser.YulBlockContext),
		)
	}

	for _, argument := range ctx.GetReturnParameters() {
		y.ReturnParameters = append(y.ReturnParameters, &YulIdentifier{
			Id:       y.GetNextID(),
			Name:     argument.GetText(),
			NodeType: ast_pb.NodeType_YUL_IDENTIFIER,
			Src: SrcNode{
				Line:        int64(argument.GetLine()),
				Column:      int64(argument.GetColumn()),
				Start:       int64(argument.GetStart()),
				End:         int64(argument.GetStop()),
				Length:      int64(argument.GetStop() - argument.GetStart() + 1),
				ParentIndex: y.GetId(),
			},
		})
	}

	return y
}
