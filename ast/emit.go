package ast

import (
	"encoding/json"

	v3 "github.com/cncf/xds/go/xds/type/v3"
	ast_pb "github.com/unpackdev/protos/dist/go/ast"
	"github.com/unpackdev/solgo/parser"
)

// Emit represents an emit statement node in the abstract syntax tree.
type Emit struct {
	*ASTBuilder

	Id         int64            `json:"id"`         // Unique identifier of the emit statement node.
	NodeType   ast_pb.NodeType  `json:"node_type"`  // Type of the node.
	Src        SrcNode          `json:"src"`        // Source location information.
	Arguments  []Node[NodeType] `json:"arguments"`  // List of arguments for the emit statement.
	Expression Node[NodeType]   `json:"expression"` // Expression node associated with the emit statement.
}

// NewEmitStatement creates a new instance of Emit with the provided ASTBuilder.
func NewEmitStatement(b *ASTBuilder) *Emit {
	return &Emit{
		ASTBuilder: b,
		Id:         b.GetNextID(),
		NodeType:   ast_pb.NodeType_EMIT_STATEMENT,
		Arguments:  make([]Node[NodeType], 0),
	}
}

// SetReferenceDescriptor sets the reference descriptors of the Emit node.
func (e *Emit) SetReferenceDescriptor(refId int64, refDesc *TypeDescription) bool {
	return false
}

// GetId returns the unique identifier of the emit statement node.
func (e *Emit) GetId() int64 {
	return e.Id
}

// GetType returns the type of the node.
func (e *Emit) GetType() ast_pb.NodeType {
	return e.NodeType
}

// GetSrc returns the source location information of the emit statement node.
func (e *Emit) GetSrc() SrcNode {
	return e.Src
}

// GetArguments returns the list of arguments associated with the emit statement.
func (e *Emit) GetArguments() []Node[NodeType] {
	return e.Arguments
}

// GetExpression returns the expression node associated with the emit statement.
func (e *Emit) GetExpression() Node[NodeType] {
	return e.Expression
}

// GetTypeDescription returns the type description of the emit statement.
func (e *Emit) GetTypeDescription() *TypeDescription {
	return &TypeDescription{
		TypeString:     "emit",
		TypeIdentifier: "$_t_emit",
	}
}

// GetNodes returns a list of nodes associated with the emit statement (arguments and expression).
func (e *Emit) GetNodes() []Node[NodeType] {
	toReturn := make([]Node[NodeType], 0)
	toReturn = append(toReturn, e.Arguments...)
	toReturn = append(toReturn, e.Expression)
	return toReturn
}

func (e *Emit) UnmarshalJSON(data []byte) error {
	var tempMap map[string]json.RawMessage
	if err := json.Unmarshal(data, &tempMap); err != nil {
		return err
	}

	if id, ok := tempMap["id"]; ok {
		if err := json.Unmarshal(id, &e.Id); err != nil {
			return err
		}
	}

	if nodeType, ok := tempMap["node_type"]; ok {
		if err := json.Unmarshal(nodeType, &e.NodeType); err != nil {
			return err
		}
	}

	if src, ok := tempMap["src"]; ok {
		if err := json.Unmarshal(src, &e.Src); err != nil {
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
			e.Arguments = append(e.Arguments, node)
		}
	}

	if expression, ok := tempMap["expression"]; ok {
		if err := json.Unmarshal(expression, &e.Expression); err != nil {
			var tempNodeMap map[string]json.RawMessage
			if err := json.Unmarshal(expression, &tempNodeMap); err != nil {
				return err
			}

			var tempNodeType ast_pb.NodeType
			if err := json.Unmarshal(tempNodeMap["node_type"], &tempNodeType); err != nil {
				return err
			}

			node, err := unmarshalNode(expression, tempNodeType)
			if err != nil {
				return err
			}
			e.Expression = node
		}
	}

	return nil
}

// ToProto converts the Emit node to its corresponding protobuf representation.
func (e *Emit) ToProto() NodeType {
	proto := ast_pb.Emit{
		Id:         e.GetId(),
		NodeType:   e.GetType(),
		Src:        e.GetSrc().ToProto(),
		Arguments:  make([]*v3.TypedStruct, 0),
		Expression: e.GetExpression().ToProto().(*v3.TypedStruct),
	}

	for _, argument := range e.GetArguments() {
		proto.Arguments = append(proto.Arguments, argument.ToProto().(*v3.TypedStruct))
	}

	return NewTypedStruct(&proto, "Emit")
}

// Parse parses the emit statement context and populates the Emit fields.
func (e *Emit) Parse(unit *SourceUnit[Node[ast_pb.SourceUnit]],
	contractNode Node[NodeType],
	fnNode Node[NodeType],
	bodyNode *BodyNode,
	ctx *parser.EmitStatementContext,
) Node[NodeType] {
	e.Src = SrcNode{
		Id:          e.GetNextID(),
		Line:        int64(ctx.GetStart().GetLine()),
		Column:      int64(ctx.GetStart().GetColumn()),
		Start:       int64(ctx.GetStart().GetStart()),
		End:         int64(ctx.GetStop().GetStop()),
		Length:      int64(ctx.GetStop().GetStop() - ctx.GetStart().GetStart() + 1),
		ParentIndex: fnNode.GetId(),
	}

	expression := NewExpression(e.ASTBuilder)

	for _, argumentCtx := range ctx.CallArgumentList().AllExpression() {
		argument := expression.Parse(unit, contractNode, fnNode, bodyNode, nil, nil, argumentCtx)
		e.Arguments = append(e.Arguments, argument)
	}

	e.Expression = expression.Parse(unit, contractNode, fnNode, bodyNode, nil, nil, ctx.Expression())
	return e
}
