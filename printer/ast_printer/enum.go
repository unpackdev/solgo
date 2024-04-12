package ast_printer

import (
	"strings"

	"github.com/unpackdev/solgo/ast"
)

func printEnumDefinition(node *ast.EnumDefinition, sb *strings.Builder, depth int) bool {
	success := true
	sb.WriteString("enum ")
	sb.WriteString(node.GetName())
	sb.WriteString(" {")
	members := []string{}
	for _, member := range node.GetMembers() {
		s, ok := Print(member)
		success = ok && success
		members = append(members, s)
	}
	writeSeperatedList(sb, ", ", members)
	sb.WriteString("}\n")
	return success
}
