package ast

import (
	v3 "github.com/cncf/xds/go/xds/type/v3"
	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/parser"
)

// TryStatement represents a try-catch statement in the AST.
type TryStatement struct {
	*ASTBuilder

	Id         int64            `json:"id"`         // Unique identifier for the TryStatement node.
	NodeType   ast_pb.NodeType  `json:"node_type"`  // Type of the AST node.
	Src        SrcNode          `json:"src"`        // Source location information.
	Body       *BodyNode        `json:"body"`       // Body of the try block.
	Kind       ast_pb.NodeType  `json:"kind"`       // Kind of try statement.
	Expression Node[NodeType]   `json:"expression"` // Expression within the try block.
	Clauses    []Node[NodeType] `json:"clauses"`    // List of catch clauses.
}

// NewTryStatement creates a new TryStatement node with a given ASTBuilder.
func NewTryStatement(b *ASTBuilder) *TryStatement {
	return &TryStatement{
		ASTBuilder: b,
		Id:         b.GetNextID(),
		NodeType:   ast_pb.NodeType_TRY_STATEMENT,
		Kind:       ast_pb.NodeType_TRY,
	}
}

// SetReferenceDescriptor sets the reference descriptions of the TryStatement node.
func (t *TryStatement) SetReferenceDescriptor(refId int64, refDesc *TypeDescription) bool {
	return false
}

// GetId returns the ID of the TryStatement node.
func (t *TryStatement) GetId() int64 {
	return t.Id
}

// GetType returns the NodeType of the TryStatement node.
func (t *TryStatement) GetType() ast_pb.NodeType {
	return t.NodeType
}

// GetSrc returns the SrcNode of the TryStatement node.
func (t *TryStatement) GetSrc() SrcNode {
	return t.Src
}

// GetBody returns the body of the TryStatement node.
func (t *TryStatement) GetBody() *BodyNode {
	return t.Body
}

// GetKind returns the kind of the try statement.
func (t *TryStatement) GetKind() ast_pb.NodeType {
	return t.Kind
}

// GetImplemented returns true if the try statement is implemented.
func (t *TryStatement) GetImplemented() bool {
	return true
}

// GetTypeDescription returns the TypeDescription of the TryStatement node.
func (t *TryStatement) GetTypeDescription() *TypeDescription {
	return &TypeDescription{
		TypeString:     "try",
		TypeIdentifier: "$_t_try",
	}
}

// GetNodes returns the child nodes of the TryStatement node.
func (t *TryStatement) GetNodes() []Node[NodeType] {
	toReturn := make([]Node[NodeType], 0)
	toReturn = append(toReturn, t.Body)
	toReturn = append(toReturn, t.Expression)
	toReturn = append(toReturn, t.Clauses...)
	return toReturn
}

// GetExpression returns the expression within the try block.
func (t *TryStatement) GetExpression() Node[NodeType] {
	return t.Expression
}

// GetClauses returns the list of catch clauses.
func (t *TryStatement) GetClauses() []Node[NodeType] {
	return t.Clauses
}

// ToProto returns a protobuf representation of the TryStatement node.
func (t *TryStatement) ToProto() NodeType {
	proto := ast_pb.Try{
		Id:       t.GetId(),
		NodeType: t.GetType(),
		Kind:     t.GetKind(),
		Src:      t.GetSrc().ToProto(),
	}

	if t.GetExpression() != nil {
		proto.Expression = t.GetExpression().ToProto().(*v3.TypedStruct)
	}

	if t.GetClauses() != nil {
		for _, clause := range t.GetClauses() {
			proto.Clauses = append(proto.Clauses, clause.ToProto().(*v3.TypedStruct))
		}
	}

	if t.GetBody() != nil {
		proto.Body = t.GetBody().ToProto().(*ast_pb.Body)
	}

	return NewTypedStruct(&proto, "Try")
}

// Parse parses a try-catch statement context into the TryStatement node.
func (t *TryStatement) Parse(
	unit *SourceUnit[Node[ast_pb.SourceUnit]],
	contractNode Node[NodeType],
	fnNode Node[NodeType],
	bodyNode *BodyNode,
	ctx *parser.TryStatementContext,
) Node[NodeType] {
	t.Src = SrcNode{
		Id:          t.GetNextID(),
		Line:        int64(ctx.GetStart().GetLine()),
		Start:       int64(ctx.GetStart().GetStart()),
		End:         int64(ctx.GetStop().GetStop()),
		Length:      int64(ctx.GetStop().GetStop() - ctx.GetStart().GetStart() + 1),
		ParentIndex: bodyNode.Id,
	}

	expression := NewExpression(t.ASTBuilder)
	t.Expression = expression.Parse(unit, contractNode, fnNode, bodyNode, nil, t, ctx.Expression())

	if ctx.Block() != nil && !ctx.Block().IsEmpty() {
		bodyNode := NewBodyNode(t.ASTBuilder)
		bodyNode.ParseBlock(unit, contractNode, t, ctx.Block())
		t.Body = bodyNode

		if ctx.Block().AllUncheckedBlock() != nil {
			for _, uncheckedCtx := range ctx.Block().AllUncheckedBlock() {
				bodyNode := NewBodyNode(t.ASTBuilder)
				bodyNode.ParseUncheckedBlock(unit, contractNode, t, uncheckedCtx)
				t.Body.Statements = append(t.Body.Statements, bodyNode)
			}
		}
	}

	for _, clauseCtx := range ctx.AllCatchClause() {
		clause := NewCatchClauseStatement(t.ASTBuilder)
		t.Clauses = append(t.Clauses, clause.Parse(
			unit, contractNode, fnNode, bodyNode, t, clauseCtx.(*parser.CatchClauseContext),
		))
	}

	return t
}
