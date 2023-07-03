package ast

import (
	"github.com/txpull/solgo/parser"
)

type VisibilityNode struct {
	Visibility string `json:"value"`
}

func (v *VisibilityNode) Children() []Node {
	return nil
}

type MutabilityNode struct {
	Mutability string `json:"value"`
}

func (m *MutabilityNode) Children() []Node {
	return nil
}

type ModifierNode struct {
	Modifier string `json:"value"`
}

func (m *ModifierNode) Children() []Node {
	return nil
}

type FunctionNode struct {
	Name             string            `json:"name"`
	Parameters       []*VariableNode   `json:"parameters"`
	ReturnParameters []*VariableNode   `json:"return_parameters"`
	Body             []*StatementNode  `json:"body"`
	Visibility       []*VisibilityNode `json:"visibility"`
	Mutability       []*MutabilityNode `json:"mutability"`
	Modifiers        []*ModifierNode   `json:"modifiers"`
	IsVirtual        bool              `json:"is_virtual"`
	IsReceive        bool              `json:"is_receive"`
	IsFallback       bool              `json:"is_fallback"`
	Overrides        bool              `json:"overrides"`
}

func (f *FunctionNode) Children() []Node {
	var nodes []Node

	// Append Parameters
	for _, param := range f.Parameters {
		nodes = append(nodes, param)
	}

	// Append ReturnParameters
	for _, retParam := range f.ReturnParameters {
		nodes = append(nodes, retParam)
	}

	// Append Body
	for _, stmt := range f.Body {
		nodes = append(nodes, stmt)
	}

	// Append Visibility
	for _, vis := range f.Visibility {
		nodes = append(nodes, vis)
	}

	// Append Mutability
	for _, mut := range f.Mutability {
		nodes = append(nodes, mut)
	}

	// Append Modifiers
	for _, mod := range f.Modifiers {
		nodes = append(nodes, mod)
	}

	return nodes
}

func (b *ASTBuilder) CreateFunction(ctx *parser.FunctionDefinitionContext) *FunctionNode {
	// Create a new FunctionNode and set it as the current function.
	toReturn := &FunctionNode{
		Name:             ctx.Identifier().GetText(), // Get the function name from the context
		Parameters:       make([]*VariableNode, 0),
		ReturnParameters: make([]*VariableNode, 0),
		Body:             make([]*StatementNode, 0),
		Visibility:       make([]*VisibilityNode, 0),
		Mutability:       make([]*MutabilityNode, 0),
		Modifiers:        make([]*ModifierNode, 0),
	}

	if ctx.GetVisibilitySet() {
		for _, visibilityCtx := range ctx.AllVisibility() {
			toReturn.Visibility = append(toReturn.Visibility, &VisibilityNode{
				Visibility: visibilityCtx.GetText(),
			})
		}
	}

	if ctx.GetMutabilitySet() {
		for _, mutabilityCtx := range ctx.AllStateMutability() {
			toReturn.Mutability = append(toReturn.Mutability, &MutabilityNode{
				Mutability: mutabilityCtx.GetText(),
			})
		}
	}

	if len(toReturn.Mutability) == 0 {
		toReturn.Mutability = append(toReturn.Mutability, &MutabilityNode{
			Mutability: "nonpayable",
		})
	}

	if ctx.GetVirtualSet() {
		toReturn.IsVirtual = true
	}

	if ctx.GetOverrideSpecifierSet() {
		for _, overrideCtx := range ctx.AllOverrideSpecifier() {
			toReturn.Overrides = overrideCtx.GetText() == "override"
		}
	}

	for _, modifierCtx := range ctx.AllModifierInvocation() {
		toReturn.Modifiers = append(toReturn.Modifiers, &ModifierNode{
			Modifier: modifierCtx.GetText(),
		})
	}

	// Handle function parameters
	if arguments := ctx.GetArguments(); arguments != nil {
		for _, parameterCtx := range arguments.AllParameterDeclaration() {
			parameterType := parameterCtx.TypeName().GetText()
			toReturn.Parameters = append(toReturn.Parameters, &VariableNode{
				Name: func() string {
					if parameterCtx.Identifier() != nil {
						return parameterCtx.Identifier().GetText()
					}
					return ""
				}(),
				Type: parameterType,
			})
		}
	}

	if returnParameters := ctx.GetReturnParameters(); returnParameters != nil {
		for _, parameterCtx := range returnParameters.AllParameterDeclaration() {
			parameterType := parameterCtx.TypeName().GetText()
			toReturn.ReturnParameters = append(toReturn.ReturnParameters, &VariableNode{
				Name: func() string {
					if parameterCtx.Identifier() != nil {
						return parameterCtx.Identifier().GetText()
					}
					return ""
				}(),
				Type: parameterType,
			})
		}
	}

	if body := ctx.GetBody(); body != nil {
		// This is a simplified example. In a real implementation, you would need to handle
		// all the different kinds of statements and expressions that can appear in a function body.
		for _, statementCtx := range body.AllStatement() {
			// Create a new StatementNode with the text of the statement.
			// Apparently whitespaces are stripped from the statementCtx.GetText() result...
			// This whole statement node will be replaced by a more complex one in the future and there
			// will be a dedicated function to parse each statement ctx.
			statement := &StatementNode{
				Raw:     b.getTextSliceWithOriginalFormatting(statementCtx),
				TextRaw: statementCtx.GetText(),
			}

			// Add the statement to the current function.
			toReturn.Body = append(toReturn.Body, statement)
		}
	}

	return toReturn
}
