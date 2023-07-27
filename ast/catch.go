package ast

import (
	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/parser"
)

type CatchStatement struct {
	*ASTBuilder

	Id         int64           `json:"id"`
	Name       string          `json:"name, omitempty"`
	NodeType   ast_pb.NodeType `json:"node_type"`
	Kind       ast_pb.NodeType `json:"kind"`
	Src        SrcNode         `json:"src"`
	Body       *BodyNode       `json:"body"`
	Parameters *ParameterList  `json:"parameters"`
}

func NewCatchClauseStatement(b *ASTBuilder) *CatchStatement {
	return &CatchStatement{
		ASTBuilder: b,
		NodeType:   ast_pb.NodeType_TRY_CATCH_CLAUSE,
		Kind:       ast_pb.NodeType_CATCH,
	}
}

func (t *CatchStatement) GetId() int64 {
	return t.Id
}

func (t *CatchStatement) GetType() ast_pb.NodeType {
	return t.NodeType
}

func (t *CatchStatement) GetSrc() SrcNode {
	return t.Src
}

func (t *CatchStatement) GetBody() *BodyNode {
	return t.Body
}

func (t *CatchStatement) GetKind() ast_pb.NodeType {
	return t.Kind
}

func (t *CatchStatement) GetParameters() *ParameterList {
	return t.Parameters
}

func (t *CatchStatement) GetTypeDescription() *TypeDescription {
	return nil
}

func (t *CatchStatement) GetNodes() []Node[NodeType] {
	return t.Body.Statements
}

func (t *CatchStatement) ToProto() NodeType {
	return ast_pb.Catch{}
}

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
