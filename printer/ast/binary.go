package printer

import (
	"strings"

	ast_pb "github.com/unpackdev/protos/dist/go/ast"
	"github.com/unpackdev/solgo/ast"
)

func getBinaryOperatorString(op ast_pb.Operator) string {
	switch op {
	case ast_pb.Operator_ADDITION:
		return "+"
	case ast_pb.Operator_SUBTRACTION:
		return "-"
	case ast_pb.Operator_MULTIPLICATION:
		return "*"
	case ast_pb.Operator_DIVISION:
		return "/"
	case ast_pb.Operator_MODULO:
		return "%"
	case ast_pb.Operator_EQUAL:
		return "=="
	case ast_pb.Operator_NOT_EQUAL:
		return "!="
	case ast_pb.Operator_GREATER_THAN:
		return ">"
	case ast_pb.Operator_GREATER_THAN_OR_EQUAL:
		return ">="
	case ast_pb.Operator_LESS_THAN:
		return "<"
	case ast_pb.Operator_LESS_THAN_OR_EQUAL:
		return "<="
	case ast_pb.Operator_OR:
		return "||"
		// not sure where is the AND operator in ast_pb?
	default:
		return ""
	}
}

func printBinaryOperation(node *ast.BinaryOperation, sb *strings.Builder, depth int) bool {
	ok := true
	if node.LeftExpression == nil || node.RightExpression == nil {
		return false
	}
	op := getBinaryOperatorString(node.Operator)
	if op == "" {
		ok = false
	}
	ok = PrintRecursive(node.LeftExpression, sb, depth) && ok
	sb.WriteString(" ")
	sb.WriteString(op)
	sb.WriteString(" ")
	ok = PrintRecursive(node.RightExpression, sb, depth) && ok
	return ok
}
