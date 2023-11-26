// Package inspector provides tools for analyzing Solidity smart contracts,
// specifically focusing on detecting unconventional patterns in minting functions.
package inspector

import (
	"context"
	"fmt"

	ast_pb "github.com/unpackdev/protos/dist/go/ast"
	"github.com/unpackdev/solgo/ast"
	"github.com/unpackdev/solgo/utils"
)

// MintResults encapsulates the results of the mint function analysis.
// It includes various flags and data points indicating the presence
// and characteristics of minting functionality within a smart contract.
type MintResults struct {
	Detected                   bool                              `json:"detected"`
	Safe                       bool                              `json:"safe"`
	FunctionName               string                            `json:"function_name"`
	Visibility                 ast_pb.Visibility                 `json:"visibility"`
	ExternallyCallable         bool                              `json:"externally_callable"`
	ExternallCallableLocations []ast.Node[ast.NodeType]          `json:"externally_callable_locations"`
	UsedAtConstructor          bool                              `json:"used_at_constructor"`
	Constructor                *ast.Constructor                  `json:"constructor"`
	Statement                  *ast.Function                     `json:"statement"`
	Detectors                  map[DetectionType]*DetectorResult `json:"detectors"`
}

// IsDetected returns true if a mint function was detected in the contract.
func (m MintResults) IsDetected() bool {
	return m.Detected
}

// IsVisible returns true if the detected mint function is either public or external,
// making it callable from outside the contract.
func (m MintResults) IsVisible() bool {
	return m.Visibility == ast_pb.Visibility_PUBLIC || m.Visibility == ast_pb.Visibility_EXTERNAL
}

// MintDetector is a structure responsible for detecting mint functions
// and analyzing their characteristics in Solidity smart contracts.
type MintDetector struct {
	ctx             context.Context
	inspector       *Inspector
	functionNames   []string
	allowancesNames []string
	results         *MintResults
}

// NewMintDetector creates a new instance of MintDetector with the provided context and inspector.
// It initializes the detector with a predefined list of function names associated with minting.
func NewMintDetector(ctx context.Context, inspector *Inspector) Detector {
	return &MintDetector{
		ctx:       ctx,
		inspector: inspector,
		// Function names typically associated with minting tokens in ERC20 contracts.
		functionNames: []string{
			"mint", "mintFor", "mintTo", "mintWithTokenURI", "mintBatch", "mintBatchFor", "mintBatchTo", "mintBatchWithTokenURI",
			"_mint", "_mintFor", "_mintTo", "_mintWithTokenURI", "_mintBatch", "_mintBatchFor", "_mintBatchTo", "_mintBatchWithTokenURI",
		},
		// Common variable names for allowances in ERC20 contracts.
		allowancesNames: []string{
			"_allowances", "allowance", "allowanceFor", "allowanceMap", "allowanceMapping", "approvedTransfers",
			"spenderAllowance", "spenderAllowed", "tokenAllowances", "delegatedBalances", "allowedTransfers",
			"authorizationMap", "authorizedAmounts",
		},
		results: &MintResults{
			Safe:                       true,
			ExternallCallableLocations: make([]ast.Node[ast.NodeType], 0),
			Detectors:                  make(map[DetectionType]*DetectorResult),
		},
	}
}

// Name returns the name of the detector, which is "Mint Detector".
func (m *MintDetector) Name() string {
	return "Mintable Detector"
}

// Type returns the type of the detector, which is MintDetectorType.
func (m *MintDetector) Type() DetectorType {
	return MintDetectorType
}

// RegisterFunctionName adds a new function name to the list of mint function names
// if it is not already present. Returns true if the name was added successfully.
func (m *MintDetector) RegisterFunctionName(fnName string) bool {
	if !utils.StringInSlice(fnName, m.functionNames) {
		m.functionNames = append(m.functionNames, fnName)
		return true
	}

	return false
}

// GetFunctionNames returns a slice of registered function names associated with minting.
func (m *MintDetector) GetFunctionNames() []string {
	return m.functionNames
}

// FunctionNameExists checks if the provided function name is in the list of registered mint function names.
func (m *MintDetector) FunctionNameExists(fnName string) bool {
	return utils.StringInSlice(fnName, m.functionNames)
}

// Enter prepares the detector for analysis but currently does nothing. It may be extended in the future.
func (m *MintDetector) Enter(ctx context.Context) map[ast_pb.NodeType]func(node ast.Node[ast.NodeType]) (bool, error) {
	return map[ast_pb.NodeType]func(node ast.Node[ast.NodeType]) (bool, error){}
}

// Detect initiates the detection process for mint functions within a smart contract.
// It returns a map of node types to handler functions for further analysis.
func (m *MintDetector) Detect(ctx context.Context) map[ast_pb.NodeType]func(node ast.Node[ast.NodeType]) (bool, error) {
	return map[ast_pb.NodeType]func(node ast.Node[ast.NodeType]) (bool, error){
		ast_pb.NodeType_FUNCTION_DEFINITION: func(node ast.Node[ast.NodeType]) (bool, error) {
			if nodeCtx, ok := node.(*ast.Function); ok {
				if m.FunctionNameExists(nodeCtx.GetName()) {
					if err := m.analyzeFunctionBody(nodeCtx); err != nil {
						return true, err
					}
				}
			}

			return true, nil
		},
	}
}

// analyzeFunctionBody analyzes the body of a function to detect unconventional patterns or potential honeypots.
// It specifically looks for assignments to known allowance variables within mint functions.
func (m *MintDetector) analyzeFunctionBody(fnCtx *ast.Function) error {
	m.results.Detected = true
	m.results.Visibility = fnCtx.GetVisibility()
	m.results.FunctionName = fnCtx.GetName()
	//m.results.Statement = fnCtx

	m.inspector.GetTree().ExecuteCustomTypeVisit(fnCtx.GetNodes(), ast_pb.NodeType_ASSIGNMENT, func(node ast.Node[ast.NodeType]) (bool, error) {
		m.inspector.GetTree().ExecuteCustomTypeVisit(node.GetNodes(), ast_pb.NodeType_INDEX_ACCESS, func(node ast.Node[ast.NodeType]) (bool, error) {
			if indexCtx, ok := node.(*ast.IndexAccess); ok {
				var detector *DetectorResult

				if nameCtx, ok := indexCtx.GetIndexExpression().(*ast.PrimaryExpression); ok {
					if utils.StringInSlice(nameCtx.GetName(), m.allowancesNames) {
						detector = &DetectorResult{
							DetectionType:       DetectionType("allowance_assignment_in_mint"),
							SeverityType:        SeverityCritical,
							ConfidenceLevelType: ConfidenceLevelHigh,
							//Statement:           indexCtx,
						}
						m.results.Safe = false
					}
				}

				// @TODO: Add allowance_literal_address_in_mint detector

				if detector != nil {
					m.results.Detectors[detector.DetectionType] = detector
				}

			}

			return true, nil
		})

		return true, nil
	})

	return nil
}

// Exit finalizes the detection process by performing any necessary cleanup and additional analysis on if discovered
// mint function is used anywhere else in the contract.
func (m *MintDetector) Exit(ctx context.Context) map[ast_pb.NodeType]func(node ast.Node[ast.NodeType]) (bool, error) {
	return map[ast_pb.NodeType]func(node ast.Node[ast.NodeType]) (bool, error){

		// Problem is that mint function can be discovered at any point in time so we need to go one more time
		// through whole process in case that mint is discovered in order to get all of the reference locations where
		// mint function is being called out...
		// Mint function can exist and never be used or it can be announced as not being used where we in fact see that
		// it can be used....
		ast_pb.NodeType_FUNCTION_DEFINITION: func(fnNode ast.Node[ast.NodeType]) (bool, error) {
			// We do not want to continue if we did not discover mint function...
			if !m.results.Detected {
				return false, nil
			}

			m.inspector.GetTree().ExecuteCustomTypeVisit(fnNode.GetNodes(), ast_pb.NodeType_MEMBER_ACCESS, func(node ast.Node[ast.NodeType]) (bool, error) {
				nodeCtx, ok := node.(*ast.MemberAccessExpression)
				if !ok {
					return true, fmt.Errorf("unable to convert node to MemberAccessExpression type in MintDetector.Exit: %T", node)
				}

				if nodeCtx.GetMemberName() == m.results.FunctionName {
					switch fnCtx := fnNode.(type) {
					case *ast.Function:
						if fnCtx.GetVisibility() == ast_pb.Visibility_PUBLIC || fnCtx.GetVisibility() == ast_pb.Visibility_EXTERNAL {
							m.results.ExternallyCallable = true
							m.results.ExternallCallableLocations = append(m.results.ExternallCallableLocations, fnNode)
						} else {
							// TODO: This should recursively look for other functions when internal or private function is visibility type
						}
					}
				}

				return true, nil
			})

			m.inspector.GetTree().ExecuteCustomTypeVisit(fnNode.GetNodes(), ast_pb.NodeType_FUNCTION_CALL, func(node ast.Node[ast.NodeType]) (bool, error) {
				nodeCtx, ok := node.(*ast.FunctionCall)
				if !ok {
					return true, fmt.Errorf("unable to convert node to FunctionCall type in MintDetector.Exit: %T", node)
				}

				expressionCtx, ok := nodeCtx.GetExpression().(*ast.PrimaryExpression)
				if !ok {
					return true, fmt.Errorf("unable to convert node to PrimaryExpression type in MintDetector.Exit: %T", nodeCtx.GetExpression())
				}

				if expressionCtx.GetName() == m.results.FunctionName {
					switch fnCtx := fnNode.(type) {
					case *ast.Constructor:
						m.results.UsedAtConstructor = true
						//m.results.Constructor = fnCtx
					case *ast.Function:
						if fnCtx.GetVisibility() == ast_pb.Visibility_PUBLIC || fnCtx.GetVisibility() == ast_pb.Visibility_EXTERNAL {
							m.results.ExternallyCallable = true
							m.results.ExternallCallableLocations = append(m.results.ExternallCallableLocations, fnNode)
						} else {
							// TODO: This should recursively look for other functions when internal or private function is visibility type
						}
					}
				}
				return true, nil
			})

			return true, nil
		},
	}
}

// Results returns the results of the mint function detection and analysis.
func (m *MintDetector) Results() any {
	return m.results
}
