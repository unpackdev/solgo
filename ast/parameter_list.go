package ast

import (
	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/parser"
)

type ParameterList[T NodeType] struct {
	*ASTBuilder

	Id         int64           `json:"id"`
	NodeType   ast_pb.NodeType `json:"node_type"`
	Src        SrcNode         `json:"src"`
	Parameters []*Parameter    `json:"parameters"`
}

func NewParameterList[T any](b *ASTBuilder) *ParameterList[T] {
	return &ParameterList[T]{
		ASTBuilder: b,
		NodeType:   ast_pb.NodeType_PARAMETER_LIST,
		Parameters: make([]*Parameter, 0),
	}
}

func (p *ParameterList[T]) GetId() int64 {
	return p.Id
}

func (p *ParameterList[T]) GetType() ast_pb.NodeType {
	return p.NodeType
}

func (p *ParameterList[T]) GetSrc() SrcNode {
	return p.Src
}

func (p *ParameterList[T]) GetTypeDescription() *TypeDescription {
	return nil
}

func (p *ParameterList[T]) GetParameters() []*Parameter {
	return p.Parameters
}

func (p *ParameterList[T]) GetNodes() []Node[NodeType] {
	return nil
}

func (p *ParameterList[T]) ToProto() ast_pb.ParametersList {
	return ast_pb.ParametersList{}
}

func (p *ParameterList[T]) Parse(unit *SourceUnit[Node[ast_pb.SourceUnit]], fNode Node[NodeType], ctx parser.IParameterListContext) {
	p.Id = p.GetNextID()
	p.Src = SrcNode{
		Id:          p.GetNextID(),
		Line:        int64(ctx.GetStart().GetLine()),
		Column:      int64(ctx.GetStart().GetColumn()),
		Start:       int64(ctx.GetStart().GetStart()),
		End:         int64(ctx.GetStop().GetStop()),
		Length:      int64(ctx.GetStop().GetStop() - ctx.GetStart().GetStart() + 1),
		ParentIndex: fNode.GetId(),
	}

	// No need to move forwards as there are no parameters to parse in this context.
	if ctx == nil || ctx.IsEmpty() {
		return
	}

	for _, paramCtx := range ctx.AllParameterDeclaration() {
		param := NewParameter(p.ASTBuilder)
		param.Parse(unit, fNode, p, paramCtx.(*parser.ParameterDeclarationContext))
		p.Parameters = append(p.Parameters, param)
	}
}

func (p *ParameterList[T]) ParseEventParameters(unit *SourceUnit[Node[ast_pb.SourceUnit]], eNode Node[NodeType], ctx []parser.IEventParameterContext) {
	p.Id = p.GetNextID()
	p.Src = eNode.GetSrc()
	p.Src.ParentIndex = eNode.GetId()

	for _, paramCtx := range ctx {
		param := NewParameter(p.ASTBuilder)
		param.ParseEventParameter(unit, eNode, p, paramCtx)
		p.Parameters = append(p.Parameters, param)
	}
}

func (p *ParameterList[T]) ParseErrorParameters(unit *SourceUnit[Node[ast_pb.SourceUnit]], eNode Node[NodeType], ctx []parser.IErrorParameterContext) {
	p.Id = p.GetNextID()
	p.Src = eNode.GetSrc()
	p.Src.ParentIndex = eNode.GetId()

	for _, paramCtx := range ctx {
		param := NewParameter(p.ASTBuilder)
		param.ParseErrorParameter(unit, eNode, p, paramCtx)
		p.Parameters = append(p.Parameters, param)
	}
}
