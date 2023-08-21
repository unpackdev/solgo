package ir

import (
	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/ast"
)

// byFunction searches for a function by its name in the contract's AST and returns a function if found.
func (r *Builder) byFunction(name string) *Function {
	for _, unit := range r.astBuilder.GetRoot().GetSourceUnits() {
		contract := unit.GetContract()
		for _, node := range contract.GetNodes() {
			if function, ok := node.(*ast.Function); ok {
				if function.GetName() == name {
					return r.processFunction(function, false)
				}
			}
		}
	}

	return nil
}

// LookupReferencedFunctionsByNode searches for referenced functions in the given AST nodes and returns a slice of functions.
// It searches for referenced functions in member access expressions and function calls within the AST nodes recursively.
func (r *Builder) LookupReferencedFunctionsByNode(nodes ast.Node[ast.NodeType]) []*Function {
	var toReturn []*Function

	for _, node := range nodes.GetNodes() {
		if node.GetType() == ast_pb.NodeType_MEMBER_ACCESS {
			expr := node.(*ast.MemberAccessExpression)
			if expr.GetMemberName() != "" {
				if refFn := r.byFunction(expr.GetMemberName()); refFn != nil {
					toReturn = append(toReturn, refFn)
					continue
				}
			}
		}

		if node.GetType() == ast_pb.NodeType_FUNCTION_CALL {
			expr := node.(*ast.FunctionCall)
			if identifier, ok := expr.GetExpression().(*ast.PrimaryExpression); ok {
				if identifier.GetName() != "" {
					if refFn := r.byFunction(identifier.GetName()); refFn != nil {
						toReturn = append(toReturn, refFn)
						continue
					}
				}
			}
		}

		if len(node.GetNodes()) > 0 {
			for _, subnodes := range node.GetNodes() {
				foundFuncs := r.LookupReferencedFunctionsByNode(subnodes)
				toReturn = append(toReturn, foundFuncs...)
			}
		}
	}

	return toReturn
}
