package ast

import (
	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/parser"
)

type IndexAccess struct {
	*ASTBuilder

	Id               int64              `json:"id"`
	NodeType         ast_pb.NodeType    `json:"node_type"`
	Src              SrcNode            `json:"src"`
	IndexExpression  Node[NodeType]     `json:"index_expression"`
	BaseExpression   Node[NodeType]     `json:"base_expression"`
	TypeDescriptions []*TypeDescription `json:"type_descriptions"`
}

func NewIndexAccess(b *ASTBuilder) *IndexAccess {
	return &IndexAccess{
		ASTBuilder: b,
		Id:         b.GetNextID(),
		NodeType:   ast_pb.NodeType_INDEX_ACCESS,
	}
}

func (i *IndexAccess) GetId() int64 {
	return i.Id
}

func (i *IndexAccess) GetType() ast_pb.NodeType {
	return i.NodeType
}

func (i *IndexAccess) GetSrc() SrcNode {
	return i.Src
}

func (i *IndexAccess) GetIndexExpression() Node[NodeType] {
	return i.IndexExpression
}

func (i *IndexAccess) GetBaseExpression() Node[NodeType] {
	return i.BaseExpression
}

func (i *IndexAccess) GetTypeDescription() *TypeDescription {
	return nil
}

func (i *IndexAccess) GetTypeDescriptions() []*TypeDescription {
	return i.TypeDescriptions
}

func (i *IndexAccess) GetNodes() []Node[NodeType] {
	return nil
}

func (i *IndexAccess) ToProto() NodeType {
	return ast_pb.IndexAccessExpression{}
}

func (i *IndexAccess) Parse(
	unit *SourceUnit[Node[ast_pb.SourceUnit]],
	contractNode Node[NodeType],
	fnNode Node[NodeType],
	bodyNode *BodyNode,
	vDeclar *VariableDeclaration,
	expNode Node[NodeType],
	ctx *parser.IndexAccessContext,
) {
	i.Src = SrcNode{
		Id:     i.GetNextID(),
		Line:   int64(ctx.GetStart().GetLine()),
		Column: int64(ctx.GetStart().GetColumn()),
		Start:  int64(ctx.GetStart().GetStart()),
		End:    int64(ctx.GetStop().GetStop()),
		Length: int64(ctx.GetStop().GetStop() - ctx.GetStart().GetStart() + 1),
		ParentIndex: func() int64 {
			if vDeclar != nil {
				return vDeclar.GetId()
			}

			if expNode != nil {
				return expNode.GetId()
			}

			return bodyNode.GetId()
		}(),
	}

	expression := NewExpression(i.ASTBuilder)

	i.IndexExpression = expression.Parse(
		unit, contractNode, fnNode, bodyNode, vDeclar, i, ctx.Expression(0),
	)

	i.BaseExpression = expression.Parse(
		unit, contractNode, fnNode, bodyNode, vDeclar, i, ctx.Expression(1),
	)

	i.TypeDescriptions = []*TypeDescription{
		i.IndexExpression.GetTypeDescription(),
		i.BaseExpression.GetTypeDescription(),
	}

	/* 	toReturn.NodeType = ast_pb.NodeType_INDEX_ACCESS

	   	toReturn.BaseExpression = b.parseExpression(
	   		sourceUnit, fnNode, bodyNode, toReturn, toReturn.Id, childCtx.Expression(0),
	   	)
	   	toReturn.IndexExpression = b.parseExpression(
	   		sourceUnit, fnNode, bodyNode, toReturn, toReturn.Id, childCtx.Expression(1),
	   	) */
	i.dumpNode(i)
}
