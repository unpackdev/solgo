package ast

import (
	ast_pb "github.com/unpackdev/protos/dist/go/ast"
	"github.com/unpackdev/solgo/parser"
)

// Yul represents an assembly statement in a Solidity source file.
type Yul struct {
	*ASTBuilder // Embedded ASTBuilder provides building functionalities for AST nodes.

	Id       int64           `json:"id"`        // Id uniquely identifies the assembly statement.
	NodeType ast_pb.NodeType `json:"node_type"` // NodeType specifies the type of the node.
	Src      SrcNode         `json:"src"`       // Src contains source location details of the node.
	Body     *BodyNode       `json:"body"`      // Body represents the content of the assembly statement.
}

// NewYul creates a new Yul and initializes its fields.
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

// GetType returns the type of the node. For Yul, this is always NodeType_ASSEMBLY_STATEMENT.
func (a *Yul) GetType() ast_pb.NodeType {
	return a.NodeType
}

// GetSrc provides source location details of the node.
func (a *Yul) GetSrc() SrcNode {
	return a.Src
}

// GetBody returns the content of the assembly statement.
func (a *Yul) GetBody() *BodyNode {
	return a.Body
}

// GetNodes retrieves the list of statements present in the assembly statement's body.
func (a *Yul) GetNodes() []Node[NodeType] {
	return a.Body.Statements
}

// GetTypeDescription provides a description of the assembly statement's type.
// For Yul, this always returns an empty TypeDescription.
func (a *Yul) GetTypeDescription() *TypeDescription {
	return &TypeDescription{}
}

// ToProto converts the assembly statement into its protobuf representation.
// Note: Complete implementation is yet to be provided.
func (a *Yul) ToProto() NodeType {
	proto := ast_pb.AssemblyStatement{
		Id:       a.GetId(),
		NodeType: a.GetType(),
		Src:      a.GetSrc().ToProto(),
		Body:     a.GetBody().ToProto().(*ast_pb.Body),
	}

	return NewTypedStruct(&proto, "AssemblyStatement")
}

// SetReferenceDescriptor is a placeholder method which currently always returns false.
// Future implementations might set a reference descriptor based on provided arguments.
func (a *Yul) SetReferenceDescriptor(refId int64, refDesc *TypeDescription) bool {
	return false
}

// Parse populates the Yul fields by parsing the provided AssemblyStatementContext.
// It processes the context to generate the body of the assembly statement and its constituent statements.
func (a *Yul) Parse(
	unit *SourceUnit[Node[ast_pb.SourceUnit]],
	contractNode Node[NodeType],
	fnNode Node[NodeType],
	bodyNode *BodyNode,
	ctx *parser.AssemblyStatementContext,
) Node[NodeType] {
	a.Src = SrcNode{
		Line:        int64(ctx.GetStart().GetLine()),
		Column:      int64(ctx.GetStart().GetColumn()),
		Start:       int64(ctx.GetStart().GetStart()),
		End:         int64(ctx.GetStop().GetStop()),
		Length:      int64(ctx.GetStop().GetStop() - ctx.GetStart().GetStart() + 1),
		ParentIndex: bodyNode.GetId(),
	}

	a.Body = NewBodyNode(a.ASTBuilder, false)
	a.Body.Src = a.Src
	a.Body.Src.ParentIndex = a.Id
	a.Body.NodeType = ast_pb.NodeType_YUL_BLOCK
	a.Body.Statements = make([]Node[NodeType], 0)

	yulStatement := NewYulStatement(a.ASTBuilder)

	for _, yulCtx := range ctx.AllYulStatement() {
		a.Body.Statements = append(a.Body.Statements,
			yulStatement.Parse(
				unit, contractNode, fnNode, a.Body, a, a, yulCtx.(*parser.YulStatementContext),
			),
		)
	}

	return a
}
