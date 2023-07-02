package common

type AST struct {
	Contracts map[string]*Contract
}

type Contract struct {
	Name       string
	Functions  map[string][]*Function // Changed to a slice to support function overloading
	Variables  map[string]*Variable
	Structs    map[string]*Struct
	Events     map[string]*Event
	Interfaces map[string]*Interface
	Modifiers  map[string]*Modifier
	Inherits   []*Contract // Changed to a slice of contracts to represent inheritance hierarchy
	IsAbstract bool
}

type Function struct {
	Name            string
	Params          []*Variable
	Return          []*Variable
	Body            []Statement
	Visibility      Visibility
	StateMutability StateMutability
	IsConstructor   bool
	IsFallback      bool
	IsReceive       bool
	Modifiers       []*ModifierInvocation
}

type Variable struct {
	Name            string
	Type            Type
	Visibility      Visibility
	IsConstant      bool
	IsStateVariable bool
	IsImmutable     bool
}

type Struct struct {
	Name     string
	Elements []*Variable
}

type Event struct {
	Name        string
	Params      []*Variable
	IsAnonymous bool
}

type Interface struct {
	Name      string
	Functions map[string]*Function
}

type Modifier struct {
	Name   string
	Params []*Variable
	Body   []Statement
}

type ModifierInvocation struct {
	Modifier *Modifier
	Args     []Expression
}

type Statement struct {
	AssemblyBlock *AssemblyBlock
	// Represents a statement in the function body
}

type AssemblyBlock struct {
	Statements []*AssemblyStatement
}

type AssemblyStatement struct {
	Instruction        *AssemblyInstruction
	Assignment         *AssemblyAssignment
	Label              *AssemblyLabel
	If                 *AssemblyIf
	For                *AssemblyFor
	Switch             *AssemblySwitch
	Continue           *AssemblyContinue
	Break              *AssemblyBreak
	FunctionDefinition *AssemblyFunctionDefinition
	SubAssembly        *AssemblySubAssembly
}

type AssemblyInstruction struct {
	OpCode OpCode
	Args   []*AssemblyExpression
}

type AssemblyAssignment struct {
	Variable string
	Value    *AssemblyExpression
}

type AssemblyLabel struct {
	Name string
}

type AssemblyIf struct {
	Condition *AssemblyExpression
	Block     *AssemblyBlock
}

type AssemblyFor struct {
	Pre  *AssemblyBlock
	Cond *AssemblyExpression
	Post *AssemblyBlock
	Body *AssemblyBlock
}

type AssemblySwitch struct {
	Expression *AssemblyExpression
	Cases      []*AssemblyCase
}

type AssemblyCase struct {
	Value *AssemblyExpression // nil for default case
	Block *AssemblyBlock
}

type AssemblyContinue struct {
	// Represents a continue statement in a loop
}

type AssemblyBreak struct {
	// Represents a break statement in a loop
}

type AssemblyFunctionDefinition struct {
	Name      string
	Params    []*Variable                   // Changed to a slice of variables to include types
	Return    []*Variable                   // Added to represent return parameters
	Modifiers []*AssemblyModifierInvocation // Changed to represent assembly-specific modifier invocations
	Block     *AssemblyBlock
}

type AssemblyModifierInvocation struct {
	Modifier *Modifier
	Args     []*AssemblyExpression // Changed to represent assembly-specific expressions
}

type AssemblySubAssembly struct {
	Name  string
	Block *AssemblyBlock
}

type AssemblyExpression struct {
	// Represents an assembly expression, which could be a literal, a variable, a function call, etc.
	Kind  AssemblyExpressionKind
	Value string                // The value of the expression, if it's a literal or a variable
	Call  *AssemblyFunctionCall // The function call, if it's a function call
}

type AssemblyFunctionCall struct {
	Function string
	Args     []*AssemblyExpression
}

type OpCode string

type AssemblyExpressionKind string

type Type struct {
	BasicType   BasicType
	UserDefined string
	IsArray     bool
	ArrayLength int
}

type BasicType string
type Visibility string
type StateMutability string

type Expression struct {
	Kind               ExpressionKind
	Literal            *Literal
	Variable           string
	Call               *FunctionCall
	Unary              *UnaryOperation
	Binary             *BinaryOperation
	Assignment         *Assignment
	TypeConversion     *TypeConversion
	MemberAccess       *MemberAccess
	ArrayAccess        *ArrayAccess
	Ternary            *TernaryOperation
	New                *NewExpression
	FunctionDefinition *FunctionDefinition
	ControlStructure   *ControlStructure
	InlineAssembly     *InlineAssemblyBlock
}

type Literal struct {
	Kind  LiteralKind
	Value string
}

type FunctionCall struct {
	Function string
	Args     []*Expression
}

type UnaryOperation struct {
	Operator UnaryOperator
	Operand  *Expression
}

type BinaryOperation struct {
	Operator BinaryOperator
	Left     *Expression
	Right    *Expression
}

type Assignment struct {
	Variable string
	Value    *Expression
}

type TypeConversion struct {
	Type       Type
	Expression *Expression
}

type MemberAccess struct {
	Expression *Expression
	Member     string
}

type ArrayAccess struct {
	Array *Expression
	Index *Expression
}

type TernaryOperation struct {
	Condition *Expression
	TrueExpr  *Expression
	FalseExpr *Expression
}

type NewExpression struct {
	Type Type
}

type FunctionDefinition struct {
	Name       string
	Parameters []*Variable
	ReturnType *Type
	Body       []Statement
}

type ControlStructure struct {
	Kind          ControlStructureKind
	Condition     *Expression
	Block         []Statement
	ElseBlock     []Statement
	InitStatement *Statement
	PostStatement *Statement
}

type InlineAssemblyBlock struct {
	Instructions []string
}

// ... Other types ...

type ExpressionKind string
type LiteralKind string
type UnaryOperator string
type BinaryOperator string
type ControlStructureKind string
