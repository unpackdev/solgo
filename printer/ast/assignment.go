package printer

import (
	"strings"

	ast_pb "github.com/unpackdev/protos/dist/go/ast"
	"github.com/unpackdev/solgo/ast"
)

func getAssignOperatorString(op ast_pb.Operator) string {
	switch op {
	case ast_pb.Operator_EQUAL:
		return "="
	case ast_pb.Operator_PLUS_EQUAL:
		return "+="
	case ast_pb.Operator_MINUS_EQUAL:
		return "-="
	case ast_pb.Operator_MUL_EQUAL:
		return "*="
	case ast_pb.Operator_DIV_EQUAL:
		return "/="
	case ast_pb.Operator_MOD_EQUAL:
		return "%="
	case ast_pb.Operator_AND_EQUAL:
		return "&="
	case ast_pb.Operator_OR_EQUAL:
		return "|="
	case ast_pb.Operator_XOR_EQUAL:
		return "^="
	case ast_pb.Operator_SHIFT_LEFT_EQUAL:
		return "<<="
	case ast_pb.Operator_SHIFT_RIGHT_EQUAL:
		return ">>="
	case ast_pb.Operator_BIT_AND_EQUAL:
		return "&="
	case ast_pb.Operator_BIT_OR_EQUAL:
		return "|="
	case ast_pb.Operator_BIT_XOR_EQUAL:
		return "^="
	case ast_pb.Operator_POW_EQUAL:
		return "**="
	default:
		return ""
	}
}

func printAssignment(node *ast.Assignment, sb *strings.Builder, depth int) bool {
	success := true
	if node.Expression != nil {
		return PrintRecursive(node.Expression, sb, depth)
	}
	if node.LeftExpression == nil || node.RightExpression == nil {
		return false
	}
	op := getAssignOperatorString(node.Operator)
	if op == "" {
		success = false
	}
	success = PrintRecursive(node.LeftExpression, sb, depth) && success
	sb.WriteString(" ")
	sb.WriteString(op)
	sb.WriteString(" ")
	success = PrintRecursive(node.RightExpression, sb, depth) && success
	return success
}
