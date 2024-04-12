package ast_printer

import (
	"strings"

	ast_pb "github.com/unpackdev/protos/dist/go/ast"
	"github.com/unpackdev/solgo/ast"
	"go.uber.org/zap"
)

const INDENT_SIZE = 2

// Print is a function that prints the AST nodes to source code
func Print(node ast.Node[ast.NodeType]) (string, bool) {
	sb := strings.Builder{}
	success := PrintRecursive(node, &sb, 0)
	return sb.String(), success
}

// PrintRecursive is a function that prints the AST nodes to source code recursively
func PrintRecursive(node ast.Node[ast.NodeType], sb *strings.Builder, depth int) bool {
	if node == nil {
		zap.S().Error("Node is nil")
		return false
	}
	switch node := node.(type) {
	case *ast.AndOperation:
		return printAndOperation(node, sb, depth)
	case *ast.BodyNode:
		return printBody(node, sb, depth)
	case *ast.Conditional:
		return printConditional(node, sb, depth)
	case *ast.Constructor:
		return printConstructor(node, sb, depth)
	case *ast.Pragma:
		return printPragma(node, sb, depth)
	case *ast.Contract:
		return printContract(node, sb, depth)
	case *ast.Function:
		return printFunction(node, sb, depth)
	case *ast.Parameter:
		return printParameter(node, sb, depth)
	case *ast.Assignment:
		return printAssignment(node, sb, depth)
	case *ast.TypeName:
		return printTypeName(node, sb, depth)
	case *ast.BinaryOperation:
		return printBinaryOperation(node, sb, depth)
	case *ast.StateVariableDeclaration:
		return printStateVariableDeclaration(node, sb, depth)
	case *ast.Emit:
		return printEmit(node, sb, depth)
	case *ast.ForStatement:
		return printFor(node, sb, depth)
	case *ast.PrimaryExpression:
		return printPrimaryExpression(node, sb, depth)
	case *ast.FunctionCall:
		return printFunctionCall(node, sb, depth)
	case *ast.Import:
		return printImport(node, sb, depth)
	case *ast.MemberAccessExpression:
		return printMemberAccessExpression(node, sb, depth)
	case *ast.VariableDeclaration:
		return printVariableDeclaration(node, sb, depth)
	case *ast.Declaration:
		return printDeclaration(node, sb, depth)
	case *ast.UnaryPrefix:
		return printUnaryPrefix(node, sb, depth)
	case *ast.UnarySuffix:
		return printUnarySuffix(node, sb, depth)
	case *ast.IndexAccess:
		return printIndexAccess(node, sb, depth)
	case *ast.ReturnStatement:
		return printReturn(node, sb, depth)
	case *ast.TupleExpression:
		return printTupleExpression(node, sb, depth)
	case *ast.StructDefinition:
		return printStructDefinition(node, sb, depth)
	case *ast.IfStatement:
		return printIfStatement(node, sb, depth)
	case *ast.EnumDefinition:
		return printEnumDefinition(node, sb, depth)
	case *ast.ModifierDefinition:
		return printModifierDefinition(node, sb, depth)
	case *ast.EventDefinition:
		return printEventDefinition(node, sb, depth)
	case *ast.ErrorDefinition:
		return printErrorDefinition(node, sb, depth)
	case *ast.PayableConversion:
		return printPayableConversion(node, sb, depth)
	case *ast.RevertStatement:
		return printRevertStatement(node, sb, depth)
	case *ast.ContinueStatement:
		return printContinueStatement(node, sb, depth)
	default:
		if node.GetType() == ast_pb.NodeType_SOURCE_UNIT {
			return printSourceUnit(node, sb, depth)
		}
		zap.S().Errorf("Unknown node type: %T\n", node)
		return false
	}
}

func writeSeperatedStrings(sb *strings.Builder, seperator string, s ...string) {
	count := 0
	for _, item := range s {
		// Skip empty strings
		if item == "" {
			continue
		}

		if count > 0 {
			sb.WriteString(seperator)
			sb.WriteString(item)
		} else {
			sb.WriteString(item)
		}
		count++
	}
}

func writeSeperatedList(sb *strings.Builder, seperator string, s []string) {
	count := 0
	for _, item := range s {
		// Skip empty strings
		if item == "" {
			continue
		}

		if count > 0 {
			sb.WriteString(seperator)
			sb.WriteString(item)
		} else {
			sb.WriteString(item)
		}
		count++
	}
}

func writeStrings(sb *strings.Builder, s ...string) {
	for _, item := range s {
		sb.WriteString(item)
	}
}

func indentString(s string, depth int) string {
	return strings.Repeat(" ", depth*INDENT_SIZE) + s
}

func getStorageLocationString(storage ast_pb.StorageLocation) string {
	switch storage {
	case ast_pb.StorageLocation_DEFAULT:
		return ""
	case ast_pb.StorageLocation_MEMORY:
		return "memory"
	case ast_pb.StorageLocation_STORAGE:
		return "storage"
	case ast_pb.StorageLocation_CALLDATA:
		return "calldata"
	default:
		return ""
	}
}

func getVisibilityString(visibility ast_pb.Visibility) string {
	switch visibility {
	case ast_pb.Visibility_INTERNAL:
		return "internal"
	case ast_pb.Visibility_PUBLIC:
		return "public"
	case ast_pb.Visibility_EXTERNAL:
		return "external"
	case ast_pb.Visibility_PRIVATE:
		return "private"
	default:
		return ""
	}
}

func getStateMutabilityString(mut ast_pb.Mutability) string {
	switch mut {
	case ast_pb.Mutability_PURE:
		return "pure"
	case ast_pb.Mutability_VIEW:
		return "view"
	case ast_pb.Mutability_NONPAYABLE:
		return ""
	case ast_pb.Mutability_PAYABLE:
		return "payable"
	default:
		return ""
	}
}
