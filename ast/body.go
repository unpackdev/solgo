package ast

import (
	"fmt"
	"reflect"

	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/parser"
)

type BodyNode struct {
	*ASTBuilder

	Id          int64            `json:"id"`
	NodeType    ast_pb.NodeType  `json:"node_type"`
	Kind        ast_pb.NodeType  `json:"kind,omitempty"`
	Src         SrcNode          `json:"src"`
	Implemented bool             `json:"implemented"`
	Statements  []Node[NodeType] `json:"statements"`
}

func NewBodyNode(b *ASTBuilder) *BodyNode {
	return &BodyNode{
		ASTBuilder: b,
		Statements: make([]Node[NodeType], 0),
	}
}

func (b *BodyNode) GetId() int64 {
	return b.Id
}

func (b *BodyNode) GetType() ast_pb.NodeType {
	return b.NodeType
}

func (b *BodyNode) GetSrc() SrcNode {
	return b.Src
}

func (b *BodyNode) GetStatements() []Node[NodeType] {
	return b.Statements
}

func (b *BodyNode) GetKind() ast_pb.NodeType {
	return b.Kind
}

func (b *BodyNode) GetImplemented() bool {
	return b.Implemented
}

func (b *BodyNode) GetTypeDescription() *TypeDescription {
	return nil
}

func (b *BodyNode) GetNodes() []Node[NodeType] {
	return b.Statements
}

func (b *BodyNode) ToProto() NodeType {
	return ast_pb.Body{}
}

func (b *BodyNode) Parse(
	unit *SourceUnit[Node[ast_pb.SourceUnit]],
	contractNode Node[NodeType],
	bodyCtx parser.IContractBodyElementContext,
) Node[NodeType] {
	for _, bodyChildCtx := range bodyCtx.GetChildren() {
		switch childCtx := bodyChildCtx.(type) {
		case *parser.UsingDirectiveContext:
			using := NewUsingDirective(b.ASTBuilder)
			using.Parse(unit, contractNode, bodyCtx, childCtx)
			return using
		case *parser.StateVariableDeclarationContext:
			stateVar := NewStateVariableDeclaration(b.ASTBuilder)
			stateVar.Parse(unit, contractNode, bodyCtx, childCtx)
			return stateVar
		case *parser.EventDefinitionContext:
			event := NewEventDefinition(b.ASTBuilder)
			return event.Parse(unit, contractNode, bodyCtx, childCtx)
		case *parser.EnumDefinitionContext:
			enum := NewEnumDefinition(b.ASTBuilder)
			return enum.Parse(unit, contractNode, bodyCtx, childCtx)
		case *parser.StructDefinitionContext:
			structDef := NewStructDefinition(b.ASTBuilder)
			return structDef.Parse(unit, contractNode, bodyCtx, childCtx)
		case *parser.ErrorDefinitionContext:
			errorDef := NewErrorDefinition(b.ASTBuilder)
			return errorDef.Parse(unit, contractNode, bodyCtx, childCtx)
		case *parser.ConstructorDefinitionContext:
			statement := NewConstructor(b.ASTBuilder)
			return statement.Parse(unit, contractNode, childCtx)
		case *parser.FunctionDefinitionContext:
			statement := NewFunctionNode(b.ASTBuilder)
			return statement.Parse(unit, contractNode, bodyCtx, childCtx)
		case *parser.ModifierDefinitionContext:
			statement := NewModifierDefinition(b.ASTBuilder)
			return statement.ParseDefinition(unit, contractNode, bodyCtx, childCtx)
		case *parser.FallbackFunctionDefinitionContext:
			statement := NewFallbackDefinition(b.ASTBuilder)
			return statement.Parse(unit, contractNode, bodyCtx, childCtx)
		default:
			panic(fmt.Sprintf("Unknown body child type @ BodyNode.Parse: %s", reflect.TypeOf(childCtx)))
		}
	}

	// Could not find any function definitions so we'll just return the body node.
	b.Id = b.GetNextID()
	b.Src = SrcNode{
		Id:          b.GetNextID(),
		Line:        int64(bodyCtx.GetStart().GetLine()),
		Column:      int64(bodyCtx.GetStart().GetColumn()),
		Start:       int64(bodyCtx.GetStart().GetStart()),
		End:         int64(bodyCtx.GetStop().GetStop()),
		Length:      int64(bodyCtx.GetStop().GetStop() - bodyCtx.GetStart().GetStart() + 1),
		ParentIndex: contractNode.GetId(),
	}
	return b
}

func (b *BodyNode) ParseBlock(
	unit *SourceUnit[Node[ast_pb.SourceUnit]],
	contractNode Node[NodeType],
	fnNode Node[NodeType],
	bodyCtx parser.IBlockContext,
) Node[NodeType] {
	// Could not find any function definitions so we'll just return the body node.
	b.Id = b.GetNextID()
	b.NodeType = ast_pb.NodeType_BLOCK
	b.Src = SrcNode{
		Id:          b.GetNextID(),
		Line:        int64(bodyCtx.GetStart().GetLine()),
		Column:      int64(bodyCtx.GetStart().GetColumn()),
		Start:       int64(bodyCtx.GetStart().GetStart()),
		End:         int64(bodyCtx.GetStop().GetStop()),
		Length:      int64(bodyCtx.GetStop().GetStop() - bodyCtx.GetStart().GetStart() + 1),
		ParentIndex: contractNode.GetId(),
	}

	for _, statementCtx := range bodyCtx.AllStatement() {
		for _, child := range statementCtx.GetChildren() {
			switch childCtx := child.(type) {
			case *parser.ConstructorDefinitionContext:
				statement := NewConstructor(b.ASTBuilder)
				b.Statements = append(b.Statements, statement.Parse(
					unit, contractNode, childCtx,
				))
			case *parser.SimpleStatementContext:
				statement := NewSimpleStatement[NodeType](b.ASTBuilder)
				b.Statements = append(b.Statements, statement.Parse(
					unit, contractNode, fnNode, b, childCtx,
				))
			case *parser.EmitStatementContext:
				statement := NewEmitStatement(b.ASTBuilder)
				b.Statements = append(b.Statements, statement.Parse(
					unit, contractNode, fnNode, b, childCtx,
				))
			case *parser.IfStatementContext:
				statement := NewIfStatement(b.ASTBuilder)
				b.Statements = append(b.Statements, statement.Parse(
					unit, contractNode, fnNode, b, childCtx,
				))
			case *parser.ReturnStatementContext:
				statement := NewReturnStatement(b.ASTBuilder)
				b.Statements = append(b.Statements, statement.Parse(
					unit, contractNode, fnNode, b, childCtx,
				))
			case *parser.RevertStatementContext:
				statement := NewRevertStatement(b.ASTBuilder)
				b.Statements = append(b.Statements, statement.Parse(
					unit, contractNode, fnNode, b, childCtx,
				))
			default:
				panic(fmt.Sprintf("Unknown statement type @ BodyNode.ParseBlock: %s", reflect.TypeOf(childCtx)))
			}
		}
	}

	return b
}

func (b *BodyNode) ParseUncheckedBlock(
	unit *SourceUnit[Node[ast_pb.SourceUnit]],
	contractNode Node[NodeType],
	fnNode Node[NodeType],
	bodyCtx parser.IUncheckedBlockContext,
) Node[NodeType] {
	// Could not find any function definitions so we'll just return the body node.
	b.Id = b.GetNextID()
	b.NodeType = ast_pb.NodeType_UNCHECKED_BLOCK
	b.Src = SrcNode{
		Id:          b.GetNextID(),
		Line:        int64(bodyCtx.GetStart().GetLine()),
		Column:      int64(bodyCtx.GetStart().GetColumn()),
		Start:       int64(bodyCtx.GetStart().GetStart()),
		End:         int64(bodyCtx.GetStop().GetStop()),
		Length:      int64(bodyCtx.GetStop().GetStop() - bodyCtx.GetStart().GetStart() + 1),
		ParentIndex: contractNode.GetId(),
	}

	for _, statementCtx := range bodyCtx.Block().AllStatement() {
		for _, child := range statementCtx.GetChildren() {
			switch childCtx := child.(type) {
			case *parser.ConstructorDefinitionContext:
				statement := NewConstructor(b.ASTBuilder)
				b.Statements = append(b.Statements, statement.Parse(
					unit, contractNode, childCtx,
				))
			case *parser.SimpleStatementContext:
				statement := NewSimpleStatement[NodeType](b.ASTBuilder)
				b.Statements = append(b.Statements, statement.Parse(
					unit, contractNode, fnNode, b, childCtx,
				))
			case *parser.EmitStatementContext:
				statement := NewEmitStatement(b.ASTBuilder)
				b.Statements = append(b.Statements, statement.Parse(
					unit, contractNode, fnNode, b, childCtx,
				))
			case *parser.IfStatementContext:
				statement := NewIfStatement(b.ASTBuilder)
				b.Statements = append(b.Statements, statement.Parse(
					unit, contractNode, fnNode, b, childCtx,
				))
			case *parser.ReturnStatementContext:
				statement := NewReturnStatement(b.ASTBuilder)
				b.Statements = append(b.Statements, statement.Parse(
					unit, contractNode, fnNode, b, childCtx,
				))
			default:
				panic(fmt.Sprintf("Unknown statement type @ BodyNode.ParseUncheckedBlock: %s", reflect.TypeOf(childCtx)))
			}
		}
	}

	return b
}
