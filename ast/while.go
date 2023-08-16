package ast

import (
	v3 "github.com/cncf/xds/go/xds/type/v3"
	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/parser"
)

// WhileStatement represents a while loop statement in the AST.
type WhileStatement struct {
	*ASTBuilder

	Id        int64           `json:"id"`        // Unique identifier for the WhileStatement node.
	NodeType  ast_pb.NodeType `json:"node_type"` // Type of the AST node.
	Kind      ast_pb.NodeType `json:"kind"`      // Kind of while loop.
	Src       SrcNode         `json:"src"`       // Source location information.
	Condition Node[NodeType]  `json:"condition"` // Condition expression of the while loop.
	Body      *BodyNode       `json:"body"`      // Body of the while loop.
}

// NewWhileStatement creates a new WhileStatement node with a given ASTBuilder.
func NewWhileStatement(b *ASTBuilder) *WhileStatement {
	return &WhileStatement{
		ASTBuilder: b,
		NodeType:   ast_pb.NodeType_WHILE_STATEMENT,
		Kind:       ast_pb.NodeType_WHILE,
	}
}

// SetReferenceDescriptor sets the reference descriptions of the WhileStatement node.
func (w *WhileStatement) SetReferenceDescriptor(refId int64, refDesc *TypeDescription) bool {
	return false
}

// GetId returns the ID of the WhileStatement node.
func (w *WhileStatement) GetId() int64 {
	return w.Id
}

// GetType returns the NodeType of the WhileStatement node.
func (w *WhileStatement) GetType() ast_pb.NodeType {
	return w.NodeType
}

// GetSrc returns the SrcNode of the WhileStatement node.
func (w *WhileStatement) GetSrc() SrcNode {
	return w.Src
}

// GetCondition returns the condition expression of the WhileStatement node.
func (w *WhileStatement) GetCondition() Node[NodeType] {
	return w.Condition
}

// GetBody returns the body of the WhileStatement node.
func (w *WhileStatement) GetBody() *BodyNode {
	return w.Body
}

// GetKind returns the kind of the while loop.
func (w *WhileStatement) GetKind() ast_pb.NodeType {
	return w.Kind
}

// GetTypeDescription returns the TypeDescription of the WhileStatement node.
func (w *WhileStatement) GetTypeDescription() *TypeDescription {
	return &TypeDescription{
		TypeString:     "while",
		TypeIdentifier: "$_t_while",
	}
}

// GetNodes returns the child nodes of the WhileStatement node.
func (w *WhileStatement) GetNodes() []Node[NodeType] {
	toReturn := make([]Node[NodeType], 0)
	toReturn = append(toReturn, w.Condition)
	toReturn = append(toReturn, w.Body.GetNodes()...)
	return toReturn
}

// ToProto returns a protobuf representation of the WhileStatement node.
func (w *WhileStatement) ToProto() NodeType {
	proto := ast_pb.While{
		Id:        w.GetId(),
		NodeType:  w.GetType(),
		Kind:      w.GetKind(),
		Src:       w.GetSrc().ToProto(),
		Condition: w.GetCondition().ToProto().(*v3.TypedStruct),
		Body:      w.Body.ToProto().(*ast_pb.Body),
	}

	return NewTypedStruct(&proto, "While")
}

// Parse parses a while loop statement context into the WhileStatement node.
func (w *WhileStatement) Parse(
	unit *SourceUnit[Node[ast_pb.SourceUnit]],
	contractNode Node[NodeType],
	fnNode Node[NodeType],
	bodyNode *BodyNode,
	ctx *parser.WhileStatementContext,
) Node[NodeType] {
	// Setting the source location information.
	w.Src = SrcNode{
		Id:          w.GetNextID(),
		Line:        int64(ctx.GetStart().GetLine()),
		Start:       int64(ctx.GetStart().GetStart()),
		End:         int64(ctx.GetStop().GetStop()),
		Length:      int64(ctx.GetStop().GetStop() - ctx.GetStart().GetStart() + 1),
		ParentIndex: bodyNode.Id,
	}

	// Parsing the condition expression.
	expression := NewExpression(w.ASTBuilder)
	w.Condition = expression.Parse(unit, contractNode, fnNode, bodyNode, nil, w, ctx.Expression())

	// Parsing the body of the while loop.
	if ctx.Statement() != nil && ctx.Statement().Block() != nil && !ctx.Statement().Block().IsEmpty() {
		bodyNode := NewBodyNode(w.ASTBuilder)
		bodyNode.ParseBlock(unit, contractNode, w, ctx.Statement().Block())
		w.Body = bodyNode

		// Parsing unchecked blocks within the body.
		if ctx.Statement().Block() != nil && ctx.Statement().Block().AllUncheckedBlock() != nil {
			for _, uncheckedCtx := range ctx.Statement().Block().AllUncheckedBlock() {
				bodyNode := NewBodyNode(w.ASTBuilder)
				bodyNode.ParseUncheckedBlock(unit, contractNode, w, uncheckedCtx)
				w.Body.Statements = append(w.Body.Statements, bodyNode)
			}
		}
	}

	return w
}
