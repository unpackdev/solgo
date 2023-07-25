package ast

import (
	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/parser"
)

type Modifier struct {
	*ASTBuilder

	Id       int64           `json:"id"`
	Name     string          `json:"name"`
	NodeType ast_pb.NodeType `json:"node_type"`
	Src      SrcNode         `json:"src"`
}

func NewModifier(b *ASTBuilder) *Modifier {
	return &Modifier{
		ASTBuilder: b,
	}
}

func (m *Modifier) GetId() int64 {
	return m.Id
}

func (m *Modifier) GetType() ast_pb.NodeType {
	return m.NodeType
}

func (m *Modifier) GetSrc() SrcNode {
	return m.Src
}

func (m *Modifier) Parse(unit *SourceUnit[Node[ast_pb.SourceUnit]], fnNode Node[NodeType], ctx parser.IModifierInvocationContext) {
	m.Id = m.GetNextID()
	m.Src = SrcNode{
		Id:          m.GetNextID(),
		Line:        int64(ctx.GetStart().GetLine()),
		Column:      int64(ctx.GetStart().GetColumn()),
		Start:       int64(ctx.GetStart().GetStart()),
		End:         int64(ctx.GetStop().GetStop()),
		Length:      int64(ctx.GetStop().GetStop() - ctx.GetStart().GetStart() + 1),
		ParentIndex: fnNode.GetId(),
	}
	m.NodeType = ast_pb.NodeType_MODIFIER_INVOCATION

	/**
	modifier := &ast_pb.Modifier{
			Id: atomic.AddInt64(&b.nextID, 1) - 1,
			Src: &ast_pb.Src{
				Line:        int64(modifierCtx.GetStart().GetLine()),
				Column:      int64(modifierCtx.GetStart().GetColumn()),
				Start:       int64(modifierCtx.GetStart().GetStart()),
				End:         int64(modifierCtx.GetStop().GetStop()),
				Length:      int64(modifierCtx.GetStop().GetStop() - modifierCtx.GetStart().GetStart() + 1),
				ParentIndex: node.Id,
			},
			NodeType: ast_pb.NodeType_MODIFIER_INVOCATION,
		}

		if modifierCtx.CallArgumentList() != nil {
			for _, argumentCtx := range modifierCtx.CallArgumentList().AllExpression() {
				argument := b.parseExpression(
					sourceUnit, nil, nil, nil, modifier.Id, argumentCtx,
				)
				modifier.Arguments = append(modifier.Arguments, argument)
			}
		}

		identifierCtx := modifierCtx.IdentifierPath()
		if identifierCtx != nil {
			modifier.ModifierName = &ast_pb.ModifierName{
				Id: atomic.AddInt64(&b.nextID, 1) - 1,
				Src: &ast_pb.Src{
					Line:        int64(identifierCtx.GetStart().GetLine()),
					Column:      int64(identifierCtx.GetStart().GetColumn()),
					Start:       int64(identifierCtx.GetStart().GetStart()),
					End:         int64(identifierCtx.GetStop().GetStop()),
					Length:      int64(identifierCtx.GetStop().GetStop() - identifierCtx.GetStart().GetStart() + 1),
					ParentIndex: modifier.Id,
				},
				NodeType: ast_pb.NodeType_IDENTIFIER,
				Name:     identifierCtx.GetText(),
			}
		}
		**/
}
