package ast

import (
	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/parser"
)

type IfStatement struct {
	*ASTBuilder

	Id        int64           `json:"id"`
	NodeType  ast_pb.NodeType `json:"node_type"`
	Src       SrcNode         `json:"src"`
	Condition Node[NodeType]  `json:"condition"`
	Body      Node[NodeType]  `json:"body"`
}

func NewIfStatement(b *ASTBuilder) *IfStatement {
	return &IfStatement{
		ASTBuilder: b,
		Id:         b.GetNextID(),
		NodeType:   ast_pb.NodeType_IF_STATEMENT,
	}
}

// SetReferenceDescriptor sets the reference descriptions of the IfStatement node.
func (i *IfStatement) SetReferenceDescriptor(refId int64, refDesc *TypeDescription) bool {
	return false
}

func (i *IfStatement) GetId() int64 {
	return i.Id
}

func (i *IfStatement) GetType() ast_pb.NodeType {
	return i.NodeType
}

func (i *IfStatement) GetSrc() SrcNode {
	return i.Src
}

func (i *IfStatement) GetCondition() Node[NodeType] {
	return i.Condition
}

func (i *IfStatement) GetTypeDescription() *TypeDescription {
	return nil
}

func (i *IfStatement) GetNodes() []Node[NodeType] {
	return nil
}

func (i *IfStatement) ToProto() NodeType {
	return ast_pb.Statement{}
}

func (i *IfStatement) Parse(
	unit *SourceUnit[Node[ast_pb.SourceUnit]],
	contractNode Node[NodeType],
	fnNode Node[NodeType],
	bodyNode *BodyNode,
	ctx *parser.IfStatementContext,
) Node[NodeType] {
	i.Src = SrcNode{
		Id:          i.GetNextID(),
		Line:        int64(ctx.GetStart().GetLine()),
		Start:       int64(ctx.GetStart().GetStart()),
		End:         int64(ctx.GetStop().GetStop()),
		Length:      int64(ctx.GetStop().GetStop() - ctx.GetStart().GetStart() + 1),
		ParentIndex: bodyNode.Id,
	}

	expression := NewExpression(i.ASTBuilder)

	i.Condition = expression.Parse(unit, contractNode, fnNode, bodyNode, nil, i, ctx.Expression())

	if len(ctx.AllStatement()) > 0 {
		for _, statementCtx := range ctx.AllStatement() {
			if statementCtx.IsEmpty() {
				continue
			}

			if statementCtx.Block() != nil {
				body := NewBodyNode(i.ASTBuilder)
				body.ParseBlock(unit, contractNode, fnNode, statementCtx.Block())
				i.Body = body
				break
			}
		}
	}

	return i
}

/**
func (b *ASTBuilder) parseIfStatement(sourceUnit *ast_pb.SourceUnit, node *ast_pb.Node, bodyNode *ast_pb.Body, ifCtx *parser.IfStatementContext) *ast_pb.Statement {
	statement := &ast_pb.Statement{
		Id: atomic.AddInt64(&b.nextID, 1) - 1,
		Src: &ast_pb.Src{
			Line:        int64(ifCtx.GetStart().GetLine()),
			Start:       int64(ifCtx.GetStart().GetStart()),
			End:         int64(ifCtx.GetStop().GetStop()),
			Length:      int64(ifCtx.GetStop().GetStop() - ifCtx.GetStart().GetStart() + 1),
			ParentIndex: node.Id,
		},
		NodeType: ast_pb.NodeType_IF_STATEMENT,
	}

	condition := b.parseExpression(sourceUnit, node, bodyNode, nil, statement.Id, ifCtx.Expression())
	statement.Condition = condition

	if !ifCtx.IsEmpty() {
		if len(ifCtx.AllStatement()) > 0 {
			for _, statementCtx := range ifCtx.AllStatement() {
				if statementCtx.IsEmpty() {
					continue
				}

				if statementCtx.Block() != nil {
					blockCtx := statementCtx.Block()
					statement.TrueBody = &ast_pb.Statement{
						Id: atomic.AddInt64(&b.nextID, 1) - 1,
						Src: &ast_pb.Src{
							Line:        int64(blockCtx.GetStart().GetLine()),
							Start:       int64(blockCtx.GetStart().GetStart()),
							End:         int64(blockCtx.GetStop().GetStop()),
							Length:      int64(blockCtx.GetStop().GetStop() - blockCtx.GetStart().GetStart() + 1),
							ParentIndex: statement.Id,
						},
						NodeType: ast_pb.NodeType_BLOCK,
					}

					for _, stmtCtx := range statementCtx.Block().AllStatement() {
						statement.TrueBody.Statements = append(
							statement.TrueBody.Statements,
							b.parseStatement(
								sourceUnit, node, bodyNode, statement.TrueBody,
								stmtCtx,
							),
						)
					}
				}
			}
		}
	}

	return statement
}
**/
