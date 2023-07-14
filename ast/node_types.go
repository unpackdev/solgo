package ast

type NodeType string

const (
	NodeTypeSourceUnit NodeType = "SourceUnit"
	NodeTypeContract   NodeType = "Contract"

	NodeTypeModifier           NodeType = "Modifier"
	NodeTypeVariable           NodeType = "Variable"
	NodeTypeEnum               NodeType = "Enum"
	NodeTypeStruct             NodeType = "Struct"
	NodeTypeEvent              NodeType = "Event"
	NodeTypeError              NodeType = "Error"
	NodeTypeUsing              NodeType = "Using"
	NodeTypePragmaDirective    NodeType = "PragmaDirective"
	NodeTypeConstructor        NodeType = "Constructor"
	NodeTypeReturn             NodeType = "Return"
	NodeTypeMapping            NodeType = "Mapping"
	NodeTypeArray              NodeType = "Array"
	NodeTypeEnumValue          NodeType = "EnumValue"
	NodeTypeIdentifier         NodeType = "Identifier"
	NodeTypeLiteral            NodeType = "Literal"
	NodeTypeUnary              NodeType = "UnaryOperation"
	NodeTypeBinary             NodeType = "BinaryOperation"
	NodeTypeTernary            NodeType = "TernaryOperation"
	NodeTypeTuple              NodeType = "Tuple"
	NodeTypeIndexAccess        NodeType = "IndexAccess"
	NodeTypeMemberAccess       NodeType = "MemberAccess"
	NodeTypeFunctionCall       NodeType = "FunctionCall"
	NodeTypeNewExpression      NodeType = "NewExpression"
	NodeTypeConditional        NodeType = "Conditional"
	NodeTypeAssignment         NodeType = "Assignment"
	NodeTypeEmit               NodeType = "Emit"
	NodeTypeImport             NodeType = "Import"
	NodeTypeElementaryTypeName NodeType = "ElementaryTypeName"

	// General comment types...
	NodeTypeCommentLine      NodeType = "Comment"
	NodeTypeCommentMultiLine NodeType = "CommentMultiline"
	NodeTypeLicense          NodeType = "License"

	// Contract definition types...
	NodeTypeLibraryDefinition  NodeType = "LibraryDefinition"
	NodeTypeContractDefinition NodeType = "ContractDefinition"

	// Contract types...
	NodeTypeKindContract  NodeType = "contract"
	NodeTypeKindLibrary   NodeType = "library"
	NodeTypeKindInterface NodeType = "interface"
	NodeTypeKindStruct    NodeType = "struct"
	NodeTypeKindEnum      NodeType = "enum"
	NodeTypeKindFunction  NodeType = "function"

	// Body types...
	NodeTypeFunctionDefinition NodeType = "FunctionDefinition"
	NodeTypeParameterList      NodeType = "ParameterList"
)

func (n NodeType) String() string {
	return string(n)
}
