package ast_printer

import (
	"strings"

	ast_pb "github.com/unpackdev/protos/dist/go/ast"
	"github.com/unpackdev/solgo/ast"
)

func printPrimaryExpression(node *ast.PrimaryExpression, sb *strings.Builder, depth int) bool {
	s := ""
	if node.GetKind() == ast_pb.NodeType_UNICODE_STRING_LITERAL {
		s = "\"" + node.GetValue() + "\""
		sb.WriteString(s)
		return true
	}

	if node.GetValue() == "" {
		s = node.GetName()
	} else {
		s = node.GetValue()
	}
	sb.WriteString(s)
	return true
}
