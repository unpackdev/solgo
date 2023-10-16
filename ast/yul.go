package ast

import (
	ast_pb "github.com/unpackdev/protos/dist/go/ast"
	"github.com/unpackdev/solgo/parser"
)

// Yul represents an assembly statement in a Solidity source file.
// @WARN: Yul is not yet implemented.
type Yul struct {
	*ASTBuilder

	// Id is the unique identifier of the assembly statement.
	Id int64 `json:"id"`
	// NodeType is the type of the node.
	// For an Yul, this is always NodeType_ASSEMBLY_STATEMENT.
	NodeType ast_pb.NodeType `json:"node_type"`
	// Src contains source information about the node, such as its line and column numbers in the source file.
	Src SrcNode `json:"src"`
	// Body is the body of the assembly statement, represented as a BodyNode.
	Body *BodyNode `json:"body"`
}

// NewYul creates a new Yul.
func NewYul(b *ASTBuilder) *Yul {
	return &Yul{
		ASTBuilder: b,
		Id:         b.GetNextID(),
		NodeType:   ast_pb.NodeType_ASSEMBLY_STATEMENT,
	}
}

// GetId returns the unique identifier of the assembly statement.
func (a *Yul) GetId() int64 {
	return a.Id
}

// GetType returns the type of the node.
// For an Yul, this is always NodeType_ASSEMBLY_STATEMENT.
func (a *Yul) GetType() ast_pb.NodeType {
	return a.NodeType
}

// GetSrc returns source information about the node, such as its line and column numbers in the source file.
func (a *Yul) GetSrc() SrcNode {
	return a.Src
}

// GetBody returns the body of the assembly statement, represented as a BodyNode.
func (a *Yul) GetBody() *BodyNode {
	return a.Body
}

// GetNodes returns the statements in the body of the assembly statement.
func (a *Yul) GetNodes() []Node[NodeType] {
	return a.Body.Statements
}

// GetTypeDescription returns the type description of the assembly statement.
// For an Yul, this is always nil.
func (a *Yul) GetTypeDescription() *TypeDescription {
	return &TypeDescription{}
}

// ToProto returns the protobuf representation of the assembly statement.
// @TODO: Implement body type...
func (a *Yul) ToProto() NodeType {
	proto := ast_pb.Assembly{
		Id:       a.GetId(),
		NodeType: a.GetType(),
		Src:      a.GetSrc().ToProto(),
	}

	return NewTypedStruct(&proto, "Assembly")
}

func (a *Yul) SetReferenceDescriptor(refId int64, refDesc *TypeDescription) bool {
	return false
}

// Parse parses an YulContext to populate the fields of the Yul.
func (a *Yul) Parse(
	unit *SourceUnit[Node[ast_pb.SourceUnit]],
	contractNode Node[NodeType],
	fnNode Node[NodeType],
	bodyNode *BodyNode,
	ctx *parser.AssemblyStatementContext,
) Node[NodeType] {
	a.Src = SrcNode{
		Id:          a.GetNextID(),
		Line:        int64(ctx.GetStart().GetLine()),
		Column:      int64(ctx.GetStart().GetColumn()),
		Start:       int64(ctx.GetStart().GetStart()),
		End:         int64(ctx.GetStop().GetStop()),
		Length:      int64(ctx.GetStop().GetStop() - ctx.GetStart().GetStart() + 1),
		ParentIndex: bodyNode.GetId(),
	}

	yulStatement := NewYulStatement(a.ASTBuilder)

	a.Body = NewBodyNode(a.ASTBuilder, false)
	a.Body.Src = a.Src
	a.Body.Src.ParentIndex = a.Id
	a.Body.NodeType = ast_pb.NodeType_AST
	a.Body.Statements = make([]Node[NodeType], 0)

	for _, yulCtx := range ctx.AllYulStatement() {
		a.Body.Statements = append(a.Body.Statements,
			yulStatement.Parse(
				unit, contractNode, fnNode, a.Body, a, yulCtx.(*parser.YulStatementContext),
			),
		)
	}

	return a
}
