package ast

import (
	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/parser"
)

type ReceiveDefinition struct {
	*ASTBuilder

	Id               int64             `json:"id"`
	NodeType         ast_pb.NodeType   `json:"node_type"`
	Kind             ast_pb.NodeType   `json:"kind"`
	Src              SrcNode           `json:"src"`
	Implemented      bool              `json:"implemented"`
	Visibility       ast_pb.Visibility `json:"visibility"`
	StateMutability  ast_pb.Mutability `json:"state_mutability"`
	Parameters       *ParameterList    `json:"parameters"`
	ReturnParameters *ParameterList    `json:"return_parameters"`
	Body             *BodyNode         `json:"body"`
	Virtual          bool              `json:"virtual"`
}

func NewReceiveDefinition(b *ASTBuilder) *ReceiveDefinition {
	return &ReceiveDefinition{
		ASTBuilder:      b,
		Id:              b.GetNextID(),
		NodeType:        ast_pb.NodeType_FUNCTION_DEFINITION,
		Kind:            ast_pb.NodeType_RECEIVE,
		StateMutability: ast_pb.Mutability_NONPAYABLE,
	}
}

func (f *ReceiveDefinition) GetId() int64 {
	return f.Id
}

func (f *ReceiveDefinition) GetSrc() SrcNode {
	return f.Src
}

func (f *ReceiveDefinition) GetType() ast_pb.NodeType {
	return f.NodeType
}

func (f *ReceiveDefinition) GetNodes() []Node[NodeType] {
	return f.Body.Statements
}

func (f *ReceiveDefinition) GetTypeDescription() *TypeDescription {
	return nil
}

func (f *ReceiveDefinition) GetParameters() *ParameterList {
	return f.Parameters
}

func (f *ReceiveDefinition) GetReturnParameters() *ParameterList {
	return f.ReturnParameters
}

func (f *ReceiveDefinition) ToProto() NodeType {
	return &ast_pb.Receive{}
}

func (f *ReceiveDefinition) Parse(
	unit *SourceUnit[Node[ast_pb.SourceUnit]],
	contractNode Node[NodeType],
	bodyCtx parser.IContractBodyElementContext,
	ctx *parser.ReceiveFunctionDefinitionContext,
) Node[NodeType] {
	f.Src = SrcNode{
		Id:          f.GetNextID(),
		Line:        int64(ctx.GetStart().GetLine()),
		Column:      int64(ctx.GetStart().GetColumn()),
		Start:       int64(ctx.GetStart().GetStart()),
		End:         int64(ctx.GetStop().GetStop()),
		Length:      int64(ctx.GetStop().GetStop() - ctx.GetStart().GetStart() + 1),
		ParentIndex: contractNode.GetId(),
	}
	f.Implemented = ctx.Block() != nil && !ctx.Block().IsEmpty()

	for _, virtual := range ctx.AllVirtual() {
		if virtual.GetText() == "virtual" {
			f.Virtual = true
		}
	}

	f.Visibility = f.getVisibilityFromCtx(ctx)
	f.StateMutability = f.getStateMutabilityFromCtx(ctx)

	params := NewParameterList(f.ASTBuilder)
	params.Src = f.Src
	params.Src.ParentIndex = f.Id
	f.Parameters = params

	returnParams := NewParameterList(f.ASTBuilder)
	returnParams.Src = f.Src
	returnParams.Src.ParentIndex = f.Id
	f.ReturnParameters = returnParams

	if ctx.Block() != nil && !ctx.Block().IsEmpty() {
		bodyNode := NewBodyNode(f.ASTBuilder)
		bodyNode.ParseBlock(unit, contractNode, f, ctx.Block())
		f.Body = bodyNode

		if ctx.Block().AllUncheckedBlock() != nil {
			for _, uncheckedCtx := range ctx.Block().AllUncheckedBlock() {
				bodyNode := NewBodyNode(f.ASTBuilder)
				bodyNode.ParseUncheckedBlock(unit, contractNode, f, uncheckedCtx)
				f.Body.Statements = append(f.Body.Statements, bodyNode)
			}
		}
	}

	return f
}

func (f *ReceiveDefinition) getVisibilityFromCtx(ctx *parser.ReceiveFunctionDefinitionContext) ast_pb.Visibility {
	for _, visibility := range ctx.AllExternal() {
		if visibility.GetText() == "external" {
			f.Visibility = ast_pb.Visibility_EXTERNAL
		}
	}

	return ast_pb.Visibility_INTERNAL
}

func (f *ReceiveDefinition) getStateMutabilityFromCtx(ctx *parser.ReceiveFunctionDefinitionContext) ast_pb.Mutability {
	for _, stateMutability := range ctx.AllPayable() {
		if stateMutability.GetText() == "payable" {
			f.StateMutability = ast_pb.Mutability_PAYABLE
		}
	}

	return ast_pb.Mutability_NONPAYABLE
}
