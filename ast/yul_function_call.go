package ast

import (
	"encoding/json"

	v3 "github.com/cncf/xds/go/xds/type/v3"
	ast_pb "github.com/unpackdev/protos/dist/go/ast"
	"github.com/unpackdev/solgo/parser"
)

// YulFunctionCallStatement represents a YUL function call statement in the abstract syntax tree.
type YulFunctionCallStatement struct {
	*ASTBuilder

	// Id is the unique identifier of the YUL function call statement.
	Id int64 `json:"id"`

	// NodeType is the type of the YUL function call statement node.
	NodeType ast_pb.NodeType `json:"node_type"`

	// Src is the source location information of the YUL function call statement.
	Src SrcNode `json:"src"`

	// FunctionName is the name of the function being called.
	FunctionName *YulIdentifier `json:"function_name"`

	// Arguments is a list of argument nodes for the function call.
	Arguments []Node[NodeType] `json:"arguments"`
}

// NewYulFunctionCallStatement creates a new YulFunctionCallStatement instance.
func NewYulFunctionCallStatement(b *ASTBuilder) *YulFunctionCallStatement {
	return &YulFunctionCallStatement{
		ASTBuilder: b,
		Id:         b.GetNextID(),
		NodeType:   ast_pb.NodeType_YUL_FUNCTION_CALL,
		Arguments:  make([]Node[NodeType], 0),
	}
}

// SetReferenceDescriptor sets the reference descriptions of the YulFunctionCallStatement node.
func (y *YulFunctionCallStatement) SetReferenceDescriptor(refId int64, refDesc *TypeDescription) bool {
	return false
}

// GetId returns the ID of the YulFunctionCallStatement.
func (y *YulFunctionCallStatement) GetId() int64 {
	return y.Id
}

// GetType returns the NodeType of the YulFunctionCallStatement.
func (y *YulFunctionCallStatement) GetType() ast_pb.NodeType {
	return y.NodeType
}

// GetSrc returns the source information of the YulFunctionCallStatement.
func (y *YulFunctionCallStatement) GetSrc() SrcNode {
	return y.Src
}

// GetNodes returns a list containing the argument nodes.
func (y *YulFunctionCallStatement) GetNodes() []Node[NodeType] {
	toReturn := make([]Node[NodeType], 0)
	if y.FunctionName != nil {
		toReturn = append(toReturn, y.FunctionName)
	}

	toReturn = append(toReturn, y.Arguments...)
	return toReturn
}

// GetTypeDescription returns the type description of the YulFunctionCallStatement.
func (y *YulFunctionCallStatement) GetTypeDescription() *TypeDescription {
	return &TypeDescription{}
}

// GetFunctionName returns the name of the function being called.
func (y *YulFunctionCallStatement) GetFunctionName() *YulIdentifier {
	return y.FunctionName
}

// GetArguments returns the list of argument nodes for the function call.
func (y *YulFunctionCallStatement) GetArguments() []Node[NodeType] {
	return y.Arguments
}

// UnmarshalJSON unmarshals a given JSON byte array into a YulFunctionCallStatement node.
func (y *YulFunctionCallStatement) UnmarshalJSON(data []byte) error {
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

	if functionName, ok := tempMap["function_name"]; ok {
		if err := json.Unmarshal(functionName, &y.FunctionName); err != nil {
			return err
		}
	}

	if arguments, ok := tempMap["arguments"]; ok {
		var nodes []json.RawMessage
		if err := json.Unmarshal(arguments, &nodes); err != nil {
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
			y.Arguments = append(y.Arguments, node)
		}
	}

	return nil
}

// ToProto converts the YulFunctionCallStatement to its protocol buffer representation.
func (y *YulFunctionCallStatement) ToProto() NodeType {
	toReturn := ast_pb.YulFunctionCallStatement{
		Id:           y.GetId(),
		NodeType:     y.GetType(),
		Src:          y.GetSrc().ToProto(),
		FunctionName: y.GetFunctionName().ToProto().(*v3.TypedStruct),
		Arguments:    make([]*v3.TypedStruct, 0),
	}

	for _, ycase := range y.GetArguments() {
		toReturn.Arguments = append(toReturn.Arguments, ycase.ToProto().(*v3.TypedStruct))
	}

	return NewTypedStruct(&toReturn, "YulFunctionCallStatement")
}

// Parse parses a YUL function call statement.
func (y *YulFunctionCallStatement) Parse(
	unit *SourceUnit[Node[ast_pb.SourceUnit]],
	contractNode Node[NodeType],
	fnNode Node[NodeType],
	bodyNode *BodyNode,
	assemblyNode *Yul,
	statementNode *YulStatement,
	parentNode Node[NodeType],
	ctx *parser.YulFunctionCallContext,
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

	if ctx.YulEVMBuiltin() != nil {
		builtin := ctx.YulEVMBuiltin()
		y.FunctionName = &YulIdentifier{
			Id:       y.GetNextID(),
			NodeType: ast_pb.NodeType_YUL_IDENTIFIER,
			Src: SrcNode{
				Id:          y.GetNextID(),
				Line:        int64(builtin.GetSymbol().GetLine()),
				Column:      int64(builtin.GetSymbol().GetColumn()),
				Start:       int64(builtin.GetSymbol().GetStart()),
				End:         int64(builtin.GetSymbol().GetStop()),
				Length:      int64(builtin.GetSymbol().GetStop() - builtin.GetSymbol().GetStart() + 1),
				ParentIndex: y.GetId(),
			},
			Name: builtin.GetText(),
		}
	} else if ctx.YulIdentifier() != nil {
		identifier := ctx.YulIdentifier()
		y.FunctionName = &YulIdentifier{
			Id:       y.GetNextID(),
			NodeType: ast_pb.NodeType_YUL_IDENTIFIER,
			Src: SrcNode{
				Id:          y.GetNextID(),
				Line:        int64(identifier.GetSymbol().GetLine()),
				Column:      int64(identifier.GetSymbol().GetColumn()),
				Start:       int64(identifier.GetSymbol().GetStart()),
				End:         int64(identifier.GetSymbol().GetStop()),
				Length:      int64(identifier.GetSymbol().GetStop() - identifier.GetSymbol().GetStart() + 1),
				ParentIndex: y.GetId(),
			},
			Name: identifier.GetText(),
		}
	}

	if ctx.AllYulExpression() != nil {
		for _, expression := range ctx.AllYulExpression() {
			if expression.YulPath() != nil {
				path := expression.YulPath()
				for _, identifier := range path.AllYulIdentifier() {
					y.Arguments = append(y.Arguments, &YulIdentifier{
						Id:       y.GetNextID(),
						NodeType: ast_pb.NodeType_YUL_IDENTIFIER,
						Name:     identifier.GetText(),
						Src: SrcNode{
							Id:          y.GetNextID(),
							Line:        int64(identifier.GetSymbol().GetLine()),
							Column:      int64(identifier.GetSymbol().GetColumn()),
							Start:       int64(identifier.GetSymbol().GetStart()),
							End:         int64(identifier.GetSymbol().GetStop()),
							Length:      int64(identifier.GetSymbol().GetStop() - identifier.GetSymbol().GetStart() + 1),
							ParentIndex: y.GetId(),
						},
					})
				}
			}

			if expression.YulFunctionCall() != nil {
				fc := expression.YulFunctionCall()
				fcStatement := NewYulFunctionCallStatement(y.ASTBuilder)
				y.Arguments = append(y.Arguments, fcStatement.Parse(
					unit, contractNode, fnNode, bodyNode, assemblyNode, statementNode, y,
					fc.(*parser.YulFunctionCallContext),
				))
			}

			if expression.YulLiteral() != nil {
				literal := expression.YulLiteral()
				literalStatement := NewYulLiteralStatement(y.ASTBuilder)
				y.Arguments = append(y.Arguments, literalStatement.Parse(
					unit, contractNode, fnNode, bodyNode, assemblyNode, statementNode, y,
					literal.(*parser.YulLiteralContext),
				))
			}
		}
	}

	return y
}
