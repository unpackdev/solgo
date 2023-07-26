package ast

import (
	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/parser"
)

type FunctionNode[T ast_pb.Function] struct {
	*ASTBuilder

	Id               int64               `json:"id"`
	Name             string              `json:"name"`
	NodeType         ast_pb.NodeType     `json:"node_type"`
	Kind             ast_pb.NodeType     `json:"kind"`
	Src              SrcNode             `json:"src"`
	Body             *BodyNode           `json:"body"`
	Implemented      bool                `json:"implemented"`
	Visibility       ast_pb.Visibility   `json:"visibility"`
	StateMutability  ast_pb.Mutability   `json:"state_mutability"`
	Virtual          bool                `json:"virtual"`
	Modifiers        []Modifier          `json:"modifiers"`
	Overrides        []OverrideSpecifier `json:"overrides"`
	Parameters       *ParameterList[T]   `json:"parameters"`
	ReturnParameters *ParameterList[T]   `json:"return_parameters"`
	Scope            int64               `json:"scope"`
}

func NewFunctionNode[T ast_pb.Function](b *ASTBuilder) *FunctionNode[T] {
	return &FunctionNode[T]{
		ASTBuilder: b,
		NodeType:   ast_pb.NodeType_FUNCTION_DEFINITION,
		Kind:       ast_pb.NodeType_KIND_FUNCTION,
		Modifiers:  make([]Modifier, 0),
		Overrides:  make([]OverrideSpecifier, 0),
	}
}

func (f FunctionNode[T]) GetId() int64 {
	return f.Id
}

func (f FunctionNode[T]) GetType() ast_pb.NodeType {
	return f.NodeType
}

func (f FunctionNode[T]) GetSrc() SrcNode {
	return f.Src
}

func (f FunctionNode[T]) GetParameters() *ParameterList[T] {
	return f.Parameters
}

func (f FunctionNode[T]) GetReturnParameters() *ParameterList[T] {
	return f.ReturnParameters
}

func (f FunctionNode[T]) GetBody() *BodyNode {
	return f.Body
}

func (f FunctionNode[T]) GetKind() ast_pb.NodeType {
	return f.Kind
}

func (f FunctionNode[T]) IsImplemented() bool {
	return f.Implemented
}

func (f FunctionNode[T]) GetModifiers() []Modifier {
	return f.Modifiers
}

func (f FunctionNode[T]) GetOverrides() []OverrideSpecifier {
	return f.Overrides
}

func (f FunctionNode[T]) GetVisibility() ast_pb.Visibility {
	return f.Visibility
}

func (f FunctionNode[T]) GetStateMutability() ast_pb.Mutability {
	return f.StateMutability
}

func (f FunctionNode[T]) IsVirtual() bool {
	return f.Virtual
}

func (f FunctionNode[T]) GetScope() int64 {
	return f.Scope
}

func (f FunctionNode[T]) GetName() string {
	return f.Name
}

func (f FunctionNode[T]) GetTypeDescription() *TypeDescription {
	return nil
}

func (f FunctionNode[T]) GetNodes() []Node[NodeType] {
	return f.Body.GetNodes()
}

func (f FunctionNode[T]) ToProto() NodeType {
	return ast_pb.Function{}
}

func (f FunctionNode[T]) Parse(
	unit *SourceUnit[Node[ast_pb.SourceUnit]],
	contractNode Node[NodeType],
	bodyCtx parser.IContractBodyElementContext,
	ctx *parser.FunctionDefinitionContext,
) Node[NodeType] {
	f.Id = f.GetNextID()
	f.Scope = contractNode.GetId()
	if ctx.Identifier() != nil {
		f.Name = ctx.Identifier().GetText()
	}
	f.Implemented = ctx.Block() != nil && !ctx.Block().IsEmpty()
	f.Src = SrcNode{
		Id:          f.GetNextID(),
		Line:        int64(ctx.GetStart().GetLine()),
		Column:      int64(ctx.GetStart().GetColumn()),
		Start:       int64(ctx.GetStart().GetStart()),
		End:         int64(ctx.GetStop().GetStop()),
		Length:      int64(ctx.GetStop().GetStop() - ctx.GetStart().GetStart() + 1),
		ParentIndex: contractNode.GetId(),
	}

	// Set function visibility state.
	f.Visibility = f.getVisibilityFromCtx(ctx)

	// Set function state mutability.
	f.StateMutability = f.getStateMutabilityFromCtx(ctx)

	// Set if function is virtual.
	for _, virtual := range ctx.AllVirtual() {
		if virtual.GetText() == "virtual" {
			f.Virtual = true
			break
		}
	}

	// Set function modifiers.
	for _, modifierCtx := range ctx.AllModifierInvocation() {
		modifier := NewModifier(f.ASTBuilder)
		modifier.Parse(unit, f, modifierCtx)
		f.Modifiers = append(f.Modifiers, *modifier)
	}

	// Set function override specifier.
	for _, overrideCtx := range ctx.AllOverrideSpecifier() {
		overrideSpecifier := NewOverrideSpecifier(f.ASTBuilder)
		overrideSpecifier.Parse(unit, f, overrideCtx)
		f.Overrides = append(f.Overrides, *overrideSpecifier)
	}

	// Set function parameters if they exist.
	if len(ctx.AllParameterList()) > 0 {
		params := NewParameterList[T](f.ASTBuilder)
		params.Parse(unit, f, ctx.AllParameterList()[0])
		f.Parameters = params
	}

	// Set function return parameters if they exist.
	// @TODO: Consider traversing through body to discover name of the return parameters even
	// if they are not defined in (name uint) format.
	if ctx.GetReturnParameters() != nil {
		returnParams := NewParameterList[T](f.ASTBuilder)
		returnParams.Parse(unit, f, ctx.GetReturnParameters())
		f.ReturnParameters = returnParams
	}

	// And now we are going to the big league. We are going to traverse the function body.
	if ctx.Block() != nil && !ctx.Block().IsEmpty() {

		bodyNode := NewBodyNode(f.ASTBuilder)
		bodyNode.ParseBlock(unit, contractNode, f, ctx.Block())
		f.Body = bodyNode

		/* 		bodyNode := &ast_pb.Body{
		   			Id: atomic.AddInt64(&b.nextID, 1) - 1,
		   			Src: &ast_pb.Src{
		   				Line:        int64(fd.Block().GetStart().GetLine()),
		   				Column:      int64(fd.Block().GetStart().GetColumn()),
		   				Start:       int64(fd.Block().GetStart().GetStart()),
		   				End:         int64(fd.Block().GetStop().GetStop()),
		   				Length:      int64(fd.Block().GetStop().GetStop() - fd.Block().GetStart().GetStart() + 1),
		   				ParentIndex: node.Id,
		   			},
		   			NodeType: ast_pb.NodeType_BLOCK,
		   		}

		   		for _, statement := range fd.Block().AllStatement() {
		   			if statement.IsEmpty() {
		   				continue
		   			}

		   			// Parent index statement in this case is used only to be able provide
		   			// index to the parent node. It is not used for anything else.
		   			parentIndexStmt := &ast_pb.Statement{Id: bodyNode.Id}

		   			bodyNode.Statements = append(bodyNode.Statements, b.parseStatement(
		   				sourceUnit, node, bodyNode, parentIndexStmt, statement,
		   			))
		   		}

		   		node.Body = bodyNode */
	}

	/* 	if fd.Block() != nil && len(fd.Block().AllUncheckedBlock()) > 0 {
		for _, uncheckedBlockCtx := range fd.Block().AllUncheckedBlock() {
			bodyNode := &ast_pb.Body{
				Id: atomic.AddInt64(&b.nextID, 1) - 1,
				Src: &ast_pb.Src{
					Line:        int64(fd.Block().GetStart().GetLine()),
					Column:      int64(fd.Block().GetStart().GetColumn()),
					Start:       int64(fd.Block().GetStart().GetStart()),
					End:         int64(fd.Block().GetStop().GetStop()),
					Length:      int64(fd.Block().GetStop().GetStop() - fd.Block().GetStart().GetStart() + 1),
					ParentIndex: node.Id,
				},
				NodeType: ast_pb.NodeType_UNCHECKED_BLOCK,
			}

			if uncheckedBlockCtx.Block() != nil && !uncheckedBlockCtx.Block().IsEmpty() {
				for _, statement := range uncheckedBlockCtx.Block().AllStatement() {
					if statement.IsEmpty() {
						continue
					}

					// Parent index statement in this case is used only to be able provide
					// index to the parent node. It is not used for anything else.
					parentIndexStmt := &ast_pb.Statement{Id: bodyNode.Id}

					bodyNode.Statements = append(bodyNode.Statements, b.parseStatement(
						sourceUnit, node, bodyNode, parentIndexStmt, statement,
					))
				}
			}

			node.Body = bodyNode
		}
	} */

	return f
}

func (f FunctionNode[T]) getVisibilityFromCtx(ctx *parser.FunctionDefinitionContext) ast_pb.Visibility {
	visibilityMap := map[string]ast_pb.Visibility{
		"public":   ast_pb.Visibility_PUBLIC,
		"private":  ast_pb.Visibility_PRIVATE,
		"internal": ast_pb.Visibility_INTERNAL,
		"external": ast_pb.Visibility_EXTERNAL,
	}

	for _, visibility := range ctx.AllVisibility() {
		if v, ok := visibilityMap[visibility.GetText()]; ok {
			return v
		}
	}

	return ast_pb.Visibility_INTERNAL
}

func (f FunctionNode[T]) getStateMutabilityFromCtx(ctx *parser.FunctionDefinitionContext) ast_pb.Mutability {
	mutabilityMap := map[string]ast_pb.Mutability{
		"payable": ast_pb.Mutability_PAYABLE,
		"pure":    ast_pb.Mutability_PURE,
		"view":    ast_pb.Mutability_VIEW,
	}

	for _, stateMutability := range ctx.AllStateMutability() {
		if m, ok := mutabilityMap[stateMutability.GetText()]; ok {
			return m
		}
	}

	return ast_pb.Mutability_NONPAYABLE
}

/**
func (b *ASTBuilder) parseFunctionDefinition(sourceUnit *ast_pb.SourceUnit, node *ast_pb.Node, fd *parser.FunctionDefinitionContext) *ast_pb.Node {

	// And now we are going to the big league. We are going to traverse the function body.
	if fd.Block() != nil && !fd.Block().IsEmpty() {
		bodyNode := &ast_pb.Body{
			Id: atomic.AddInt64(&b.nextID, 1) - 1,
			Src: &ast_pb.Src{
				Line:        int64(fd.Block().GetStart().GetLine()),
				Column:      int64(fd.Block().GetStart().GetColumn()),
				Start:       int64(fd.Block().GetStart().GetStart()),
				End:         int64(fd.Block().GetStop().GetStop()),
				Length:      int64(fd.Block().GetStop().GetStop() - fd.Block().GetStart().GetStart() + 1),
				ParentIndex: node.Id,
			},
			NodeType: ast_pb.NodeType_BLOCK,
		}

		for _, statement := range fd.Block().AllStatement() {
			if statement.IsEmpty() {
				continue
			}

			// Parent index statement in this case is used only to be able provide
			// index to the parent node. It is not used for anything else.
			parentIndexStmt := &ast_pb.Statement{Id: bodyNode.Id}

			bodyNode.Statements = append(bodyNode.Statements, b.parseStatement(
				sourceUnit, node, bodyNode, parentIndexStmt, statement,
			))
		}

		node.Body = bodyNode
	}

	if fd.Block() != nil && len(fd.Block().AllUncheckedBlock()) > 0 {
		for _, uncheckedBlockCtx := range fd.Block().AllUncheckedBlock() {
			bodyNode := &ast_pb.Body{
				Id: atomic.AddInt64(&b.nextID, 1) - 1,
				Src: &ast_pb.Src{
					Line:        int64(fd.Block().GetStart().GetLine()),
					Column:      int64(fd.Block().GetStart().GetColumn()),
					Start:       int64(fd.Block().GetStart().GetStart()),
					End:         int64(fd.Block().GetStop().GetStop()),
					Length:      int64(fd.Block().GetStop().GetStop() - fd.Block().GetStart().GetStart() + 1),
					ParentIndex: node.Id,
				},
				NodeType: ast_pb.NodeType_UNCHECKED_BLOCK,
			}

			if uncheckedBlockCtx.Block() != nil && !uncheckedBlockCtx.Block().IsEmpty() {
				for _, statement := range uncheckedBlockCtx.Block().AllStatement() {
					if statement.IsEmpty() {
						continue
					}

					// Parent index statement in this case is used only to be able provide
					// index to the parent node. It is not used for anything else.
					parentIndexStmt := &ast_pb.Statement{Id: bodyNode.Id}

					bodyNode.Statements = append(bodyNode.Statements, b.parseStatement(
						sourceUnit, node, bodyNode, parentIndexStmt, statement,
					))
				}
			}

			node.Body = bodyNode
		}
	}

	return node
}
**/
