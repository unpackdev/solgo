package contracts

import (
	"context"
	"strings"

	"github.com/unpackdev/solgo/ast"
	"github.com/unpackdev/solgo/inspector"
	"github.com/unpackdev/solgo/standards"
	"github.com/unpackdev/solgo/utils"
	"go.uber.org/zap"
)

/* var mintFunctions = []string{
	"mint", "mintWithTokenURI", "mintBatch", "mintBatchWithTokenURI",
	"_mint", "_mintWithTokenURI", "_mintBatch", "_mintBatchWithTokenURI",
}

var burnFunctions = []string{
	"burn", "_burn", "burnFrom", "_burnFrom",
} */

func (c *Contract) Inspect(ctx context.Context) (*SafetyDescriptor, error) {
	descriptor := c.GetDescriptor()
	detector := descriptor.Detector

	inspector, err := inspector.NewInspector(detector, c.GetAddress())
	if err != nil {
		zap.L().Error("Error creating inspector", zap.Error(err))
	}

	// If contract does not have any source code available we don't want to check it here.
	// In that case we will in the future go towards the opcodes...
	if !inspector.IsReady() {
		return nil, nil
	}

	// First we don't want to do any type of inspections if contract is not ERC20
	if !inspector.HasStandard(standards.ERC20) {
		return nil, nil
	} else {
		// It can be that we're not able to successfully get the standard but it is still doing trading...
		if !inspector.UsesTransfers() {
			return nil, nil
		}
	}

	// Alright now we're at the point that we know contract should be checked for any type of malicious activity
	if err := inspector.Inspect(); err != nil {
		zap.L().Error("Error inspecting contract", zap.Error(err))
		return nil, err
	}

	utils.DumpNodeNoExit(inspector.GetReport())

	/* 	shouldDump := false
	   	if descriptor.IsERC20() {

	   		detector.GetIR().GetRoot().Walk(ast.NodeVisitor{
	   			TypeVisit: map[ast_pb.NodeType]func(node ast.Node[ast.NodeType]) bool{
	   				ast_pb.NodeType_FUNCTION_DEFINITION: func(node ast.Node[ast.NodeType]) bool {
	   					switch nodeCtx := node.(type) {
	   					case *ast.Function:

	   						// Check for mint functions...
	   						if stringInSlice(nodeCtx.GetName(), mintFunctions) {
	   							descriptor.Safety.Mintable = true
	   							shouldDump = true
	   						}

	   						// Check for mint functions...
	   						if stringInSlice(nodeCtx.GetName(), mintFunctions) {
	   							descriptor.Safety.Burnable = true
	   							shouldDump = true
	   						}

	   						if isRenounceOwnershipFunction(nodeCtx) {
	   							descriptor.Safety.CanRenounceOwnership = true
	   							shouldDump = true
	   						}
	   					}

	   					return true // Continue walking
	   				},
	   				ast_pb.NodeType_FUNCTION_CALL: func(node ast.Node[ast.NodeType]) bool {
	   					nodeCtx := node.(*ast.FunctionCall)
	   					utils.DumpNodeNoExit(nodeCtx)
	   					return true // Continue walking
	   				}, */

	/* 				ast_pb.NodeType_MEMBER_ACCESS: func(node ast.Node[ast.NodeType]) bool {
		nodeCtx := node.(*ast.MemberAccessExpression)

		switch memberName := strings.ToLower(nodeCtx.GetMemberName()); memberName {
		case "transfer", "transferfrom":
			fmt.Println("Found transfer function", memberName)
		case "approve", "increaseallowance", "decreaseallowance":
			fmt.Println("Found approve function", memberName)
		case "burn", "burnfrom":
			fmt.Println("Found burn function", memberName)
		case "renounceownership":
			fmt.Println("Found renounceOwnership function", memberName)
		case "transferownership":
			fmt.Println("Found transferOwnership function", memberName)
		case "pause", "unpause":
			fmt.Println("Found pause function", memberName)
		case "call", "delegatecall", "staticcall":
			fmt.Println("Found CALL FUNCTION", memberName)

			if nodeCtx.ToText() == "target.call" {

				if nodeCtx.GetExpression() != nil && nodeCtx.GetExpression().GetType() == ast_pb.NodeType_IDENTIFIER {
					expr := nodeCtx.GetExpression().(*ast.PrimaryExpression)
					if expr.GetReferencedDeclaration() != 0 {
						exprNode := detector.GetAST().GetTree().GetById(expr.GetReferencedDeclaration())
						utils.DumpNodeNoExit(exprNode)
						fmt.Println("----------------")

					}
				}
				spew.Dump(c.GetAddress())
				utils.DumpNodeWithExit(nodeCtx)

			}

			utils.DumpNodeNoExit(nodeCtx.ToText())
		}

		return true // Continue walking
	}, */
	/* 			},
		})
	}

	if shouldDump {
		utils.DumpNodeNoExit(descriptor.Safety)
	}
	*/
	return descriptor.Safety, nil
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func isRenounceOwnershipFunction(functionNode *ast.Function) bool {
	// Check the function name
	if functionNode.GetName() != "renounceOwnership" {
		return false
	}

	// Check if the function body contains logic to set the owner to zero address
	// This is simplified; you would need to parse and understand the function body
	if strings.Contains(functionNode.GetBody().ToString(), "owner = address(0)") {
		return true
	}

	return false
}
