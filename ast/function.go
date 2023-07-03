package ast

import (
	"github.com/txpull/solgo/parser"
)

// FunctionNode represents a function definition in Solidity.
type FunctionNode struct {
	Name             string           `json:"name"`
	Parameters       []*VariableNode  `json:"parameters"`
	ReturnParameters []*VariableNode  `json:"return_parameters"`
	Body             []*StatementNode `json:"body"`
	Visibility       []string         `json:"visibility"`
	Mutability       []string         `json:"mutability"`
	Modifiers        []string         `json:"modifiers"`
	IsVirtual        bool             `json:"is_virtual"`
	IsReceive        bool             `json:"is_receive"`
	IsFallback       bool             `json:"is_fallback"`
	Overrides        bool             `json:"overrides"`
}

func (f *FunctionNode) Children() []Node {
	nodes := make([]Node, len(f.Parameters)+len(f.ReturnParameters)+len(f.Body))
	for i, parameter := range f.Parameters {
		nodes[i] = parameter
	}
	for i, returnParameter := range f.ReturnParameters {
		nodes[i+len(f.Parameters)] = returnParameter
	}
	for i, statement := range f.Body {
		nodes[i+len(f.Parameters)+len(f.ReturnParameters)] = statement
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
		Visibility:       make([]string, 0),
		Mutability:       make([]string, 0),
		Modifiers:        make([]string, 0),
	}

	if ctx.GetVisibilitySet() {
		for _, visibilityCtx := range ctx.AllVisibility() {
			visibility := visibilityCtx.GetText()
			toReturn.Visibility = append(toReturn.Visibility, visibility)
		}
	}

	if ctx.GetMutabilitySet() {
		for _, mutabilityCtx := range ctx.AllStateMutability() {
			mutability := mutabilityCtx.GetText()
			toReturn.Mutability = append(toReturn.Mutability, mutability)
		}
	}

	if len(toReturn.Mutability) == 0 {
		toReturn.Mutability = append(toReturn.Mutability, "nonpayable")
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
		modifier := modifierCtx.GetText()
		toReturn.Modifiers = append(toReturn.Modifiers, modifier)
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
