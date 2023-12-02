package inspector

import (
	"context"
	"fmt"

	ast_pb "github.com/unpackdev/protos/dist/go/ast"
	"github.com/unpackdev/solgo/ast"
	"github.com/unpackdev/solgo/standards"
	"github.com/unpackdev/solgo/utils"
)

type Statement struct {
}

type Function struct {
	ContractName           string                       `json:"contract_name"`
	ContractType           ast_pb.NodeType              `json:"contract_type"`
	ContractKind           ast_pb.NodeType              `json:"contract_kind"`
	Name                   string                       `json:"name"`
	SignatureCompatibility standards.ConfidenceLevel    `json:"signature_compatibility"`
	Standard               *standards.FunctionDiscovery `json:"standard"`
	Visibility             ast_pb.Visibility            `json:"visibility"`
	StateMutability        ast_pb.Mutability            `json:"mutability"`
	Unit                   *ast.Function                `json:"unit"`
	Detectors              []DetectorResult             `json:"detectors"`
}

type TransferResults struct {
	Detected  bool       `json:"detected"`
	Safe      bool       `json:"safe"`
	Functions []Function `json:"functions"`
}

type TransferDetector struct {
	ctx           context.Context
	inspector     *Inspector
	functionNames []string
	results       *TransferResults
}

func NewTransferDetector(ctx context.Context, inspector *Inspector) Detector {
	return &TransferDetector{
		ctx:       ctx,
		inspector: inspector,
		functionNames: []string{
			"transfer", "transferFrom", "_transfer", "_transferFrom",
		},
		results: &TransferResults{
			Safe:      true,
			Functions: make([]Function, 0),
		},
	}
}

func (m *TransferDetector) Name() string {
	return "Transfer Detector"
}

func (m *TransferDetector) Type() DetectorType {
	return TransferDetectorType
}

func (m *TransferDetector) Enter(ctx context.Context) (DetectorFn, error) {

	standard, err := standards.GetContractByStandard(standards.ERC20)
	if err != nil {
		fmt.Println(err)
	}

	return map[ast_pb.NodeType]func(node ast.Node[ast.NodeType]) (bool, error){
		ast_pb.NodeType_FUNCTION_DEFINITION: func(node ast.Node[ast.NodeType]) (bool, error) {
			if fnCtx, ok := node.(*ast.Function); ok {
				if utils.StringInSlice(fnCtx.GetName(), m.functionNames) {
					var discoveredFn Function
					discoveredFn.Name = fnCtx.GetName()
					discoveredFn.Detectors = make([]DetectorResult, 0)
					//discoveredFn.Statement = fnCtx

					if contract := m.inspector.GetDetector().GetAST().GetTree().GetById(fnCtx.GetScope()); contract != nil {
						discoveredFn.ContractType = contract.GetType()
						switch contractCtx := contract.(type) {
						case *ast.Contract:
							discoveredFn.ContractName = contractCtx.GetName()
							discoveredFn.ContractKind = contractCtx.GetKind()
						case *ast.Interface:
							discoveredFn.ContractName = contractCtx.GetName()
							discoveredFn.ContractKind = contractCtx.GetKind()
						case *ast.Library:
							discoveredFn.ContractName = contractCtx.GetName()
							discoveredFn.ContractKind = contractCtx.GetKind()
						}
					}

					if standardFn := m.getStandardFunction(standard, fnCtx.GetName()); standardFn != nil {
						m.results.Detected = true
						newStandardFn := m.buildStandardFunction(fnCtx)
						if check, found := standards.FunctionConfidenceCheck(standard, &newStandardFn); found {
							discoveredFn.Standard = &check
							discoveredFn.SignatureCompatibility = check.Confidence
						}
					}

					m.analyzeERC20Function(fnCtx, &discoveredFn)
					m.results.Functions = append(m.results.Functions, discoveredFn)
				}
			}
			return true, nil
		},
	}, nil
}

func (m *TransferDetector) analyzeERC20Function(fnCtx *ast.Function, function *Function) {
	function.Visibility = fnCtx.GetVisibility()
	function.StateMutability = fnCtx.GetStateMutability()

	if fnCtx.GetName() == "transfer" || fnCtx.GetName() == "transferFrom" {
		m.checkForOwnerVariable(fnCtx, function)
		m.checkForInternalTransferCall(fnCtx, function)
		m.checkForBalanceUpdate(fnCtx, function)
		m.checkForEventEmission(fnCtx, function)
		m.checkForAccessControl(fnCtx, function)
	}

	// Additional checks can be added here based on your requirements
}

// Example of a new check: Verify that the function correctly updates balances
func (m *TransferDetector) checkForBalanceUpdate(fnCtx *ast.Function, function *Function) {
	senderBalanceUpdated := false
	recipientBalanceUpdated := false

	for _, node := range fnCtx.GetNodes() {
		// Check for assignments that update the balance mapping
		if assignStmt, ok := node.(*ast.Assignment); ok {
			_ = assignStmt
			/* 			if mappingAccess, ok := assignStmt.GetLeftHandSide().(*ast.MappingAccess); ok {
				if mappingAccess.GetMapping().GetName() == "_balances" {
					if isSenderBalanceUpdate(mappingAccess, assignStmt) {
						senderBalanceUpdated = true
					}
					if isRecipientBalanceUpdate(mappingAccess, assignStmt) {
						recipientBalanceUpdated = true
					}
				}
			} */
		}
	}

	if !senderBalanceUpdated || !recipientBalanceUpdated {
		function.Detectors = append(function.Detectors, DetectorResult{
			DetectionType:       "balance_update_missing",
			SeverityType:        SeverityMedium,
			ConfidenceLevelType: ConfidenceLevelHigh,
			Description:         "Balance update logic is missing or incorrect.",
		})
	} else {
		function.Detectors = append(function.Detectors, DetectorResult{
			DetectionType:       "balance_update_detected",
			SeverityType:        SeverityInfo,
			ConfidenceLevelType: ConfidenceLevelHigh,
			Description:         "Balance update logic is correctly implemented.",
		})
	}
}

// Example of a new check: Verify that the correct events are emitted
func (m *TransferDetector) checkForEventEmission(fnCtx *ast.Function, function *Function) {
	expectedEvent := "Transfer" // Default event for transfer functions
	if fnCtx.GetName() == "transferFrom" {
		expectedEvent = "Approval" // Approval event is also expected in transferFrom
	}

	eventEmitted := false
	for _, node := range fnCtx.GetNodes() {
		if emitStmt, ok := node.(*ast.Emit); ok {
			if eventCall, ok := emitStmt.GetExpression().(*ast.FunctionCall); ok {
				if exprCtx, ok := eventCall.GetExpression().(*ast.PrimaryExpression); ok {
					if exprCtx.GetName() == expectedEvent {
						eventEmitted = true
						break
					}
				}
			}
		}
	}

	if !eventEmitted {
		function.Detectors = append(function.Detectors, DetectorResult{
			DetectionType:       "event_emission_missing",
			SeverityType:        SeverityMedium,
			ConfidenceLevelType: ConfidenceLevelHigh,
			Description:         fmt.Sprintf("Expected %s event emission is missing.", expectedEvent),
		})
	} else {
		function.Detectors = append(function.Detectors, DetectorResult{
			DetectionType:       "event_emission_detected",
			SeverityType:        SeverityInfo,
			ConfidenceLevelType: ConfidenceLevelHigh,
			Description:         fmt.Sprintf("%s event emission is correctly implemented.", expectedEvent),
		})
	}
}

func (m *TransferDetector) checkForAccessControl(fnCtx *ast.Function, function *Function) {
	accessControlImplemented := false

	// Check if the function has modifiers
	for _, modifier := range fnCtx.GetModifiers() {
		if isAccessControlModifier(modifier) {
			accessControlImplemented = true
			break
		}
	}

	if !accessControlImplemented {
		function.Detectors = append(function.Detectors, DetectorResult{
			DetectionType:       "access_control_missing",
			SeverityType:        SeverityMedium,
			ConfidenceLevelType: ConfidenceLevelHigh,
			Description:         "Access control is missing or not properly implemented.",
		})
	} else {
		function.Detectors = append(function.Detectors, DetectorResult{
			DetectionType:       "access_control_detected",
			SeverityType:        SeverityInfo,
			ConfidenceLevelType: ConfidenceLevelHigh,
			Description:         "Access control is properly implemented.",
		})
	}
}

func isAccessControlModifier(modifier *ast.ModifierInvocation) bool {
	// Implement logic to identify access control modifiers
	// Common examples include 'onlyOwner' or custom-defined access control modifiers
	return modifier.GetName() == "onlyOwner" // Extend this check as needed
}

func (m *TransferDetector) checkForOwnerVariable(fnCtx *ast.Function, function *Function) {
	// Iterate through all variable declarations in the function
	for _, node := range fnCtx.GetNodes() {
		if varCtx, ok := node.(*ast.VariableDeclaration); ok {
			for _, declaration := range varCtx.GetDeclarations() {
				if declaration.GetName() == "owner" {
					// Add a result indicating the presence of 'owner' variable
					function.Detectors = append(function.Detectors, DetectorResult{
						DetectionType:       "owner_variable_detected",
						SeverityType:        SeverityInfo,
						ConfidenceLevelType: ConfidenceLevelHigh,
						Description:         "Owner variable is present in the function.",
					})
				}
			}
		}
	}
}

func (m *TransferDetector) checkForInternalTransferCall(fnCtx *ast.Function, function *Function) {
	// Iterate through all function call expressions in the function
	for _, node := range fnCtx.GetNodes() {
		if fcCtx, ok := node.(*ast.FunctionCall); ok {
			if exprCtx, ok := fcCtx.GetExpression().(*ast.PrimaryExpression); ok {
				if exprCtx.GetName() == "_transfer" {
					// Add a result indicating the use of internal _transfer call
					function.Detectors = append(function.Detectors, DetectorResult{
						DetectionType:       "internal_transfer_call_detected",
						SeverityType:        SeverityInfo,
						ConfidenceLevelType: ConfidenceLevelHigh,
						Description:         "Internal _transfer call is present in the function.",
					})
				}
			}
		}
	}
}

func (m *TransferDetector) Detect(ctx context.Context) (DetectorFn, error) {
	return map[ast_pb.NodeType]func(node ast.Node[ast.NodeType]) (bool, error){}, nil
}

func (m *TransferDetector) Exit(ctx context.Context) (DetectorFn, error) {
	return map[ast_pb.NodeType]func(node ast.Node[ast.NodeType]) (bool, error){}, nil
}

func (m *TransferDetector) Results() any {
	return m.results
}

func (m *TransferDetector) getStandardFunction(standard standards.EIP, fnName string) *standards.Function {
	standardFunctions := standard.GetFunctions()
	for _, fn := range standardFunctions {
		if fn.Name == fnName {
			return &fn
		}
	}
	return nil
}

func (m *TransferDetector) buildStandardFunction(fnCtx *ast.Function) standards.Function {
	var inputs []standards.Input
	var outputs []standards.Output

	if parametersList := fnCtx.GetParameters(); parametersList != nil {
		if parameters := parametersList.GetParameters(); parameters != nil {
			for _, param := range parameters {
				inputs = append(inputs, standards.Input{
					Type: param.GetTypeName().GetName(),
				})
			}
		}
	}

	if returnsList := fnCtx.GetReturnParameters(); returnsList != nil {
		if returns := returnsList.GetParameters(); returns != nil {
			for _, ret := range returns {
				outputs = append(outputs, standards.Output{
					Type: ret.GetTypeName().GetName(),
				})
			}
		}
	}

	return standards.Function{
		Name:    fnCtx.GetName(),
		Inputs:  inputs,
		Outputs: outputs,
	}
}
