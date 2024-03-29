package ast

import (
	"github.com/goccy/go-json"

	v3 "github.com/cncf/xds/go/xds/type/v3"
	ast_pb "github.com/unpackdev/protos/dist/go/ast"
	"github.com/unpackdev/solgo/parser"
)

// DoWhileStatement represents a do-while loop statement node in the abstract syntax tree (AST).
// It encapsulates information about the condition and body of the loop.
type DoWhileStatement struct {
	*ASTBuilder                 // Embedded ASTBuilder for building the AST.
	Id          int64           `json:"id"`        // Unique identifier for the DoWhileStatement node.
	NodeType    ast_pb.NodeType `json:"node_type"` // Type of the AST node.
	Src         SrcNode         `json:"src"`       // Source location information.
	Condition   Node[NodeType]  `json:"condition"` // Condition expression for the do-while loop.
	Body        *BodyNode       `json:"body"`      // Body of the do-while loop.
}

// NewDoWhileStatement creates a new DoWhileStatement node with default values and returns it.
func NewDoWhileStatement(b *ASTBuilder) *DoWhileStatement {
	return &DoWhileStatement{
		ASTBuilder: b,
		Id:         b.GetNextID(),
		NodeType:   ast_pb.NodeType_DO_WHILE_STATEMENT,
	}
}

// SetReferenceDescriptor sets the reference descriptions of the DoWhileStatement node.
// This function currently returns false, as no reference description updates are performed.
func (d *DoWhileStatement) SetReferenceDescriptor(refId int64, refDesc *TypeDescription) bool {
	return false
}

// GetId returns the unique identifier of the DoWhileStatement node.
func (d *DoWhileStatement) GetId() int64 {
	return d.Id
}

// GetType returns the type of the AST node, which is NodeType_DO_WHILE_STATEMENT for a do-while loop.
func (d *DoWhileStatement) GetType() ast_pb.NodeType {
	return d.NodeType
}

// GetSrc returns the source location information of the DoWhileStatement node.
func (d *DoWhileStatement) GetSrc() SrcNode {
	return d.Src
}

// GetCondition returns the condition expression of the do-while loop.
func (d *DoWhileStatement) GetCondition() Node[NodeType] {
	return d.Condition
}

// GetBody returns the body of the do-while loop.
func (d *DoWhileStatement) GetBody() *BodyNode {
	return d.Body
}

// GetNodes returns a slice of child nodes within the do-while loop.
func (d *DoWhileStatement) GetNodes() []Node[NodeType] {
	toReturn := []Node[NodeType]{d.Condition}
	toReturn = append(toReturn, d.Body.GetNodes()...)
	return toReturn
}

// GetTypeDescription returns the type description associated with the DoWhileStatement node.
func (d *DoWhileStatement) GetTypeDescription() *TypeDescription {
	return &TypeDescription{
		TypeString:     "dowhile",
		TypeIdentifier: "$_t_do_while",
	}
}

// MarshalJSON marshals the DoWhileStatement node into a JSON-encoded byte slice.
func (d *DoWhileStatement) UnmarshalJSON(data []byte) error {
	var tempMap map[string]json.RawMessage
	if err := json.Unmarshal(data, &tempMap); err != nil {
		return err
	}

	if id, ok := tempMap["id"]; ok {
		if err := json.Unmarshal(id, &d.Id); err != nil {
			return err
		}
	}

	if nodeType, ok := tempMap["node_type"]; ok {
		if err := json.Unmarshal(nodeType, &d.NodeType); err != nil {
			return err
		}
	}

	if src, ok := tempMap["src"]; ok {
		if err := json.Unmarshal(src, &d.Src); err != nil {
			return err
		}
	}

	if condition, ok := tempMap["condition"]; ok {
		if err := json.Unmarshal(condition, &d.Condition); err != nil {
			var tempNodeMap map[string]json.RawMessage
			if err := json.Unmarshal(condition, &tempNodeMap); err != nil {
				return err
			}

			var tempNodeType ast_pb.NodeType
			if err := json.Unmarshal(tempNodeMap["node_type"], &tempNodeType); err != nil {
				return err
			}

			node, err := unmarshalNode(condition, tempNodeType)
			if err != nil {
				return err
			}
			d.Condition = node
		}
	}

	if body, ok := tempMap["body"]; ok {
		if err := json.Unmarshal(body, &d.Body); err != nil {
			return err
		}
	}

	return nil
}

// ToProto converts the DoWhileStatement node to its corresponding protocol buffer representation.
func (d *DoWhileStatement) ToProto() NodeType {
	proto := ast_pb.Do{
		Id:        d.GetId(),
		NodeType:  d.GetType(),
		Src:       d.GetSrc().ToProto(),
		Condition: d.GetCondition().ToProto().(*v3.TypedStruct),
		Body:      d.GetBody().ToProto().(*ast_pb.Body),
	}

	return NewTypedStruct(&proto, "Do")
}

// Parse is responsible for parsing the do-while loop statement from the context and populating the DoWhileStatement node.
func (d *DoWhileStatement) Parse(
	unit *SourceUnit[Node[ast_pb.SourceUnit]],
	contractNode Node[NodeType],
	fnNode Node[NodeType],
	bodyNode *BodyNode,
	ctx *parser.DoWhileStatementContext,
) Node[NodeType] {
	d.Src = SrcNode{
		Line:        int64(ctx.GetStart().GetLine()),
		Start:       int64(ctx.GetStart().GetStart()),
		End:         int64(ctx.GetStop().GetStop()),
		Length:      int64(ctx.GetStop().GetStop() - ctx.GetStart().GetStart() + 1),
		ParentIndex: bodyNode.Id,
	}

	expression := NewExpression(d.ASTBuilder)
	d.Condition = expression.Parse(unit, contractNode, fnNode, bodyNode, nil, d, d.GetId(), ctx.Expression())

	if ctx.Statement() != nil && ctx.Statement().Block() != nil && !ctx.Statement().Block().IsEmpty() {
		bodyNode := NewBodyNode(d.ASTBuilder, false)
		bodyNode.ParseBlock(unit, contractNode, d, ctx.Statement().Block())
		d.Body = bodyNode

		if ctx.Statement().Block() != nil && ctx.Statement().Block().AllUncheckedBlock() != nil {
			for _, uncheckedCtx := range ctx.Statement().Block().AllUncheckedBlock() {
				bodyNode := NewBodyNode(d.ASTBuilder, false)
				bodyNode.ParseUncheckedBlock(unit, contractNode, d, uncheckedCtx)
				d.Body.Statements = append(d.Body.Statements, bodyNode)
			}
		}
	}

	return d
}
