package ast

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/antlr4-go/antlr/v4"
	"github.com/txpull/solgo/parser"
)

type ASTBuilder struct {
	*parser.BaseSolidityParserListener
	currentContract *ContractNode
	astRoot         *RootNode
	errors          []error
	parseTime       time.Time
	pragmas         [][]string
}

func NewAstBuilder() *ASTBuilder {
	return &ASTBuilder{
		pragmas: make([][]string, 0),
	}
}

func (b *ASTBuilder) EnterSourceUnit(ctx *parser.SourceUnitContext) {
	b.currentContract = nil

	b.astRoot = &RootNode{Contracts: make([]*ContractNode, 0)}

	b.errors = nil
	b.parseTime = time.Now()

	for _, child := range ctx.GetChildren() {
		if contractCtx, ok := child.(*parser.ContractDefinitionContext); ok {
			// Found a contract definition
			contractPragmas := b.findPragmasForContract(ctx, contractCtx)
			if len(contractPragmas) > 0 {
				contract := &ContractNode{
					Name:    contractCtx.Identifier().GetText(),
					Pragmas: contractPragmas,
					Line:    contractCtx.GetStart().GetLine(),
				}
				b.astRoot.Contracts = append(b.astRoot.Contracts, contract)
			}
		}
	}
}

func (b *ASTBuilder) EnterContractDefinition(ctx *parser.ContractDefinitionContext) {
	// Look if contract already exists in the AST root
	for _, contract := range b.astRoot.Contracts {
		if contract.Name == ctx.Identifier().GetText() {
			b.currentContract = contract
			return
		}
	}

	if b.currentContract == nil {
		b.currentContract = &ContractNode{
			Name:           ctx.Identifier().GetText(),
			StateVariables: make([]*StateVariableNode, 0),
			Kind:           "contract",
			Line:           ctx.GetStart().GetLine(),
		}
	}

	if ctx.InheritanceSpecifierList() != nil {
		for _, inheritance := range ctx.InheritanceSpecifierList().AllInheritanceSpecifier() {
			b.currentContract.Inherits = append(b.currentContract.Inherits, inheritance.GetText())
		}
	}

	if ctx.Abstract() != nil {
		b.currentContract.Kind = "abstract"
	}

	b.astRoot.Contracts = append(b.astRoot.Contracts, b.currentContract)
}

func (b *ASTBuilder) ExitContractDefinition(ctx *parser.ContractDefinitionContext) {
	b.currentContract = nil
}

func (b *ASTBuilder) EnterInterfaceDefinition(ctx *parser.InterfaceDefinitionContext) {
	// Look if contract already exists in the AST root
	for _, contract := range b.astRoot.Contracts {
		if contract.Name == ctx.Identifier().GetText() {
			b.currentContract = contract
			return
		}
	}

	if b.currentContract == nil {
		b.currentContract = &ContractNode{
			Name:           ctx.Identifier().GetText(),
			StateVariables: make([]*StateVariableNode, 0),
			Kind:           "interface",
		}
	}

	if ctx.InheritanceSpecifierList() != nil {
		for _, inheritance := range ctx.InheritanceSpecifierList().AllInheritanceSpecifier() {
			b.currentContract.Inherits = append(b.currentContract.Inherits, inheritance.GetText())
		}
	}
}

func (b *ASTBuilder) EnterLibraryDefinition(ctx *parser.LibraryDefinitionContext) {
	// Look if contract already exists in the AST root
	for _, contract := range b.astRoot.Contracts {
		if contract.Name == ctx.Identifier().GetText() {
			b.currentContract = contract
			return
		}
	}

	if b.currentContract == nil {
		b.currentContract = &ContractNode{
			Name:           ctx.Identifier().GetText(),
			StateVariables: make([]*StateVariableNode, 0),
			Kind:           "library",
		}
	}
}

func (b *ASTBuilder) EnterUsingDirective(ctx *parser.UsingDirectiveContext) {
	usingDirective := &UsingDirectiveNode{
		Type:       ctx.TypeName().GetText(),
		IsWildcard: ctx.Mul() != nil,
		IsGlobal:   ctx.Global() != nil,
	}

	if ctx.AllUserDefinableOperator() != nil {
		for _, operator := range ctx.AllUserDefinableOperator() {
			if operator.GetText() == "userDefined" {
				usingDirective.IsUserDef = true
			}
		}
	}

	if ctx.AllIdentifierPath() != nil {
		for _, identifier := range ctx.AllIdentifierPath() {
			usingDirective.Alias = identifier.GetText()
		}
	}

	b.currentContract.Using = append(b.currentContract.Using, usingDirective)
}

func (b *ASTBuilder) EnterStateVariableDeclaration(ctx *parser.StateVariableDeclarationContext) {
	variable := &StateVariableNode{
		Name: func() string {
			if ctx.Identifier() != nil {
				return ctx.Identifier().GetText()
			}
			return ""
		}(),
		Type: ctx.GetType_().GetText(),
	}

	if ctx.GetVisibilitySet() {
		if ctx.AllInternal() != nil {
			variable.Visibility = "internal"
		} else if ctx.AllPrivate() != nil {
			variable.Visibility = "private"
		} else if ctx.AllPublic() != nil {
			variable.Visibility = "public"
		}
	}

	if ctx.GetConstantnessSet() {
		for _, constant := range ctx.AllConstant() {
			if constant.GetText() == "constant" {
				variable.IsConstant = true
			}
		}
	}

	if ctx.AllImmutable() != nil {
		for _, modifier := range ctx.AllImmutable() {
			if modifier.GetText() == "immutable" {
				variable.IsImmutable = true
			}
		}
	}

	if initialValue := ctx.GetInitialValue(); initialValue != nil {
		variable.InitialValue = initialValue.GetText()
	}

	b.currentContract.StateVariables = append(b.currentContract.StateVariables, variable)
}

func (b *ASTBuilder) EnterEnumDefinition(ctx *parser.EnumDefinitionContext) {
	enum := &EnumNode{
		Name:         ctx.GetName().GetText(),
		MemberValues: make([]*EnumMemberNode, len(ctx.GetEnumValues())),
	}

	for i, valueCtx := range ctx.GetEnumValues() {
		enum.MemberValues[i] = &EnumMemberNode{Name: valueCtx.GetText()}
	}

	b.currentContract.Enums = append(b.currentContract.Enums, enum)
}

func (b *ASTBuilder) EnterConstructorDefinition(ctx *parser.ConstructorDefinitionContext) {
	constructor := &ConstructorNode{
		Parameters: make([]*VariableNode, 0),
		Body:       make([]*StatementNode, 0),
	}

	if arguments := ctx.GetArguments(); arguments != nil {
		for _, parameterCtx := range arguments.AllParameterDeclaration() {
			constructor.Parameters = append(constructor.Parameters, &VariableNode{
				Name: func() string {
					if parameterCtx.Identifier() != nil {
						return parameterCtx.Identifier().GetText()
					}
					return ""
				}(),
				Type: parameterCtx.TypeName().GetText(),
			})
		}
	}

	if body := ctx.GetBody(); body != nil {
		if !body.IsEmpty() {
			statements := b.TraverseStatements(constructor.Body, constructor.Parameters, body)
			constructor.Body = append(constructor.Body, statements...)
		}
	}

	b.currentContract.Constructor = constructor
}

func (b *ASTBuilder) TraverseStatements(body []*StatementNode, fnArguments []*VariableNode, node antlr.Tree) []*StatementNode {
	statements := []*StatementNode{}

	switch node := node.(type) {

	case *parser.AssignmentContext:
		tokens := b.collectTokens(node)
		tokensNode := []TokenNode{}

		for _, token := range tokens {
			tokenNode := TokenNode{
				Name:           token.GetText(),
				LexerTypeIndex: token.GetTokenType(),
				LexerType:      getTokenTypeName(token),
			}

			for _, fnArg := range fnArguments {
				if token.GetText() == fnArg.Name {
					tokenNode.IsFunctionArgument = true
					tokenNode.Type = fnArg.Type
				}
			}

			for _, stateVar := range b.currentContract.StateVariables {
				if token.GetText() == stateVar.Name {
					tokenNode.IsStateVariable = true
					tokenNode.Type = stateVar.Type
				}
			}

			tokensNode = append(tokensNode, tokenNode)
		}

		if len(tokensNode) > 0 && tokensNode[len(tokensNode)-1:][0].LexerTypeIndex != parser.SolidityParserSemicolon {
			tokensNode = append(tokensNode, TokenNode{
				Name:           ";",
				LexerTypeIndex: parser.SolidityParserSemicolon,
				LexerType:      "semicolon",
			})
		}

		statementNode := &StatementNode{
			Expression: func() string {
				toReturn := []string{}

				for _, token := range tokens {
					if token.GetTokenType() != parser.SolidityParserSemicolon {
						toReturn = append(toReturn, token.GetText())
					}
				}

				return strings.TrimSpace(strings.Join(toReturn, " ")) + ";"
			}(),
			Line:   node.GetStart().GetLine(),
			Type:   "assignment",
			Tokens: tokensNode,
		}

		statements = append(statements, statementNode)

	case *parser.VariableDeclarationStatementContext:
		tokens := b.collectTokens(node)
		tokensNode := []TokenNode{}

		for _, token := range tokens {
			tokenNode := TokenNode{
				Name:           token.GetText(),
				LexerTypeIndex: token.GetTokenType(),
				LexerType:      getTokenTypeName(token),
			}

			if tokenNode.LexerTypeIndex == parser.SolidityParserIdentifier {
				tokenNode.IsFunctionArgument = true
			}

			tokensNode = append(tokensNode, tokenNode)
		}

		if len(tokensNode) > 0 && tokensNode[len(tokensNode)-1:][0].LexerTypeIndex != parser.SolidityParserSemicolon {
			tokensNode = append(tokensNode, TokenNode{
				Name:           ";",
				LexerTypeIndex: parser.SolidityParserSemicolon,
				LexerType:      "semicolon",
			})
		}

		statementNode := &StatementNode{
			Expression: func() string {
				toReturn := []string{}

				for _, token := range tokens {
					if token.GetTokenType() != parser.SolidityParserSemicolon {
						toReturn = append(toReturn, token.GetText())
					}
				}

				return strings.TrimSpace(strings.Join(toReturn, " ")) + ";"
			}(),
			Line:   node.GetStart().GetLine(),
			Type:   "variable_declaration",
			Tokens: tokensNode,
		}

		statements = append(statements, statementNode)

	case *parser.FunctionCallContext:
		tokens := b.collectTokens(node)
		tokensNode := []TokenNode{}

		for _, token := range tokens {
			tokensNode = append(tokensNode, TokenNode{
				Name:           token.GetText(),
				LexerTypeIndex: token.GetTokenType(),
				LexerType:      getTokenTypeName(token),
			})
		}

		if len(tokensNode) > 0 && tokensNode[len(tokensNode)-1:][0].LexerTypeIndex != parser.SolidityParserSemicolon {
			tokensNode = append(tokensNode, TokenNode{
				Name:           ";",
				LexerTypeIndex: parser.SolidityParserSemicolon,
				LexerType:      "semicolon",
			})
		}

		statementNode := &StatementNode{
			Expression: func() string {
				var toReturn string

				for _, token := range tokensNode {
					if token.LexerTypeIndex != parser.SolidityParserLParen {
						toReturn += token.Name + ""
					} else if token.LexerTypeIndex != parser.SolidityParserIdentifier {
						toReturn += token.Name + ""
					} else if token.LexerTypeIndex != parser.SolidityParserComma {
						toReturn += token.Name + " "
					} else if token.LexerTypeIndex != parser.SolidityParserSemicolon {
						toReturn += ""
					}
				}

				return toReturn
			}(),
			Line:   node.GetStart().GetLine(),
			Type:   "function_call",
			Tokens: tokensNode,
		}

		statements = append(statements, statementNode)

	default:
		// This node is not a statement, so we recurse on its children.
		for i := 0; i < node.GetChildCount(); i++ {
			childStatements := b.TraverseStatements(body, fnArguments, node.GetChild(i))
			if childStatements != nil {
				statements = append(statements, childStatements...)
			}
		}
	}

	return statements
}

func (b *ASTBuilder) collectTokens(node antlr.Tree) []antlr.Token {
	tokens := []antlr.Token{}

	switch node := node.(type) {
	case antlr.TerminalNode:
		tokens = append(tokens, node.GetSymbol())
	default:
		// This node is not a token, so we recurse on its children.
		for i := 0; i < node.GetChildCount(); i++ {
			childTokens := b.collectTokens(node.GetChild(i))
			tokens = append(tokens, childTokens...)
		}
	}

	return tokens
}

func (b *ASTBuilder) EnterStructDefinition(ctx *parser.StructDefinitionContext) {
	structNode := &StructNode{
		Name: ctx.GetName().GetText(),
	}

	for _, memberCtx := range ctx.AllStructMember() {
		structNode.Members = append(structNode.Members, &StructMemberNode{
			Name: func() string {
				if memberCtx.Identifier() != nil {
					return memberCtx.Identifier().GetText()
				}
				return ""
			}(),
			Type: memberCtx.GetType_().GetText(),
		})
	}

	b.currentContract.Structs = append(b.currentContract.Structs, structNode)
}

func (b *ASTBuilder) EnterErrorDefinition(ctx *parser.ErrorDefinitionContext) {
	errorNode := &ErrorNode{
		Name:   ctx.GetName().GetText(),
		Values: make([]*ErrorValueNode, 0),
	}

	for _, errorParamCtx := range ctx.GetParameters() {
		errorValue := &ErrorValueNode{
			Name: func() string {
				if errorParamCtx.Identifier() != nil {
					return errorParamCtx.Identifier().GetText()
				}
				return ""
			}(),
			Type: errorParamCtx.GetType_().GetText(),
			Code: 0,
		}
		errorNode.Values = append(errorNode.Values, errorValue)
	}

	b.currentContract.Errors = append(b.currentContract.Errors, errorNode)
}

func (b *ASTBuilder) EnterFunctionDefinition(ctx *parser.FunctionDefinitionContext) {
	currentFunction := b.CreateFunction(ctx)
	b.currentContract.Functions = append(b.currentContract.Functions, currentFunction)
}

func (b *ASTBuilder) EnterFallbackFunctionDefinition(ctx *parser.FallbackFunctionDefinitionContext) {
	fallbackFunction := &FunctionNode{
		Name:             "fallback",
		Parameters:       make([]*VariableNode, 0),
		ReturnParameters: make([]*VariableNode, 0),
		Body:             make([]*StatementNode, 0),
		Visibility:       make([]*VisibilityNode, 0),
		Mutability:       make([]*MutabilityNode, 0),
		Modifiers:        make([]*ModifierNode, 0),
		Overrides:        false,
		IsVirtual:        false,
		IsFallback:       true,
	}

	// Handle virtual modifier
	if ctx.GetVirtualSet() {
		fallbackFunction.IsVirtual = true
	}

	if ctx.GetVisibilitySet() {
		for _, externalCtx := range ctx.AllExternal() {
			fallbackFunction.Visibility = append(fallbackFunction.Visibility, &VisibilityNode{
				Visibility: externalCtx.GetText(),
			})
		}
	}

	if ctx.GetMutabilitySet() {
		for _, mutabilityCtx := range ctx.AllStateMutability() {
			fallbackFunction.Mutability = append(fallbackFunction.Mutability, &MutabilityNode{
				Mutability: mutabilityCtx.GetText(),
			})
		}
	}

	if len(fallbackFunction.Mutability) == 0 {
		fallbackFunction.Mutability = append(fallbackFunction.Mutability, &MutabilityNode{
			Mutability: "nonpayable",
		})
	}

	if ctx.GetOverrideSpecifierSet() {
		for _, overrideCtx := range ctx.AllOverrideSpecifier() {
			fallbackFunction.Overrides = !overrideCtx.IsEmpty()
		}
	}

	for _, modifierCtx := range ctx.AllModifierInvocation() {
		fallbackFunction.Modifiers = append(fallbackFunction.Modifiers, &ModifierNode{
			Modifier: modifierCtx.GetText(),
		})
	}

	if parameters := ctx.AllParameterList(); parameters != nil {
		for _, parameterCtx := range parameters {
			for _, param := range parameterCtx.AllParameterDeclaration() {
				fallbackFunction.Parameters = append(fallbackFunction.Parameters, &VariableNode{
					Name: func() string {
						if param.Identifier() != nil {
							return param.Identifier().GetText()
						}
						return ""
					}(),
					Type: param.TypeName().GetText(),
				})
			}
		}
	}

	if body := ctx.GetBody(); body != nil {
		for _, statementCtx := range body.AllStatement() {
			statement := &StatementNode{
				Expression: statementCtx.GetText(),
			}
			fallbackFunction.Body = append(fallbackFunction.Body, statement)
		}
	}

	b.currentContract.Functions = append(b.currentContract.Functions, fallbackFunction)
}

func (b *ASTBuilder) EnterReceiveFunctionDefinition(ctx *parser.ReceiveFunctionDefinitionContext) {
	receiveFn := &FunctionNode{
		Name:             "receive",
		Parameters:       make([]*VariableNode, 0),
		ReturnParameters: make([]*VariableNode, 0),
		Body:             make([]*StatementNode, 0),
		Visibility:       make([]*VisibilityNode, 0),
		Mutability:       make([]*MutabilityNode, 0),
		Modifiers:        make([]*ModifierNode, 0),
		Overrides:        false,
		IsVirtual:        false,
		IsReceive:        true,
	}

	if ctx.GetVirtualSet() {
		receiveFn.IsVirtual = true
	}

	if ctx.GetOverrideSpecifierSet() {
		for _, overrideCtx := range ctx.AllOverrideSpecifier() {
			receiveFn.Overrides = !overrideCtx.IsEmpty()
		}
	}

	if body := ctx.GetBody(); body != nil {
		for _, statementCtx := range body.AllStatement() {
			statement := &StatementNode{
				Expression: statementCtx.GetText(),
			}
			receiveFn.Body = append(receiveFn.Body, statement)
		}
	}

	b.currentContract.Functions = append(b.currentContract.Functions, receiveFn)
}

func (b *ASTBuilder) EnterEventDefinition(ctx *parser.EventDefinitionContext) {
	event := &EventNode{
		Name:      ctx.GetName().GetText(),
		Anonymous: ctx.Anonymous() != nil,
	}

	if ctx.AllEventParameter() != nil {
		event.Parameters = make([]*EventParameterNode, 0)

		for _, parameterCtx := range ctx.AllEventParameter() {
			event.Parameters = append(event.Parameters, &EventParameterNode{
				Name: func() string {
					if parameterCtx.Identifier() != nil {
						return parameterCtx.Identifier().GetText()
					}
					return ""
				}(),
				Type:    parameterCtx.GetType_().GetText(),
				Indexed: parameterCtx.Indexed() != nil,
			})
		}
	}

	b.currentContract.Events = append(b.currentContract.Events, event)
}

func (b *ASTBuilder) GetTree() Node {
	return b.astRoot
}

func (b *ASTBuilder) GetRootNode() *RootNode {
	return b.astRoot
}

func (b *ASTBuilder) GetPragmas() [][]string {
	return b.pragmas
}

func (b *ASTBuilder) GetErrors() []error {
	return b.errors
}

func (b *ASTBuilder) GetParseTime() time.Duration {
	return time.Since(b.parseTime)
}

// ToJSON converts the ABI object into a JSON string.
func (b *ASTBuilder) ToJSON() (string, error) {
	abiJSON, err := json.Marshal(b.astRoot)
	if err != nil {
		return "", err
	}

	return string(abiJSON), nil
}
