package ast_printer

import (
	"strings"

	ast_pb "github.com/unpackdev/protos/dist/go/ast"
	"github.com/unpackdev/solgo/ast"
)

func getUnaryOperatorString(op ast_pb.Operator) string {
	switch op {
	case ast_pb.Operator_NOT:
		return "!"
	case ast_pb.Operator_BIT_NOT:
		return "~"
	case ast_pb.Operator_SUBTRACT:
		return "-"
	case ast_pb.Operator_INCREMENT:
		return "++"
	case ast_pb.Operator_DECREMENT:
		return "--"
	default:
		return ""
	}
}

func printUnarySuffix(node *ast.UnarySuffix, sb *strings.Builder, depth int) bool {
	success := PrintRecursive(node.GetExpression(), sb, depth)
	sb.WriteString(getUnaryOperatorString(node.GetOperator()))
	return success
}
