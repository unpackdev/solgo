package ast

import (
	"fmt"
	"reflect"

	"github.com/antlr4-go/antlr/v4"
	v3 "github.com/cncf/xds/go/xds/type/v3"
	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/parser"
)

// BodyNode represents a body node in the abstract syntax tree.
// It includes various attributes like id, node type, kind, source node, implemented status, and statements.
type BodyNode struct {
	*ASTBuilder

	Id          int64            `json:"id"`          // Id is the unique identifier of the body node.
	NodeType    ast_pb.NodeType  `json:"node_type"`   // NodeType is the type of the AST node.
	Kind        ast_pb.NodeType  `json:"kind"`        // Kind is the kind of the AST node.
	Src         SrcNode          `json:"src"`         // Src is the source code location.
	Implemented bool             `json:"implemented"` // Implemented indicates whether the function is implemented.
	Statements  []Node[NodeType] `json:"statements"`  // Statements is the list of AST nodes in the body.
}

// NewBodyNode creates a new BodyNode with the provided ASTBuilder.
// It returns a pointer to the created BodyNode.
func NewBodyNode(b *ASTBuilder) *BodyNode {
	return &BodyNode{
		ASTBuilder: b,
		Id:         b.GetNextID(),
		NodeType:   ast_pb.NodeType_BLOCK,
		Statements: make([]Node[NodeType], 0),
	}
}

// SetReferenceDescriptor sets the reference descriptions of the BodyNode node.
func (b *BodyNode) SetReferenceDescriptor(refId int64, refDesc *TypeDescription) bool {
	return false
}

// GetId returns the unique identifier of the body node.
func (b *BodyNode) GetId() int64 {
	return b.Id
}

// GetType returns the type of the body node.
func (b *BodyNode) GetType() ast_pb.NodeType {
	return b.NodeType
}

// GetSrc returns the source code location of the body node.
func (b *BodyNode) GetSrc() SrcNode {
	return b.Src
}

// GetStatements returns the statements associated with the body node.
func (b *BodyNode) GetStatements() []Node[NodeType] {
	return b.Statements
}

// GetKind returns the kind of the body node.
func (b *BodyNode) GetKind() ast_pb.NodeType {
	return b.Kind
}

// IsImplemented returns the implemented status of the body node.
func (b *BodyNode) IsImplemented() bool {
	return b.Implemented
}

// GetTypeDescription returns the type description of the body node.
// As BodyNode does not have a type description, it returns nil.
func (b *BodyNode) GetTypeDescription() *TypeDescription {
	return &TypeDescription{
		TypeString:     "block",
		TypeIdentifier: "$_t_block",
	}
}

// GetNodes returns the nodes associated with the body node.
func (b *BodyNode) GetNodes() []Node[NodeType] {
	return b.Statements
}

// ToProto converts the BodyNode to a protocol buffer representation.
func (b *BodyNode) ToProto() NodeType {
	proto := ast_pb.Body{
		Id:          b.GetId(),
		NodeType:    b.GetType(),
		Kind:        b.GetKind(),
		Implemented: b.IsImplemented(),
		Src:         b.GetSrc().ToProto(),
		Statements:  make([]*v3.TypedStruct, 0),
	}

	for _, statement := range b.GetStatements() {
		proto.Statements = append(proto.Statements, statement.ToProto().(*v3.TypedStruct))
	}

	switch b.NodeType {
	case ast_pb.NodeType_UNCHECKED_BLOCK:
		return NewTypedStruct(&proto, "Block")
	default:
		return &proto
	}
}

// ParseDefinitions is a method of the BodyNode struct. It parses the definitions of a contract body element context.
// It takes a source unit, a contract node, and a contract body element context as arguments.
// It iterates over the children of the body context, and based on the type of each child, it creates a new node of the corresponding type and parses it.
// It then returns the newly created and parsed node.
// If the type of the child context is unknown, it panics and prints an error message.
// Panic is here so we are forced to implement missing functionality.
// After parsing all the children, it sets the source node of the BodyNode and returns the BodyNode itself.
func (b *BodyNode) ParseDefinitions(
	unit *SourceUnit[Node[ast_pb.SourceUnit]],
	contractNode Node[NodeType],
	bodyCtx parser.IContractBodyElementContext,
) Node[NodeType] {
	// We are considering function implemented in case that there's really anything defined in the body.
	// This is a basic approach and it's not 100% correct, but it's good enough for now.
	b.Implemented = len(bodyCtx.GetChildren()) > 0

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
			statement := NewFunction(b.ASTBuilder)
			return statement.Parse(unit, contractNode, bodyCtx, childCtx)
		case *parser.ModifierDefinitionContext:
			statement := NewModifierDefinition(b.ASTBuilder)
			return statement.ParseDefinition(unit, contractNode, bodyCtx, childCtx)
		case *parser.FallbackFunctionDefinitionContext:
			statement := NewFallbackDefinition(b.ASTBuilder)
			return statement.Parse(unit, contractNode, bodyCtx, childCtx)
		case *parser.ReceiveFunctionDefinitionContext:
			statement := NewReceiveDefinition(b.ASTBuilder)
			return statement.Parse(unit, contractNode, bodyCtx, childCtx)
		default:
			panic(fmt.Sprintf("Unknown body child type @ BodyNode.Parse: %s", reflect.TypeOf(childCtx)))
		}
	}

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

// ParseBlock is a method of the BodyNode struct. It parses a block context.
// It takes a source unit, a contract node, a function node, and a block context as arguments.
// It sets the source node of the BodyNode and checks if the function is implemented by checking if there are any children in the block context.
// It then iterates over all the statements in the block context and parses each one by calling the parseStatements helper function.
// It finally returns the BodyNode itself.
func (b *BodyNode) ParseBlock(
	unit *SourceUnit[Node[ast_pb.SourceUnit]],
	contractNode Node[NodeType],
	fnNode Node[NodeType],
	bodyCtx parser.IBlockContext,
) Node[NodeType] {
	b.Src = SrcNode{
		Id:          b.GetNextID(),
		Line:        int64(bodyCtx.GetStart().GetLine()),
		Column:      int64(bodyCtx.GetStart().GetColumn()),
		Start:       int64(bodyCtx.GetStart().GetStart()),
		End:         int64(bodyCtx.GetStop().GetStop()),
		Length:      int64(bodyCtx.GetStop().GetStop() - bodyCtx.GetStart().GetStart() + 1),
		ParentIndex: fnNode.GetId(),
	}

	// We are considering function implemented in case that there's really anything defined in the body.
	// This is a basic approach and it's not 100% correct, but it's good enough for now.
	b.Implemented = len(bodyCtx.GetChildren()) > 0

	for _, statementCtx := range bodyCtx.AllStatement() {
		for _, child := range statementCtx.GetChildren() {
			b.parseStatements(unit, contractNode, fnNode, child)
		}
	}

	return b
}

// ParseUncheckedBlock is a method of the BodyNode struct. It parses an unchecked block context.
// It takes a source unit, a contract node, a function node, and an unchecked block context as arguments.
// It sets the node type of the BodyNode to UNCHECKED_BLOCK and sets its source node.
// It then iterates over all the statements in the block context and parses each one by calling the parseStatements helper function.
// It finally returns the BodyNode itself.
func (b *BodyNode) ParseUncheckedBlock(
	unit *SourceUnit[Node[ast_pb.SourceUnit]],
	contractNode Node[NodeType],
	fnNode Node[NodeType],
	bodyCtx parser.IUncheckedBlockContext,
) Node[NodeType] {
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
			b.parseStatements(unit, contractNode, fnNode, child)
		}
	}

	return b
}

// parseStatements is a helper function for the ParseBlock and ParseUncheckedBlock methods.
// It takes a source unit, a contract node, a function node, and a child context as arguments.
// It checks the type of the child context and based on its type, it creates a new node of the corresponding type and parses it.
// If the type of the child context is unknown, it panics and prints an error message.
// Panic is here so we are forced to implement missing functionality.
func (b *BodyNode) parseStatements(
	unit *SourceUnit[Node[ast_pb.SourceUnit]],
	contractNode Node[NodeType],
	fnNode Node[NodeType],
	childCtx antlr.Tree,
) {
	switch childCtx := childCtx.(type) {
	case *parser.ConstructorDefinitionContext:
		statement := NewConstructor(b.ASTBuilder)
		b.Statements = append(b.Statements, statement.Parse(
			unit, contractNode, childCtx,
		))
	case *parser.SimpleStatementContext:
		statement := NewSimpleStatement(b.ASTBuilder)
		b.Statements = append(b.Statements, statement.Parse(
			unit, contractNode, fnNode, b, nil, childCtx,
		))
	case *parser.EmitStatementContext:
		statement := NewEmitStatement(b.ASTBuilder)
		b.Statements = append(b.Statements, statement.Parse(
			unit, contractNode, fnNode, b, childCtx,
		))
	case *parser.ForStatementContext:
		statement := NewForStatement(b.ASTBuilder)
		b.Statements = append(b.Statements, statement.Parse(
			unit, contractNode, fnNode, b, childCtx,
		))
	case *parser.IfStatementContext:
		statement := NewIfStatement(b.ASTBuilder)
		b.Statements = append(b.Statements, statement.Parse(
			unit, contractNode, fnNode, b, childCtx,
		))
	case *parser.DoWhileStatementContext:
		statement := NewDoWhileStatement(b.ASTBuilder)
		b.Statements = append(b.Statements, statement.Parse(
			unit, contractNode, fnNode, b, childCtx,
		))
	case *parser.TryStatementContext:
		statement := NewTryStatement(b.ASTBuilder)
		b.Statements = append(b.Statements, statement.Parse(
			unit, contractNode, fnNode, b, childCtx,
		))
	case *parser.WhileStatementContext:
		statement := NewWhileStatement(b.ASTBuilder)
		b.Statements = append(b.Statements, statement.Parse(
			unit, contractNode, fnNode, b, childCtx,
		))
	case *parser.BreakStatementContext:
		statement := NewBreakStatement(b.ASTBuilder)
		b.Statements = append(b.Statements, statement.Parse(
			unit, contractNode, fnNode, b, childCtx,
		))
	case *parser.ContinueStatementContext:
		statement := NewContinueStatement(b.ASTBuilder)
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
	case *parser.AssemblyStatementContext:
		statement := NewAssemblyStatement(b.ASTBuilder)
		b.Statements = append(b.Statements, statement.Parse(
			unit, contractNode, fnNode, b, childCtx,
		))
	default:
		panic(fmt.Sprintf("Unknown body statement type @ BodyNode.ParseUncheckedBlock: %s", reflect.TypeOf(childCtx)))
	}
}
