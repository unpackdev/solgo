package ast

import (
	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/parser"
)

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

func NewExprExpression(b *ASTBuilder) *NewExpr {
	return &NewExpr{
		ASTBuilder:    b,
		Id:            b.GetNextID(),
		NodeType:      ast_pb.NodeType_NEW_EXPRESSION,
		ArgumentTypes: make([]*TypeDescription, 0),
	}
}

// SetReferenceDescriptor sets the reference descriptions of the NewExpr node.
func (m *NewExpr) SetReferenceDescriptor(refId int64, refDesc *TypeDescription) bool {
	m.ReferencedDeclaration = refId
	m.TypeDescription = refDesc
	return false
}

func (n *NewExpr) GetId() int64 {
	return n.Id
}

func (n *NewExpr) GetType() ast_pb.NodeType {
	return n.NodeType
}

func (n *NewExpr) GetSrc() SrcNode {
	return n.Src
}

func (n *NewExpr) GetArgumentTypes() []*TypeDescription {
	return n.ArgumentTypes
}

func (n *NewExpr) GetTypeName() *TypeName {
	return n.TypeName
}

func (n *NewExpr) GetTypeDescription() *TypeDescription {
	return n.TypeDescription
}

func (n *NewExpr) GetNodes() []Node[NodeType] {
	return nil
}

func (n *NewExpr) ToProto() NodeType {
	return ast_pb.Node{} // @TODO: Lazy right now...
}

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

	typeName := NewTypeName(n.ASTBuilder)
	typeName.Parse(unit, fnNode, n.GetId(), ctx.TypeName())
	n.TypeName = typeName
	n.TypeDescription = typeName.GetTypeDescription()

	return n
}
