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
