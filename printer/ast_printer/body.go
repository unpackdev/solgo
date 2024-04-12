package ast_printer

import (
	"strings"

	ast_pb "github.com/unpackdev/protos/dist/go/ast"
	"github.com/unpackdev/solgo/ast"
)

var blocks []ast_pb.NodeType = []ast_pb.NodeType{
	ast_pb.NodeType_IF_STATEMENT,
	ast_pb.NodeType_FOR_STATEMENT,
	ast_pb.NodeType_WHILE_STATEMENT,
	ast_pb.NodeType_DO_WHILE_STATEMENT,
	ast_pb.NodeType_FUNCTION_DEFINITION,
	ast_pb.NodeType_MODIFIER_DEFINITION,
	ast_pb.NodeType_CONTRACT_DEFINITION,
	ast_pb.NodeType_STRUCT_DEFINITION,
	ast_pb.NodeType_ENUM_DEFINITION,
}

func isBlock(nodeType ast_pb.NodeType) bool {
	for _, block := range blocks {
		if block == nodeType {
			return true
		}
	}
	return false
}

func printBody(node *ast.BodyNode, sb *strings.Builder, depth int) bool {
	success := true
	sb.WriteString("{\n")
	for _, stmt := range node.GetStatements() {
		sb.WriteString(indentString("", depth+1))
		success = PrintRecursive(stmt, sb, depth+1) && success
		if !isBlock(stmt.GetType()) {
			writeStrings(sb, ";\n")
		}
	}
	sb.WriteString(indentString("}\n", depth))
	return success
}
