package ast

import (
	ast_pb "github.com/unpackdev/protos/dist/go/ast"
	"github.com/unpackdev/solgo/parser"
)

// The CatchStatement struct represents a 'catch' clause in a 'try-catch' statement in Solidity.
type CatchStatement struct {
	// Embedding the ASTBuilder to provide common functionality
	*ASTBuilder

	// The unique identifier for the 'catch' clause
	Id int64 `json:"id"`

	// The name of the exception variable in the 'catch' clause, if any
	Name string `json:"name,omitempty"`

	// The type of the node, which is 'TRY_CATCH_CLAUSE' for a 'catch' clause
	NodeType ast_pb.NodeType `json:"node_type"`

	// The kind of the node, which is 'CATCH' for a 'catch' clause
	Kind ast_pb.NodeType `json:"kind"`

	// The source information about the 'catch' clause, such as its line and column numbers in the source file
	Src SrcNode `json:"src"`

	// The body of the 'catch' clause, which is a block of statements
	Body *BodyNode `json:"body"`

	// The parameters of the 'catch' clause, if any
	Parameters *ParameterList `json:"parameters"`
}

// NewCatchClauseStatement creates a new CatchStatement instance.
func NewCatchClauseStatement(b *ASTBuilder) *CatchStatement {
	return &CatchStatement{
		ASTBuilder: b,
		NodeType:   ast_pb.NodeType_TRY_CATCH_CLAUSE,
		Kind:       ast_pb.NodeType_CATCH,
	}
}

// SetReferenceDescriptor sets the reference descriptions of the CatchStatement node.
func (t *CatchStatement) SetReferenceDescriptor(refId int64, refDesc *TypeDescription) bool {
	return false
}

// GetId returns the unique identifier of the 'catch' clause.
func (t *CatchStatement) GetId() int64 {
	return t.Id
}

// GetType returns the type of the node, which is 'TRY_CATCH_CLAUSE' for a 'catch' clause.
func (t *CatchStatement) GetType() ast_pb.NodeType {
	return t.NodeType
}

// GetSrc returns the source information about the 'catch' clause.
func (t *CatchStatement) GetSrc() SrcNode {
	return t.Src
}

// GetBody returns the body of the 'catch' clause.
func (t *CatchStatement) GetBody() *BodyNode {
	return t.Body
}

// GetKind returns the kind of the node, which is 'CATCH' for a 'catch' clause.
func (t *CatchStatement) GetKind() ast_pb.NodeType {
	return t.Kind
}

// GetParameters returns the parameters of the 'catch' clause.
func (t *CatchStatement) GetParameters() *ParameterList {
	return t.Parameters
}

// GetTypeDescription returns the type description of the 'catch' clause, which is nil as 'catch' clauses do not have a type description.
func (t *CatchStatement) GetTypeDescription() *TypeDescription {
	return &TypeDescription{
		TypeString:     "catch",
		TypeIdentifier: "$_t_catch",
	}
}

// GetNodes returns the statements in the body of the 'catch' clause.
func (t *CatchStatement) GetNodes() []Node[NodeType] {
	return t.Body.Statements
}

// GetName returns the name of the exception variable in the 'catch' clause, if any.
func (t *CatchStatement) GetName() string {
	return t.Name
}

// ToProto returns the protobuf representation of the 'catch' clause.
func (t *CatchStatement) ToProto() NodeType {
	proto := ast_pb.Catch{
		Id:         t.GetId(),
		Name:       t.GetName(),
		NodeType:   t.GetType(),
		Kind:       t.GetKind(),
		Src:        t.GetSrc().ToProto(),
		Parameters: t.GetParameters().ToProto(),
		Body:       t.GetBody().ToProto().(*ast_pb.Body),
	}

	return NewTypedStruct(&proto, "Catch")
}

// Parse parses a 'catch' clause from the provided parser.CatchClauseContext and returns the corresponding CatchStatement.
func (t *CatchStatement) Parse(
	unit *SourceUnit[Node[ast_pb.SourceUnit]],
	contractNode Node[NodeType],
	fnNode Node[NodeType],
	bodyNode *BodyNode,
	tryNode *TryStatement,
	ctx *parser.CatchClauseContext,
) Node[NodeType] {
	t.Src = SrcNode{
		Id:          t.GetNextID(),
		Line:        int64(ctx.GetStart().GetLine()),
		Start:       int64(ctx.GetStart().GetStart()),
		End:         int64(ctx.GetStop().GetStop()),
		Length:      int64(ctx.GetStop().GetStop() - ctx.GetStart().GetStart() + 1),
		ParentIndex: tryNode.Id,
	}

	if ctx.Identifier() != nil {
		t.Name = ctx.Identifier().GetText()
	}

	pList := NewParameterList(t.ASTBuilder)
	if ctx.ParameterList() != nil {
		pList.Parse(unit, t, ctx.ParameterList())
	} else {
		pList.Src = t.Src
		pList.Src.ParentIndex = t.Id
	}
	t.Parameters = pList

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

	return t
}
