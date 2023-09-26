package ast

import (
	ast_pb "github.com/unpackdev/protos/dist/go/ast"
	"github.com/unpackdev/solgo/parser"
)

// NewExpr represents a new expression node in the AST.
// It contains information about the type being instantiated, type description, and related metadata.
type NewExpr struct {
	*ASTBuilder

	Id                    int64              `json:"id"`
	NodeType              ast_pb.NodeType    `json:"node_type"`
	Src                   SrcNode            `json:"src"`
	ArgumentTypes         []*TypeDescription `json:"argument_types"`
	TypeName              *TypeName          `json:"type_name"`
	ReferencedDeclaration int64              `json:"referenced_declaration,omitempty"`
	TypeDescription       *TypeDescription   `json:"type_description"`
}

// NewExprExpression creates a new NewExpr instance with initial values.
func NewExprExpression(b *ASTBuilder) *NewExpr {
	return &NewExpr{
		ASTBuilder:    b,
		Id:            b.GetNextID(),
		NodeType:      ast_pb.NodeType_NEW_EXPRESSION,
		ArgumentTypes: make([]*TypeDescription, 0),
	}
}

// SetReferenceDescriptor sets the reference descriptions of the NewExpr node.
func (n *NewExpr) SetReferenceDescriptor(refId int64, refDesc *TypeDescription) bool {
	n.ReferencedDeclaration = refId
	n.TypeDescription = refDesc
	return false
}

// GetId returns the ID of the NewExpr node.
func (n *NewExpr) GetId() int64 {
	return n.Id
}

// GetType returns the NodeType of the NewExpr node.
func (n *NewExpr) GetType() ast_pb.NodeType {
	return n.NodeType
}

// GetSrc returns the source information of the NewExpr node.
func (n *NewExpr) GetSrc() SrcNode {
	return n.Src
}

// GetArgumentTypes returns the type descriptions of arguments in the new expression.
func (n *NewExpr) GetArgumentTypes() []*TypeDescription {
	return n.ArgumentTypes
}

// GetTypeName returns the type name associated with the new expression.
func (n *NewExpr) GetTypeName() *TypeName {
	return n.TypeName
}

// GetTypeDescription returns the type description associated with the new expression.
func (n *NewExpr) GetTypeDescription() *TypeDescription {
	return n.TypeDescription
}

// GetNodes returns the list of child nodes of the NewExpr node.
func (n *NewExpr) GetNodes() []Node[NodeType] {
	return []Node[NodeType]{
		n.GetTypeName(),
	}
}

// GetReferencedDeclaration returns the ID of the referenced declaration in the context of new expression.
func (n *NewExpr) GetReferencedDeclaration() int64 {
	return n.ReferencedDeclaration
}

// ToProto converts the NewExpr node to its corresponding protobuf representation.
func (n *NewExpr) ToProto() NodeType {
	protos := ast_pb.NewExpression{
		Id:                    n.GetId(),
		NodeType:              n.GetType(),
		Src:                   n.GetSrc().ToProto(),
		ReferencedDeclaration: n.GetReferencedDeclaration(),
		TypeName:              n.GetTypeName().ToProto().(*ast_pb.TypeName),
		TypeDescription:       n.GetTypeDescription().ToProto(),
		ArgumentTypes:         make([]*ast_pb.TypeDescription, 0),
	}

	for _, arguments := range n.GetArgumentTypes() {
		protos.ArgumentTypes = append(protos.ArgumentTypes, arguments.ToProto())
	}

	return NewTypedStruct(&protos, "NewExpression")
}

// Parse populates the NewExpr node based on the provided context and other information.
func (n *NewExpr) Parse(
	unit *SourceUnit[Node[ast_pb.SourceUnit]],
	contractNode Node[NodeType],
	fnNode Node[NodeType],
	bodyNode *BodyNode,
	vDeclar *VariableDeclaration,
	exprNode Node[NodeType],
	ctx *parser.NewExprContext,
) Node[NodeType] {
	n.Src = SrcNode{
		Id:     n.GetNextID(),
		Line:   int64(ctx.GetStart().GetLine()),
		Column: int64(ctx.GetStart().GetColumn()),
		Start:  int64(ctx.GetStart().GetStart()),
		End:    int64(ctx.GetStop().GetStop()),
		Length: int64(ctx.GetStop().GetStop() - ctx.GetStart().GetStart() + 1),
		ParentIndex: func() int64 {
			if exprNode != nil {
				return exprNode.GetId()
			}

			if fnNode != nil {
				return fnNode.GetId()
			}

			return bodyNode.GetId()
		}(),
	}

	// Parsing the type name associated with the new expression.
	typeName := NewTypeName(n.ASTBuilder)
	typeName.WithParentNode(contractNode)
	typeName.WithBodyNode(bodyNode)
	typeName.WithParentNode(exprNode)
	typeName.Parse(unit, fnNode, n.GetId(), ctx.TypeName())
	n.TypeName = typeName
	n.TypeDescription = typeName.GetTypeDescription()
	return n
}
