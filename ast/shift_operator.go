package ast

import (
	"github.com/goccy/go-json"

	v3 "github.com/cncf/xds/go/xds/type/v3"
	ast_pb "github.com/unpackdev/protos/dist/go/ast"
	"github.com/unpackdev/solgo/parser"
)

// ShiftOperation represents an 'bit and' operation in an abstract syntax tree.
type ShiftOperation struct {
	*ASTBuilder

	Id               int64              `json:"id"`
	NodeType         ast_pb.NodeType    `json:"node_type"`
	Src              SrcNode            `json:"src"`
	Operator         ast_pb.NodeType    `json:"operator"`
	Expressions      []Node[NodeType]   `json:"expressions"`
	TypeDescriptions []*TypeDescription `json:"type_descriptions"`
	TypeDescription  *TypeDescription   `json:"type_description"`
}

// NewShiftOperationExpression creates a new ShiftOperation instance.
func NewShiftOperationExpression(b *ASTBuilder) *ShiftOperation {
	return &ShiftOperation{
		ASTBuilder:       b,
		Id:               b.GetNextID(),
		NodeType:         ast_pb.NodeType_SHIFT_OPERATION,
		TypeDescriptions: make([]*TypeDescription, 0),
	}
}

// SetReferenceDescriptor sets the reference descriptions of the ShiftOperation node.
// This function always returns false for now.
func (f *ShiftOperation) SetReferenceDescriptor(refId int64, refDesc *TypeDescription) bool {
	f.TypeDescriptions = []*TypeDescription{}
	for _, expr := range f.Expressions {
		f.TypeDescriptions = append(f.TypeDescriptions, expr.GetTypeDescription())
	}
	f.TypeDescription = f.TypeDescriptions[1]
	return false
}

// GetId returns the ID of the ShiftOperation.
func (f *ShiftOperation) GetId() int64 {
	return f.Id
}

// GetType returns the NodeType of the ShiftOperation.
func (f *ShiftOperation) GetType() ast_pb.NodeType {
	return f.NodeType
}

// GetSrc returns the source information of the ShiftOperation.
func (f *ShiftOperation) GetSrc() SrcNode {
	return f.Src
}

// GetTypeDescription returns the type description associated with the ShiftOperation.
func (f *ShiftOperation) GetTypeDescription() *TypeDescription {
	return f.TypeDescription
}

// GetNodes returns the child nodes of the ShiftOperation.
func (f *ShiftOperation) GetNodes() []Node[NodeType] {
	toReturn := []Node[NodeType]{}
	toReturn = append(toReturn, f.Expressions...)
	return toReturn
}

// GetExpressions returns the expressions within the ShiftOperation.
func (f *ShiftOperation) GetExpressions() []Node[NodeType] {
	return f.Expressions
}

func (f *ShiftOperation) UnmarshalJSON(data []byte) error {
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

	if operator, ok := tempMap["operator"]; ok {
		if err := json.Unmarshal(operator, &f.Operator); err != nil {
			return err
		}
	}

	if src, ok := tempMap["src"]; ok {
		if err := json.Unmarshal(src, &f.Src); err != nil {
			return err
		}
	}

	if expressions, ok := tempMap["expressions"]; ok {
		var nodes []json.RawMessage
		if err := json.Unmarshal(expressions, &nodes); err != nil {
			return err
		}

		for _, node := range nodes {
			var tempNode map[string]json.RawMessage
			if err := json.Unmarshal(node, &tempNode); err != nil {
				return err
			}

			var tempNodeType ast_pb.NodeType
			if err := json.Unmarshal(tempNode["node_type"], &tempNodeType); err != nil {
				return err
			}

			parsedNode, err := unmarshalNode(node, tempNodeType)
			if err != nil {
				return err
			}

			f.Expressions = append(f.Expressions, parsedNode)
		}
	}

	if typeDescriptions, ok := tempMap["type_descriptions"]; ok {
		if err := json.Unmarshal(typeDescriptions, &f.TypeDescriptions); err != nil {
			return err
		}
	}

	return nil
}

// ToProto converts the ShiftOperation to its corresponding protocol buffer representation.
func (f *ShiftOperation) ToProto() NodeType {
	proto := ast_pb.ShiftOperation{
		Id:               f.GetId(),
		NodeType:         f.GetType(),
		Src:              f.GetSrc().ToProto(),
		Expressions:      make([]*v3.TypedStruct, 0),
		TypeDescriptions: make([]*ast_pb.TypeDescription, 0),
	}

	for _, exp := range f.GetExpressions() {
		proto.Expressions = append(proto.Expressions, exp.ToProto().(*v3.TypedStruct))
	}

	for _, typeDesc := range f.TypeDescriptions {
		proto.TypeDescriptions = append(proto.TypeDescriptions, typeDesc.ToProto())
	}

	return NewTypedStruct(&proto, "ShiftOperation")
}

// Parse parses the ShiftOperation node from the parsing context and associates it with other nodes.
func (f *ShiftOperation) Parse(
	unit *SourceUnit[Node[ast_pb.SourceUnit]],
	contractNode Node[NodeType],
	fnNode Node[NodeType],
	bodyNode *BodyNode,
	vDeclar *VariableDeclaration,
	expNode Node[NodeType],
	parentNodeId int64,
	ctx *parser.ShiftOperationContext,
) Node[NodeType] {
	f.Id = f.GetNextID()
	f.Src = SrcNode{
		Line:        int64(ctx.GetStart().GetLine()),
		Column:      int64(ctx.GetStart().GetColumn()),
		Start:       int64(ctx.GetStart().GetStart()),
		End:         int64(ctx.GetStop().GetStop()),
		Length:      int64(ctx.GetStop().GetStop() - ctx.GetStart().GetStart() + 1),
		ParentIndex: parentNodeId,
	}

	if ctx.Shr() != nil {
		f.Operator = ast_pb.NodeType_SHIFT_RIGHT_OPERATION
	} else if ctx.Shl() != nil {
		f.Operator = ast_pb.NodeType_SHIFT_LEFT_OPERATION
	}

	expression := NewExpression(f.ASTBuilder)

	for _, expr := range ctx.AllExpression() {
		parsedExp := expression.Parse(unit, contractNode, fnNode, bodyNode, vDeclar, f, f.GetId(), expr)
		f.Expressions = append(
			f.Expressions,
			parsedExp,
		)
		f.TypeDescriptions = append(f.TypeDescriptions, parsedExp.GetTypeDescription())
	}

	f.TypeDescription = f.TypeDescriptions[1]

	return f
}
