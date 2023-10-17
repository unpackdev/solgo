package ast

import (
	"github.com/antlr4-go/antlr/v4"
	ast_pb "github.com/unpackdev/protos/dist/go/ast"
)

// YulLeaveStatement represents a YUL Leave statement in the AST.
type YulLeaveStatement struct {
	*ASTBuilder // Embedded ASTBuilder for utility functions.

	Id       int64           `json:"id"`        // Unique identifier for the statement.
	NodeType ast_pb.NodeType `json:"node_type"` // Type of the node in the AST.
	Src      SrcNode         `json:"src"`       // Source code location information.
}

// NewYulLeaveStatement creates and initializes a new YulLeaveStatement.
func NewYulLeaveStatement(b *ASTBuilder) *YulLeaveStatement {
	return &YulLeaveStatement{
		ASTBuilder: b,                         // Reference to ASTBuilder.
		Id:         b.GetNextID(),             // Generate new ID for the node.
		NodeType:   ast_pb.NodeType_YUL_LEAVE, // Set node type as YUL_LEAVE.
	}
}

// SetReferenceDescriptor sets the reference descriptions of the YulLeaveStatement node.
// Currently, it always returns false as there are no reference descriptions for YulLeaveStatement.
func (y *YulLeaveStatement) SetReferenceDescriptor(refId int64, refDesc *TypeDescription) bool {
	return false
}

// GetId retrieves the unique identifier of the YulLeaveStatement.
func (y *YulLeaveStatement) GetId() int64 {
	return y.Id
}

// GetType retrieves the node type of the YulLeaveStatement.
func (y *YulLeaveStatement) GetType() ast_pb.NodeType {
	return y.NodeType
}

// GetSrc retrieves the source code location information for the YulLeaveStatement.
func (y *YulLeaveStatement) GetSrc() SrcNode {
	return y.Src
}

// GetNodes retrieves child nodes of the YulLeaveStatement.
// For YulLeaveStatement, it always returns an empty slice as it doesn't have child nodes.
func (y *YulLeaveStatement) GetNodes() []Node[NodeType] {
	toReturn := make([]Node[NodeType], 0)
	return toReturn
}

// GetTypeDescription retrieves the type description of the YulLeaveStatement.
// This currently returns an empty TypeDescription.
func (y *YulLeaveStatement) GetTypeDescription() *TypeDescription {
	return &TypeDescription{}
}

// ToProto converts the YulLeaveStatement to its Protocol Buffer representation.
func (y *YulLeaveStatement) ToProto() NodeType {
	toReturn := ast_pb.YulLeaveStatement{
		Id:       y.GetId(),
		NodeType: y.GetType(),
		Src:      y.GetSrc().ToProto(),
	}

	return NewTypedStruct(&toReturn, "YulLeaveStatement")
}

// Parse populates the YulLeaveStatement fields based on the provided context.
// This method is typically used during the AST construction phase.
func (y *YulLeaveStatement) Parse(
	unit *SourceUnit[Node[ast_pb.SourceUnit]],
	contractNode Node[NodeType],
	fnNode Node[NodeType],
	bodyNode *BodyNode,
	assemblyNode *Yul,
	statementNode *YulStatement,
	ctx *antlr.TerminalNodeImpl, // Context from the ANTLR parse tree.
) Node[NodeType] {
	// Populate source location information based on the provided context.
	y.Src = SrcNode{
		Id:          y.GetNextID(),
		Line:        int64(ctx.GetSymbol().GetLine()),
		Column:      int64(ctx.GetSymbol().GetColumn()),
		Start:       int64(ctx.GetSymbol().GetStart()),
		End:         int64(ctx.GetSymbol().GetStop()),
		Length:      int64(ctx.GetSymbol().GetStop() - ctx.GetSymbol().GetStart() + 1),
		ParentIndex: statementNode.GetId(),
	}

	return y
}
